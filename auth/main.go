package main

import (
	"bytedance_hertz_mod/auth/aes_util"
	"fmt"
)

func main() {
	msg := "Vr3Qaq9PGnpg7r0Q6Q6RA7TlhGIraP9oKsp1ae5fT89nGf+JkV9UQp6RNGBoWS0lU22XVaIKgo72D+Ri27Fcu1r2nNBms3bF9SAI3QkfOaTxWqjc6U7x0jIZV6lB0T2KeWoqP75Jphpez81iI/ubzLhesJY4151wMPqhoklOu3DtT4Lf0ycHiqnBHahIkB5faBwPofGnJW9fDUqAss6vMVJSJijZTZPf0qBx4wQ4OdpJ4EzdyJHnDnXvyZ335Jao"
	au := aes_util.AesUtil{}

	decrpt, err := au.Decrypt(msg, "Kl69MKUF7OXHHkOrjyKC6j3DG1sgkx7QLLkuTj9B2Zw")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(decrpt)

}
