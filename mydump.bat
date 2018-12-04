rem c:\ozoerp\mysql\bin\mysqldump.exe --user=root --password=Lxz410928 --databases ozoerp > c:\ozoerp\backup\data\mydump\sql\ozoerp_%1.sql
rem "C:\Program Files\WinRAR\Rar.exe" a c:\ozoerp\backup\data\mydump\ozoerp_%1.rar c:\ozoerp\backup\data\mydump\sql\ozoerp_%1.sql
rem del c:\ozoerp\backup\data\mydump\sql\*.sql


c:\mysql\bin\mysqldump.exe --user=root --password=root --databases eems > d:\mydump\sql\eems_%1.sql
"C:\Program Files\WinRAR\Rar.exe" a d:\mydump\eems_%1.rar w d:\mydump\sql eems_%1.sql
del d:\mydump\sql\*.sql