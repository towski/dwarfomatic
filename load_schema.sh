#!/bin/bash
mysql -u root -pmysql $1 < sql/schema.sql
