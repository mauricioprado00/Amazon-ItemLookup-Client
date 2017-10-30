#!/bin/bash

curl http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemLookup.html | \
	tr -d '\n' | \
	sed 's#\s\+# #g' | \
	hxselect -s '\n' -c '#w70aab9b7c23b7b2 tr td' | \
	awk '{gsub(/<[^>]*>/,""); print }'