/*
 * Copyright (c) 2020 Gilles Chehade <gilles@poolp.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"github.com/poolpOrg/go-opensmtpd/table"
)

func check(token string, tableName string, service table.LookupService, key string) {
	//table.Boolean(token, true)
	//table.Boolean(token, false)
	table.Failure(token)
}

func lookup(token string, tableName string, service table.LookupService, key string) {
	table.Result(token, "foobar")
	//table.Failure(token)
}

func fetch(token string, tableName string, service table.LookupService) {
	table.Result(token, "foobar")
	//table.Failure(token)
}

func main() {
	table.OnCheck(check)
	table.OnLookup(lookup)
	table.OnFetch(fetch)
	table.Dispatch()
}
