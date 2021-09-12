package ldap

import (
	"fmt"
	"gopkg.in/ldap.v2"
	"log"
)

func SearchUser(user string) []*ldap.Entry {
	baseDN := "DC=explabs,DC=ru"
	filter := fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(user))

	// Filters must start and finish with ()!
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, 0, 0, 0, false, filter, []string{}, []ldap.Control{})

	l := LdapConnect()
	defer l.Close()

	result, err := l.Search(searchReq)
	if err != nil {
		log.Println("failed to query LDAP: %w", err)
	}

	log.Println("Got", len(result.Entries), "search results")
	return result.Entries
}
