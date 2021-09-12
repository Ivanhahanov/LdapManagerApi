package ldap

import (
	"gopkg.in/ldap.v2"
	"log"
)

func LdapConnect() *ldap.Conn {
	l, err := ldap.Dial("tcp", "127.0.0.1:389")
	if err != nil {
		log.Fatal("connect",err)
	}
	// TODO: add tls
	// TODO: replace to conf cred enter
	authErr := l.Bind("cn=admin,dc=explabs,dc=ru", "admin")
	if authErr != nil {
		log.Fatal("auth", err)
	}
	return l
}
