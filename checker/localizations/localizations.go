// Code generated by go-localize; DO NOT EDIT.
// This file was generated by robots at
// 2022-11-24 19:18:50.59362609 +0300 MSK m=+0.002051350

package localizations

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var localizations = map[string]string{
	"en.messages.added-required-request-body":                              "added required request body",
	"en.messages.api-deprecated-sunset-parse":                              "api sunset date can't be parsed for deprecated API",
	"en.messages.api-path-removed-before-sunset":                           "api path removed before the sunset date %s",
	"en.messages.api-path-removed-without-deprecation":                     "api path removed without deprecation",
	"en.messages.api-removed-before-sunset":                                "api removed before the sunset date %s",
	"en.messages.api-removed-without-deprecation":                          "api removed without deprecation",
	"en.messages.api-sunset-date-too-small":                                "api sunset date changed to earlier date from %s to %s, new sunset date must be not earlier than %s at least %d days from now",
	"en.messages.at":                                                       "at",
	"en.messages.in":                                                       "in",
	"en.messages.new-request-path-parameter":                               "added the new path request parameter %s",
	"en.messages.new-required-request-header-property":                     "added the new required %s request header's property %s",
	"en.messages.new-required-request-parameter":                           "added the new required %s request parameter %s",
	"en.messages.new-required-request-property":                            "added the new required request property %s",
	"en.messages.optional-response-header-removed":                         "the optional response header %s removed for the status %s",
	"en.messages.pattern-changed-warn-comment":                             "It is the warning because it is difficult to automatically analyze if the new pattern is a superset of the previous pattern(e.g. changed from '[0-9]+' to '[0-9]*')",
	"en.messages.request-allOf-modified":                                   "modified allOf for the request property %s",
	"en.messages.request-allOf-modified-comment":                           "It is a warn because it is very difficult to check that allOf changed correctly without breaking changes",
	"en.messages.request-body-became-required":                             "request body became required",
	"en.messages.request-body-max-decreased":                               "the request's body max was decreased to %s",
	"en.messages.request-body-max-length-decreased":                        "the request's body maxLength was decreased to %s",
	"en.messages.request-body-max-length-set":                              "the request's body maxLength was set to %s",
	"en.messages.request-body-max-length-set-comment":                      "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-body-max-set":                                     "the request's body max was set to %s",
	"en.messages.request-body-max-set-comment":                             "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-body-min-increased":                               "the request's body min was increased to %s",
	"en.messages.request-body-min-items-increased":                         "the request's body minItems was increased to %s",
	"en.messages.request-body-min-items-set":                               "the request's body minItems was set to %s",
	"en.messages.request-body-min-items-set-comment":                       "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-body-min-set":                                     "the request's body min was set to %s",
	"en.messages.request-body-min-set-comment":                             "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-body-type-changed":                                "the request's body type/format changed from %s/%s to %s/%s",
	"en.messages.request-header-property-became-required":                  "the %s request header's property %s became required",
	"en.messages.request-parameter-became-required":                        "the %s request parameter %s became required",
	"en.messages.request-parameter-enum-value-removed":                     "removed the enum value %s for the %s request parameter %s",
	"en.messages.request-parameter-max-decreased":                          "for the %s request parameter %s, the max was decreased from %s to %s",
	"en.messages.request-parameter-max-length-decreased":                   "for the %s request parameter %s, the maxLength was decreased from %s to %s",
	"en.messages.request-parameter-max-length-set":                         "for the %s request parameter %s, the maxLength was set to %s",
	"en.messages.request-parameter-max-length-set-comment":                 "It is warn because sometimes it is required to be set because of security reasons or current error in specification. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-parameter-max-set":                                "for the %s request parameter %s, the max was set to %s",
	"en.messages.request-parameter-max-set-comment":                        "It is warn because sometimes it is required to be set because of security reasons or current error in specification. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-parameter-min-increased":                          "for the %s request parameter %s, the min was increased from %s to %s",
	"en.messages.request-parameter-min-items-increased":                    "for the %s request parameter %s, the minItems was increased from %s to %s",
	"en.messages.request-parameter-min-items-set":                          "for the %s request parameter %s, the minItems was set to %s",
	"en.messages.request-parameter-min-items-set-comment":                  "It is warn because sometimes it is required to be set because of security reasons or current error in specification. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-parameter-min-set":                                "for the %s request parameter %s, the min was set to %s",
	"en.messages.request-parameter-min-set-comment":                        "It is warn because sometimes it is required to be set because of security reasons or current error in specification. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-parameter-pattern-added":                          "added the pattern '%s' for the %s request parameter %s",
	"en.messages.request-parameter-pattern-changed":                        "changed the pattern for the %s request parameter %s from '%s' to '%s'",
	"en.messages.request-parameter-removed":                                "deleted the %s request parameter %s",
	"en.messages.request-parameter-type-changed":                           "for the %s request parameter %s, the type/format was changed from %s/%s to %s/%s",
	"en.messages.request-parameter-x-extensible-enum-value-removed":        "removed the x-extensible-enum value %s for the %s request parameter %s",
	"en.messages.request-property-became-required":                         "the request property %s became required",
	"en.messages.request-property-enum-value-removed":                      "removed the enum value %s of the request property %s",
	"en.messages.request-property-max-decreased":                           "the %s request property's max was decreased to %s",
	"en.messages.request-property-max-length-decreased":                    "the %s request property's maxLength was decreased to %s",
	"en.messages.request-property-max-length-set":                          "the %s request property's maxLength was set to %s",
	"en.messages.request-property-max-length-set-comment":                  "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-property-max-set":                                 "the %s request property's max was set to %s",
	"en.messages.request-property-max-set-comment":                         "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-property-min-increased":                           "the %s request property's min was increased to %s",
	"en.messages.request-property-min-items-increased":                     "the %s request property's minItems was increased to %s",
	"en.messages.request-property-min-items-set":                           "the %s request property's minItems was set to %s",
	"en.messages.request-property-min-items-set-comment":                   "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-property-min-set":                                 "the %s request property's min was set to %s",
	"en.messages.request-property-min-set-comment":                         "It is warn because sometimes it is required to be set. But good clients should be checked to support this restriction before such change in specification.",
	"en.messages.request-property-pattern-added":                           "added the pattern '%s' for the request property %s",
	"en.messages.request-property-pattern-changed":                         "changed the pattern for the request property %s from '%s' to '%s'",
	"en.messages.request-property-removed":                                 "removed the request property %s",
	"en.messages.request-property-type-changed":                            "the %s request property type/format changed from %s/%s to %s/%s",
	"en.messages.request-property-x-extensible-enum-value-removed":         "removed the x-extensible-enum value '%s' of the request property %s",
	"en.messages.required-response-header-removed":                         "the mandatory response header %s removed for the status %s",
	"en.messages.response-body-max-length-increased":                       "the response's body maxLength was increased from %s to %s",
	"en.messages.response-body-max-length-unset":                           "the response's body maxLength was unset from %s",
	"en.messages.response-body-min-items-decreased":                        "the response's body minItems was decreased from %s to %s",
	"en.messages.response-body-min-items-unset":                            "the response's body minItems was unset from %s",
	"en.messages.response-body-min-length-decreased":                       "the response's body minLength was decreased from %s to %s",
	"en.messages.response-body-type-changed":                               "the response's body type/format changed from %s/%s to %s/%s for status %s",
	"en.messages.response-header-became-optional":                          "the response header %s became optional for the status %s",
	"en.messages.response-media-type-removed":                              "removed the media type %s for the response with the status %s",
	"en.messages.response-optional-property-removed":                       "removed the optional property %s from the response with the %s status",
	"en.messages.response-property-became-optional":                        "the response property %s became optional for the status %s",
	"en.messages.response-property-enum-value-added":                       "added the new '%s' enum value the %s response property for the response status %s",
	"en.messages.response-property-enum-value-added-comment":               "Adding new enum values to response could be unexpected for clients, use x-extensible-enum instead.",
	"en.messages.response-property-max-length-increased":                   "the %s response property's maxLength was increased from %s to %s for the response status %s",
	"en.messages.response-property-max-length-unset":                       "the %s response property's maxLength was unset from %s for the response status %s",
	"en.messages.response-property-min-items-decreased":                    "the %s response property's minItems was decreased from %s to %s for the response status %s",
	"en.messages.response-property-min-items-unset":                        "the %s response property's minItems was unset from %s for the response status %s",
	"en.messages.response-property-min-length-decreased":                   "the %s response property's minLength was decreased from %s to %s for the response status %s",
	"en.messages.response-property-type-changed":                           "the response's property type/format changed from %s/%s to %s/%s for status %s",
	"en.messages.response-required-property-became-not-write-only":         "the response required property %s became not write-only for the status %s",
	"en.messages.response-required-property-became-not-write-only-comment": "It is valid only if the property was always returned before the specification has been changed",
	"en.messages.response-required-property-removed":                       "removed the required property %s from the response with the %s status",
	"en.messages.response-success-status-removed":                          "removed the success response with the status %s",
	"en.messages.sunset-deleted":                                           "api sunset date deleted, but deprecated=true kept",
	"en.messages.total-errors":                                             "Backward compatibility errors (%d):\n",
	"ru.messages.added-required-request-body":                              "добавлено обязательное тело запроса",
	"ru.messages.at":                                                       "в",
	"ru.messages.in":                                                       "в",
	"ru.messages.request-parameter-pattern-added":                          "добавлен pattern '%s' для %s параметра запроса %s",
	"ru.messages.request-parameter-pattern-changed":                        "изменён pattern для %s параметра запроса %s со значения '%s' на значение '%s'",
	"ru.messages.request-parameter-removed":                                "удалён %s параметр запроса %s",
	"ru.messages.request-property-pattern-added":                           "добавлен pattern '%s' для поля запроса %s",
	"ru.messages.request-property-pattern-changed":                         "изменён pattern для поля запроса %s со значения '%s' на значение '%s'",
	"ru.messages.total-errors":                                             "Ошибки обратной совместимости (всего: %d):\n",
}

type Replacements map[string]interface{}

type Localizer struct {
	Locale         string
	FallbackLocale string
	Localizations  map[string]string
}

func New(locale string, fallbackLocale string) *Localizer {
	t := &Localizer{Locale: locale, FallbackLocale: fallbackLocale}
	t.Localizations = localizations
	return t
}

func (t Localizer) SetLocales(locale, fallback string) Localizer {
	t.Locale = locale
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) SetLocale(locale string) Localizer {
	t.Locale = locale
	return t
}

func (t Localizer) SetFallbackLocale(fallback string) Localizer {
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) GetWithLocale(locale, key string, replacements ...*Replacements) string {
	str, ok := t.Localizations[t.getLocalizationKey(locale, key)]
	if !ok {
		str, ok = t.Localizations[t.getLocalizationKey(t.FallbackLocale, key)]
		if !ok {
			return key
		}
	}

	// If the str doesn't have any substitutions, no need to
	// template.Execute.
	if strings.Index(str, "}}") == -1 {
		return str
	}

	return t.replace(str, replacements...)
}

func (t Localizer) Get(key string, replacements ...*Replacements) string {
	str := t.GetWithLocale(t.Locale, key, replacements...)
	return str
}

func (t Localizer) getLocalizationKey(locale string, key string) string {
	return fmt.Sprintf("%v.%v", locale, key)
}

func (t Localizer) replace(str string, replacements ...*Replacements) string {
	b := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(str)
	if err != nil {
		return str
	}

	replacementsMerge := Replacements{}
	for _, replacement := range replacements {
		for k, v := range *replacement {
			replacementsMerge[k] = v
		}
	}

	err = template.Must(tmpl, err).Execute(b, replacementsMerge)
	if err != nil {
		return str
	}
	buff := b.String()
	return buff
}