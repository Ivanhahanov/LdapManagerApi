package ldap

import (
	"fmt"
	"gopkg.in/ldap.v2"
	"log"
)

func AddUser(){
	l := LdapConnect()
	defer l.Close()
	name := "vlukovskaya"
	fullName := "Vlada Lukovskaya"

	addReq := ldap.NewAddRequest(fmt.Sprintf("CN=%s,OU=users,dc=explabs,dc=ru", name))
	addReq.Attribute("objectClass", []string{"top", "inetOrgPerson"})
	addReq.Attribute("sn", []string{name})
	addReq.Attribute("userPassword", []string{name})
	addReq.Attribute("uid", []string{fullName})

	if err := l.Add(addReq); err != nil {
		log.Fatal("error adding service:", addReq, err)
	}
}
