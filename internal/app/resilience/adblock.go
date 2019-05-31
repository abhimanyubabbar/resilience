/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
/* Based on code by Jun Kimura. */

package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	adblockHasAnyPrefix  = adblockCreateCheckStringSetFunc(strings.HasPrefix)
	adblockContainsAny   = adblockCreateCheckStringSetFunc(strings.Contains)
	adblockBinaryOptions = []string{
		"script",
		"image",
		"stylesheet",
		"object",
		"xmlhttprequest",
		"object-subrequest",
		"subdocument",
		"document",
		"elemhide",
		"other",
		"background",
		"xbl",
		"ping",
		"dtd",
		"media",
		"third-party",
		"match-case",
		"collapse",
		"donottrack",
		"popup",
	}
	adblockEscapeSpecialRegexp = regexp.MustCompile(`([.$+?{}()\[\]\\])`)
)

type adblockRule struct {
	raw           string
	text          string
	regexString   string
	regex         *regexp.Regexp
	isComment     bool
	isHTMLRule    bool
	isException   bool
	options       map[string]bool
	domainOptions map[string]bool
	rawOptions    []string
	optionsKeys   []string
}

type adblockRules struct {
	rules                  []*adblockRule
	opt                    *adblockRulesOption
	blacklist              []*adblockRule
	whitelist              []*adblockRule
	blacklistRe            *regexp.Regexp
	whitelistRe            *regexp.Regexp
	blacklistWithOptions   []*adblockRule
	whitelistWithOptions   []*adblockRule
	blacklistRequireDomain map[string][]*adblockRule
	whitelistRequireDomain map[string][]*adblockRule
}

type adblockRulesOption struct {
	Supports              []string
	CheckUnsupportedRules bool
}

func (rule *adblockRule) OptionsKeys() []string {
	opts := []string{}
	for opt := range rule.options {
		if opt != "match-case" {
			opts = append(opts, opt)
		}
	}
	if rule.domainOptions != nil && len(rule.domainOptions) >= 0 {
		opts = append(opts, "domain")
	}
	return opts
}

func (rule *adblockRule) HasOption(key string) bool {
	if key == "domain" {
		return rule.domainOptions != nil && len(rule.domainOptions) >= 0
	}
	_, ok := rule.options[key]
	return ok
}

func (rule *adblockRule) DomainOptions() map[string]bool {
	return rule.domainOptions
}

func (rule *adblockRule) MatchingSupported(options map[string]interface{}, reverse bool) bool {
	if rule.isComment {
		return false
	}
	if rule.isHTMLRule {
		return false
	}
	if options == nil {
		options = map[string]interface{}{}
	}
	keys := adblockMapKeys(options)
	if !adblockIsSuperSet(rule.OptionsKeys(), keys, reverse) {
		return false
	}
	return true
}

func (rule *adblockRule) MatchURL(u string, options map[string]interface{}) bool {
	for opt := range rule.options {
		if opt == "match-case" {
			continue
		}
		if _, ok := options[opt]; !ok {
			return false
		}
		v, ok := options[opt]
		if ok {
			bl, ok := v.(bool)
			if ok {
				rv, ok := rule.options[opt]
				if ok {
					if bl != rv {
						return false
					}
				}
			}
		}
	}
	if len(rule.DomainOptions()) > 0 {
		if _, ok := options["domain"]; !ok {
			panic("Rule requires option domain")
		}
	}
	v, ok := options["domain"]
	if ok {
		sv := v.(string)
		if !rule.DomainMatches(sv) {
			return false
		}
	}
	return rule.URLMatches(u)
}

func (rule *adblockRule) URLMatches(u string) bool {
	if rule.regex == nil {
		rule.regex = regexp.MustCompile(rule.regexString)
	}
	return rule.regex.MatchString(u)
}

func (rule *adblockRule) DomainMatches(domain string) bool {
	for _, dm := range adblockDomainVariants(domain) {
		if bl, ok := rule.domainOptions[dm]; ok {
			return bl
		}
	}
	for _, bl := range rule.domainOptions {
		if bl {
			return false
		}
	}
	return true
}

func (rules *adblockRules) ShouldBlock(u string, options map[string]interface{}) bool {
	if rules.IsWhiteListed(u, options) {
		return false
	}
	if rules.IsBlackListed(u, options) {
		return true
	}
	return false
}

func (rules *adblockRules) IsWhiteListed(u string, options map[string]interface{}) bool {
	return rules.matches(u, options, rules.whitelistRe, rules.whitelistRequireDomain, rules.whitelistWithOptions)
}

func (rules *adblockRules) IsBlackListed(u string, options map[string]interface{}) bool {
	return rules.matches(u, options, rules.blacklistRe, rules.blacklistRequireDomain, rules.blacklistWithOptions)
}

func (rules *adblockRules) matches(u string, options map[string]interface{}, generalRe *regexp.Regexp, domainRequiredRules map[string][]*adblockRule, rulesWithOptions []*adblockRule) bool {
	if generalRe != nil && generalRe.MatchString(u) {
		return true
	}
	rls := []*adblockRule{}
	isrcDomain, ok := options["domain"]
	srcDomain, ok2 := isrcDomain.(string)
	if ok && ok2 && len(domainRequiredRules) > 0 {
		for _, domain := range adblockDomainVariants(srcDomain) {
			if vs, ok := domainRequiredRules[domain]; ok {
				rls = append(rls, vs...)
			}
		}
	}
	rls = append(rls, rulesWithOptions...)
	if !rules.opt.CheckUnsupportedRules {
		for _, rule := range rls {
			if rule.MatchingSupported(options, true) {
				if rule.MatchURL(u, options) {
					return true
				}
			}
		}
	}
	return false
}

func (rules *adblockRules) BlackList() []*adblockRule {
	return rules.blacklist
}

func (rules *adblockRules) WhiteList() []*adblockRule {
	return rules.whitelist
}

func adblockNewRule(text string) (*adblockRule, error) {
	rule := &adblockRule{}
	rule.raw = text
	text = strings.TrimSpace(text)
	rule.isComment = adblockHasAnyPrefix(text, "!", "[Adblock")
	if rule.isComment {
		rule.isHTMLRule = false
		rule.isException = false
	} else {
		rule.isHTMLRule = adblockContainsAny(text, "##", "#@#")
		rule.isException = strings.HasPrefix(text, "@@")
		if rule.isException {
			text = text[2:]
		}
	}
	rule.options = make(map[string]bool)
	if !rule.isComment && strings.Contains(text, "$") {
		var option string
		parts := strings.SplitN(text, "$", 2)
		length := len(parts)
		if length > 0 {
			text = parts[0]
		}
		if length > 1 {
			option = parts[1]
		}
		rule.rawOptions = strings.Split(option, ",")
		for _, opt := range rule.rawOptions {
			if strings.HasPrefix(opt, "domain=") {
				rule.domainOptions = adblockParseDomainOption(opt)
			} else {
				rule.options[strings.TrimLeft(opt, "~")] = !strings.HasPrefix(opt, "~")
			}
		}
	} else {
		rule.rawOptions = []string{}
		rule.domainOptions = make(map[string]bool)
	}
	rule.optionsKeys = rule.OptionsKeys()
	rule.text = text
	if rule.isComment || rule.isHTMLRule {
		rule.regexString = ""
	} else {
		var err error
		rule.regexString, err = adblockRuleToRegexp(text)
		if err != nil {
			return nil, err
		}
	}
	return rule, nil
}

func adblockNewRules(ruleStrs []string, opt *adblockRulesOption) (*adblockRules, error) {
	rls := &adblockRules{}
	if opt == nil {
		rls.opt = &adblockRulesOption{}
	} else {
		rls.opt = opt
	}
	if rls.opt.Supports == nil {
		rls.opt.Supports = append(adblockBinaryOptions, "domain")
	}
	params := adblockSliceToMap(rls.opt.Supports)
	for _, ruleStr := range ruleStrs {
		rule, err := adblockNewRule(ruleStr)
		if err != nil {
			return nil, err
		}
		if rule.regexString != "" && rule.MatchingSupported(params, false) {
			rls.rules = append(rls.rules, rule)
		}
	}
	advancedRules, basicRules := adblockSplitRuleData(rls.rules, func(rule *adblockRule) bool {
		if (rule.options != nil && len(rule.options) > 0) || (rule.domainOptions != nil && len(rule.domainOptions) > 0) {
			return true
		}
		return false
	})
	domainRequiredRules, NonDomainRules := adblockSplitRuleData(advancedRules, func(rule *adblockRule) bool {
		return rule.HasOption("domain") && adblockAnyTrueValue(rule.domainOptions)
	})
	rls.blacklist, rls.whitelist = adblockSplitBlackWhite(basicRules)
	rls.blacklistRe = adblockCombinedRegex(rls.blacklist)
	rls.whitelistRe = adblockCombinedRegex(rls.whitelist)
	rls.blacklistWithOptions, rls.whitelistWithOptions = adblockSplitBlackWhite(NonDomainRules)
	rls.blacklistRequireDomain, rls.whitelistRequireDomain = adblockSplitBlackWhiteDomain(domainRequiredRules)
	return rls, nil
}

func adblockCreateCheckStringSetFunc(checkFunc func(string, string) bool) func(string, ...string) bool {
	return func(s string, sets ...string) bool {
		for _, set := range sets {
			if checkFunc(s, set) {
				return true
			}
		}
		return false
	}
}

func adblockParseDomainOption(text string) map[string]bool {
	domains := text[len("domain="):]
	parts := strings.Split(strings.Replace(domains, ",", "|", -1), "|")
	opts := make(map[string]bool, len(parts))
	for _, part := range parts {
		opts[strings.TrimLeft(part, "~")] = !strings.HasPrefix(part, "~")
	}
	return opts
}

func adblockRuleToRegexp(text string) (string, error) {
	if text == "" {
		return ".*", nil
	}
	length := len(text)
	if length >= 2 && text[:1] == "/" && text[length-1:] == "/" {
		return text[1 : length-1], nil
	}
	rule := adblockEscapeSpecialRegexp.ReplaceAllStringFunc(text, func(src string) string {
		return fmt.Sprintf(`\%v`, src)
	})
	rule = strings.Replace(rule, "^", `(?:[^\\w\\d_\\\-.%]|$)`, -1)
	rule = strings.Replace(rule, "*", ".*", -1)
	length = len(rule)
	if rule[length-1] == '|' {
		rule = rule[:length-1] + "$"
	}
	if rule[:2] == "||" {
		if len(rule) > 2 {
			rule = `^(?:[^:/?#]+:)?(?://(?:[^/?#]*\\.)?)?` + rule[2:]
		}
	} else if rule[0] == '|' {
		rule = "^" + rule[1:]
	}
	rule = regexp.MustCompile(`(\|)[^$]`).ReplaceAllString(rule, `\|`)
	return rule, nil
}

func adblockDomainVariants(domain string) []string {
	variants := []string{}
	parts := strings.Split(domain, ".")
	for i := len(parts); i > 1; i-- {
		p := parts[len(parts)-i:]
		variants = append(variants, strings.Join(p, "."))
	}
	return variants
}

func adblockSliceToMap(sl []string) map[string]interface{} {
	opts := make(map[string]interface{})
	for _, v := range sl {
		opts[v] = true
	}
	return opts
}

func adblockCombinedRegex(rules []*adblockRule) *regexp.Regexp {
	regexes := []string{}
	for _, rule := range rules {
		regexes = append(regexes, rule.regexString)
	}
	rs := strings.Join(regexes, "|")
	if rs == "" {
		return nil
	}
	return regexp.MustCompile(rs)
}

func adblockMapKeys(m map[string]interface{}) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func adblockIsSuperSet(a, b []string, reverse bool) bool {
	var (
		mr map[string]interface{}
		sr []string
	)
	if !reverse {
		mr = adblockSliceToMap(b)
		sr = a
	} else {
		mr = adblockSliceToMap(a)
		sr = b
	}
	for _, key := range sr {
		_, ok := mr[key]
		if !ok {
			return false
		}
	}
	return true
}

func adblockSplitRuleData(iter []*adblockRule, pred func(*adblockRule) bool) ([]*adblockRule, []*adblockRule) {
	var yes, no []*adblockRule
	for _, v := range iter {
		if pred(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return yes, no
}

func adblockSplitBlackWhite(rules []*adblockRule) ([]*adblockRule, []*adblockRule) {
	return adblockSplitRuleData(rules, func(rule *adblockRule) bool {
		return !rule.isException
	})
}

func adblockSplitBlackWhiteDomain(rules []*adblockRule) (map[string][]*adblockRule, map[string][]*adblockRule) {
	blacklist, whitelist := adblockSplitBlackWhite(rules)
	return adblockDomainIndex(blacklist), adblockDomainIndex(whitelist)
}

func adblockDomainIndex(rules []*adblockRule) map[string][]*adblockRule {
	result := make(map[string][]*adblockRule)
	for _, rule := range rules {
		for domain, required := range rule.domainOptions {
			if required {
				result[domain] = append(result[domain], rule)
			}
		}
	}
	return result
}

func adblockAnyTrueValue(mp map[string]bool) bool {
	for _, it := range mp {
		if it {
			return true
		}
	}
	return false
}
