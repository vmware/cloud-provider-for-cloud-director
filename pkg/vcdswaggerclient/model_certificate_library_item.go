/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// A certificate library item.
type CertificateLibraryItem struct {
	Id string `json:"id,omitempty"`
	Alias string `json:"alias"`
	// PEM encoded private key. Required if providing a certificate chain
	PrivateKey string `json:"privateKey,omitempty"`
	// passphrase for the private key. Required if the private key is encrypted
	PrivateKeyPassphrase string `json:"privateKeyPassphrase,omitempty"`
	// PEM encoded certificate
	Certificate string `json:"certificate"`
	// Description of the certificate library item
	Description string `json:"description,omitempty"`
}
