package zeroconf

import "strings"

func parseSubtypes(service string) (string, []string) {
	subtypes := strings.Split(service, ",")
	return subtypes[0], subtypes[1:]
}

// trimDot is used to trim the dots from the start or end of a string
func trimDot(s string) string {
	return strings.Trim(s, ".")
}

func split(s string) (instance, service, domain string) {
	all := strings.Split(s, ".")
	b, e := 0, 0
	for i, v := range all {
		if strings.HasPrefix(v, "_") {
			e = i
			if b == 0 {
				b = i
			}
		}
	}
	instance = strings.Join(all[:b], ".")
	service = strings.Join(all[b:e+1], ".") + "."
	domain = strings.Join(all[e+1:], ".")
	domain = strings.TrimRight(domain, ".")

	return
}
