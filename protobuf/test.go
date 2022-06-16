package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

var pbStr = "100718012293150aee04189be6f79306280030a8dfe56e3801409be6f793064a19636f6e766572736174696f6e5f73657474696e67735f3732395215636f6e766572736174696f6e5f73657474696e67735803620b0a075052494d415259100062190a11756e69715f6169645f7569645f637369641001100210046a87030a076774696453657412fb0235653163346465632d633737632d313165612d626430302d3030313633653034646136623a312d383231363934383231302c37643536643838332d333831322d313165632d616532322d6661313633653132326537363a312d383032313333303637382c61616535646639372d393835362d313165622d386565352d3030313633653034386130333a312d323235343337353535392c63663262393465642d313264642d313165622d613838632d6661313633656138383032393a312d353332383935313937302c65396137633634332d383133392d313165392d616438342d3030313633653030383639303a312c65643334646632642d383133392d313165392d616438342d3030313633653036613665653a312c66633763323937372d383133392d313165392d616438342d3030313633653036613939643a312d32353937373131313936302c30613539306239382d666133622d313165382d626433622d3030313633653036613963623a312d31313432303438343135386a1f0a0b6c6f6766696c654e616d6512106d7973716c2d62696e2e3030333630336a1a0a0d6c6f6766696c654f666673657412093832373535333332336a120a0b6576656e744c656e67746812033239316a0f0a087a6f6e654e616d6512035554436a0f0a0a7a6f6e654f6666736574120130129f100a97100a3d0800120269641801280030003a03696e74420932363839363830373348016a1d0a096d7973716c547970651210696e742831302920756e7369676e65640a3c080112066170705f69641800280030003a03696e7442043131383048016a1d0a096d7973716c547970651210696e742831312920756e7369676e65640a5208021207757365725f69641800280030003a06626967696e7442133639343433313031373438393938393733353048016a200a096d7973716c547970651213626967696e742832302920756e7369676e65640a6208031206636f6e5f69641800280030003a0776617263686172422b303a313a363934343331303137343839393839373335303a3730383539373037363939343834363030333848006a180a096d7973716c54797065120b76617263686172283830290a570804120c636f6e5f73686f72745f69641800280030003a06626967696e7442133730393730393538383533313737383338323448016a200a096d7973716c547970651213626967696e742832302920756e7369676e65640a4208051208636f6e5f747970651800280030003a0774696e79696e7442013148016a200a096d7973716c54797065121374696e79696e7428342920756e7369676e65640a440806120a696e626f785f747970651800280030003a0774696e79696e7442013048016a200a096d7973716c54797065121374696e79696e7428342920756e7369676e65640a42080712096d696e5f696e6465781800280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e65640a450808120c7365745f746f705f74696d651800280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e65640a4a080912117365745f6661766f726974655f74696d651800280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e65640a3c080a120b707573685f7374617475731800280030003a0774696e79696e7442013148006a170a096d7973716c54797065120a74696e79696e742834290a4f080b120b6d6f646966795f74696d651800280030003a0974696d657374616d704213323032322d30352d31332030353a35363a343348006a160a096d7973716c54797065120974696d657374616d700a41080c12086f70657261746f721800280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e65640a2c080d120565787472611800280130003a0474657874420048006a110a096d7973716c54797065120474657874123f08001202696418012000280030003a03696e74420932363839363830373348016a1d0a096d7973716c547970651210696e742831302920756e7369676e6564123e080112066170705f696418002000280030003a03696e7442043131383048016a1d0a096d7973716c547970651210696e742831312920756e7369676e6564125408021207757365725f696418002000280030003a06626967696e7442133639343433313031373438393938393733353048016a200a096d7973716c547970651213626967696e742832302920756e7369676e6564126408031206636f6e5f696418002000280030003a0776617263686172422b303a313a363934343331303137343839393839373335303a3730383539373037363939343834363030333848006a180a096d7973716c54797065120b766172636861722838302912590804120c636f6e5f73686f72745f696418002000280030003a06626967696e7442133730393730393538383533313737383338323448016a200a096d7973716c547970651213626967696e742832302920756e7369676e6564124408051208636f6e5f7479706518002000280030003a0774696e79696e7442013148016a200a096d7973716c54797065121374696e79696e7428342920756e7369676e656412460806120a696e626f785f7479706518002000280030003a0774696e79696e7442013048016a200a096d7973716c54797065121374696e79696e7428342920756e7369676e65641244080712096d696e5f696e64657818002000280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e656412470808120c7365745f746f705f74696d6518002000280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e6564124c080912117365745f6661766f726974655f74696d6518002000280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e6564123e080a120b707573685f73746174757318002000280030003a0774696e79696e7442013148006a170a096d7973716c54797065120a74696e79696e742834291251080b120b6d6f646966795f74696d6518002000280030003a0974696d657374616d704213323032322d30352d31332030353a35363a343348006a160a096d7973716c54797065120974696d657374616d701243080c12086f70657261746f7218002000280030003a06626967696e7442013048016a200a096d7973716c547970651213626967696e742832302920756e7369676e6564124b080d1205657874726118002001280130003a0474657874421d7b5c22733a73796e635f73657474696e675c223a5c22747275655c227d48006a110a096d7973716c54797065120474657874100320d202"



func main(){

	pbBytes, err := hex.DecodeString(pbStr)
	if err != nil {
		fmt.Println("hex.DecodeString err:", err)
		return
	}


	var message Message
	err = proto.Unmarshal(pbBytes, &message)
	if err != nil {
		fmt.Println("proto.Unmarshal err:", err)
		return
	}

	var entry Entry
	err = proto.Unmarshal(message.GetPayload(), &entry)
	if err != nil {
		fmt.Println("proto.Unmarshal Entry err:", err)
		return
	}

	entryStr, _ := json.Marshal(entry)
	fmt.Printf("message: %s", entryStr)


}