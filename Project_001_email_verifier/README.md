# Email Verfier

- This tool helps in verfication of the email with the below vaildations as speciied in capabilities.
- This tool is light weight and also outputs the results in tabular format.
- This tool also supports bulk email vaildation
- DEA support to identify the disposable email will be added soon


## Current Abilities

[x] validate email address
[x] verify the DNS using DNSTXT records
[x] verify the DNS MX records
[ ] verify the smtp
[ ] verify the DEA (disposable email address) 

## Example Ouputs

- Successfull email validation

```
+----------------------------------------------------------------------------------------------------+
| Email Validator : pree@gmail.com                                                                   |
+---+-----------------+---------+--------------------------------------------------------------------+
| # | VALIDATION TYPE | RESULT  | RESULT DESCRIPTION                                                 |
+---+-----------------+---------+--------------------------------------------------------------------+
| 1 | Email Validity  | SUCCESS |                                                                    |
| 2 | DNS TXT Records | SUCCESS | v=spf1 redirect=_spf.google.com                                    |
| 3 | MX TXT Records  | SUCCESS | [0xc000008120 0xc000008150 0xc000008168 0xc000008138 0xc000008180] |
+---+-----------------+---------+--------------------------------------------------------------------+
```

- Failure email validation

```
+---------------------------------------------------------------------------------------+
| Email Validator : pree@pfwrfr.com                                                     |
+---+-----------------+---------+-------------------------------------------------------+
| # | VALIDATION TYPE | RESULT  | RESULT DESCRIPTION                                    |
+---+-----------------+---------+-------------------------------------------------------+
| 1 | Email Validity  | SUCCESS |                                                       |
| 2 | DNS TXT Records | ERROR   | lookup pfwrfr.com: dnsquery: DNS name does not exist. |
| 3 | MX TXT Records  | ERROR   | lookup pfwrfr.com: dnsquery: DNS name does not exist. |
+---+-----------------+---------+-------------------------------------------------------+
```
