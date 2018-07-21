package tools

import (
	"testing"
	"strings"
	"fmt"
)

func Test1(t *testing.T) {
	a := "bill_no	tnt_inst_id	creator	last_modifier	relative_bill_no	event_no	service_target	service_type	product_code	event_code	fee_item_code	ar_type	ar_no	ar_version	price_code	price_version	rate_id	bd_product_code	sign_uid	settle_uid	payer_uid	payer_account	payer_account_type	payee_uid	payee_account	payee_account_type	inst_code	settle_type	writeoff_status	service_amount	calc_basic_amount	bill_amount	bill_currency	real_amount	nonpayment_amount	gaap_amount	pay_origin	payment_way	payment_no	gmt_pay	gmt_receive	gmt_record	gmt_service	gmt_create	gmt_modified	type	bill_source_type	memo	properties	cnl_pd_code	cnl_ev_code	cnl_no	biz_pd_code	biz_ev_code	pd_code	ev_code	rate	policy_code	dt	pt_source_type"
	t.Log(strings.Replace(a,"\\t",",",0))
}

func Test2(t *testing.T)  {
	x,y:=1,2
	x,y = y,x+y
	fmt.Println(x,y)
}