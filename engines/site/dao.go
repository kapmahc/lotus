package site

import (
	"bytes"
	"encoding/gob"

	"github.com/astaxie/beego/orm"
)

//Set save setting
func Set(k string, v interface{}, f bool) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	val := buf.Bytes()
	if f {
		val, err = Encrypt(buf.Bytes())
		if err != nil {
			return err
		}
	}

	var m Setting
	o := orm.NewOrm()
	err = o.QueryTable(&m).Filter("key", k).One(&m, "flag", "val")
	if err == nil {
		m.Flag = f
		m.Val = string(val)
		_, err = o.Update(&m, "val")
	} else if err == orm.ErrNoRows {
		m.Key = k
		m.Flag = f
		m.Val = string(val)
		_, err = o.Insert(&m)
	}

	return err
}

//Get get setting value by key
func Get(k string, v interface{}) error {
	var m Setting
	err := orm.NewOrm().QueryTable(&m).Filter("key", k).One(&m)
	if err != nil {
		return err
	}
	val := []byte(m.Val)
	if m.Flag {
		if val, err = Decrypt(val); err != nil {
			return err
		}
	}

	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(val)
	return dec.Decode(v)
}
