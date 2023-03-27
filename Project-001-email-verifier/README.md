# Email Verfier

- This tool helps in verfication of the email with the below vaildations as speciied in capabilities.
- This tool is light weight and also outputs the results in tabular format.
- This tool also supports bulk email vaildation
- DEA support to identify the disposable email will be added soon


## Current Abilities

- [x] validate email address
- [x] verify the DNS using DNSTXT records
- [x] verify the DNS MX records
- [x] verify the DEA (disposable email address)*
- [x] verify the smtp**

* only if domain is listed [here](https://raw.githubusercontent.com/ivolo/disposable-email-domains/master/index.json)
** may not work as expected and can often lead to error

## Example Ouputs

- Successfull email validation

```
+---------------------------------------------------------------------------------------------------------+
| Email Verifier : pree@gmail.com                                                                         |
+---+----------------------+---------+--------------------------------------------------------------------+
| # | VALIDATION TYPE      | RESULT  | RESULT DESCRIPTION                                                 |
+---+----------------------+---------+--------------------------------------------------------------------+
| 1 | Valid Email          | SUCCESS |                                                                    |
| 2 | DNS TXT Records      | SUCCESS | globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=   |
| 3 | MX TXT Records       | SUCCESS | [0xc000008150 0xc000008138 0xc000008120 0xc0000080a8 0xc000008168] |
| 4 | Non Disposable Email | SUCCESS | email is NOT disposable                                            |
| 5 | SMTP Verification    | SUCCESS | SMTP Verification done                                             |
+---+----------------------+---------+--------------------------------------------------------------------+
```

- Failure email validation

```
Enter email address
pree@dkhdsvh.
INFO : Here is the email address entered : pree@dkhdsvh.
ERROR : Invalid email :  pree@dkhdsvh.
```

---

```
+------------------------------------------------------------------------------------+
| Email Verifier : pree@zzzz1717.com                                                 |
+---+----------------------+---------+-----------------------------------------------+
| # | VALIDATION TYPE      | RESULT  | RESULT DESCRIPTION                            |
+---+----------------------+---------+-----------------------------------------------+
| 1 | Valid Email          | SUCCESS |                                               |
| 2 | DNS TXT Records      | SUCCESS | v=spf1 a mx include:spf.forwardemail.net -all |
| 3 | MX TXT Records       | SUCCESS | [0xc000008180 0xc0000081e0]                   |
| 4 | Non Disposable Email | ERROR   | email is disposable                           |
+---+----------------------+---------+-----------------------------------------------+
```

---

```
+------------------------------------------------------------------------------------------------+
| Email Verifier : pree@aslkhvfsihvrsh                                                           |
+---+----------------------+---------+-----------------------------------------------------------+
| # | VALIDATION TYPE      | RESULT  | RESULT DESCRIPTION                                        |
+---+----------------------+---------+-----------------------------------------------------------+
| 1 | Valid Email          | ERROR   | mail: no angle-addr                                       |
| 2 | DNS TXT Records      | ERROR   | lookup aslkhvfsihvrsh: dnsquery: DNS name does not exist. |
| 3 | MX TXT Records       | ERROR   | lookup aslkhvfsihvrsh: dnsquery: DNS name does not exist. |
| 4 | Non Disposable Email | SUCCESS | email is NOT disposable                                   |
+---+----------------------+---------+-----------------------------------------------------------+
```

---

```
+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Email Verifier : pree@gmail.com                                                                                                                                                                                                                                                |
+---+----------------------+---------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| # | VALIDATION TYPE      | RESULT  | RESULT DESCRIPTION                                                                                                                                                                                                                        |
+---+----------------------+---------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| 1 | Valid Email          | SUCCESS |                                                                                                                                                                                                                                           |
| 2 | DNS TXT Records      | SUCCESS | v=spf1 redirect=_spf.google.com                                                                                                                                                                                                           |
| 3 | MX TXT Records       | SUCCESS | [0xc00028a0a8 0xc00028a090 0xc00028a078 0xc00028a060 0xc00028a0c0]                                                                                                                                                                        |
| 4 | Non Disposable Email | SUCCESS | email is NOT disposable                                                                                                                                                                                                                   |
| 5 | SMTP Verification    | ERROR   | dial tcp [2404:6800:4003:c01::1b]:587: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond. |
+---+----------------------+---------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
```