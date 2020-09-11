NAME: vault
LAST DEPLOYED: Fri Sep 11 14:23:13 2020
NAMESPACE: vault
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Thank you for installing HashiCorp Vault!

Now that you have deployed Vault, you should look over the docs on using
Vault with Kubernetes available here:

https://www.vaultproject.io/docs/


Your release is named vault. To learn more about the release, try:

  $ helm status vault
  $ helm get vault


Unseal Key 1: DX8Mo/KrROxyEfSh370d4UGdEB50tXZTblS/FIgMefQ=

Initial Root Token: s.PdnFT2kl7FrYamylkEsLTXt6

Vault initialized with 1 key shares and a key threshold of 1. Please securely
distribute the key shares printed above. When the Vault is re-sealed,
restarted, or stopped, you must supply at least 1 of these keys to unseal it
before it can start servicing requests.

Vault does not store the generated master key. Without at least 1 key to
reconstruct the master key, Vault will remain permanently sealed!

It is possible to generate new unseal keys, provided you have a quorum of
existing unseal keys shares. See "vault operator rekey" for more information.

Key             Value
---             -----
Seal Type       shamir
Initialized     true
Sealed          false
Total Shares    1
Threshold       1
Version         1.5.2
Cluster Name    vault-cluster-37f4ac16
Cluster ID      74221bdc-9bb2-2538-0cc2-4966fe582402
HA Enabled      true
HA Cluster      https://vault-0.vault-internal:8201
HA Mode         active


Success! You are now authenticated. The token information displayed below
is already stored in the token helper. You do NOT need to run "vault login"
again. Future Vault requests will automatically use this token.

Key                  Value
---                  -----
token                s.PdnFT2kl7FrYamylkEsLTXt6
token_accessor       Gjx4CSEDajDvCrLgmaf6WiZv
token_duration       ∞
token_renewable      false
token_policies       ["root"]
identity_policies    []
policies             ["root"]


Path      Type     Accessor               Description
----      ----     --------               -----------
token/    token    auth_token_85283dab    token based credentials


Key                 Value
---                 -----
refresh_interval    768h
password            asajkjkahs
username            otus

====== Data ======
Key         Value
---         -----
password    asajkjkahs
username    otus


Path           Type          Accessor                    Description
----           ----          --------                    -----------
kubernetes/    kubernetes    auth_kubernetes_5d08fd78    n/a
token/         token         auth_token_85283dab         token based credentials



Мы не смогли обновить otus-rw/config, т.к. в политике не было capabilities "update". Но смогли создать otus-rw/config1, т.к. есть capabilities "create". 




Key                 Value
---                 -----
ca_chain            [-----BEGIN CERTIFICATE-----
MIIDnDCCAoSgAwIBAgIUe7rmqaKUBQ8PfS2KpT6rruAGhaswDQYJKoZIhvcNAQEL
BQAwFTETMBEGA1UEAxMKZXhtYXBsZS5ydTAeFw0yMDA5MTExMzQ4MTJaFw0yNTA5
MTAxMzQ4NDJaMCwxKjAoBgNVBAMTIWV4YW1wbGUucnUgSW50ZXJtZWRpYXRlIEF1
dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANv5BS/pwKFv
36bZtp75ATCljDMp06z21+P5ZKyZ2/7hMKAR2puBp8Yd3zTR6oKW63faqmbuA6iV
7LiCzw7pnGHrSw6RIZf4LQt7aUuhEELLs/tp/THUxPoLCRB0OwxSThbn1aFEAGE9
4TDFZ+rD6EXQhBiFwjjwEmGzKoXiUzj5RNed+5AOicrSuwCMzwFoOd0WytJ/5RDH
Q2F7cGIeUh/9rC61+YsdJzg6woUg/1ljtE8fee4BDic4jMGQZvJNNiINDWFe6HTS
STGS6WMlxLw3M+eZMBJwu/LFsvrEjf/6kcbl1OLdEWF52p7KAN3Dw5zJ+pekSqiF
580zFgD2R5cCAwEAAaOBzDCByTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUw
AwEB/zAdBgNVHQ4EFgQUbD/HIwNualdNlQNr+PquYnFHrSAwHwYDVR0jBBgwFoAU
coM/J/hxQ8wUCpTlFdvJmK4cxIcwNwYIKwYBBQUHAQEEKzApMCcGCCsGAQUFBzAC
hhtodHRwOi8vdmF1bHQ6ODIwMC92MS9wa2kvY2EwLQYDVR0fBCYwJDAioCCgHoYc
aHR0cDovL3ZhdWx0OjgyMDAvdjEvcGtpL2NybDANBgkqhkiG9w0BAQsFAAOCAQEA
JTpzXTTqogDAAMERsTE7+sPW1hMHchw/5+wVYaroMGobWsA6are7wug0d+mJJMEy
FpeP/S4RgknnGihq7chLEdSKAj4M0Rr0998hnb8KxROd5hyhbOr6LZ9NUFTrMuT6
8F9T17lVkrnIznHE2HN4cF3kyD/2TfralQDmJTNtJN85xdhBz9rv7XQC4Vrl0jev
qoCGMdDRKfSMT5NqFk37zQHA08FfHqG7ZxrPc6LBStZZhn2pJ6ryvmD1AffuxG2h
Of98IqQy8y+8PuF1bilDyigiGxjVqlbp2otPAFNXnaKVpn9wKEXZh8I0PEKXTIOy
v1ricup/6lWCc65sLdatZg==
-----END CERTIFICATE-----]
certificate         -----BEGIN CERTIFICATE-----
MIIDZzCCAk+gAwIBAgIUVY6HUEFT2dfLqVZyhsFcRpDaP04wDQYJKoZIhvcNAQEL
BQAwLDEqMCgGA1UEAxMhZXhhbXBsZS5ydSBJbnRlcm1lZGlhdGUgQXV0aG9yaXR5
MB4XDTIwMDkxMTEzNTE0OVoXDTIwMDkxMjEzNTIxOVowHDEaMBgGA1UEAxMRZ2l0
bGFiLmV4YW1wbGUucnUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDD
meiilG4SjniGpw0OeIdamVlfuGRhBTJydWykZUXw+RV+aqxLagu9kCmaOffy5sAy
bBsB0lsEeQTBjN+n52iLhs919t25UE3tTeb+NF68oe3WEAo/lMftuLK56qq85w7X
abJ0MjjPq7V9xCk3HZstR7hiZBlRCcgnStD7I/o0ueMCQTR13nphBFbQs3TUgcU0
zLkfTj09xkc0selcwEMPlPvUn/wQSCUCa2CLK8ydwANowxJXaQM7t2C+mIdTHIMc
j1P0FYRXPHS/1BsCaNEYmR7bWaPXo5ftB9rvwxqdMpaSq7WNYEsu3wTxSa4g2JfM
osQba6qUPr7swxuG35yzAgMBAAGjgZAwgY0wDgYDVR0PAQH/BAQDAgOoMB0GA1Ud
JQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQUvQ3jCDt+FV86V+19
BU7wYrEADfAwHwYDVR0jBBgwFoAUbD/HIwNualdNlQNr+PquYnFHrSAwHAYDVR0R
BBUwE4IRZ2l0bGFiLmV4YW1wbGUucnUwDQYJKoZIhvcNAQELBQADggEBALNhg1Wm
uFulKTPCtg4owcA1i/SWTa78K4Ke7oveDNOaQIfWkB0RQhNICTzyqMe7k9U+Zkj7
Bhv1rLgMFw0KixGFKLJVgb3ozTCM7uWeN0tx9QJEWAJQXRUjy+XtU8m1xZUkiy4E
IG71iDmPHVovKHuwqEUy8k824tTzz1NewSmOmGc6dvY9SgV+DRdRKjyqUbLIfP2O
twizcRQeZlvf7f7Pac42ooOZMwLjD73Fu6vOKXx/Dom8oPjPQD+5oUQVCGSxOD5X
JFbKpAR3ylSdr/GMnxocBXzBTXIp9iSLpQmkddHIMQUHotYvQmgt8u6fvB+NJRTM
rgoTfgjmT5ImlgY=
-----END CERTIFICATE-----
expiration          1599918739
issuing_ca          -----BEGIN CERTIFICATE-----
MIIDnDCCAoSgAwIBAgIUe7rmqaKUBQ8PfS2KpT6rruAGhaswDQYJKoZIhvcNAQEL
BQAwFTETMBEGA1UEAxMKZXhtYXBsZS5ydTAeFw0yMDA5MTExMzQ4MTJaFw0yNTA5
MTAxMzQ4NDJaMCwxKjAoBgNVBAMTIWV4YW1wbGUucnUgSW50ZXJtZWRpYXRlIEF1
dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANv5BS/pwKFv
36bZtp75ATCljDMp06z21+P5ZKyZ2/7hMKAR2puBp8Yd3zTR6oKW63faqmbuA6iV
7LiCzw7pnGHrSw6RIZf4LQt7aUuhEELLs/tp/THUxPoLCRB0OwxSThbn1aFEAGE9
4TDFZ+rD6EXQhBiFwjjwEmGzKoXiUzj5RNed+5AOicrSuwCMzwFoOd0WytJ/5RDH
Q2F7cGIeUh/9rC61+YsdJzg6woUg/1ljtE8fee4BDic4jMGQZvJNNiINDWFe6HTS
STGS6WMlxLw3M+eZMBJwu/LFsvrEjf/6kcbl1OLdEWF52p7KAN3Dw5zJ+pekSqiF
580zFgD2R5cCAwEAAaOBzDCByTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUw
AwEB/zAdBgNVHQ4EFgQUbD/HIwNualdNlQNr+PquYnFHrSAwHwYDVR0jBBgwFoAU
coM/J/hxQ8wUCpTlFdvJmK4cxIcwNwYIKwYBBQUHAQEEKzApMCcGCCsGAQUFBzAC
hhtodHRwOi8vdmF1bHQ6ODIwMC92MS9wa2kvY2EwLQYDVR0fBCYwJDAioCCgHoYc
aHR0cDovL3ZhdWx0OjgyMDAvdjEvcGtpL2NybDANBgkqhkiG9w0BAQsFAAOCAQEA
JTpzXTTqogDAAMERsTE7+sPW1hMHchw/5+wVYaroMGobWsA6are7wug0d+mJJMEy
FpeP/S4RgknnGihq7chLEdSKAj4M0Rr0998hnb8KxROd5hyhbOr6LZ9NUFTrMuT6
8F9T17lVkrnIznHE2HN4cF3kyD/2TfralQDmJTNtJN85xdhBz9rv7XQC4Vrl0jev
qoCGMdDRKfSMT5NqFk37zQHA08FfHqG7ZxrPc6LBStZZhn2pJ6ryvmD1AffuxG2h
Of98IqQy8y+8PuF1bilDyigiGxjVqlbp2otPAFNXnaKVpn9wKEXZh8I0PEKXTIOy
v1ricup/6lWCc65sLdatZg==
-----END CERTIFICATE-----
private_key         -----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAw5noopRuEo54hqcNDniHWplZX7hkYQUycnVspGVF8PkVfmqs
S2oLvZApmjn38ubAMmwbAdJbBHkEwYzfp+doi4bPdfbduVBN7U3m/jRevKHt1hAK
P5TH7biyueqqvOcO12mydDI4z6u1fcQpNx2bLUe4YmQZUQnIJ0rQ+yP6NLnjAkE0
dd56YQRW0LN01IHFNMy5H049PcZHNLHpXMBDD5T71J/8EEglAmtgiyvMncADaMMS
V2kDO7dgvpiHUxyDHI9T9BWEVzx0v9QbAmjRGJke21mj16OX7Qfa78ManTKWkqu1
jWBLLt8E8UmuINiXzKLEG2uqlD6+7MMbht+cswIDAQABAoIBAQCA+qkjI7sdc2LY
b6k1ZJbkUABWFRHjw/iK7xFPqhQfnS/mJ1Mre2b4fepg716nhi8wUIbSgbJQ1bgk
JM+KORaAAsokYD++pcxs8ZSjBUd/HUvmfrczS9k10TfRToN9guMPenwZENDTZ8eK
geInnNcpWYSlJJppDIk2F6iMIiInx0+k70xAC/n6NfmuHHdtsy2reWqWsuzkwL3V
JzORebIPF9KAaDnnbbQE8pCDGiDSHom6aIQZIyiCTTB+1wkChyWtdiW0JFp2CDYs
g96MZvqlAkYnwvyOlPdUgzjXeUF007ZaRnLerLGs9KFZVd3EhKmCiIyn9GpgU27c
AkDti5qJAoGBANBvOcMbozPLE+Y2knDqFZTM56+yUxPTDDHgDZmycmxrJi/FPfvf
P0Y8Tb0kleAu4aZ8Rzsdbd25aKxb3l73nGU5G1CspE4F29xgfAY0mcAgmIXVxTp5
Ocxdsi6U3qm03oVKm7vUb7QQoP2CIYz6+ktEBEJCcy/EpHYcqMpqXyanAoGBAPA8
9iZgwJwqrpKv/7/kFJo84YYKHtsCrqWT1OBwyyecayN4qdZxMUY2b07VY0fm7nP/
X59RrlDnkOAdkBfYIb4MbHb+hOG8DmmCpObQuefsHpnl+Bi5S042w3zjjjCfCAcU
pnt76iksVtwvKnT68QMlynkZ1ikp7lNEEt70IicVAoGBAIhzyETKmxUVJXnY0BXL
qUU4v5RuEaUwzRL5su2jvPTtUJqFbgauKCY65Emm4ddZ2a+8PIWexoYZE0WC/xj8
0Q4TTKaL3ySVAVJMd+U4Gk+FB1TWlb+qDQSVcki8fOhs2CA6r6lPcbFYEYzk3EIH
LESJyp/2EKJ9AY2xRW7+JDyRAoGASsEdczm+FZouclSzx1lIk+oH3za8/Tdrcvh6
UFCo1q2UpRgB+UmFpKyBnE1INHKGlq/LaH75OBGmgiaQ1OqbLVEeWzS90qQ689gE
ShJ/FszhhNALkF3wMelWkUEZ2MVTsSha2afoaGF0sdU3Z+lCH9GdnGtOucZjc6OG
H7xl+DUCgYEAyUoVlRdGgiw9Fyr+Lj9zR3VM1vv/KUEZsPXis/KXwHtZEpiEQTjQ
swGPXfhaXMMakVPMncPncOvRI4UKzG6/w8gZj0BgQnoez2uNdNGijZy+lSAWgsd9
NrFznHF0HYbINi8XDspnjQlOPzKwio0JZTXJM1TNpatibxZpVFkYfgQ=
-----END RSA PRIVATE KEY-----
private_key_type    rsa
serial_number       55:8e:87:50:41:53:d9:d7:cb:a9:56:72:86:c1:5c:46:90:da:3f:4e


Скриншоты с сертификатом nginx во вложении




export TMPDIR=.
export CSR_NAME=vault-certs

cat <<EOF >${TMPDIR}/csr.yaml
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: ${CSR_NAME}
spec:
  groups:
  - system:authenticated
  request: $(cat ${TMPDIR}/tls.csr | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
  - server auth
EOF
