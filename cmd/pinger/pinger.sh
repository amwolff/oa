#!/usr/bin/env bash

while true; do
    curl 'http://sip.zdzit.olsztyn.eu/PublicService.asmx' -H 'User-Agent: oaservice/1.0 (+https://ssdip.bip.gov.pl/artykuly/art-61-konstytucji-rp_75.html)' -H 'Content-Type: text/xml;charset=UTF-8' -H 'Cookie: ASP.NET_SessionId=usziyhl5fh3ypxxyf5i0aavn' -H 'SOAPAction: http://PublicService/GetStreets' --data-binary $'<?xml version=\'1.0\' encoding=\'utf-8\'?><soap:Envelope xmlns:xsi=\'http://www.w3.org/2001/XMLSchema-instance\' xmlns:xsd=\'http://www.w3.org/2001/XMLSchema\' xmlns:soap=\'http://schemas.xmlsoap.org/soap/envelope/\'><soap:Body><GetStreets xmlns=\'http://PublicService/\'><s></s></GetStreets></soap:Body></soap:Envelope>' --compressed >> pinger_history.txt
    sleep 60
done
