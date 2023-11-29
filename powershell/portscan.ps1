$IPADD = "192.168.50.121"
1..8000 | % {echo ((New-Object Net.Sockets.TcpClient).Connect($IPADD, $_)) "TCP port $_ is open"} 2>$null