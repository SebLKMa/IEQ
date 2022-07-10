export PGPASSWORD='iequser'
psql -h 'localhost' -U 'iequser' -d 'ieqdb' \
     -f setup.sql

