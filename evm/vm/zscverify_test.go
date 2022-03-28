package vm

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// Benchmarks the sample inputs from the ECRECOVER precompile.
func BenchmarkZscverifyBurn(bench *testing.B) {
	inputHex := "000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000009000e0cafd8d17c3dd011488f3151f605517dca53cd5537942b1551c8bcec151411000000000000000000000000000000000000000000000000000000000000088001d39aef1308fae84642befcdb6c07f655cc4d092f6a66f464cb9c959bff743a277420423ebed18174bd2730d4387b06c10958e564af6444333ac5b30767c59c10bb645f46b504112dffcaff3a218ccc92e7499573ca00814cbe42fbb602c58c06de103558b65a2495b7e8e5df4e6642c574c3d14f956586bd66b40f7b4363152deb65cdd8997866abd7f4d442ddc2857aad1e32ca7a10f9be6897bc6db7f42810045f9c1b42b3ec39ceed568215c244830eda34667a15b94063d4e05a25e9040c04d4cf5350dc9384440f73c4bce7adb7ee966ddbe23902eabf8231378a93d21ec94153797611b0a17050936cc48db816f9f306528b5c98e1e03261db051ae92ef67ea2c0c83907fe3d7d3c2be6d20dc743c1b433195df6d02ae9935c43f84707106973b5ff800ea76a56f545b5d2086640f2918bd7b92f6c92c8af2fbe78c12cb0547e3980b8123a831402f879e248ffd86693ccbaf5c9a4f30957431819961461f4b01525da485892b40407c2ea37c51b3b8e13674323415b948777d93d670bc1504a244fe8b620027df0e9d4488f34ff1a9a7c6c4b9fef402ebf7104765912bf5990b9714dd22fa862a714207b67c4d9a8a44c69579caf6790fbf449448906e2ebd49f54888c20337ad7b9d633fbd3aa740c7d62e70cc73200946ea650940212670d13041e854d1bf32849b80d735edc464414aca7d7ea6bcc346fd9ebe3303f283464f6a19f61b763fc6acb4d4ba08f1979fbf805e245debad3357c5ac32a32109715c54025a02664a5191d698c24f166364cc388fc79c3914faac62699030e4cca2df4154c8cd2fd7316d9a017ecdc409335746536beaadd1350dda8ff25d1ecfcf7d5c4d656e08f33a9899143cb7d71d4ebc4897dc971e8cc68925e211cd8551d68ff56f4eb69c3cd0c9ca893e97408ba9f561a73a15f06070258b15d09e1c24ae2d8e76f1d758641f8728b337290b4ca1c86ca7c4f5870473b75e41d2eaaaabae5aa6b24ca533c052173e1c06511536454e8187ba6385444f4ca684703aef40408785f63ec8896bdf608a4fd452b94f8f7d3688bc3c11f135b37c4a4065ec32c8a329a5637a07cb60e8236c396089dbdefd83a8ffd94a535764db15417fc1b254514117a38059f9456ed80fcea1ca91bb3bdea27860092266eee975508d030c0d7e1709121a4da341aacc4f44e7173689a91a90a9108379d31e8860126cee2fa85178e43403a4c40ea3e0da87d9c4ccba79358ee67b4359019dfa800237ca48ea211a9cd00ebbaf24b4c1c230936f82c596ba13b333b3e1c2e319b09036821b59a67de9ef2b555a2f390639eccedb21715aa102869a31edc12c1c0bb1917942bc574d1bf066fa787d3912cb9b0349262e3036483904f7b050808a9572b3e0fbc885a5207ea3f613395cef0708a91456318f76f074098b0231faba38209780946da36a0c2453d527b16784c3d7fb5699a8a9bc2db398ba77bd627188a0406dc46e0c041eace80ecac28c42ebc27b8f06c91f6663ba08341125310b5f5269397d20ba7bfca894fe24aeba0d83e05160859b1cc0fa6647610b08333390f14c5e3eff6b766b1d38efba711074a08940eda5b6963127218164cbeea06840a2b630e29f5e6cbb13ef0e1e29bb6d1693468b8abf52a21d6ce4b596162a16acb2a48565af4022bf953230b25ed42b47e78d70a67587021cfce04a7093b62eb1d2a9dcce09bf76010365566aff41b967986e1a15fca3a86be3a81d38b28255c1a1c270ca5f3225342f19062c25367d94f101692ddebcf4aabd422bcade27276071f279c7d657496c9e063058dec9294cb7b2226bc8c7eafa513bd802cb6acce920bf999760d8a3d508cece2249abb7b9b2eb247f29fedf02a20856da8d64853a0140b7dd1b77bec32599401089dbc5d29c0771b494ef3acc1afa1629288eafdb92aca9c7f1e1ec6d03064ddf21712e0d8bde40d3142efa4890f861813f364b05a2214c82d1f03104d23f4156587ae23cc67d7fc8a5e5f8dd79e3c69cfca5c17342fe81d3940c5c1f9ae7284653742b900384841be6d21b93d02dac7f42d1f955125f7533e0cef8d981d81cbbce5ab8545314929e1602038a00cf3efeb3254be2600033a54c4f3f644a694c68076311ca3a3b65a1ad9f05cbbc6bce085e9e452b81d6391fe588728f506fcd2edf0955242cc4d59e399385f46dfe940da2615765b1196b5e054c3188bc844ec674c1ee7295cfbd0bb497764be9e9beb9fe1f82e642be9096789ecbb415d5e6b7e6462f1e41c15cac37904226c7227b3dcbcda2420228596a794cb485a0813b46df17a9ee26823c4db4a6185a5ea7a6003129257a016d528ba8a06cc0ea02c1e11254cfc113bc4a6f3e7e520a20ea9fdef0d2466cd0f599c9f8216a5e763509e9c671d6c5a59e2d17e82cfb4c81082f839140aae741d2d00161945e96c84874b1e1656938ccb49c671da319a4955291a69b7c1f8b816ba911e1b6ad20d7459d9832e549b02626259910aafb21412ae2ca0f30361c927863ded5d38c15aa75341dd6fcab9d9f34104adcf400d2ddb8df4b80a5e6d1f0508ba2b60a92a49ca086961c736c2182792bdcec91d08f349e4faa2e4dd4336127a04071c95429ca776b01a0517c38c085b97ee764e1367592eb20d206bd09c13dc065e879692698fc2bfdd29fd680a455b6f0cb03c332e7512b60750c7ab3f13e916d1a0b085d1c023dc8c2ba8ec62e3bdde2376f2338a80390685279d684d2374fc49110497115070bd7df00a88e0e8dc6cd3245c37d18aeebace8e0290a41ee677d889a7943b772704de77403b99ad2b58ae25e9d1121b2d1c6727f590c911c75603c79bff93af453da3cb3cecaa9428bf3ba3f8ba9cf7e5f94d60870a8f24655095fa6a5ea61b7205773dbadcdbf994adc240d9389029ff25ca89eb1c8d0c5fb1b92c6fc284fb03ea69aaa79d7decac973982a42973ba0efbf6bd3b38e21d7cb8e89245c626c1603812a768f6585d73d9818fd5a21defc7c8c0ea12748d0a93b6aac47294d27e9da0509204e2705a8e6e40a173d1e1c120a04faa338bb700000000000000000000000000000000000000000000000000000000000002c02bac979f0aa2b4a288b5ea11ddcc31cd426ee39113926c2bbe3c519b4064ae031351e6c8bd716c65ef84bfc04952fa19c94d950931aa614b644e95f5939f35491f2fcbfaae6e3ea22cedebf4c50e2a3514a8df3fcca7439f40fd59415a91de1d0aa472b1e6ac289ba7d534e56dd6337f332992c92861fbcc1382afb6f98e603e1e274d2797c99113bab1c2e81cbc0f14178c6337392e887101b749fabe6b41f323070b2696d1d4dab969a39ad65975f986868ab5a41e38cef6cc889f60de0e2a12da947cbd81d886d00819b794853fdff0ac3abb39aec24e4421c7b3872482171fa22207dbfbe8b0cc510ce42c061f4698b06de734cae13ee68016a5eec730031fe5d3d667cc5c80862e24f8ef64fbe5e57a0ff9010e24af3c44973bb47c7fe30a7c9011c29a8c26c67045aee08b6e0feb26e01ac8c4eaf73734b0838d2b285a059b844a7012235cac64250a28fc73d82c3920fe495e453660e2ea9f0b43e67218cc1ac20613c3e608015fc358ba411a726a04107412466e39c653c8bbe57a6f1e377425bd821a405ec34b0c83b1c3c551adad8ad3275fea8a0a1324b903b64827738110b6e165cbe87d8c94b77f3558b098de2a6878c7c43c586ac6c2658078125266cf659dbbd67fb8f3d233545b9f3e602c512c6e62edd0e40632180fe8ac047fd8b690663f1e148cabbb10d76bcf282411fe494002bdb6a6c597a66ec9bc0f6b9740d173b4d773d64d7c4da3b592fa90a221c8aaf265a518d634255ab6961be4df91b15b763b39c65cdc86477c7c00a08b173586b74bd4969106cfc25b092d9efbc28c9eaec9af3ac6ad06b45241df896827b1d28fb3bfb28322adb2573d1d06b19e444234a84339dc2b4ed9c7a15854306eb26e7a2cf224eee3e7f195a70cb3d04193e2de981cedb233cf8c117af7144edaa34d0cceb1dc5b7b31014c602c2f50f44d420330a78590b8d9e267d828e5897a826e9af595e4d5b4ee4d8e91"
	input, _ := hex.DecodeString(inputHex)
	benchmarkZscverify(input, bench)
}

func BenchmarkZscverifyTransfer(bench *testing.B) {
	inputHex := "000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000011002c3b1a81a35e6c01d70bcbef5105e42f8a94ac8cdafa6505e9dc3bb295b0741c000000000000000000000000000000000000000000000000000000000000108001d39aef1308fae84642befcdb6c07f655cc4d092f6a66f464cb9c959bff743a277420423ebed18174bd2730d4387b06c10958e564af6444333ac5b30767c59c2e831fa947834bd5b3890ab703b3b361272a76d748df9c46b7e61a0424b67dfd2f8629fbb94ae5cdc2dfd06268e18122389f479f8cdd172e44fc6b44f5df9a45051a248f86c0e622205ec065e4d6a014cfa5c17005dbd47cb2dfa00cdf0deb37107c45538a2e4cd0c24933565dcdb6df936c979a19783b32a3e60820f6e0e5752b311dbc1031a74f3e3106978e2259ec4105b0da3d3df67c9339becb990ed9ff148b6de377473544a54dd78fabb23c68c3e688bc5cb6272494c8ab41c3a6ea2620fb252785e0d00a6ba83b21908a793b55a87c160cfcd006edaf10916cbf96e62c6147d8f5b2bdff29e6bdc09186828530aaa54c1fd4eb14eef71fb0c4602e312fb4f894448dbfb7f56f0c1946c6de2540814e411ddd66fa0f8cb6a571838c811d6f07362738a3383eaece74f32c0297f6bb01e2a26ad4cb3ac3c2d9ca2bd25c191484ff7cef3671bb474603ea6a9eea4976f938637dde202902024333c79ba9099fac376bf324c8583b9b09a83e455ea07e987004e995bebc36e3a3db98d689267d4ae17f55d0ba0e6e4b18ee13d61c4c92ca913d3b72e5f5110a35b17bbf3f1408b22f536bb6ffc3e01dc268788d1145cd60c6628ed6f9d19b0ee64a96cfc61c1bc21753ef79ada75495cd44307f6fc2cc631f232fe3a90bf01ec6852702a20adafea7a9e760e7786e70ec917cff497aba9f5d7301ff7dfdcfe3d895e4162f1595aeaceda668d6b22459e86dc450a479970c4355813ae80628f125f368ad1a08576fc2a7e32c3aad18de93d0f32388be74441edd4106a89e5d5d9a7723e1c9269f09fb03b1576948d4c7bd5d0c024bca01b29fb63cd04063b26fbf6128c7681f3e8a0ac987e9291339042a761adea79f5722eab7c6cbc47238227d01800a8e2fb15b910dbb38f550f4385dd2cb573d65230dc13f64228844339a1c86f6f5522edb9022f7e9d30b271b5b6b45759556448771e207115fc9a9e85da22b9848530b18375bc840e4d263fd017fa2bdeee847cdd8760a7b566bbd410f4a373e084a21b4c1786a2a5c417316974c99885d70000b1a8953679c8431ae8db27d8c724e2121602555ce5bd4ee44836e0c5b72b9b58b0f04b66d5ed5590ec540c58ab7152955d58ffab99932e1d0250792b2cce49ee53f0a109308b1c81748b3880c3477091301416fd9f700210eea6cabede4764e5406a887ab9e7ac64d697f773f092b025119c38ab8b4928a31e64eeb58ba0f77300a96767a5535160d5599e53a51aa0c0c0a7de3e417cf97b64f0179ba6da19be314bc6aca43c75a69aa8fe25a4c4b2b46ad27525d33ac65e77aa256cd97952c8401c2cc0d002d5c8bf549099bfe5f05a39877372cfdc820384d02c9e64055e95150cb94329063c8c89ae8aaefbac20e3a1d49b48d5218921109331644c9296be01a23daabdaf40507bc387289c09b200e5fdd686290d80d8c94d8e7264b4c9fd5fdcab91ddf3ffe2bb32d605d80831428e77db2907d4c12be5d65a114f066a539d04d8842a8d001ae291954db47a62127e7f6c3bf3decb2d25cd7155bf9ba49201c97f0b2982b8d721991e6c9c0eb14ae177a496f6ba4191b151c185e64c7ce9c9be1fafc25d58289aa33902339d40345ab5954404e3c68bb5829ece57aa97918f94dfdb24ac62752d6e22e150ab32168b1630052b3d37b1138654cdee1f1fc4c52740e30dbea5b4e6ec4af7e3d96072d1cb78f5ce52d42fee31a2a3bd97112862d4ade51f43a844284d3a41e6cb41fadb448560ace58c8417e85ec7af7daf3e4b514d26c9ce7ebb4573eb6684bfc16b6be2a5d4f31f482b33c0db93ce3d6cc54d6c5f08048aa53b6d6cdfb6b2af3216124920d2530734a0a45fe76049c414de2ef1d383cac5e5b0d057eed58303b2ed0a6cc40fd13c84298fc4757c3593a5a5b52c5b21e162d80b425afd84cb6f001be3f9782b5371bc91daa03931489f5e490b2f4ee4793f12b964cfb6819ac1d051b9d61aa7124b5f270b0d6bdcfea2e4401619e2a17eb54b316d41f82ac275e0cd74d00590825e62d420038caa540caf42fa0f3ac68990ed4b96e7bad53136b0ef378ac6a790ed65fd8a6350278745b2bd6675607eb49f4874dc2236496ac7f0bfa844f57d17ac7b485782149ebb2593986849a8d1617acc941d8ee0e565742177ecf66033287f55401c820bb9ab25700f64a88734b4edba044fc6f275eec912cfc5f242ef49e7454f52797897cc4baaad13dbda5c9d56378f63dd8d83d89db13affa66495da950801084178e054b49f9d69fb7e3668067f3de7f454267f3132bd3f5b9aa64bac2dcddd53152f5308f563c8f4d028dfc209932c80477befd350e97778be4b5222d3235ed83737424e6b8f0c8642549dddd3799cdb7c3e506c20701da0da4bd8bed6970cf30963db7aa6291f8c10dd2e554783214b877d162522b7ee912bc5dbb52e5b66a58fac9bfcee9519db749cf0510f4438a627b7a9de01171d4cfbda49e522e895ceab7e4ad6c8b0285869bdabc1ba1792c33205b84c31d42304f11f54ace47de929feba1aeb9e557516ebb23b13c6d982ce6227db9a400a0dab1e2e438fdeba25a7b6bc81acf61c1514fafda2584d1f1783a23ef3e482d7fc8622fa9e0cf882fb329765b9b6ea9326f51a4fa1bd1d650a44c11a31a7001ba0252eef501484df272eab311890e894123530cd23fd4495b62ddfb5051f30b77cb7b1bf3f8738ab2ee2362ea3cbb5bbbade8c6827e19b126e76e6a555b13245fb10a97a243d0c341ea14c37f161a8f24f2b7a64bf82dea9bdf078b8beb381d83aa93844f038127915740dbf875689ce324b7c21b4514cfab638234274563300e0e287954f945faf3c3011b489e312162a60ceb4df09ed9c9d639bac227010e8c7fae44023d41f5da2a22bd4ea2144916a3bcd384af47ddb622c2af8c3fbe2fda6a3a2335ee963922de668ece00bf55436f7e423f223651a2f00f0f37257d202b1d56d89f7cd81e7e727a64424f4a1436e9e748c444b992a09df77df2454f0f64d428d87acd57e6c1b1b1bbf7716f3e1f1a899aae345d7babc09f6fbef25a060e7770b65ed5372e675e1a5174ff3f7b14a3f5e7c2b22ecb6bb9445d75193822f22456b18046bd7a33cb3f3c3be2518c030de7ff099dfa7c04932ed65de75402f292211c91e0dae046592fa2f639e276fec34823957117c97c175fc66795892bd99fd395c1f027c85816452f959932c02c6337768cc48da3b54071ea5dfb2f11c5125a96a41c26ccb24f0c46ce625bc2cba223b4a6709722022cccaa7b7f7e24acd8214ff638303e89ba8b5941ac4124b5c55720954adfe09cfe2e70e8b54207d99d2136b2e84a981d7b0347a0f30c4bd362262d6b3e0f8dc5bec6652373e20f2264f43eb77fbf37baa5969b2900a88c574d9b23e26f7e633c6d04d1c5f82c2c7157802c150c62a75668a6f519294e9fa6976a2efbe5c57654c23e2eefc3c72f78dbabe7087bb2f3826541b52f595ac7f7cb5cbad8f492e793076c1a61c64a17a0a13a47fdff04abcd46c99f33cbb873f46dd698c800b683de425e3e0b67d61061da7e67b81f068a978c2bcfb70b8c7c05b39e5e5900171dd54ad593a97f8419d436c2fd7f84b081d7b55ebc836200ef95aba4f4db25471a0217d8e4676a9300a4bbdd4983b1b9fe11e249606fe333834d24fdf64dbdee1fa5c8d59cd197871d5beccca023b3085d3aa38cf0a8eca6a3c13c0f3acf0f3823c5944d74230fa7294fe070b1104e6babcc898668ea261b910feca9194766b7802ab70d8b0b597b0df15c68fe5ad084c4ff8d05e9ad939290a2dea34cfefec7fabedb18979a4fca22667e7ecec862a34d6b3d847bc7092e28c95dda0f5b1787d02054ce73322ca722b02fd5887561be60776446d60e1b83988500a99a1143783f6cec1c5e27a33805b58e59afadc060aab5f7106e3d3958436df1d411c1309601aa073c47e646b81aa01a38228d538e4ee799950452c5d063276a1cd75b1ac3d5bd2eaeb5dc77051a0adff229f75bb8d29e6af6822a7621393b7f0ac127b1eb33aa4a71af244dd1233e66753377238a2b2658c86afd71bd2e005a1287ce031b2d23c0d3bbcc199d2254491d498c0e26c1dadff6286f9746b177ddf121579dd01a7edb9078055b6028aff50ca3a6af7d8e2c46bbc01042f3e26fa4cb69b1802b5a93257891fbf31e04f90d2e518033ee9ce42ad770b1fd7e3b0181174bffd8f1eab768c972ee9b1105a43966d4434c4dc2c97545b64d10c9df5d611ab9ba5bb058df97e0c95e46341417ff7157e0418e09b4851933b61c712a81e33bf6f61bfbde5c459d5b949bed2db54098f1f7ad9f1de215c750990728ba555366de3a838e18da75499065114d0c0e6858b4492ad9e7951b6bf2c0c232c946532827aeaf31e5f190bb476a91e3220e9acabde85b9db20fc213023dcee243446a2e77defe482b08c5b435e9fc3d1e5fe9c9ec7371f57ba4526d79fc66b9d232318aa1ac67215527c16cfaca6aa82a4e33a5d971808d78efaca098f80ca5c328fe20172d989f039279f1567b48eb19b9ef8d8d55866e43b825df9386aeb25e86fc2036952c044be3ea3e4aeb49ed16a1ee30159f5abfeed34573e4b804d1e436c5cf588071ea33dce093317bf9a62b3625458355f58d4152f8240a94e8385d80c38a9550f3cafaafa7ab106e59631fa82450ebd6a2d6e2e7c49c61576310a94f3809c9aaaec22581121091eba6160e80fff9186f47e828dca100473202bc699f2526fbc273bbdd4db5b1c3b27d7c2c63b44e366667137f114ab4249cc22bc67e4b4e29b0072d1456cfcb583186f41ac8b74635746ec697dc3284c1c6baa0c02a6ce79fee0a73b111e546c95a896e13a84eed278843b0cea00a3f09e184820b68a9d0ecb61c7c05e56ed1fef4e523173c553f2359b60f99d26cf23dc5134b52fe46f81efb9d621aa69b5820cf70d0163acd233b7a24b26de663d928a5dec2308d34ab38d41d2beee2aff90e6c1ac819e527e1d4b72808e43d0b55dd8e8cd551e9863603f8356897a5bc11d5a5d40a2e0ab87295b59a2ddd8b03ed12b13a9f10081c0fd3e569656d12db772484b4ac048e44242d4d47700c917af62a125308d2ee6a64097f7034e16f26a8541487d72867a5ec382e179c507f8905f1b7b8179219ac07e79ad8346ef462f449c2dca3068a0968388a02081430e51e22790dffa67dc4854c31e522bf4eaf0121822c742d25f32506a64c0eed85a0571f788a5f0f58415d032f4569530c5131d878d4e11348e459c0a770649e86a0dce48c246c0d1719a5f06dfb2dec467631d808140718cdc715cae2106337b2d981e235106391bbc3734c3f5de30611839f72b73f6e145baef4b8dcd4fec50b42c72e10d8ea0188742ca163a7a93afe25469a978ee61d9e1d16e8a6853075b0f1479376dafd1b04b59651cb04ac38d8632fc78c82b80ec91cb20c7184c8348f0066b5859cbc9d783d10b3b41873024e99cd84710af51dda2c9c59355d862ad756622dceb95efb75ac82e5b031292c60ac8ee2e8b3e901b629e981273c3deeb9b41ab609bf3e498c2772bd01f7093998f0ddc4e8d4430f0f16991723eab208847b3b17c4f12c45f11fb7ef160ce412f654bed19a01391dfe1871bb5011698556ddd7b2b3b4a5e01de8973358376345f87251e4a83f441f734daaa98c1b7003880d50d5fd361e2e76922f5db5983462a9caead05eeed51ed908cd572386dd663f0bf6a0754f8e21182eb8eec53c1127b76f80e1240fab00baf0ce0124746ad3f01eb13b4111e56cd66ef86cd0746b5d3ed368d00ba6402ea7872f4d46f609f7fea30743131585f08731cc2dce30a28082c49f5cc2b3e600000000000000000000000000000000000000000000000000000000000003401a451db752bbff87c76c093329d993142b0787549c25eccc164353a29fb461bd00bf9b045f6d070c5d8aecb1a5cc8c29def1500597d9b38d37cce2667de63d8c181a8d33cb81310a690033c612af39de14b77e126e93469942623f9a05ac533f286e064293d2b11973cc4c7f7ff8f72b37de89d7c478fbc2b7a8be370cce64de074813884e50ab578e8583e4074e4d29962dae17334feffa476919671476ad19109b7fd1b036ceca8aa36744a719799d6d8095e063912ddbe0525233cfb8437e05313ddf9e3ccad360f9c3c488dc7f242310505bb9d2b3e6fa39c0f1d02f6f37225bb7947c2f1f9e13f9076d40dbd0fd0aaa347dfa3a552f2738c5b542b62ea6173f6d065826678b5bf44df499ead5af8eb1b89ae65cebc303f53121a31c85720f8adc2f8ffafcd4489238d35d8048ac68f28b9788f2bb9d508ca662ca2128ca227ba4f12010b7fe148930917129b4f8931182d4c667a1a145f30250857452e80761421f7917b5a90eb062929f0aca72a14437edf8320a66e97f1541c05c4913115d8e12af65042909e035a4adf966c52286df7ef72ff619a296cca5c2c48c740612494b868396081e9fa54e309cb55660eb053ed71247cd9512a8b55157bb2d02d5400d64b27e4658344fd9baa635f114a00bc8d7ab6345d5de5262a5e62c17134565e57862c89289125f08bb46f23a318497650254868d0caa211cc0ff24fe0b4cc9f4cba4e66f5844ce291fa71969d48f0c8473ad2debaecfcea19cf95adb023e86b2c7256520a4d60e077d8e32f6b1735c4446487644bbe7cc0ed2ea943c21a0fa98d92db6d929e812d919b3d7adc274defc9520729161835a77a6a42a9918ff0c8a288cfab2e1f819d751eac2c13f7e9668bb3810bdbd1e7bc78751a8b513ec0a28519cb03c4ab770f27fd41b04852dfdd3328a002206d26354e6c7e9bc18378f7d1c87b6b98acf094a287b43a884c23a6fb5741c6f00bfe52b1456133101cd8b4e6d1b624ede05ea6c7be29765e9037239a55c46944bca83f3c4a0a850181a15e4c2da5eb5f6cbb330d9c409e7235fc62f25d81eb13f01486a67b948c12e4957a0c24d274310b152db8a12c23749104aebe697369564d3ab06c40072e02c882979c3838d8598b9860f9772ba0b1a55a6e92d3ee4f034e4291aba0af868"
	input, _ := hex.DecodeString(inputHex)
	benchmarkZscverify(input, bench)
}

func benchmarkZscverify(input []byte, bench *testing.B) {
	var (
		res []byte
		err error
	)

	bench.Run(fmt.Sprintf("zscverify"), func(bench *testing.B) {
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			res, err = verify(input)
		}
		bench.StopTimer()
		//Check if it is correct
		if err != nil {
			bench.Error(err)
			return
		}
		if res[31] != 1 {
			bench.Error(fmt.Sprintf("Expected %v, got %v", 1, res[31]))
			return
		}
	})
}
