
c:\mysql\bin\mysqldump.exe --user=root --password=root --databases eems > d:\mydump\sql\eems.sql
"C:\Program Files\WinRAR\Rar.exe" a d:\mydump\eems.rar w d:\mydump\sql eems.sql
del d:\mydump\sql\eems.sql


c:\ozoerp\mysql\bin\mysqldump.exe --user=root --password=Lxz410928 --databases ozoerp > c:\ozoerp\backup\data\mydump\sql\ozoerp.sql
"C:\Program Files\WinRAR\Rar.exe" a c:\ozoerp\backup\data\mydump\ozoerp.rar c:\ozoerp\backup\data\mydump\sql\ozoerp.sql
del c:\ozoerp\backup\data\mydump\sql\*.sql
