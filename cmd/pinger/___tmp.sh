#!/usr/bin/env bash
while true; do
    curl 'http://sip.zdzit.olsztyn.eu/PublicService.asmx' -H 'Pragma: no-cache' -H 'Origin: http://sip.zdzit.olsztyn.eu' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: pl-PL,pl;q=0.9,en-US;q=0.8,en;q=0.7' -H 'User-Agent: PZI TARAN' -H 'Content-Type: text/xml;charset=UTF-8' -H 'Accept: */*' -H 'Cache-Control: no-cache' -H 'Referer: http://sip.zdzit.olsztyn.eu/' -H 'Cookie: ASP.NET_SessionId=usziyhl5fh3ypxxyf5i0aavn' -H 'Connection: keep-alive' -H 'SOAPAction: http://PublicService/GetStreets' --data-binary $'<?xml version=\'1.0\' encoding=\'utf-8\'?><soap:Envelope xmlns:xsi=\'http://www.w3.org/2001/XMLSchema-instance\' xmlns:xsd=\'http://www.w3.org/2001/XMLSchema\' xmlns:soap=\'http://schemas.xmlsoap.org/soap/envelope/\'><soap:Body><GetStreets xmlns=\'http://PublicService/\'><s>Bxr2XdypulnHonFAVcj1gK5sErCmhsise1dZuNY1cMDw=</s></GetStreets></soap:Body></soap:Envelope>' --compressed
    sleep 600
done
