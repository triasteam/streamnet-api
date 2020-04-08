### Streamnet-Api

#### API Description:

##### GetRank 接口

Address: http://106.3.140.187/api/QueryNodes
RequestType: POST
Content-Type: application/json;charset=UTF-8
RequestData:

    {
      "period": 5,
      "numRank": 10
    }

Description

|Param|Desc|
|----|----|
|period|大于0整数|
|numRank|大于0整数|



ResponseData:

    {
      "code": 0,
      "timestamp": 1585637566,
      "message": "Query node data successfully",
      "data": {
        "dataScore": [
          {
            "attestee": "10.0.1.390",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.391",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.392",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.393",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.394",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.395",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.396",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.397",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.398",
            "score": 0.00917431
          },
          {
            "attestee": "10.0.1.399",
            "score": 0.00917431
          }
        ],
        "dataCtx": [
          {
            "attester": "10.0.1.398",
            "attestee": "10.0.1.399",
            "score": 0,
            "time": "2213223190"
          },
          {
            "attester": "10.0.1.393",
            "attestee": "10.0.1.394",
            "score": 0,
            "time": "2213223190"
          },
          {
            "attester": "10.0.1.394",
            "attestee": "10.0.1.395",
            "score": 0,
            "time": "2213223190"
          },
          {
            "attester": "10.0.1.390",
            "attestee": "10.0.1.391",
            "score": 0,
            "time": "2213223190"
          },
          {
            "attester": "10.0.1.392",
            "attestee": "10.0.1.393",
            "score": 0,
            "time": "2213223190"
          },
          {
            "attester": "10.0.1.395",
            "attestee": "10.0.1.396",
            "score": 0,
            "time": "2213223190"
          },
          {
            "attester": "10.0.1.396",
            "attestee": "10.0.1.397",
            "score": 0,
            "time": "2213223190"
          }
        ]
      }
    }


|Param|Desc|
|----|----|
|code|0:success|
|message|success message or error infomation|
|timestamp|时间戳|
|data|data by json string|
|dataScore|分数排名|
|dataCtx|节点关系attester指向attestee|


##### DagMap 接口

Address: http://106.3.140.187/api/QueryNodeDetail
RequestType: POST
Content-Type: application/json;charset=UTF-8
RequestData:

    {"command":"getDAG","type": "JSON"}

ResponseData:

    {
      "code": 1,
      "timestamp": 0,
      "message": "Success!",
      "data": "{"dag":"{\"NLAOTOTNKEVV9ZTOTNFM9QYFLKLJDUIPJH9GFFMESRZYBUALFTNXKULJEBSTCODUVGHTOIGMZWNFCQ999\":[\"M9UKIXVLTGNCZLCGDIKXWCOHQTEXKTWOKOIXOMOZE9LPSBNMHIOPRGAGYHZWMPRDHQVSAEGDEFDQSY999\"],\"SSBADPECBMQEDKJRYDKWVYWTFLTWPMDBGY9YQGZYLYKXOFXKYUM9ILVLGDCLLVOHOANJPIWDPJSLEK999\":[\"GJIYCH9WAYCLLFAMROOVQYOKKOJDPAWCVCUTK9SVFYIVRG9YCNIHMHSZVDMDKRTBZHAYCGXWSIONRU999\"],\"DNJMWLNUXJSJIAPEAAAESNLSVXME9CDYFWZKIBRSCUPHJSPHKXIMRUEXSELMCH9IWSISJVIEDAQTSELL9\":[\"SEKMEQFYWHEWQEVUWHBHTLVOPGGRUY9IMOXKFTVWDOMGAOPAXNGVFCCSAAQBPFQCGTICSYSTT9NUERNKB\",\"LHWTYZPOCVWCGQEUAYILVWONMFXTSSFMBNKFTBKPVTZVAISGNFODOONKOMEHAAUL9RFXEMMMN9JCXFHQZ\"],\"SEKMEQFYWHEWQEVUWHBHTLVOPGGRUY9IMOXKFTVWDOMGAOPAXNGVFCCSAAQBPFQCGTICSYSTT9NUERNKB\":[\"EPTKFKSVCERRNIFWAOVKBXWBPOT99AJJZMZCTJTEJNJLXHMYZBRQ9N9GDLYJMLERBCFFORSLZZVEZLPLZ\"],\"VLWJAILOZDUCWJAVKK99NHKZMSFDJVKQNYPJLEUBAKEOOCDVXUDTZOUXRPWUVGTMTVHHAUEEYYTVKL999\":[\"U9OFIGLFNOZYUFXMJVSJCKNUUIESJJUSOULSANUIRUMBOGNFHUZNCUANYOMLCPHXCTKCFHKJPOTEWLQBC\",\"ATVRTXWUF9JY9BPMGNCKPQIAIDPFRNM9PFJVXONHBQOOQYCVCGGVWEPEZARERPOBCUWCNHLIIWJHKKVRJ\"],\"ZVSHASQWJFIUUYBREQTMUWHKW9RMBDQSVACEXDQFWUHGSHAREPNUWJDSEGEQGRESSZTJ9WVWGBPUGC999\":[\"SSBADPECBMQEDKJRYDKWVYWTFLTWPMDBGY9YQGZYLYKXOFXKYUM9ILVLGDCLLVOHOANJPIWDPJSLEK999\"],\"LHWTYZPOCVWCGQEUAYILVWONMFXTSSFMBNKFTBKPVTZVAISGNFODOONKOMEHAAUL9RFXEMMMN9JCXFHQZ\":[\"ZVSHASQWJFIUUYBREQTMUWHKW9RMBDQSVACEXDQFWUHGSHAREPNUWJDSEGEQGRESSZTJ9WVWGBPUGC999\"],\"GJIYCH9WAYCLLFAMROOVQYOKKOJDPAWCVCUTK9SVFYIVRG9YCNIHMHSZVDMDKRTBZHAYCGXWSIONRU999\":[\"NLAOTOTNKEVV9ZTOTNFM9QYFLKLJDUIPJH9GFFMESRZYBUALFTNXKULJEBSTCODUVGHTOIGMZWNFCQ999\"],\"M9UKIXVLTGNCZLCGDIKXWCOHQTEXKTWOKOIXOMOZE9LPSBNMHIOPRGAGYHZWMPRDHQVSAEGDEFDQSY999\":[\"VLWJAILOZDUCWJAVKK99NHKZMSFDJVKQNYPJLEUBAKEOOCDVXUDTZOUXRPWUVGTMTVHHAUEEYYTVKL999\"],\"DAYALBLFUHONLUTPOXPPCJHICZMNRWMMZUDMWX99VGNDXVWRSIUSOBHDBXSZIP9XWILJCESMPGAISPQKB\":[\"SEKMEQFYWHEWQEVUWHBHTLVOPGGRUY9IMOXKFTVWDOMGAOPAXNGVFCCSAAQBPFQCGTICSYSTT9NUERNKB\"],\"EPTKFKSVCERRNIFWAOVKBXWBPOT99AJJZMZCTJTEJNJLXHMYZBRQ9N9GDLYJMLERBCFFORSLZZVEZLPLZ\":[\"ZVSHASQWJFIUUYBREQTMUWHKW9RMBDQSVACEXDQFWUHGSHAREPNUWJDSEGEQGRESSZTJ9WVWGBPUGC999\"]}"}"
    }

解析参考如下代码：

    let resdata = JSON.parse(res.data["data"]);
    data = JSON.parse(resdata["dag"]);

|Param|Desc|
|----|----|
|code|0:success|
|message|success message or error infomation|
|timestamp|时间戳|
|data|data by json string|
|dag|dag关系参考key指向对应数组内的值，例如:DNJMWLNUX……指向SEKMEQFY……和LHWTYZPOC……|

##### TotalOrder接口


Address: http://106.3.140.187/streamnet-api/getTotalOrder
RequestType: POST
Content-Type: application/json;charset=UTF-8
RequestData:

    {"requestData":"","requestMethod":"GET"}


ResponseData:

    {
      "code": 1,
      "message": "success",
      "data": {
        "totalOrder": "[\"DKOXPEFSAXT9NRPRA9ERYBXAYHBJVFZWPFRGOONRVDDKCCMEPTZOHKEJ9AMRIW9KTUOIRGQJTTRDPF999\",\"VWQGRYXLVSETTXBLWYKHTSYHJKQZUDZBKBEKEQHUGVXKHXGKKYEOAGVRAAEHHHIHLB9QZZXZLLDOCZ999\",\"DEPCGDGOYIUUQJMYGYQHRYQLR9EIFNCVHKVISOIXTMJYMABQNROJCQIWXEZGXKOIPZIRGCKDRBTTSD999\",\"MUOPLQQTFHYMQHVUFGGAAYJ9FTCZZXMN99IHVOOFUJICSEQUNPQWLHLASMOOMIOORR9ST9ROWWBMOD999\",\"QCHPSWXIWIGBGBROWMJKEPNIYUJJSWPRJWNRUYQTAEPCTESSBUZTQKUZWHC9PWVGADRZULQREPHZYK999\",\"PYKEAZRWOZGXXNZUCSWLECCRAXGQMECOTORYNWOODY9WBTVWAVWPJBIAFJSNTME9KICUJCVQBVBZGK999\",\"MEWYQPCCKRUBZVGUQJERDKDZ9DELCAQAILBWEPLBKNPVFONSKHRKR99CGLDJYQBZIEKFGQUEWPECWB999\",\"LYPBGHZMFYERRLXHJILVBVAOFYWWLUGTJHEJSAFFCYLKBHZWWAEDUQHXACSEMFSJWASKJZBJPIVRAC999\",\"KCQJXMQSAMHRKIIAYKGAQOGYUJUPPH9WUNSWNOOMUKCXTBVIMDJXNQZLUZUGOPMVUTKSTMHFFFQXJB999\",\"VMRXHWGOQDZNFX9NGZQQIJNWZPKQ9BYCCCCUJLXJCFHRNITVWUSQVWNUZFEWGQ9PRDIGLINODMEKJU999\",\"TQVNDVRVUOVOFHIVFQWGBESRQSBDDMWUVLKNYVMTJ9QPNUHUREDMJYLDJCJQRA9EGACTKLEDYEGTOB999\",\"FRZUT99KMIJAMUYNOWWGWWFXOHPFAVFNBEKEAJAIMPTDGJWCBVEGYALTWIZDFJBFPACSELDERVCFVX999\",\"VNJCYHVOQHIRQEAKVTHI9FOFHSNMSDPDH9HGHBTYCRV9SRETVVHJMKTXLFRMZDJDVNHJOOMMIJSPYR999\",\"XVSRDEBJXEVLOTFBLDFEVOBGPTYKHTSSKBNIMBMPTUNHZWYDSAE9ZRTEMGAP9NZVCWAEL9TEEUDYFV999\",\"KVXW9TLWRZHLEIMIAF9PGYQ9XAJEJOVDOJFSFEPLECYGZQLKKGBLETUYMTVR9DNMAAV9EYHKIATGT9999\",\"VZDEB9YEURGTNNRWKSF9VEWZBAJEAKRFUUZHQFCWXMNAGVGXACAAJMYDSSXEGEI9LWWCHBCNAFFGIX999\"]"
      }
    }

|Param|Desc|
|----|----|
|code|0:success|
|message|success message or error infomation|
|timestamp|时间戳|
|data|data by json string|
|totalOrder|节点数组，已经排列好，前一个指向后一个即可|

##### getBlockContent接口

Address: http://106.3.140.187/streamnet-api/getBlockContent
RequestType: POST
Content-Type: application/json;charset=UTF-8
RequestData:

    {"requestData":"{\"hashes\":[\"VWQGRYXLVSETTXBLWYKHTSYHJKQZUDZBKBEKEQHUGVXKHXGKKYEOAGVRAAEHHHIHLB9QZZXZLLDOCZ999\"],\"command\":\"getBlockContent\"}","requestMethod":"GET"}

Description

|Param|Desc|
|----|----|
|hashes|Dagmap和totalOrder返回的hash值|

ResponseData:

    {
      "code": 1,
      "message": "success",
      "data": "{\"trytes\":[\"{\\\"tee\\\":{\\\"tee_num\\\":1,\\\"tee_content\\\":[{\\\"attester\\\":\\\"192.168.130.12\\\",\\\"attestee\\\":\\\"192.168.130.13\\\",\\\"score\\\":12,\\\"nonce\\\":0}]},\\\"score\\\":574.0,\\\"pScore\\\":574.0}\"],\"duration\":1}"
    }

|Param|Desc|
|----|----|
|code|0:success|
|message|success message or error infomation|
|timestamp|时间戳|
|data|节点内容，展示的时候拿trytes字段内容即可|