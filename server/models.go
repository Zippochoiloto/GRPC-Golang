package main

import (
	"grpc-course/contact/contactpb"
	"log"

	"github.com/astaxie/beego/orm"
)

type ContactInfo struct {
	PhoneNumber string `orm:"size(15);pk"`
	Name        string `orm:"size(400)"`
	Address     string `orm:"type(text)"`
}

func (c *ContactInfo) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	if err != nil {
		log.Printf("Insert contact %+v error: %v\n", c, err)
		return err
	}

	log.Printf("Insert %+v succesful", c)
	return nil
}

func ConvertPbContact2ContactInfo(pbContact *contactpb.Contact) *ContactInfo {
	return &ContactInfo{
		PhoneNumber: pbContact.PhoneNumber,
		Name:        pbContact.Name,
		Address:     pbContact.Address,
	}
}
