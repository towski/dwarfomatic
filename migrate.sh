#!/bin/bash
mysql -u root -pmysql dwarfomatic < sql/$1.sql
