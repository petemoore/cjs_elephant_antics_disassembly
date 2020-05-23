# Full listing

```
5B00  00                 nop
5B01  01 00 00           ld bc,0
5B04  00                 nop
5B05  03                 inc bc
5B06  20 00              jr nz,$5B08
5B08  00                 nop
5B09  07                 rlca
5B0A  E6 00              and $00
5B0C  00                 nop
5B0D  05                 dec b
5B0E  B6                 or (hl)
5B0F  00                 nop
5B10  00                 nop
5B11  07                 rlca
5B12  FC 40 00           call m,$0040
5B15  10 EA              djnz $5B01
5B17  80                 add a,b
5B18  00                 nop
5B19  20 3D              jr nz,$5B58
5B1B  00                 nop
5B1C  00                 nop
5B1D  00                 nop
5B1E  0A                 ld a,(bc)
5B1F  80                 add a,b
5B20  00                 nop
5B21  04                 inc b
5B22  01 30 00           ld bc,48
5B25  DA 20 5C           jp c,$5C20
5B28  00                 nop
5B29  2C                 inc l
5B2A  20 3C              jr nz,$5B68
5B2C  00                 nop
5B2D  2B                 dec hl
5B2E  68                 ld l,b
5B2F  5E                 ld e,(hl)
5B30  00                 nop
5B31  D5                 push de
5B32  BC                 cp h
5B33  3E 00              ld a,0
5B35  34                 inc (hl)
5B36  B8                 cp b
5B37  5E                 ld e,(hl)
5B38  00                 nop
5B39  F8                 ret m
5B3A  70                 ld (hl),b
5B3B  3E 00              ld a,0
5B3D  D4 F3 1E           call nc,$1EF3
5B40  00                 nop
5B41  FF                 rst 38h
5B42  EF                 rst 28h
5B43  BE                 cp (hl)
5B44  00                 nop
5B45  9F                 sbc a,a
5B46  DF                 rst 18h
5B47  9C                 sbc a,h
5B48  00                 nop
5B49  C0                 ret nz
5B4A  BF                 cp a
5B4B  FC 00 FF           call m,$FF00
5B4E  7F                 ld a,a
5B4F  FC 00 7E           call m,$7E00
5B52  FF                 rst 38h
5B53  FC 00 9E           call m,$9E00
5B56  FF                 rst 38h
5B57  F8                 ret m
5B58  7F                 ld a,a
5B59  E1                 pop hl
5B5A  FF                 rst 38h
5B5B  B8                 cp b
5B5C  7F                 ld a,a
5B5D  F2 7F 78           jp p,$787F
5B60  FF                 rst 38h
5B61  FF                 rst 38h
5B62  3C                 inc a
5B63  F0                 ret p
5B64  FF                 rst 38h
5B65  FF                 rst 38h
5B66  F8                 ret m
5B67  F0                 ret p
5B68  3F                 ccf
5B69  81                 add a,c
5B6A  E0                 ret po
5B6B  20 0E              jr nz,$5B7B
5B6D  5E                 ld e,(hl)
5B6E  10 10              djnz $5B80
5B70  04                 inc b
5B71  E7                 rst 20h
5B72  FF                 rst 38h
5B73  10 01              djnz $5B76
5B75  D1                 pop de
5B76  FF                 rst 38h
5B77  90                 sub b
5B78  31 EB FF           ld sp,$FFEB
5B7B  E0                 ret po
5B7C  19                 add hl,de
5B7D  D0                 ret nc
5B7E  FF                 rst 38h
5B7F  F0                 ret p
5B80  68                 ld l,b
5B81  EB                 ex de,hl
5B82  FF                 rst 38h
5B83  F0                 ret p
5B84  6C                 ld l,h
5B85  73                 ld (hl),e
5B86  7F                 ld a,a
5B87  F0                 ret p
5B88  0C                 inc c
5B89  A8                 xor b
5B8A  FF                 rst 38h
5B8B  F0                 ret p
5B8C  64                 ld h,h
5B8D  50                 ld d,b
5B8E  CF                 rst 08h
5B8F  D0                 ret nc
5B90  08                 ex af,af'
5B91  A8                 xor b
5B92  0F                 rrca
5B93  D0                 ret nc
5B94  19                 add hl,de
5B95  40                 ld b,b
5B96  1C                 inc e
5B97  C0                 ret nz
5B98  12                 ld (de),a
5B99  A0                 and b
5B9A  0C                 inc c
5B9B  00                 nop
5B9C  11 40 00           ld de,64
5B9F  02                 ld (bc),a
5BA0  08                 ex af,af'
5BA1  80                 add a,b
5BA2  60                 ld h,b
5BA3  00                 nop
5BA4  82                 add a,d
5BA5  01 89 04           ld bc,1161
5BA8  C0                 ret nz
5BA9  03                 inc bc
5BAA  54                 ld d,h
5BAB  10 E4              djnz $5B91
5BAD  03                 inc bc
5BAE  A0                 and b
5BAF  40                 ld b,b
5BB0  70                 ld (hl),b
5BB1  00                 nop
5BB2  05                 dec b
5BB3  00                 nop
5BB4  20 03              jr nz,$5BB9
5BB6  F0                 ret p
5BB7  00                 nop
5BB8  01 EB FF           ld bc,-21
5BBB  E0                 ret po
5BBC  01 D0 FF           ld bc,-48
5BBF  F0                 ret p
5BC0  00                 nop
5BC1  EB                 ex de,hl
5BC2  FF                 rst 38h
5BC3  F0                 ret p
5BC4  00                 nop
5BC5  75                 ld (hl),l
5BC6  7F                 ld a,a
5BC7  F0                 ret p
5BC8  00                 nop
5BC9  AA                 xor d
5BCA  FF                 rst 38h
5BCB  F0                 ret p
5BCC  00                 nop
5BCD  54                 ld d,h
5BCE  CF                 rst 08h
5BCF  D0                 ret nc
5BD0  00                 nop
5BD1  2A 0F D0           ld hl,($D00F)
5BD4  00                 nop
5BD5  54                 ld d,h
5BD6  1C                 inc e
5BD7  C0                 ret nz
5BD8  00                 nop
5BD9  2A 0C 00           ld hl,($000C)
5BDC  3C                 inc a
5BDD  04                 inc b
5BDE  00                 nop
5BDF  00                 nop
5BE0  CE 02              adc a,2
5BE2  00                 nop
5BE3  00                 nop
5BE4  83                 add a,e
5BE5  C0                 ret nz
5BE6  00                 nop
5BE7  00                 nop
5BE8  38 FE              jr c,$5BE8
5BEA  00                 nop
5BEB  00                 nop
5BEC  38 07              jr c,$5BF5
5BEE  00                 nop
5BEF  00                 nop
5BF0  00                 nop
5BF1  00                 nop
5BF2  00                 nop
5BF3  00                 nop
5BF4  FF                 rst 38h
5BF5  FF                 rst 38h
5BF6  00                 nop
5BF7  00                 nop
5BF8  C1                 pop bc
5BF9  EB                 ex de,hl
5BFA  FF                 rst 38h
5BFB  E0                 ret po
5BFC  E1                 pop hl
5BFD  D0                 ret nc
5BFE  FF                 rst 38h
5BFF  F0                 ret p
5C00  60                 ld h,b
5C01  EA 3F F0           jp pe,$F03F
5C04  A8                 xor b
5C05  75                 ld (hl),l
5C06  3F                 ccf
5C07  F0                 ret p
5C08  50                 ld d,b
5C09  2A 9F F0           ld hl,($F09F)
5C0C  11 15 CF           ld de,$CF15
5C0F  D0                 ret nc
5C10  20 0B              jr nz,$5C1D
5C12  D3 D0              out ($D0),a
5C14  22 05 D0           ld ($D005),hl
5C17  C0                 ret nz
5C18  00                 nop
5C19  00                 nop
5C1A  A0                 and b
5C1B  00                 nop
5C1C  08                 ex af,af'
5C1D  0E 43              ld c,67
5C1F  00                 nop
5C20  80                 add a,b
5C21  3F                 ccf
5C22  06 00              ld b,0
5C24  50                 ld d,b
5C25  41                 ld b,c
5C26  8E                 adc a,(hl)
5C27  00                 nop
5C28  00                 nop
5C29  80                 add a,b
5C2A  18 00              jr $5C2C
5C2C  00                 nop
5C2D  80                 add a,b
5C2E  30 00              jr nc,$5C30
5C30  00                 nop
5C31  00                 nop
5C32  E0                 ret po
5C33  00                 nop
5C34  00                 nop
5C35  FF                 rst 38h
5C36  80                 add a,b
5C37  00                 nop
5C38  E8                 ret pe
5C39  00                 nop
5C3A  00                 nop
5C3B  00                 nop
5C3C  FA 00 FB           jp m,$FB00
5C3F  00                 nop
5C40  03                 inc bc
5C41  00                 nop
5C42  F3                 di
5C43  00                 nop
5C44  F2 00 00           jp p,$0000
5C47  80                 add a,b
5C48  F9                 ld sp,hl
5C49  C0                 ret nz
5C4A  FB                 ei
5C4B  A0                 and b
5C4C  03                 inc bc
5C4D  D0                 ret nc
5C4E  33                 inc sp
5C4F  44                 ld b,h
5C50  E8                 ret pe
5C51  81                 add a,c
5C52  32 01 00           ld ($0001),a
5C55  00                 nop
5C56  F1                 pop af
5C57  03                 inc bc
5C58  F8                 ret m
5C59  43                 ld b,e
5C5A  78                 ld a,b
5C5B  07                 rlca
5C5C  FD 97              sub a
5C5E  FC C3 7C           call m,$7CC3
5C61  73                 ld (hl),e
5C62  FC 3D F8           call m,$F83D
5C65  BE                 cp (hl)
5C66  F0                 ret p
5C67  FF                 rst 38h
5C68  E8                 ret pe
5C69  00                 nop
5C6A  00                 nop
5C6B  00                 nop
5C6C  FB                 ei
5C6D  00                 nop
5C6E  FB                 ei
5C6F  50                 ld d,b
5C70  03                 inc bc
5C71  03                 inc bc
5C72  F3                 di
5C73  01 F3 00           ld bc,243
5C76  00                 nop
5C77  03                 inc bc
5C78  F8                 ret m
5C79  43                 ld b,e
5C7A  78                 ld a,b
5C7B  07                 rlca
5C7C  FD 97              sub a
5C7E  FC C3 7C           call m,$7CC3
5C81  73                 ld (hl),e
5C82  FC 3D F8           call m,$F83D
5C85  BE                 cp (hl)
5C86  F0                 ret p
5C87  FF                 rst 38h
5C88  DA 20 8C           jp c,$8C20
5C8B  20 8B              jr nz,$5C18
5C8D  68                 ld l,b
5C8E  D5                 push de
5C8F  BC                 cp h
5C90  34                 inc (hl)
5C91  B8                 cp b
5C92  C0                 ret nz
5C93  BF                 cp a
5C94  E0                 ret po
5C95  7F                 ld a,a
5C96  78                 ld a,b
5C97  FF                 rst 38h
5C98  9E                 sbc a,(hl)
5C99  FF                 rst 38h
5C9A  03                 inc bc
5C9B  40                 ld b,b
5C9C  0E A8              ld c,-88
5C9E  1F                 rra
5C9F  D4 3F FA           call nc,$FA3F
5CA2  7B                 ld a,e
5CA3  DD 7D              ld a,ixl
5CA5  BA                 cp d
5CA6  F3                 di
5CA7  4D                 ld c,l
5CA8  F3                 di
5CA9  8E                 adc a,(hl)
5CAA  FF                 rst 38h
5CAB  5D                 ld e,l
5CAC  FF                 rst 38h
5CAD  B6                 or (hl)
5CAE  7F                 ld a,a
5CAF  FD 7F              ld a,a
5CB1  FA 1F 80           jp m,$801F
5CB4  3F                 ccf
5CB5  E0                 ret po
5CB6  3F                 ccf
5CB7  F0                 ret p
5CB8  7F                 ld a,a
5CB9  FC 7F FE           call m,$FE7F
5CBC  3F                 ccf
5CBD  FE 3F              cp 63
5CBF  FF                 rst 38h
5CC0  1F                 rra
5CC1  FF                 rst 38h
5CC2  0F                 rrca
5CC3  FB                 ei
5CC4  03                 inc bc
5CC5  F7                 rst 30h
5CC6  01 F7 00           ld bc,247
5CC9  EF                 rst 28h
5CCA  00                 nop
5CCB  6F                 ld l,a
5CCC  00                 nop
5CCD  1F                 rra
5CCE  00                 nop
5CCF  1F                 rra
5CD0  00                 nop
5CD1  1F                 rra
5CD2  00                 nop
5CD3  1F                 rra
5CD4  00                 nop
5CD5  1F                 rra
5CD6  00                 nop
5CD7  1F                 rra
5CD8  00                 nop
5CD9  1F                 rra
5CDA  00                 nop
5CDB  0F                 rrca
5CDC  00                 nop
5CDD  0F                 rrca
5CDE  00                 nop
5CDF  07                 rlca
5CE0  00                 nop
5CE1  0B                 dec bc
5CE2  00                 nop
5CE3  0E 00              ld c,0
5CE5  FE 07              cp 7
5CE7  FF                 rst 38h
5CE8  0F                 rrca
5CE9  FF                 rst 38h
5CEA  1F                 rra
5CEB  FB                 ei
5CEC  3F                 ccf
5CED  F7                 rst 30h
5CEE  3F                 ccf
5CEF  F7                 rst 30h
5CF0  7F                 ld a,a
5CF1  EF                 rst 28h
5CF2  7F                 ld a,a
5CF3  EF                 rst 28h
5CF4  7F                 ld a,a
5CF5  9F                 sbc a,a
5CF6  3E 1F              ld a,31
5CF8  1C                 inc e
5CF9  1F                 rra
5CFA  00                 nop
5CFB  1F                 rra
5CFC  00                 nop
5CFD  1F                 rra
5CFE  00                 nop
5CFF  1F                 rra
5D00  00                 nop
5D01  0F                 rrca
5D02  00                 nop
5D03  0F                 rrca
5D04  00                 nop
5D05  07                 rlca
5D06  00                 nop
5D07  07                 rlca
5D08  00                 nop
5D09  0B                 dec bc
5D0A  00                 nop
5D0B  02                 ld (bc),a
5D0C  00                 nop
5D0D  07                 rlca
5D0E  00                 nop
5D0F  1F                 rra
5D10  00                 nop
5D11  3B                 dec sp
5D12  00                 nop
5D13  37                 scf
5D14  00                 nop
5D15  77                 ld (hl),a
5D16  00                 nop
5D17  6F                 ld l,a
5D18  00                 nop
5D19  EF                 rst 28h
5D1A  00                 nop
5D1B  DF                 rst 18h
5D1C  00                 nop
5D1D  DF                 rst 18h
5D1E  00                 nop
5D1F  DF                 rst 18h
5D20  01 DF 01           ld bc,479
5D23  DF                 rst 18h
5D24  01 DF 01           ld bc,479
5D27  DF                 rst 18h
5D28  01 EF 01           ld bc,495
5D2B  EF                 rst 28h
5D2C  01 F7 00           ld bc,247
5D2F  EB                 ex de,hl
5D30  B8                 cp b
5D31  04                 inc b
5D32  DF                 rst 18h
5D33  FD FF              rst 38h
5D35  EA FE FF           jp pe,$FFFE
5D38  FD 3F              ccf
5D3A  FE 3F              cp 63
5D3C  FD 7D              ld a,iyl
5D3E  FE FF              cp -1
5D40  FD 3F              ccf
5D42  FE 3F              cp 63
5D44  FD 7F              ld a,a
5D46  FE FF              cp -1
5D48  FD 3B              dec sp
5D4A  FE 3E              cp 62
5D4C  FD 7F              ld a,a
5D4E  FE ED              cp -19
5D50  80                 add a,b
5D51  00                 nop
5D52  10 00              djnz $5D54
5D54  18 00              jr $5D56
5D56  28 00              jr z,$5D58
5D58  44                 ld b,h
5D59  00                 nop
5D5A  A4                 and h
5D5B  00                 nop
5D5C  D2 00 62           jp nc,$6200
5D5F  00                 nop
5D60  D2 00 E9           jp nc,$E900
5D63  00                 nop
5D64  B1                 or c
5D65  00                 nop
5D66  E9                 jp (hl)
5D67  00                 nop
5D68  D0                 ret nc
5D69  80                 add a,b
5D6A  68                 ld l,b
5D6B  80                 add a,b
5D6C  D2 80 A1           jp nc,$A180
5D6F  00                 nop
5D70  C3 00 8A           jp $8A00
5D73  00                 nop
5D74  50                 ld d,b
5D75  00                 nop
5D76  08                 ex af,af'
5D77  00                 nop
5D78  9F                 sbc a,a
5D79  80                 add a,b
5D7A  3F                 ccf
5D7B  E0                 ret po
5D7C  5F                 ld e,a
5D7D  F8                 ret m
5D7E  AD                 xor l
5D7F  F8                 ret m
5D80  A7                 and a
5D81  FC A2 BE           call m,$BEA2
5D84  D5                 push de
5D85  FE 52              cp 82
5D87  AF                 xor a
5D88  D1                 pop de
5D89  7F                 ld a,a
5D8A  A8                 xor b
5D8B  AF                 xor a
5D8C  C8                 ret z
5D8D  56                 ld d,(hl)
5D8E  A8                 xor b
5D8F  0C                 inc c
5D90  C8                 ret z
5D91  00                 nop
5D92  68                 ld l,b
5D93  00                 nop
5D94  C0                 ret nz
5D95  00                 nop
5D96  A0                 and b
5D97  00                 nop
5D98  C0                 ret nz
5D99  00                 nop
5D9A  80                 add a,b
5D9B  00                 nop
5D9C  50                 ld d,b
5D9D  00                 nop
5D9E  01 FC 07           ld bc,2044
5DA1  FE 8F              cp -113
5DA3  FE 9F              cp -97
5DA5  FE 96              cp -106
5DA7  FA AF DE           jp m,$DEAF
5DAA  1D                 dec e
5DAB  FC 57 50           call m,$5057
5DAE  AA                 xor d
5DAF  A0                 and b
5DB0  25                 dec h
5DB1  40                 ld b,b
5DB2  12                 ld (de),a
5DB3  80                 add a,b
5DB4  91                 sub c
5DB5  00                 nop
5DB6  D0                 ret nc
5DB7  00                 nop
5DB8  88                 adc a,b
5DB9  00                 nop
5DBA  C8                 ret z
5DBB  00                 nop
5DBC  68                 ld l,b
5DBD  00                 nop
5DBE  C8                 ret z
5DBF  00                 nop
5DC0  A8                 xor b
5DC1  00                 nop
5DC2  C8                 ret z
5DC3  00                 nop
5DC4  90                 sub b
5DC5  00                 nop
5DC6  D0                 ret nc
5DC7  00                 nop
5DC8  90                 sub b
5DC9  00                 nop
5DCA  C0                 ret nz
5DCB  00                 nop
5DCC  90                 sub b
5DCD  00                 nop
5DCE  0D                 dec c
5DCF  FD 1F              rra
5DD1  FE 1F              cp 31
5DD3  F7                 rst 30h
5DD4  1F                 rra
5DD5  D9                 exx
5DD6  3F                 ccf
5DD7  60                 ld h,b
5DD8  3F                 ccf
5DD9  C8                 ret z
5DDA  3F                 ccf
5DDB  A0                 and b
5DDC  7E                 ld a,(hl)
5DDD  D0                 ret nc
5DDE  7F                 ld a,a
5DDF  80                 add a,b
5DE0  7F                 ld a,a
5DE1  50                 ld d,b
5DE2  FE 80              cp -128
5DE4  FB                 ei
5DE5  20 FE              jr nz,$5DE5
5DE7  00                 nop
5DE8  FD 40              ld b,b
5DEA  F8                 ret m
5DEB  00                 nop
5DEC  7D                 ld a,l
5DED  00                 nop
5DEE  0B                 dec bc
5DEF  FF                 rst 38h
5DF0  0F                 rrca
5DF1  FF                 rst 38h
5DF2  0F                 rrca
5DF3  FF                 rst 38h
5DF4  07                 rlca
5DF5  FF                 rst 38h
5DF6  07                 rlca
5DF7  FB                 ei
5DF8  07                 rlca
5DF9  D7                 rst 10h
5DFA  07                 rlca
5DFB  E9                 jp (hl)
5DFC  03                 inc bc
5DFD  F4 03 FA           call p,$FA03
5E00  01 FC 01           ld bc,508
5E03  F6 00              or $00
5E05  FD 00              nop
5E07  7E                 ld a,(hl)
5E08  00                 nop
5E09  7D                 ld a,l
5E0A  00                 nop
5E0B  3A 00 0C           ld a,($0C00)
5E0E  3A 28 14           ld a,($1428)
5E11  F4 A9 7A           call p,$7AA9
5E14  40                 ld b,b
5E15  FC 07 F6           call m,$F607
5E18  07                 rlca
5E19  FD 03              inc bc
5E1B  FE 03              cp 3
5E1D  FB                 ei
5E1E  01 FE 01           ld bc,510
5E21  FF                 rst 38h
5E22  01 FB 00           ld bc,251
5E25  FE 00              cp 0
5E27  7F                 ld a,a
5E28  00                 nop
5E29  7F                 ld a,a
5E2A  00                 nop
5E2B  3F                 ccf
5E2C  00                 nop
5E2D  0F                 rrca
5E2E  3F                 ccf
5E2F  90                 sub b
5E30  3F                 ccf
5E31  50                 ld d,b
5E32  F6 90              or $90
5E34  FF                 rst 38h
5E35  50                 ld d,b
5E36  F5                 push af
5E37  A0                 and b
5E38  EB                 ex de,hl
5E39  20 C6              jr nz,$5E01
5E3B  A0                 and b
5E3C  0B                 dec bc
5E3D  20 16              jr nz,$5E55
5E3F  40                 ld b,b
5E40  2F                 cpl
5E41  40                 ld b,b
5E42  5E                 ld e,(hl)
5E43  40                 ld b,b
5E44  34                 inc (hl)
5E45  80                 add a,b
5E46  5E                 ld e,(hl)
5E47  80                 add a,b
5E48  35                 dec (hl)
5E49  00                 nop
5E4A  5B                 ld e,e
5E4B  00                 nop
5E4C  2C                 inc l
5E4D  00                 nop
5E4E  00                 nop
5E4F  00                 nop
5E50  7F                 ld a,a
5E51  FF                 rst 38h
5E52  06 30              ld b,48
5E54  FA 2F 12           jp m,$122F
5E57  24                 inc h
5E58  ED DB              ???
5E5A  22 22 E9           ld ($E922),hl
5E5D  CB 58              bit 3,b
5E5F  0D                 dec c
5E60  DB BD              in a,($BD)
5E62  5F                 ld e,a
5E63  FD DE 3D           sbc a,61
5E66  5B                 ld e,e
5E67  ED D6              ???
5E69  35                 dec (hl)
5E6A  5C                 ld e,h
5E6B  1D                 dec e
5E6C  DF                 rst 18h
5E6D  FD 00              nop
5E6F  00                 nop
5E70  7F                 ld a,a
5E71  FF                 rst 38h
5E72  06 30              ld b,48
5E74  FA 2F 12           jp m,$122F
5E77  24                 inc h
5E78  ED DB              ???
5E7A  22 22 E9           ld ($E922),hl
5E7D  CB 58              bit 3,b
5E7F  0D                 dec c
5E80  DD DD 5F           ld e,a
5E83  FD DE 3D           sbc a,61
5E86  5B                 ld e,e
5E87  ED D6              ???
5E89  35                 dec (hl)
5E8A  5D                 ld e,l
5E8B  DD CF              rst 08h
5E8D  F9                 ld sp,hl
5E8E  00                 nop
5E8F  00                 nop
5E90  7F                 ld a,a
5E91  FF                 rst 38h
5E92  06 30              ld b,48
5E94  FA 2F 12           jp m,$122F
5E97  24                 inc h
5E98  ED DB              ???
5E9A  22 22 E9           ld ($E922),hl
5E9D  CB 58              bit 3,b
5E9F  0D                 dec c
5EA0  DE ED              sbc a,-19
5EA2  5F                 ld e,a
5EA3  FD DE 3D           sbc a,61
5EA6  5B                 ld e,e
5EA7  ED D6              ???
5EA9  35                 dec (hl)
5EAA  5C                 ld e,h
5EAB  1D                 dec e
5EAC  DF                 rst 18h
5EAD  FD 01 83 01        ld bc,$0183
5EB1  80                 add a,b
5EB2  01 8F 01           ld bc,399
5EB5  00                 nop
5EB6  00                 nop
5EB7  9F                 sbc a,a
5EB8  01 00 00           ld bc,0
5EBB  9F                 sbc a,a
5EBC  01 00 01           ld bc,256
5EBF  BF                 cp a
5EC0  01 80 01           ld bc,384
5EC3  BF                 cp a
5EC4  00                 nop
5EC5  80                 add a,b
5EC6  01 3F 00           ld bc,63
5EC9  80                 add a,b
5ECA  01 3F 00           ld bc,63
5ECD  80                 add a,b
5ECE  0D                 dec c
5ECF  80                 add a,b
5ED0  09                 add hl,bc
5ED1  A2                 and d
5ED2  1A                 ld a,(de)
5ED3  03                 inc bc
5ED4  1B                 dec de
5ED5  F1                 pop af
5ED6  17                 rla
5ED7  F0                 ret p
5ED8  16 00              ld d,0
5EDA  17                 rla
5EDB  F8                 ret m
5EDC  1B                 dec de
5EDD  F8                 ret m
5EDE  0A                 ld a,(bc)
5EDF  00                 nop
5EE0  03                 inc bc
5EE1  F0                 ret p
5EE2  01 F0 00           ld bc,240
5EE5  02                 ld (bc),a
5EE6  01 9A 00           ld bc,154
5EE9  03                 inc bc
5EEA  00                 nop
5EEB  00                 nop
5EEC  30 0F              jr nc,$5EFD
5EEE  30 00              jr nc,$5EF0
5EF0  30 1F              jr nc,$5F11
5EF2  28 00              jr z,$5EF4
5EF4  10 1F              djnz $5F15
5EF6  0C                 inc c
5EF7  00                 nop
5EF8  14                 inc d
5EF9  3F                 ccf
5EFA  0C                 inc c
5EFB  00                 nop
5EFC  0C                 inc c
5EFD  3F                 ccf
5EFE  0A                 ld a,(bc)
5EFF  00                 nop
5F00  04                 inc b
5F01  3F                 ccf
5F02  02                 ld (bc),a
5F03  00                 nop
5F04  05                 dec b
5F05  3F                 ccf
5F06  03                 inc bc
5F07  00                 nop
5F08  03                 inc bc
5F09  00                 nop
5F0A  08                 ex af,af'
5F0B  E2 0B C3           jp po,$C30B
5F0E  1B                 dec de
5F0F  31 17 F0           ld sp,$F017
5F12  17                 rla
5F13  C8                 ret z
5F14  14                 inc d
5F15  38 1B              jr c,$5F32
5F17  F8                 ret m
5F18  0B                 dec bc
5F19  C4 02 3C           call nz,$3C02
5F1C  01 F8 01           ld bc,504
5F1F  C2 00 1A           jp nz,$1A00
5F22  40                 ld b,b
5F23  01 C0 01           ld bc,448
5F26  03                 inc bc
5F27  60                 ld h,b
5F28  71                 ld (hl),c
5F29  C7                 rst 00h
5F2A  38 0E              jr c,$5F3A
5F2C  1C                 inc e
5F2D  1C                 inc e
5F2E  88                 adc a,b
5F2F  80                 add a,b
5F30  01 C0 BF           ld bc,$BFC0
5F33  FE 9F              cp -97
5F35  FC 00 00           call m,$0000
5F38  FF                 rst 38h
5F39  FF                 rst 38h
5F3A  00                 nop
5F3B  00                 nop
5F3C  4E                 ld c,(hl)
5F3D  39                 add hl,sp
5F3E  CE 39              adc a,57
5F40  CE 39              adc a,57
5F42  CE 39              adc a,57
5F44  CE 39              adc a,57
5F46  CE 39              adc a,57
5F48  CE 39              adc a,57
5F4A  8E                 adc a,(hl)
5F4B  38 04              jr c,$5F51
5F4D  10 70              djnz $5FBF
5F4F  07                 rlca
5F50  0F                 rrca
5F51  F8                 ret m
5F52  E0                 ret po
5F53  00                 nop
5F54  00                 nop
5F55  00                 nop
5F56  F8                 ret m
5F57  06 00              ld b,0
5F59  06 FC              ld b,-4
5F5B  06 00              ld b,0
5F5D  0A                 ld a,(bc)
5F5E  FC 04 00           call m,$0004
5F61  18 FE              jr $5F61
5F63  14                 inc d
5F64  00                 nop
5F65  18 FE              jr $5F65
5F67  18 00              jr $5F69
5F69  28 FE              jr z,$5F69
5F6B  10 00              djnz $5F6D
5F6D  20 FE              jr nz,$5F6D
5F6F  50                 ld d,b
5F70  00                 nop
5F71  60                 ld h,b
5F72  00                 nop
5F73  60                 ld h,b
5F74  23                 inc hl
5F75  88                 adc a,b
5F76  61                 ld h,c
5F77  E8                 ret pe
5F78  C6 6C              add a,108
5F7A  07                 rlca
5F7B  F4 89 F4           call p,$F489
5F7E  8E                 adc a,(hl)
5F7F  14                 inc d
5F80  0F                 rrca
5F81  EC 91 E8           call pe,$E891
5F84  1E 20              ld e,32
5F86  0F                 rrca
5F87  C0                 ret nz
5F88  A1                 and c
5F89  C0                 ret nz
5F8A  AC                 xor h
5F8B  00                 nop
5F8C  E0                 ret po
5F8D  C0                 ret nz
5F8E  00                 nop
5F8F  C0                 ret nz
5F90  F8                 ret m
5F91  C0                 ret nz
5F92  00                 nop
5F93  40                 ld b,b
5F94  FC 80 00           call m,$0080
5F97  40                 ld b,b
5F98  FC 80 00           call m,$0080
5F9B  40                 ld b,b
5F9C  FE C0              cp -64
5F9E  00                 nop
5F9F  C0                 ret nz
5FA0  FE C0              cp -64
5FA2  00                 nop
5FA3  80                 add a,b
5FA4  FE 40              cp 64
5FA6  00                 nop
5FA7  80                 add a,b
5FA8  FE 40              cp 64
5FAA  00                 nop
5FAB  80                 add a,b
5FAC  00                 nop
5FAD  D8                 ret c
5FAE  22 C8 60           ld ($60C8),hl
5FB1  2C                 inc l
5FB2  C7                 rst 00h
5FB3  EC 07 F4           call pe,$F407
5FB6  80                 add a,b
5FB7  34                 inc (hl)
5FB8  8F                 adc a,a
5FB9  F4 0F EC           call p,$EC0F
5FBC  80                 add a,b
5FBD  28 07              jr z,$5FC6
5FBF  E0                 ret po
5FC0  07                 rlca
5FC1  C0                 ret nz
5FC2  A0                 and b
5FC3  00                 nop
5FC4  AC                 xor h
5FC5  C0                 ret nz
5FC6  03                 inc bc
5FC7  FC 0F F8           call m,$F80F
5FCA  12                 ld (de),a
5FCB  44                 ld b,h
5FCC  2D                 dec l
5FCD  BA                 cp d
5FCE  2D                 dec l
5FCF  BA                 cp d
5FD0  00                 nop
5FD1  00                 nop
5FD2  3F                 ccf
5FD3  FE 00              cp 0
5FD5  1A                 ld a,(de)
5FD6  00                 nop
5FD7  66                 ld h,(hl)
5FD8  00                 nop
5FD9  00                 nop
5FDA  00                 nop
5FDB  7E                 ld a,(hl)
5FDC  00                 nop
5FDD  00                 nop
5FDE  00                 nop
5FDF  AA                 xor d
5FE0  00                 nop
5FE1  00                 nop
5FE2  00                 nop
5FE3  AA                 xor d
5FE4  00                 nop
5FE5  00                 nop
5FE6  03                 inc bc
5FE7  FC 0F F8           call m,$F80F
5FEA  12                 ld (de),a
5FEB  44                 ld b,h
5FEC  2D                 dec l
5FED  BA                 cp d
5FEE  2D                 dec l
5FEF  BA                 cp d
5FF0  00                 nop
5FF1  00                 nop
5FF2  3F                 ccf
5FF3  FE 2C              cp 44
5FF5  00                 nop
5FF6  33                 inc sp
5FF7  00                 nop
5FF8  00                 nop
5FF9  00                 nop
5FFA  3F                 ccf
5FFB  00                 nop
5FFC  00                 nop
5FFD  00                 nop
5FFE  2A 80 00           ld hl,($0080)
6001  00                 nop
6002  2A 80 00           ld hl,($0080)
6005  00                 nop
6006  1F                 rra
6007  E0                 ret po
6008  0F                 rrca
6009  F8                 ret m
600A  11 24 2E           ld de,$2E24
600D  DA 2E DA           jp c,$DA2E
6010  00                 nop
6011  00                 nop
6012  3F                 ccf
6013  FE 1F              cp 31
6015  E0                 ret po
6016  0F                 rrca
6017  F8                 ret m
6018  11 24 2E           ld de,$2E24
601B  DA 2E DA           jp c,$DA2E
601E  00                 nop
601F  00                 nop
6020  3F                 ccf
6021  FE 32              cp 50
6023  60                 ld h,b
6024  32 60 5E           ld ($5E60),a
6027  65                 ld h,l
6028  7F                 ld a,a
6029  6D                 ld l,l
602A  F0                 ret p
602B  03                 inc bc
602C  2C                 inc l
602D  05                 dec b
602E  21 08 8B           ld hl,$8B08
6031  06 80              ld b,-128
6033  17                 rla
6034  08                 ex af,af'
6035  02                 ld (bc),a
6036  10 11              djnz $6049
6038  09                 add hl,bc
6039  16 1D              ld d,29
603B  1E 0A              ld e,10
603D  19                 add hl,de
603E  0F                 rrca
603F  0C                 inc c
6040  38 0D              jr c,$604F
6042  23                 inc hl
6043  14                 inc d
6044  0E 35              ld c,53
6046  34                 inc (hl)
6047  15                 dec d
6048  37                 scf
6049  22 36 0B           ld ($0B36),hl
604C  06 21              ld b,33
604E  1F                 rra
604F  1A                 ld a,(de)
6050  C7                 rst 00h
6051  F4 02 E5           call p,$E502
6054  B9                 cp c
6055  77                 ld (hl),a
6056  F8                 ret m
6057  01 7D 90           ld bc,$907D
605A  97                 sub a
605B  FE 0F              cp 15
605D  A4                 and h
605E  05                 dec b
605F  CB 5D              bit 3,l
6061  FC 01 1F           call m,$1F01
6064  F8                 ret m
6065  3E 40              ld a,64
6067  3F                 ccf
6068  23                 inc hl
6069  E4 C0 0B           call po,$0BC0
606C  91                 sub c
606D  B5                 or l
606E  C0                 ret nz
606F  7A                 ld a,d
6070  47                 ld b,a
6071  FC 01 7D           call m,$7D01
6074  90                 sub b
6075  97                 sub a
6076  F0                 ret p
6077  02                 ld (bc),a
6078  E4 6D 77           call po,$776D
607B  FE 0F              cp 15
607D  C4 18 FF           call nz,$FF18
6080  C0                 ret nz
6081  0F                 rrca
6082  E1                 pop hl
6083  E4 1D F9           call po,$F91D
6086  31 F6 01           ld sp,$01F6
6089  E3                 ex (sp),hl
608A  1A                 ld a,(de)
608B  BF                 cp a
608C  DC 01 7D           call c,$7D01
608F  90                 sub b
6090  97                 sub a
6091  F8                 ret m
6092  07                 rlca
6093  A4                 and h
6094  7D                 ld a,l
6095  80                 add a,b
6096  2E 46              ld l,70
6098  EF                 rst 28h
6099  84                 add a,h
609A  05                 dec b
609B  C8                 ret z
609C  DA EF C4           jp c,$C4EF
609F  18 FF              jr $60A0
60A1  C1                 pop bc
60A2  FC 02 2F           call m,$2F02
60A5  D8                 ret c
60A6  F9                 ld sp,hl
60A7  31 F6 00           ld sp,$00F6
60AA  FA DF DC           jp m,$DCDF
60AD  01 7D 90           ld bc,$907D
60B0  97                 sub a
60B1  D8                 ret c
60B2  07                 rlca
60B3  8C                 adc a,h
60B4  6A                 ld l,d
60B5  FF                 rst 38h
60B6  C0                 ret nz
60B7  09                 add hl,bc
60B8  42                 ld b,d
60B9  12                 ld (de),a
60BA  84                 add a,h
60BB  F9                 ld sp,hl
60BC  01 4F E0           ld bc,$E04F
60BF  14                 inc d
60C0  21 29 F5           ld hl,$F529
60C3  83                 add a,e
60C4  1F                 rra
60C5  F8                 ret m
60C6  01 1A 87           ld bc,$871A
60C9  BF                 cp a
60CA  26 1A              ld h,26
60CC  EF                 rst 28h
60CD  E0                 ret po
60CE  0F                 rrca
60CF  18 FE              jr $60CF
60D1  35                 dec (hl)
60D2  7F                 ld a,a
60D3  D0                 ret nc
60D4  0F                 rrca
60D5  48                 ld c,b
60D6  00                 nop
60D7  0B                 dec bc
60D8  97                 sub a
60D9  00                 nop
60DA  02                 ld (bc),a
60DB  FB                 ei
60DC  21 2F B0           ld hl,$B02F
60DF  07                 rlca
60E0  D6 C0              sub -64
60E2  0F                 rrca
60E3  89                 adc a,c
60E4  00                 nop
60E5  07                 rlca
60E6  F0                 ret p
60E7  80                 add a,b
60E8  03                 inc bc
60E9  E4 47 C2           call po,$C247
60EC  0A                 ld a,(bc)
60ED  10 00              djnz $60EF
60EF  A1                 and c
60F0  3E 40              ld a,64
60F2  63                 ld h,e
60F3  EF                 rst 28h
60F4  04                 inc b
60F5  6A                 ld l,d
60F6  10 03              djnz $60FB
60F8  BF                 cp a
60F9  26 3F              ld h,63
60FB  40                 ld b,b
60FC  1F                 rra
60FD  F8                 ret m
60FE  D6 C6              sub -58
6100  D7                 rst 10h
6101  01 E3 1A           ld bc,$1AE3
6104  BF                 cp a
6105  D8                 ret c
6106  01 7D 90           ld bc,$907D
6109  97                 sub a
610A  C8                 ret z
610B  0A                 ld a,(bc)
610C  12                 ld (de),a
610D  8F                 adc a,a
610E  04                 inc b
610F  27                 daa
6110  FE 00              cp 0
6112  00                 nop
6113  94                 sub h
6114  07                 rlca
6115  C8                 ret z
6116  85                 add a,l
6117  21 1F 20           ld hl,$201F
611A  31 FA 80           ld sp,$80FA
611D  B9                 cp c
611E  1B                 dec de
611F  5C                 ld e,h
6120  00                 nop
6121  46                 ld b,(hl)
6122  A1                 and c
6123  F9                 ld sp,hl
6124  00                 nop
6125  F1                 pop af
6126  AF                 xor a
6127  93                 sub e
6128  00                 nop
6129  7A                 ld a,d
612A  47                 ld b,a
612B  F0                 ret p
612C  03                 inc bc
612D  FF                 rst 38h
612E  1A                 ld a,(de)
612F  DF                 rst 18h
6130  C8                 ret z
6131  03                 inc bc
6132  EB                 ex de,hl
6133  7F                 ld a,a
6134  60                 ld h,b
6135  05                 dec b
6136  F6 42              or $42
6138  44                 ld b,h
6139  A1                 and c
613A  01 FA 22           ld bc,$22FA
613D  DE FF              sbc a,-1
613F  C0                 ret nz
6140  84                 add a,h
6141  7D                 ld a,l
6142  80                 add a,b
6143  C7                 rst 00h
6144  E8                 ret pe
6145  0C                 inc c
6146  7D                 ld a,l
6147  60                 ld h,b
6148  8D                 adc a,l
6149  43                 ld b,e
614A  FA 01 E3           jp m,$E301
614D  4F                 ld c,a
614E  E0                 ret po
614F  F9                 ld sp,hl
6150  30 F1              jr nc,$6143
6152  8D                 adc a,l
6153  5C                 ld e,h
6154  9F                 sbc a,a
6155  C5                 push bc
6156  1E 08              ld e,8
6158  0A                 ld a,(bc)
6159  12                 ld (de),a
615A  8F                 adc a,a
615B  48                 ld c,b
615C  A1                 and c
615D  1E 09              ld e,9
615F  42                 ld b,d
6160  53                 ld d,e
6161  F8                 ret m
6162  01 05 F6           ld bc,$F605
6165  42                 ld b,d
6166  5F                 ld e,a
6167  C2 00 01           jp nz,$0100
616A  7F                 ld a,a
616B  10 32              djnz $619F
616D  6D                 ld l,l
616E  EF                 rst 28h
616F  9C                 sbc a,h
6170  04                 inc b
6171  04                 inc b
6172  FE C0              cp -64
6174  60                 ld h,b
6175  94                 sub h
6176  FE 01              cp 1
6178  8F                 adc a,a
6179  8C                 adc a,h
617A  11 A8 7E           ld de,$7EA8
617D  20 2F              jr nz,$61AE
617F  E4 C7 F2           call po,$F2C7
6182  8F                 adc a,a
6183  E0                 ret po
6184  08                 ex af,af'
6185  17                 rla
6186  D9                 exx
6187  09                 add hl,bc
6188  7E                 ld a,(hl)
6189  80                 add a,b
618A  17                 rla
618B  E9                 jp (hl)
618C  03                 inc bc
618D  26 DE              ld h,-34
618F  FB                 ei
6190  00                 nop
6191  5C                 ld e,h
6192  8D                 adc a,l
6193  AE                 xor (hl)
6194  FF                 rst 38h
6195  00                 nop
6196  8F                 adc a,a
6197  B0                 or b
6198  08                 ex af,af'
6199  47                 ld b,a
619A  C4 0C 7F           call nz,$7F0C
619D  08                 ex af,af'
619E  00                 nop
619F  31 F6 82           ld sp,$82F6
61A2  35                 dec (hl)
61A3  0F                 rrca
61A4  D4 05 FC           call nc,$FC05
61A7  98                 sbc a,b
61A8  FA 61 2F           jp m,$2F61
61AB  90                 sub b
61AC  02                 ld (bc),a
61AD  FB                 ei
61AE  21 2F E0           ld hl,$E02F
61B1  08                 ex af,af'
61B2  00                 nop
61B3  BF                 cp a
61B4  C8                 ret z
61B5  19                 add hl,de
61B6  36 F7              ld (hl),-9
61B8  C2 0C 7E           jp nz,$7E0C
61BB  40                 ld b,b
61BC  25                 dec h
61BD  28 40              jr z,$61FF
61BF  50                 ld d,b
61C0  98                 sbc a,b
61C1  FD 01 8F E4        ld bc,$E48F
61C5  0F                 rrca
61C6  C8                 ret z
61C7  8D                 adc a,l
61C8  43                 ld b,e
61C9  FA 03 F8           jp m,$F803
61CC  50                 ld d,b
61CD  01 42 2F           ld bc,$2F42
61D0  E4 C7 D3           call po,$D3C7
61D3  09                 add hl,bc
61D4  00                 nop
61D5  05                 dec b
61D6  70                 ld (hl),b
61D7  BE                 cp (hl)
61D8  C8                 ret z
61D9  48                 ld c,b
61DA  94                 sub h
61DB  20 00              jr nz,$61DD
61DD  2F                 cpl
61DE  CA 06 4D           jp z,$4D06
61E1  BD                 cp l
61E2  F7                 rst 30h
61E3  03                 inc bc
61E4  1F                 rra
61E5  C0                 ret nz
61E6  0A                 ld a,(bc)
61E7  E0                 ret po
61E8  03                 inc bc
61E9  EC 42 90           call pe,$9042
61EC  8C                 adc a,h
61ED  7F                 ld a,a
61EE  00                 nop
61EF  A1                 and c
61F0  31 F1 00           ld sp,$00F1
61F3  B9                 cp c
61F4  1B                 dec de
61F5  5D                 ld e,l
61F6  FA 02 2F           jp m,$2F02
61F9  D8                 ret c
61FA  0F                 rrca
61FB  C8                 ret z
61FC  0F                 rrca
61FD  E1                 pop hl
61FE  5F                 ld e,a
61FF  C2 00 0F           jp nz,$0F00
6202  E1                 pop hl
6203  1F                 rra
6204  26 3E              ld h,62
6206  98                 sbc a,b
6207  49                 ld c,c
6208  1F                 rra
6209  C0                 ret nz
620A  05                 dec b
620B  F6 42              or $42
620D  5F                 ld e,a
620E  22 00 2F           ld ($2F00),hl
6211  EA 06 4D           jp pe,$4D06
6214  BD                 cp l
6215  F6 00              or $00
6217  B9                 cp c
6218  1B                 dec de
6219  80                 add a,b
621A  0C                 inc c
621B  7C                 ld a,h
621C  C0                 ret nz
621D  98                 sbc a,b
621E  00                 nop
621F  4C                 ld c,h
6220  00                 nop
6221  1F                 rra
6222  C2 31 F0           jp nz,$F031
6225  82                 add a,d
6226  35                 dec (hl)
6227  0F                 rrca
6228  7F                 ld a,a
6229  80                 add a,b
622A  28 1F              jr z,$624B
622C  C2 31 FE           jp nz,$FE31
622F  03                 inc bc
6230  E4 C7 D3           call po,$D3C7
6233  09                 add hl,bc
6234  7C                 ld a,h
6235  80                 add a,b
6236  17                 rla
6237  D9                 exx
6238  09                 add hl,bc
6239  00                 nop
623A  14                 inc d
623B  21 05 FB           ld hl,-1275
623E  40                 ld b,b
623F  C9                 ret
6240  B7                 or a
6241  BE                 cp (hl)
6242  A0                 and b
6243  63                 ld h,e
6244  FE 02              cp 2
6246  31 FA 03           ld sp,$03FA
6249  1F                 rra
624A  70                 ld (hl),b
624B  23                 inc hl
624C  50                 ld d,b
624D  80                 add a,b
624E  1D                 dec e
624F  FE 01              cp 1
6251  30 63              jr nc,$62B6
6253  EA 07 C9           jp pe,$C907
6256  8F                 adc a,a
6257  A6                 and (hl)
6258  12                 ld (de),a
6259  C8                 ret z
625A  DA E1 7D           jp c,$7DE1
625D  90                 sub b
625E  90                 sub b
625F  3E 44              ld a,68
6261  0B                 dec bc
6262  FE 81              cp -127
6264  93                 sub e
6265  6F                 ld l,a
6266  7C                 ld a,h
6267  40                 ld b,b
6268  C1                 pop bc
6269  3E 4A              ld a,74
626B  10 80              djnz $61ED
626D  25                 dec h
626E  28 42              jr z,$62B2
6270  02                 ld (bc),a
6271  84                 add a,h
6272  C1                 pop bc
6273  29                 add hl,hl
6274  FC 03 1F           call m,$1F03
6277  C0                 ret nz
6278  0B                 dec bc
6279  96                 sub (hl)
627A  BB                 cp e
627B  F4 04 6A           call p,$6A04
627E  1F                 rra
627F  90                 sub b
6280  1D                 dec e
6281  FD 01 8B 7B        ld bc,$7B8B
6285  E2 07 C9           jp po,$C907
6288  8F                 adc a,a
6289  A6                 and (hl)
628A  12                 ld (de),a
628B  FE 00              cp 0
628D  81                 add a,c
628E  7D                 ld a,l
628F  90                 sub b
6290  91                 sub c
6291  09                 add hl,bc
6292  40                 ld b,b
6293  00                 nop
6294  5F                 ld e,a
6295  8C                 adc a,h
6296  0C                 inc c
6297  9B                 sbc a,e
6298  7B                 ld a,e
6299  EC 06 3E           call pe,$3E06
629C  14                 inc d
629D  63                 ld h,e
629E  F8                 ret m
629F  40                 ld b,b
62A0  01 8F A8           ld bc,$A88F
62A3  11 A8 7E           ld de,$7EA8
62A6  C0                 ret nz
62A7  75                 ld (hl),l
62A8  5F                 ld e,a
62A9  20 30              jr nz,$62DB
62AB  64                 ld h,h
62AC  DB DF              in a,($DF)
62AE  60                 ld h,b
62AF  3E 4C              ld a,76
62B1  7D                 ld a,l
62B2  30 97              jr nc,$624B
62B4  C8                 ret z
62B5  01 7D 90           ld bc,$907D
62B8  97                 sub a
62B9  C8                 ret z
62BA  80                 add a,b
62BB  0B                 dec bc
62BC  F9                 ld sp,hl
62BD  81                 add a,c
62BE  93                 sub e
62BF  6F                 ld l,a
62C0  7C                 ld a,h
62C1  80                 add a,b
62C2  C7                 rst 00h
62C3  DA 0C 7C           jp c,$7C0C
62C6  40                 ld b,b
62C7  8D                 adc a,l
62C8  42                 ld b,d
62C9  F4 8F F0           call p,$F08F
62CC  1D                 dec e
62CD  9F                 sbc a,a
62CE  C8                 ret z
62CF  0C                 inc c
62D0  42                 ld b,d
62D1  06 4D              ld b,77
62D3  BD                 cp l
62D4  F2 03 E4           jp p,$E403
62D7  C7                 rst 00h
62D8  D3 CC              out ($CC),a
62DA  7C                 ld a,h
62DB  80                 add a,b
62DC  17                 rla
62DD  D9                 exx
62DE  09                 add hl,bc
62DF  00                 nop
62E0  04                 inc b
62E1  A5                 and l
62E2  05                 dec b
62E3  FA C0 C9           jp m,$C9C0
62E6  B7                 or a
62E7  80                 add a,b
62E8  0C                 inc c
62E9  7E                 ld a,(hl)
62EA  A0                 and b
62EB  26 1E              ld h,30
62ED  01 FA 02           ld bc,762
62F0  35                 dec (hl)
62F1  08                 ex af,af'
62F2  1E 31              ld e,49
62F4  AB                 xor e
62F5  FC 80 20           call m,$2080
62F8  3B                 dec sp
62F9  FA 03 1F           jp m,$1F03
62FC  24                 inc h
62FD  0C                 inc c
62FE  9B                 sbc a,e
62FF  7F                 ld a,a
6300  00                 nop
6301  F9                 ld sp,hl
6302  31 FE 80           ld sp,$80FE
6305  5F                 ld e,a
6306  64                 ld h,h
6307  24                 inc h
6308  0F                 rrca
6309  91                 sub c
630A  02                 ld (bc),a
630B  FF                 rst 38h
630C  60                 ld h,b
630D  64                 ld h,h
630E  80                 add a,b
630F  18 FE              jr $630F
6311  41                 ld b,c
6312  FC 23 0B           call m,$0B23
6315  00                 nop
6316  1F                 rra
6317  91                 sub c
6318  1A                 ld a,(de)
6319  87                 add a,a
631A  F8                 ret m
631B  01 F5 BF           ld bc,$BFF5
631E  E0                 ret po
631F  08                 ex af,af'
6320  47                 ld b,a
6321  F8                 ret m
6322  AC                 xor h
6323  7D                 ld a,l
6324  90                 sub b
6325  37                 scf
6326  F0                 ret p
6327  0F                 rrca
6328  93                 sub e
6329  1F                 rra
632A  20 28              jr nz,$6354
632C  4F                 ld c,a
632D  E0                 ret po
632E  14                 inc d
632F  20 01              jr nz,$6332
6331  4A                 ld c,d
6332  12                 ld (de),a
6333  84                 add a,h
6334  FD 00              nop
6336  2F                 cpl
6337  B2                 or d
6338  12                 ld (de),a
6339  25                 dec h
633A  3F                 ccf
633B  80                 add a,b
633C  0B                 dec bc
633D  E3                 ex (sp),hl
633E  84                 add a,h
633F  80                 add a,b
6340  0C                 inc c
6341  7D                 ld a,l
6342  80                 add a,b
6343  20 00              jr nz,$6345
6345  4F                 ld c,a
6346  98                 sbc a,b
6347  18 7F              jr $63C8
6349  03                 inc bc
634A  F8                 ret m
634B  01 D0 FF           ld bc,-48
634E  00                 nop
634F  A1                 and c
6350  28 F0              jr z,$6342
6352  42                 ld b,d
6353  7F                 ld a,a
6354  00                 nop
6355  14                 inc d
6356  3F                 ccf
6357  C4 63 EC           call nz,$EC63
635A  46                 ld b,(hl)
635B  00                 nop
635C  05                 dec b
635D  7C                 ld a,h
635E  98                 sbc a,b
635F  FE 01              cp 1
6361  FC 20 00           call m,$0020
6364  FE 10              cp 16
6366  00                 nop
6367  7D                 ld a,l
6368  88                 adc a,b
6369  F9                 ld sp,hl
636A  14                 inc d
636B  2F                 cpl
636C  B2                 or d
636D  12                 ld (de),a
636E  FE 10              cp 16
6370  00                 nop
6371  0B                 dec bc
6372  E3                 ex (sp),hl
6373  84                 add a,h
6374  80                 add a,b
6375  0C                 inc c
6376  7F                 ld a,a
6377  00                 nop
6378  F8                 ret m
6379  91                 sub c
637A  F5                 push af
637B  00                 nop
637C  98                 sbc a,b
637D  5A                 ld e,d
637E  C1                 pop bc
637F  DF                 rst 18h
6380  90                 sub b
6381  1F                 rra
6382  A2                 and d
6383  3F                 ccf
6384  85                 add a,l
6385  0A                 ld a,(bc)
6386  10 97              djnz $631F
6388  FC 0F D3           call m,$D30F
638B  1F                 rra
638C  50                 ld d,b
638D  26 00              ld h,0
638F  7F                 ld a,a
6390  58                 ld e,b
6391  FB                 ei
6392  39                 add hl,sp
6393  8F                 adc a,a
6394  90                 sub b
6395  14                 inc d
6396  17                 rla
6397  C7                 rst 00h
6398  09                 add hl,bc
6399  00                 nop
639A  18 FA              jr $6396
639C  81                 add a,c
639D  84                 add a,h
639E  7F                 ld a,a
639F  80                 add a,b
63A0  FE 11              cp 17
63A2  87                 add a,a
63A3  9C                 sbc a,h
63A4  2C                 inc l
63A5  88                 adc a,b
63A6  F9                 ld sp,hl
63A7  00                 nop
63A8  2F                 cpl
63A9  AA                 xor d
63AA  12                 ld (de),a
63AB  FE 01              cp 1
63AD  42                 ld b,d
63AE  52                 ld d,d
63AF  84                 add a,h
63B0  07                 rlca
63B1  C9                 ret
63B2  42                 ld b,d
63B3  53                 ld d,e
63B4  F4 C7 D8           call p,$D8C7
63B7  02                 ld (bc),a
63B8  53                 ld d,e
63B9  F4 06 00           call p,$0006
63BC  09                 add hl,bc
63BD  FD 40              ld b,b
63BF  98                 sbc a,b
63C0  01 8F E0           ld bc,$E08F
63C3  08                 ex af,af'
63C4  40                 ld b,b
63C5  BE                 cp (hl)
63C6  38 48              jr c,$6410
63C8  00                 nop
63C9  C7                 rst 00h
63CA  F0                 ret p
63CB  02                 ld (bc),a
63CC  7D                 ld a,l
63CD  80                 add a,b
63CE  C1                 pop bc
63CF  3E A0              ld a,-96
63D1  61                 ld h,c
63D2  6B                 ld l,e
63D3  42                 ld b,d
63D4  28 00              jr z,$63D6
63D6  50                 ld d,b
63D7  50                 ld d,b
63D8  97                 sub a
63D9  C5                 push bc
63DA  CC 03 EE           call z,$EE03
63DD  47                 ld b,a
63DE  E9                 jp (hl)
63DF  8F                 adc a,a
63E0  B0                 or b
63E1  08                 ex af,af'
63E2  47                 ld b,a
63E3  E8                 ret pe
63E4  0C                 inc c
63E5  7C                 ld a,h
63E6  20 C5              jr nz,$63AD
63E8  3E 40              ld a,64
63EA  63                 ld h,e
63EB  E3                 ex (sp),hl
63EC  E6 00              and $00
63EE  0C                 inc c
63EF  00                 nop
63F0  1F                 rra
63F1  62                 ld h,d
63F2  00                 nop
63F3  0C                 inc c
63F4  23                 inc hl
63F5  FC 01 00           call m,$0001
63F8  0C                 inc c
63F9  7E                 ld a,(hl)
63FA  88                 adc a,b
63FB  00                 nop
63FC  10 28              djnz $6426
63FE  4B                 ld c,e
63FF  EA 06 3E           jp pe,$3E06
6402  E0                 ret po
6403  7C                 ld a,h
6404  98                 sbc a,b
6405  FD 81              add a,c
6407  8F                 adc a,a
6408  E0                 ret po
6409  04                 inc b
640A  00                 nop
640B  09                 add hl,bc
640C  FA 00 9F           jp m,$9F00
640F  C0                 ret nz
6410  30 84              jr nc,$6396
6412  00                 nop
6413  04                 inc b
6414  C7                 rst 00h
6415  F6 0C              or $0C
6417  23                 inc hl
6418  F4 06 3F           call p,$3F06
641B  80                 add a,b
641C  60                 ld h,b
641D  9F                 sbc a,a
641E  E0                 ret po
641F  3F                 ccf
6420  84                 add a,h
6421  63                 ld h,e
6422  E2 00 A1           jp po,$A100
6425  2F                 cpl
6426  E0                 ret po
6427  04                 inc b
6428  FD 00              nop
642A  4C                 ld c,h
642B  7D                 ld a,l
642C  C0                 ret nz
642D  F9                 ld sp,hl
642E  31 FC 00           ld sp,$00FC
6431  94                 sub h
6432  FE 01              cp 1
6434  42                 ld b,d
6435  50                 ld d,b
6436  00                 nop
6437  C7                 rst 00h
6438  F0                 ret p
6439  04                 inc b
643A  00                 nop
643B  08                 ex af,af'
643C  00                 nop
643D  10 00              djnz $643F
643F  23                 inc hl
6440  F8                 ret m
6441  06 3E              ld b,62
6443  40                 ld b,b
6444  23                 inc hl
6445  1F                 rra
6446  A0                 and b
6447  3E 4A              ld a,74
6449  10 80              djnz $63CB
644B  20 29              jr nz,$6476
644D  F2 03 F8           jp p,$F803
6450  A1                 and c
6451  08                 ex af,af'
6452  0A                 ld a,(bc)
6453  53                 ld d,e
6454  1F                 rra
6455  60                 ld h,b
6456  31 FC 03           ld sp,$03FC
6459  08                 ex af,af'
645A  42                 ld b,d
645B  90                 sub b
645C  00                 nop
645D  13                 inc de
645E  E4 06 3F           call po,$3F06
6461  C0                 ret nz
6462  20 50              jr nz,$64B4
6464  90                 sub b
6465  3F                 ccf
6466  24                 inc h
6467  63                 ld h,e
6468  F6 02              or $02
646A  3E 4C              ld a,76
646C  7F                 ld a,a
646D  00                 nop
646E  FE 10              cp 16
6470  00                 nop
6471  7F                 ld a,a
6472  08                 ex af,af'
6473  52                 ld d,d
6474  B1                 or c
6475  F6 83              or $83
6477  1F                 rra
6478  20 3F              jr nz,$64B9
647A  D4 63 F8           call nc,$F863
647D  47                 ld b,a
647E  F0                 ret p
647F  0C                 inc c
6480  7F                 ld a,a
6481  00                 nop
6482  C7                 rst 00h
6483  E8                 ret pe
6484  04                 inc b
6485  23                 inc hl
6486  E4 06 3F           call po,$3F06
6489  C0                 ret nz
648A  50                 ld d,b
648B  50                 ld d,b
648C  90                 sub b
648D  01 3F F0           ld bc,$F03F
6490  18 FB              jr $648D
6492  81                 add a,c
6493  F2 63 FA           jp p,$FA63
6496  07                 rlca
6497  F1                 pop af
6498  8F                 adc a,a
6499  E0                 ret po
649A  08                 ex af,af'
649B  FC 40 8C           call m,$8C40
649E  7F                 ld a,a
649F  E0                 ret po
64A0  00                 nop
64A1  11 8F E0           ld de,$E08F
64A4  18 FE              jr $64A4
64A6  99                 sbc a,c
64A7  8F                 adc a,a
64A8  91                 sub c
64A9  88                 adc a,b
64AA  00                 nop
64AB  10 28              djnz $64D5
64AD  4B                 ld c,e
64AE  EA 06 3F           jp pe,$3F06
64B1  60                 ld h,b
64B2  23                 inc hl
64B3  E4 C7 D8           call po,$D8C7
64B6  02                 ld (bc),a
64B7  53                 ld d,e
64B8  FB                 ei
64B9  01 31 FF           ld bc,-207
64BC  83                 add a,e
64BD  F8                 ret m
64BE  06 3F              ld b,63
64C0  80                 add a,b
64C1  63                 ld h,e
64C2  FA 46 3E           jp m,$3E46
64C5  44                 ld b,h
64C6  63                 ld h,e
64C7  F8                 ret m
64C8  00                 nop
64C9  A1                 and c
64CA  24                 inc h
64CB  5B                 ld e,e
64CC  DF                 rst 18h
64CD  E0                 ret po
64CE  31 F7 03           ld sp,$03F7
64D1  E4 C7 D8           call po,$D8C7
64D4  04                 inc b
64D5  23                 inc hl
64D6  F2 02 3F           jp p,$3F02
64D9  10 23              djnz $64FE
64DB  00                 nop
64DC  05                 dec b
64DD  09                 add hl,bc
64DE  42                 ld b,d
64DF  7C                 ld a,h
64E0  80                 add a,b
64E1  A1                 and c
64E2  3F                 ccf
64E3  80                 add a,b
64E4  10 00              djnz $64E6
64E6  A1                 and c
64E7  3E 40              ld a,64
64E9  50                 ld d,b
64EA  94                 sub h
64EB  A0                 and b
64EC  00                 nop
64ED  80                 add a,b
64EE  01 18 FE           ld bc,-488
64F1  81                 add a,c
64F2  3F                 ccf
64F3  90                 sub b
64F4  13                 inc de
64F5  FB                 ei
64F6  00                 nop
64F7  81                 add a,c
64F8  42                 ld b,d
64F9  48                 ld c,b
64FA  32 6D EF           ld ($EF6D),a
64FD  E4 08 F9           call po,$F908
6500  31 42 12           ld sp,$1242
6503  80                 add a,b
6504  25                 dec h
6505  3F                 ccf
6506  80                 add a,b
6507  20 00              jr nz,$6509
6509  A1                 and c
650A  3F                 ccf
650B  E0                 ret po
650C  12                 ld (de),a
650D  94                 sub h
650E  20 01              jr nz,$6511
6510  85                 add a,l
6511  7F                 ld a,a
6512  88                 adc a,b
6513  00                 nop
6514  3F                 ccf
6515  84                 add a,h
6516  00                 nop
6517  08                 ex af,af'
6518  00                 nop
6519  3F                 ccf
651A  84                 add a,h
651B  00                 nop
651C  1F                 rra
651D  A2                 and d
651E  3E C0              ld a,-64
6520  62                 ld h,d
6521  9F                 sbc a,a
6522  A0                 and b
6523  29                 add hl,hl
6524  F2 02 9F           jp p,$9F02
6527  20 29              jr nz,$6552
6529  4F                 ld c,a
652A  90                 sub b
652B  14                 inc d
652C  20 A1              jr nz,$64CF
652E  24                 inc h
652F  40                 ld b,b
6530  C6 4D              add a,77
6532  BD                 cp l
6533  FC 13 FE           call m,$FE13
6536  07                 rlca
6537  C9                 ret
6538  8F                 adc a,a
6539  F1                 pop af
653A  1F                 rra
653B  63                 ld h,e
653C  10 8F              djnz $64CD
653E  B9                 cp c
653F  8A                 adc a,d
6540  F9                 ld sp,hl
6541  10 C6              djnz $6509
6543  63                 ld h,e
6544  FF                 rst 38h
6545  67                 ld h,a
6546  F0                 ret p
6547  8C                 adc a,h
6548  7F                 ld a,a
6549  08                 ex af,af'
654A  63                 ld h,e
654B  3F                 ccf
654C  84                 add a,h
654D  31 9F C2           ld sp,$C29F
6550  18 CF              jr $6521
6552  91                 sub c
6553  0C                 inc c
6554  67                 ld h,a
6555  F0                 ret p
6556  81                 add a,c
6557  42                 ld b,d
6558  5F                 ld e,a
6559  9A                 sbc a,d
655A  3F                 ccf
655B  8C                 adc a,h
655C  7E                 ld a,(hl)
655D  00                 nop
655E  80                 add a,b
655F  69                 ld l,c
6560  02                 ld (bc),a
6561  55                 ld d,l
6562  63                 ld h,e
6563  5A                 ld e,d
6564  6C                 ld l,h
6565  5B                 ld e,e
6566  67                 ld h,a
6567  66                 ld h,(hl)
6568  6A                 ld l,d
6569  57                 ld d,a
656A  56                 ld d,(hl)
656B  53                 ld d,e
656C  68                 ld l,b
656D  58                 ld e,b
656E  59                 ld e,c
656F  34                 inc (hl)
6570  38 37              jr c,$65A9
6572  64                 ld h,h
6573  35                 dec (hl)
6574  11 6B 36           ld de,$366B
6577  65                 ld h,l
6578  54                 ld d,h
6579  60                 ld h,b
657A  5D                 ld e,l
657B  85                 add a,l
657C  42                 ld b,d
657D  D7                 rst 10h
657E  F3                 di
657F  02                 ld (bc),a
6580  C0                 ret nz
6581  57                 ld d,a
6582  CE 83              adc a,-125
6584  1B                 dec de
6585  FA 02 3F           jp m,$3F02
6588  F0                 ret p
6589  01 16 B7           ld bc,$B716
658C  FE 08              cp 8
658E  A5                 and l
658F  1F                 rra
6590  E0                 ret po
6591  23                 inc hl
6592  51                 ld d,c
6593  E5                 push hl
6594  3F                 ccf
6595  20 20              jr nz,$65B7
6597  22 42 04           ld ($0442),hl
659A  84                 add a,h
659B  09                 add hl,bc
659C  08                 ex af,af'
659D  12                 ld (de),a
659E  10 24              djnz $65C4
65A0  20 48              jr nz,$65EA
65A2  3A 08 F9           ld a,($F908)
65A5  40                 ld b,b
65A6  8F                 adc a,a
65A7  FC 00 47           call m,$4700
65AA  FE 0F              cp 15
65AC  90                 sub b
65AD  08                 ex af,af'
65AE  F8                 ret m
65AF  41                 ld b,c
65B0  01 10 52           ld bc,$5210
65B3  24                 inc h
65B4  84                 add a,h
65B5  41                 ld b,c
65B6  48                 ld c,b
65B7  85                 add a,l
65B8  11 04 52           ld de,$5204
65BB  4A                 ld c,d
65BC  72                 ld (hl),d
65BD  40                 ld b,b
65BE  F2 9F 84           jp p,$849F
65C1  11 46 53           ld de,$5346
65C4  F7                 rst 30h
65C5  81                 add a,c
65C6  1F                 rra
65C7  C0                 ret nz
65C8  2D                 dec l
65C9  FB                 ei
65CA  03                 inc bc
65CB  C1                 pop bc
65CC  F9                 ld sp,hl
65CD  81                 add a,c
65CE  E0                 ret po
65CF  80                 add a,b
65D0  08                 ex af,af'
65D1  F8                 ret m
65D2  41                 ld b,c
65D3  01 10 52           ld bc,$5210
65D6  20 A4              jr nz,$657C
65D8  49                 ld c,c
65D9  17                 rla
65DA  E2 C4 A4           jp po,$A4C4
65DD  8F                 adc a,a
65DE  94                 sub h
65DF  08                 ex af,af'
65E0  FF                 rst 38h
65E1  C0                 ret nz
65E2  04                 inc b
65E3  7C                 ld a,h
65E4  20 F2              jr nz,$65D8
65E6  10 1F              djnz $6607
65E8  A8                 xor b
65E9  1E 0F              ld e,15
65EB  82                 add a,d
65EC  04                 inc b
65ED  7D                 ld a,l
65EE  C0                 ret nz
65EF  8D                 adc a,l
65F0  60                 ld h,b
65F1  22 08 44           ld ($4408),hl
65F4  11 48 2C           ld de,$2C48
65F7  21 12 24           ld hl,$2412
65FA  A4                 and h
65FB  42                 ld b,d
65FC  CA 48 FE           jp z,$FE48
65FF  41                 ld b,c
6600  64                 ld h,h
6601  5B                 ld e,e
6602  F9                 ld sp,hl
6603  05                 dec b
6604  80                 add a,b
6605  0B                 dec bc
6606  58                 ld e,b
6607  16 B0              ld d,-80
6609  01 6B 22           ld bc,$226B
660C  D6 FD              sub -3
660E  81                 add a,c
660F  E4 18 C8           call po,$C818
6612  08                 ex af,af'
6613  01 1A E5           ld bc,$E51A
6616  20 07              jr nz,$661F
6618  BE                 cp (hl)
6619  44                 ld b,h
661A  05                 dec b
661B  1F                 rra
661C  08                 ex af,af'
661D  20 22              jr nz,$6641
661F  08                 ex af,af'
6620  44                 ld b,h
6621  9C                 sbc a,h
6622  89                 adc a,c
6623  84                 add a,h
6624  12                 ld (de),a
6625  52                 ld d,d
6626  24                 inc h
6627  A4                 and h
6628  4F                 ld c,a
6629  E2 48 FC           jp po,$FC48
662C  81                 add a,c
662D  FC 5B E2           call m,$E25B
6630  02                 ld (bc),a
6631  3F                 ccf
6632  F0                 ret p
6633  01 1F 70           ld bc,$701F
6636  3C                 inc a
6637  83                 add a,e
6638  40                 ld b,b
6639  86                 add a,(hl)
663A  40                 ld b,b
663B  7F                 ld a,a
663C  C0                 ret nz
663D  23                 inc hl
663E  E4 82 10           call po,$1082
6641  00                 nop
6642  2D                 dec l
6643  60                 ld h,b
6644  02                 ld (bc),a
6645  D6 00              sub 0
6647  2D                 dec l
6648  60                 ld h,b
6649  07                 rlca
664A  F5                 push af
664B  84                 add a,h
664C  20 85              jr nz,$65D3
664E  97                 sub a
664F  D9                 exx
6650  8D                 adc a,l
6651  23                 inc hl
6652  E9                 jp (hl)
6653  05                 dec b
6654  AC                 xor h
6655  8B                 adc a,e
6656  5B                 ld e,e
6657  FA 07 F1           jp m,$F107
665A  60                 ld h,b
665B  03                 inc bc
665C  1F                 rra
665D  B0                 or b
665E  11 FC 02           ld de,764
6661  DF                 rst 18h
6662  90                 sub b
6663  3C                 inc a
6664  83                 add a,e
6665  42                 ld b,d
6666  04                 inc b
6667  81                 add a,c
6668  90                 sub b
6669  1F                 rra
666A  B0                 or b
666B  08                 ex af,af'
666C  F9                 ld sp,hl
666D  20 84              jr nz,$65F3
666F  7C                 ld a,h
6670  20 42              jr nz,$66B4
6672  3E CA              ld a,-54
6674  7F                 ld a,a
6675  08                 ex af,af'
6676  24                 inc h
6677  A4                 and h
6678  49                 ld c,c
6679  48                 ld c,b
667A  92                 sub d
667B  70                 ld (hl),b
667C  91                 sub c
667D  FF                 rst 38h
667E  80                 add a,b
667F  00                 nop
6680  8D                 adc a,l
6681  72                 ld (hl),d
6682  90                 sub b
6683  01 1F B0           ld bc,$B01F
6686  11 6B 7C           ld de,$7C6B
6689  80                 add a,b
668A  8E                 adc a,(hl)
668B  53                 ld d,e
668C  FC 03 C8           call m,$C803
668F  31 21 02           ld sp,$0221
6692  42                 ld b,d
6693  06 43              ld b,67
6695  C2 BE 40           jp nz,$40BE
6698  59                 ld e,c
6699  1F                 rra
669A  24                 inc h
669B  10 8F              djnz $662C
669D  C4 08 FF           call nz,$FF08
66A0  28 84              jr z,$6626
66A2  12                 ld (de),a
66A3  53                 ld d,e
66A4  24                 inc h
66A5  84                 add a,h
66A6  41                 ld b,c
66A7  16 48              ld d,72
66A9  FF                 rst 38h
66AA  C1                 pop bc
66AB  FC 02 DF           call m,$DF02
66AE  20 11              jr nz,$66C1
66B0  FB                 ei
66B1  01 1F D0           ld bc,$D01F
66B4  3C                 inc a
66B5  83                 add a,e
66B6  40                 ld b,b
66B7  90                 sub b
66B8  81                 add a,c
66B9  21 02 1A           ld hl,$1A02
66BC  0F                 rrca
66BD  0A                 ld a,(bc)
66BE  F9                 ld sp,hl
66BF  00                 nop
66C0  8F                 adc a,a
66C1  92                 sub d
66C2  08                 ex af,af'
66C3  47                 ld b,a
66C4  E2 04 7C           jp po,$7C04
66C7  54                 ld d,h
66C8  47                 ld b,a
66C9  E9                 jp (hl)
66CA  8D                 adc a,l
66CB  63                 ld h,e
66CC  12                 ld (de),a
66CD  12                 ld (de),a
66CE  3E C0              ld a,-64
66D0  46                 ld b,(hl)
66D1  A3                 and e
66D2  CA 7F 40           jp z,$407F
66D5  63                 ld h,e
66D6  1A                 ld a,(de)
66D7  D6 35              sub 53
66D9  AC                 xor h
66DA  63                 ld h,e
66DB  EE 02              xor $02
66DD  3F                 ccf
66DE  60                 ld h,b
66DF  23                 inc hl
66E0  EA 07 90           jp pe,$9007
66E3  68                 ld l,b
66E4  40                 ld b,b
66E5  90                 sub b
66E6  81                 add a,c
66E7  21 02 42           ld hl,$4202
66EA  04                 inc b
66EB  78                 ld a,b
66EC  50                 ld d,b
66ED  2D                 dec l
66EE  6C                 ld l,h
66EF  7C                 ld a,h
66F0  90                 sub b
66F1  42                 ld b,d
66F2  3E 40              ld a,64
66F4  46                 ld b,(hl)
66F5  A9                 xor c
66F6  FA 80 8F           jp m,$8F80
66F9  AA                 xor d
66FA  88                 adc a,b
66FB  81                 add a,c
66FC  21 02 48           ld hl,$4802
66FF  52                 ld d,d
6700  47                 ld b,a
6701  E8                 ret pe
6702  0B                 dec bc
6703  5B                 ld e,e
6704  FB                 ei
6705  06 3E              ld b,62
6707  40                 ld b,b
6708  23                 inc hl
6709  FC 04 6A           call m,$6A04
670C  3C                 inc a
670D  A4                 and h
670E  08                 ex af,af'
670F  F9                 ld sp,hl
6710  01 6F 90           ld bc,$906F
6713  1E 41              ld e,65
6715  89                 adc a,c
6716  08                 ex af,af'
6717  62                 ld h,d
6718  10 24              djnz $673E
671A  20 48              jr nz,$6764
671C  40                 ld b,b
671D  90                 sub b
671E  78                 ld a,b
671F  50                 ld d,b
6720  00                 nop
6721  8F                 adc a,a
6722  93                 sub e
6723  08                 ex af,af'
6724  47                 ld b,a
6725  E2 0F F1           jp po,$F10F
6728  1F                 rra
6729  25                 dec h
672A  11 02 42           ld de,$4202
672D  04                 inc b
672E  FE 24              cp 36
6730  80                 add a,b
6731  02                 ld (bc),a
6732  D6 FD              sub -3
6734  C1                 pop bc
6735  84                 add a,h
6736  7C                 ld a,h
6737  80                 add a,b
6738  47                 ld b,a
6739  EC 04 02           call pe,$0204
673C  DF                 rst 18h
673D  60                 ld h,b
673E  3C                 inc a
673F  83                 add a,e
6740  40                 ld b,b
6741  90                 sub b
6742  82                 add a,d
6743  21 0F EB           ld hl,$EB0F
6746  08                 ex af,af'
6747  22 21 04           ld ($0421),hl
674A  23                 inc hl
674B  F1                 pop af
674C  02                 ld (bc),a
674D  BF                 cp a
674E  48                 ld c,b
674F  23                 inc hl
6750  E4 A2 20           call po,$20A2
6753  47                 ld b,a
6754  40                 ld b,b
6755  9F                 sbc a,a
6756  C6 11              add a,17
6758  6F                 ld l,a
6759  EC 11 4F           call pe,$4F11
675C  C8                 ret z
675D  04                 inc b
675E  7C                 ld a,h
675F  80                 add a,b
6760  47                 ld b,a
6761  EC 04 7F           call pe,$7F04
6764  80                 add a,b
6765  F2 0D 08           jp p,$080D
6768  12                 ld (de),a
6769  10 44              djnz $67AF
676B  20 87              jr nz,$66F4
676D  40                 ld b,b
676E  8E                 adc a,(hl)
676F  81                 add a,c
6770  1D                 dec e
6771  02                 ld (bc),a
6772  3A 04 44           ld a,($4404)
6775  20 84              jr nz,$66FB
6777  7F                 ld a,a
6778  C0                 ret nz
6779  8D                 adc a,l
677A  7E                 ld a,(hl)
677B  C8                 ret z
677C  23                 inc hl
677D  E4 A2 20           call po,$20A2
6780  48                 ld c,b
6781  40                 ld b,b
6782  92                 sub d
6783  94                 sub h
6784  D1                 pop de
6785  F4 82 39           call p,$3982
6788  4F                 ld c,a
6789  D4 0F E3           call nc,$E30F
678C  08                 ex af,af'
678D  F9                 ld sp,hl
678E  00                 nop
678F  80                 add a,b
6790  46                 ld b,(hl)
6791  B9                 cp c
6792  4F                 ld c,a
6793  C4 04 7F           call nz,$7F04
6796  00                 nop
6797  B0                 or b
6798  01 E4 18           ld bc,$18E4
679B  90                 sub b
679C  81                 add a,c
679D  31 04 42           ld sp,$4204
67A0  08                 ex af,af'
67A1  84                 add a,h
67A2  09                 add hl,bc
67A3  08                 ex af,af'
67A4  12                 ld (de),a
67A5  10 24              djnz $67CB
67A7  20 44              jr nz,$67ED
67A9  42                 ld b,d
67AA  08                 ex af,af'
67AB  46                 ld b,(hl)
67AC  3E 10              ld a,16
67AE  41                 ld b,c
67AF  DF                 rst 18h
67B0  C4 30 8F           call nz,$8F30
67B3  92                 sub d
67B4  88                 adc a,b
67B5  81                 add a,c
67B6  21 02 08           ld hl,$0802
67B9  A2                 and d
67BA  45                 ld b,l
67BB  BF                 cp a
67BC  80                 add a,b
67BD  45                 ld b,l
67BE  31 CA 7E           ld sp,$7ECA
67C1  30 45              jr nc,$6808
67C3  03                 inc bc
67C4  CA 40 8F           jp z,$8F40
67C7  D8                 ret c
67C8  08                 ex af,af'
67C9  FD 01 E4 1A        ld bc,$1AE4
67CD  04                 inc b
67CE  84                 add a,h
67CF  08                 ex af,af'
67D0  88                 adc a,b
67D1  22 10 24           ld ($2410),hl
67D4  20 48              jr nz,$681E
67D6  63                 ld h,e
67D7  10 86              djnz $675F
67D9  21 0C 62           ld hl,$620C
67DC  10 42              djnz $6820
67DE  3E C0              ld a,-64
67E0  46                 ld b,(hl)
67E1  BF                 cp a
67E2  C0                 ret nz
67E3  F2 9F 90           jp p,$909F
67E6  1F                 rra
67E7  64                 ld h,h
67E8  11 F2 51           ld de,$51F2
67EB  1F                 rra
67EC  A6                 and (hl)
67ED  3F                 ccf
67EE  88                 adc a,b
67EF  23                 inc hl
67F0  FF                 rst 38h
67F1  07                 rlca
67F2  F0                 ret p
67F3  04                 inc b
67F4  7C                 ld a,h
67F5  80                 add a,b
67F6  47                 ld b,a
67F7  EC 04 02           call pe,$0204
67FA  C0                 ret nz
67FB  07                 rlca
67FC  90                 sub b
67FD  68                 ld l,b
67FE  40                 ld b,b
67FF  90                 sub b
6800  86                 add a,(hl)
6801  11 04 42           ld de,$4204
6804  04                 inc b
6805  87                 add a,a
6806  F1                 pop af
6807  88                 adc a,b
6808  12                 ld (de),a
6809  10 24              djnz $682F
680B  20 44              jr nz,$6851
680D  42                 ld b,d
680E  08                 ex af,af'
680F  47                 ld b,a
6810  E2 0F D2           jp po,$D20F
6813  18 47              jr $685C
6815  E5                 push hl
6816  44                 ld b,h
6817  7F                 ld a,a
6818  10 45              djnz $685F
681A  BF                 cp a
681B  30 78              jr nc,$6895
681D  3F                 ccf
681E  A0                 and b
681F  2C                 inc l
6820  8F                 adc a,a
6821  90                 sub b
6822  08                 ex af,af'
6823  F9                 ld sp,hl
6824  01 E0 FF           ld bc,-32
6827  80                 add a,b
6828  40                 ld b,b
6829  40                 ld b,b
682A  0F                 rrca
682B  20 C4              jr nz,$67F1
682D  84                 add a,h
682E  09                 add hl,bc
682F  FC 41 10           call m,$1041
6832  44                 ld b,h
6833  20 48              jr nz,$687D
6835  40                 ld b,b
6836  90                 sub b
6837  81                 add a,c
6838  21 02 42           ld hl,$4202
683B  04                 inc b
683C  44                 ld b,h
683D  20 84              jr nz,$67C3
683F  7E                 ld a,(hl)
6840  20 EF              jr nz,$6831
6842  E9                 jp (hl)
6843  04                 inc b
6844  7E                 ld a,(hl)
6845  54                 ld d,h
6846  47                 ld b,a
6847  F1                 pop af
6848  04                 inc b
6849  7D                 ld a,l
684A  C0                 ret nz
684B  8D                 adc a,l
684C  47                 ld b,a
684D  94                 sub h
684E  FD 80              add a,b
6850  F5                 push af
6851  38 1F              jr c,$6872
6853  90                 sub b
6854  16 47              ld d,71
6856  C8                 ret z
6857  04                 inc b
6858  7F                 ld a,a
6859  00                 nop
685A  F5                 push af
685B  38 1F              jr c,$687C
685D  D0                 ret nc
685E  1B                 dec de
685F  42                 ld b,d
6860  03                 inc bc
6861  E4 1A 04           call po,$041A
6864  84                 add a,h
6865  09                 add hl,bc
6866  0C                 inc c
6867  61                 ld h,c
6868  10 24              djnz $688E
686A  21 FB 61           ld hl,$61FB
686D  10 82              djnz $67F1
686F  11 8F E0           ld de,$E08F
6872  1F                 rra
6873  B6                 or (hl)
6874  11 FA 61           ld de,$61FA
6877  1F                 rra
6878  95                 sub l
6879  11 FC 41           ld de,$41FC
687C  00                 nop
687D  06 DF              ld b,-33
687F  0C                 inc c
6880  F4 6D BE           call p,$BE6D
6883  19                 add hl,de
6884  E8                 ret pe
6885  DF                 rst 18h
6886  A0                 and b
6887  36 F8              ld (hl),-8
6889  67                 ld h,a
688A  A3                 and e
688B  7C                 ld a,h
688C  40                 ld b,b
688D  F5                 push af
688E  29                 add hl,hl
688F  CE 07              adc a,7
6891  F8                 ret m
6892  07                 rlca
6893  90                 sub b
6894  1F                 rra
6895  E0                 ret po
6896  08                 ex af,af'
6897  00                 nop
6898  3D                 dec a
6899  4A                 ld c,d
689A  73                 ld (hl),e
689B  81                 add a,c
689C  FE 01              cp 1
689E  E0                 ret po
689F  9C                 sbc a,h
68A0  C8                 ret z
68A1  41                 ld b,c
68A2  A1                 and c
68A3  02                 ld (bc),a
68A4  42                 ld b,d
68A5  04                 inc b
68A6  FE 20              cp 32
68A8  88                 adc a,b
68A9  12                 ld (de),a
68AA  10 44              djnz $68F0
68AC  20 48              jr nz,$68F6
68AE  40                 ld b,b
68AF  90                 sub b
68B0  81                 add a,c
68B1  21 02 7F           ld hl,$7F02
68B4  10 42              djnz $68F8
68B6  3F                 ccf
68B7  10 23              djnz $68DC
68B9  FE A2              cp -94
68BB  20 E8              jr nz,$68A5
68BD  21 E1 00           ld hl,225
68C0  3B                 dec sp
68C1  DE 9F              sbc a,-97
68C3  F4 77 BD           call p,$BD77
68C6  3F                 ccf
68C7  E9                 jp (hl)
68C8  BE                 cp (hl)
68C9  0B                 dec bc
68CA  40                 ld b,b
68CB  77                 ld (hl),a
68CC  BD                 cp l
68CD  3F                 ccf
68CE  E8                 ret pe
68CF  0D                 dec c
68D0  6F                 ld l,a
68D1  82                 add a,d
68D2  D6 FB              sub -5
68D4  F8                 ret m
68D5  A7                 and a
68D6  F1                 pop af
68D7  C0                 ret nz
68D8  B7                 or a
68D9  D8                 ret c
68DA  47                 ld b,a
68DB  03                 inc bc
68DC  E0                 ret po
68DD  A1                 and c
68DE  BE                 cp (hl)
68DF  FE 29              cp 41
68E1  FC 70 2D           call m,$2D70
68E4  F5                 push af
68E5  38 1F              jr c,$6906
68E7  42                 ld b,d
68E8  04                 inc b
68E9  84                 add a,h
68EA  09                 add hl,bc
68EB  08                 ex af,af'
68EC  12                 ld (de),a
68ED  18 C2              jr $68B1
68EF  20 48              jr nz,$6939
68F1  41                 ld b,c
68F2  10 71              djnz $6965
68F4  21 07 42           ld hl,$4207
68F7  04                 inc b
68F8  74                 ld (hl),h
68F9  09                 add hl,bc
68FA  07                 rlca
68FB  41                 ld b,c
68FC  08                 ex af,af'
68FD  FC 40 8F           call m,$8F40
6900  B2                 or d
6901  9F                 sbc a,a
6902  52                 ld d,d
6903  3F                 ccf
6904  88                 adc a,b
6905  23                 inc hl
6906  F8                 ret m
6907  E3                 ex (sp),hl
6908  CE B9              adc a,-71
690A  C7                 rst 00h
690B  9D                 sbc a,l
690C  7F                 ld a,a
690D  9C                 sbc a,h
690E  79                 ld a,c
690F  D7                 rst 10h
6910  F3                 di
6911  C4 7C 9C           call nz,$9C7C
6914  47                 ld b,a
6915  ED CF              ???
6917  12                 ld (de),a
6918  10 24              djnz $693E
691A  20 48              jr nz,$6964
691C  40                 ld b,b
691D  9F                 sbc a,a
691E  C4 11 0C           call nz,$0C11
6921  42                 ld b,d
6922  08                 ex af,af'
6923  84                 add a,h
6924  09                 add hl,bc
6925  08                 ex af,af'
6926  12                 ld (de),a
6927  10 24              djnz $694D
6929  20 4F              jr nz,$697A
692B  E2 08 40           jp po,$4008
692E  01 1A 8F           ld bc,$8F1A
6931  29                 add hl,hl
6932  FA 80 8F           jp m,$8F80
6935  B2                 or d
6936  88                 adc a,b
6937  FC 81 FC           call m,$FC81
693A  41                 ld b,c
693B  1F                 rra
693C  26 3F              ld h,63
693E  85                 add a,l
693F  7C                 ld a,h
6940  F8                 ret m
6941  F9                 ld sp,hl
6942  21 FE E3           ld hl,$E3FE
6945  5A                 ld e,d
6946  C6 B4              add a,-76
6948  8D                 adc a,l
6949  23                 inc hl
694A  18 43              jr $698F
694C  E0                 ret po
694D  6F                 ld l,a
694E  87                 add a,a
694F  C0                 ret nz
6950  DF                 rst 18h
6951  0F                 rrca
6952  81                 add a,c
6953  BE                 cp (hl)
6954  1F                 rra
6955  03                 inc bc
6956  7F                 ld a,a
6957  93                 sub e
6958  08                 ex af,af'
6959  47                 ld b,a
695A  E2 04 52           jp po,$5204
695D  9F                 sbc a,a
695E  A2                 and d
695F  3F                 ccf
6960  20 7F              jr nz,$69E1
6962  10 44              djnz $69A8
6964  09                 add hl,bc
6965  04                 inc b
6966  40                 ld b,b
6967  8A                 adc a,d
6968  44                 ld b,h
6969  09                 add hl,bc
696A  04                 inc b
696B  40                 ld b,b
696C  90                 sub b
696D  47                 ld b,a
696E  F1                 pop af
696F  44                 ld b,h
6970  40                 ld b,b
6971  90                 sub b
6972  44                 ld b,h
6973  09                 add hl,bc
6974  04                 inc b
6975  7D                 ld a,l
6976  94                 sub h
6977  47                 ld b,a
6978  C9                 ret
6979  04                 inc b
697A  7E                 ld a,(hl)
697B  94                 sub h
697C  FE 11              cp 17
697E  08                 ex af,af'
697F  12                 ld (de),a
6980  08                 ex af,af'
6981  84                 add a,h
6982  09                 add hl,bc
6983  08                 ex af,af'
6984  22 10 24           ld ($2410),hl
6987  20 48              jr nz,$69D1
6989  40                 ld b,b
698A  90                 sub b
698B  81                 add a,c
698C  20 88              jr nz,$6916
698E  41                 ld b,c
698F  10 82              djnz $6913
6991  21 04 42           ld hl,$4204
6994  08                 ex af,af'
6995  FE 20              cp 32
6997  84                 add a,h
6998  7E                 ld a,(hl)
6999  20 45              jr nz,$69E0
699B  11 FE 03           ld de,1022
699E  FC C7 F1           call m,$F1C7
69A1  04                 inc b
69A2  40                 ld b,b
69A3  90                 sub b
69A4  44                 ld b,h
69A5  09                 add hl,bc
69A6  04                 inc b
69A7  40                 ld b,b
69A8  90                 sub b
69A9  44                 ld b,h
69AA  09                 add hl,bc
69AB  04                 inc b
69AC  7F                 ld a,a
69AD  14                 inc d
69AE  44                 ld b,h
69AF  09                 add hl,bc
69B0  04                 inc b
69B1  40                 ld b,b
69B2  90                 sub b
69B3  47                 ld b,a
69B4  D9                 exx
69B5  44                 ld b,h
69B6  7C                 ld a,h
69B7  8E                 adc a,(hl)
69B8  47                 ld b,a
69B9  C9                 ret
69BA  4F                 ld c,a
69BB  E1                 pop hl
69BC  04                 inc b
69BD  83                 add a,e
69BE  89                 adc a,c
69BF  04                 inc b
69C0  42                 ld b,d
69C1  0E 84              ld c,-124
69C3  10 21              djnz $69E6
69C5  10 42              djnz $6A09
69C7  20 A4              jr nz,$696D
69C9  49                 ld c,c
69CA  08                 ex af,af'
69CB  82                 add a,d
69CC  44                 ld b,h
69CD  20 88              jr nz,$6957
69CF  41                 ld b,c
69D0  10 82              djnz $6954
69D2  21 04 42           ld hl,$4204
69D5  18 42              jr $6A19
69D7  3F                 ccf
69D8  20 47              jr nz,$6A21
69DA  29                 add hl,hl
69DB  FD 00              nop
69DD  8A                 adc a,d
69DE  23                 inc hl
69DF  F6 02              or $02
69E1  3F                 ccf
69E2  88                 adc a,b
69E3  22 04 82           ld ($8204),hl
69E6  20 45              jr nz,$6A2D
69E8  22 04 82           ld ($8204),hl
69EB  20 48              jr nz,$6A35
69ED  23                 inc hl
69EE  F8                 ret m
69EF  A2                 and d
69F0  20 48              jr nz,$6A3A
69F2  22 04 82           ld ($8204),hl
69F5  3E CA              ld a,-54
69F7  23                 inc hl
69F8  E4 82 3F           call po,$3F82
69FB  8A                 adc a,d
69FC  21 10 81           ld hl,$8110
69FF  21 02 41           ld hl,$4102
6A02  10 81              djnz $6985
6A04  21 04 4A           ld hl,$4A04
6A07  44                 ld b,h
6A08  94                 sub h
6A09  89                 adc a,c
6A0A  21 10 52           ld hl,$5210
6A0D  24                 inc h
6A0E  91                 sub c
6A0F  08                 ex af,af'
6A10  22 10 44           ld ($4410),hl
6A13  20 88              jr nz,$699D
6A15  40                 ld b,b
6A16  9F                 sbc a,a
6A17  C4 10 8F           call nz,$8F10
6A1A  C4 08 A2           call nz,$A208
6A1D  3F                 ccf
6A1E  C0                 ret nz
6A1F  63                 ld h,e
6A20  F4 02 3F           call p,$3F02
6A23  88                 adc a,b
6A24  22 04 82           ld ($8204),hl
6A27  20 48              jr nz,$6A71
6A29  22 04 82           ld ($8204),hl
6A2C  20 48              jr nz,$6A76
6A2E  23                 inc hl
6A2F  F8                 ret m
6A30  A2                 and d
6A31  20 45              jr nz,$6A78
6A33  22 04 52           ld ($5204),hl
6A36  3E CA              ld a,-54
6A38  23                 inc hl
6A39  E4 82 29           call po,$2982
6A3C  44                 ld b,h
6A3D  20 90              jr nz,$69CF
6A3F  87                 add a,a
6A40  E8                 ret pe
6A41  A4                 and h
6A42  42                 ld b,d
6A43  18 84              jr $69C9
6A45  11 08 76           ld de,$7608
6A48  10 EC              djnz $6A36
6A4A  21 D8 40           ld hl,$40D8
6A4D  90                 sub b
6A4E  44                 ld b,h
6A4F  20 88              jr nz,$69D9
6A51  41                 ld b,c
6A52  10 82              djnz $69D6
6A54  21 0C 42           ld hl,$420C
6A57  18 42              jr $6A9B
6A59  3F                 ccf
6A5A  10 22              djnz $6A7E
6A5C  88                 adc a,b
6A5D  FD 80              add a,b
6A5F  88                 adc a,b
6A60  3A 08 81           ld a,($8108)
6A63  14                 inc d
6A64  88                 adc a,b
6A65  11 48 81           ld de,$8148
6A68  14                 inc d
6A69  88                 adc a,b
6A6A  13                 inc de
6A6B  A8                 xor b
6A6C  FE 28              cp 40
6A6E  88                 adc a,b
6A6F  12                 ld (de),a
6A70  08                 ex af,af'
6A71  81                 add a,c
6A72  20 8F              jr nz,$6A03
6A74  B2                 or d
6A75  88                 adc a,b
6A76  F9                 ld sp,hl
6A77  20 8A              jr nz,$6A03
6A79  21 10 24           ld hl,$2410
6A7C  3A 48 38           ld a,($3848)
6A7F  90                 sub b
6A80  44                 ld b,h
6A81  20 88              jr nz,$6A0B
6A83  41                 ld b,c
6A84  0E 82              ld c,-126
6A86  1D                 dec e
6A87  04                 inc b
6A88  3A 08 74           ld a,($7408)
6A8B  14                 inc d
6A8C  A4                 and h
6A8D  6B                 ld l,e
6A8E  48                 ld c,b
6A8F  84                 add a,h
6A90  11 AD 23           ld de,$23AD
6A93  5A                 ld e,d
6A94  47                 ld b,a
6A95  F1                 pop af
6A96  04                 inc b
6A97  23                 inc hl
6A98  1F                 rra
6A99  C0                 ret nz
6A9A  3F                 ccf
6A9B  2C                 inc l
6A9C  7F                 ld a,a
6A9D  00                 nop
6A9E  C2 28 8F           jp nz,$8F28
6AA1  F0                 ret p
6AA2  18 FD              jr $6AA1
6AA4  00                 nop
6AA5  8F                 adc a,a
6AA6  E2 08 81           jp po,$8108
6AA9  20 88              jr nz,$6A33
6AAB  12                 ld (de),a
6AAC  08                 ex af,af'
6AAD  81                 add a,c
6AAE  20 88              jr nz,$6A38
6AB0  12                 ld (de),a
6AB1  08                 ex af,af'
6AB2  FE 28              cp 40
6AB4  88                 adc a,b
6AB5  11 48 81           ld de,$8148
6AB8  14                 inc d
6AB9  8F                 adc a,a
6ABA  B2                 or d
6ABB  88                 adc a,b
6ABC  F9                 ld sp,hl
6ABD  34                 inc (hl)
6ABE  8A                 adc a,d
6ABF  22 10 C4           ld ($C410),hl
6AC2  20 48              jr nz,$6B0C
6AC4  40                 ld b,b
6AC5  90                 sub b
6AC6  46                 ld b,(hl)
6AC7  20 8C              jr nz,$6A55
6AC9  41                 ld b,c
6ACA  10 82              djnz $6A4E
6ACC  21 04 42           ld hl,$4204
6ACF  08                 ex af,af'
6AD0  F9                 ld sp,hl
6AD1  20 88              jr nz,$6A5B
6AD3  41                 ld b,c
6AD4  D0                 ret nc
6AD5  83                 add a,e
6AD6  A1                 and c
6AD7  07                 rlca
6AD8  42                 ld b,d
6AD9  08                 ex af,af'
6ADA  FE 20              cp 32
6ADC  84                 add a,h
6ADD  7E                 ld a,(hl)
6ADE  20 45              jr nz,$6B25
6AE0  11 F2 03           ld de,1010
6AE3  F2 C2 3F           jp p,$3FC2
6AE6  88                 adc a,b
6AE7  22 04 52           ld ($5204),hl
6AEA  20 45              jr nz,$6B31
6AEC  22 04 52           ld ($5204),hl
6AEF  20 4E              jr nz,$6B3F
6AF1  A3                 and e
6AF2  F8                 ret m
6AF3  A2                 and d
6AF4  20 48              jr nz,$6B3E
6AF6  22 04 82           ld ($8204),hl
6AF9  3E CA              ld a,-54
6AFB  7D                 ld a,l
6AFC  88                 adc a,b
6AFD  A2                 and d
6AFE  21 82 41           ld hl,$4182
6B01  C4 54 09           call nz,$0954
6B04  84                 add a,h
6B05  2A 1F 22           ld hl,($221F)
6B08  A1                 and c
6B09  84                 add a,h
6B0A  43                 ld b,e
6B0B  08                 ex af,af'
6B0C  86                 add a,(hl)
6B0D  10 AC              djnz $6ABB
6B0F  29                 add hl,hl
6B10  48                 ld c,b
6B11  84                 add a,h
6B12  09                 add hl,bc
6B13  08                 ex af,af'
6B14  12                 ld (de),a
6B15  10 24              djnz $6B3B
6B17  20 8F              jr nz,$6AA8
6B19  E2 08 47           jp po,$4708
6B1C  E2 04 51           jp po,$5104
6B1F  1F                 rra
6B20  C6 00              add a,0
6B22  88                 adc a,b
6B23  40                 ld b,b
6B24  90                 sub b
6B25  81                 add a,c
6B26  21 02 7F           ld hl,$7F02
6B29  10 44              djnz $6B6F
6B2B  09                 add hl,bc
6B2C  04                 inc b
6B2D  40                 ld b,b
6B2E  90                 sub b
6B2F  44                 ld b,h
6B30  09                 add hl,bc
6B31  04                 inc b
6B32  40                 ld b,b
6B33  90                 sub b
6B34  47                 ld b,a
6B35  F1                 pop af
6B36  44                 ld b,h
6B37  40                 ld b,b
6B38  8A                 adc a,d
6B39  44                 ld b,h
6B3A  08                 ex af,af'
6B3B  A4                 and h
6B3C  7E                 ld a,(hl)
6B3D  D4 44 10           call nc,$1044
6B40  48                 ld c,b
6B41  40                 ld b,b
6B42  90                 sub b
6B43  87                 add a,a
6B44  50                 ld d,b
6B45  48                 ld c,b
6B46  60                 ld h,b
6B47  90                 sub b
6B48  C1                 pop bc
6B49  20 82              jr nz,$6ACD
6B4B  41                 ld b,c
6B4C  04                 inc b
6B4D  82                 add a,d
6B4E  08                 ex af,af'
6B4F  A4                 and h
6B50  7E                 ld a,(hl)
6B51  90                 sub b
6B52  24                 inc h
6B53  20 48              jr nz,$6B9D
6B55  40                 ld b,b
6B56  90                 sub b
6B57  82                 add a,d
6B58  3F                 ccf
6B59  88                 adc a,b
6B5A  21 1F 88           ld hl,$881F
6B5D  11 44 7C           ld de,$7C44
6B60  80                 add a,b
6B61  44                 ld b,h
6B62  20 48              jr nz,$6BAC
6B64  40                 ld b,b
6B65  90                 sub b
6B66  81                 add a,c
6B67  3F                 ccf
6B68  88                 adc a,b
6B69  22 04 52           ld ($5204),hl
6B6C  20 45              jr nz,$6BB3
6B6E  22 04 52           ld ($5204),hl
6B71  20 4E              jr nz,$6BC1
6B73  A3                 and e
6B74  F8                 ret m
6B75  A2                 and d
6B76  20 48              jr nz,$6BC0
6B78  22 04 82           ld ($8204),hl
6B7B  3F                 ccf
6B7C  6A                 ld l,d
6B7D  22 08 24           ld ($2408),hl
6B80  20 48              jr nz,$6BCA
6B82  40                 ld b,b
6B83  88                 adc a,b
6B84  24                 inc h
6B85  20 48              jr nz,$6BCF
6B87  40                 ld b,b
6B88  90                 sub b
6B89  41                 ld b,c
6B8A  21 02 42           ld hl,$4202
6B8D  04                 inc b
6B8E  82                 add a,d
6B8F  3E EC              ld a,-20
6B91  23                 inc hl
6B92  F8                 ret m
6B93  D2 10 08           jp nc,$0810
6B96  D7                 rst 10h
6B97  29                 add hl,hl
6B98  F9                 ld sp,hl
6B99  80                 add a,b
6B9A  8A                 adc a,d
6B9B  23                 inc hl
6B9C  F8                 ret m
6B9D  06 11              ld b,17
6B9F  08                 ex af,af'
6BA0  7C                 ld a,h
6BA1  58                 ld e,b
6BA2  84                 add a,h
6BA3  11 02 41           ld de,$4102
6BA6  10 24              djnz $6BCC
6BA8  11 02 41           ld de,$4102
6BAB  10 24              djnz $6BD1
6BAD  11 FC 51           ld de,$51FC
6BB0  10 22              djnz $6BD4
6BB2  91                 sub c
6BB3  02                 ld (bc),a
6BB4  29                 add hl,hl
6BB5  1F                 rra
6BB6  B5                 or l
6BB7  11 04 12           ld de,$1204
6BBA  10 27              djnz $6BE3
6BBC  F1                 pop af
6BBD  84                 add a,h
6BBE  2B                 dec hl
6BBF  0E 26              ld c,38
6BC1  14                 inc d
6BC2  4C                 ld c,h
6BC3  28 98              jr z,$6B5D
6BC5  52                 ld d,d
6BC6  B0                 or b
6BC7  A2                 and d
6BC8  41                 ld b,c
6BC9  1F                 rra
6BCA  4B                 ld c,e
6BCB  10 8F              djnz $6B5C
6BCD  C4 08 A2           call nz,$A208
6BD0  3E 40              ld a,64
6BD2  22 10 24           ld ($2410),hl
6BD5  20 48              jr nz,$6C1F
6BD7  40                 ld b,b
6BD8  88                 adc a,b
6BD9  84                 add a,h
6BDA  11 02 29           ld de,$2902
6BDD  50                 ld d,b
6BDE  22 95 02           ld ($0295),hl
6BE1  29                 add hl,hl
6BE2  50                 ld d,b
6BE3  27                 daa
6BE4  7F                 ld a,a
6BE5  44                 ld b,h
6BE6  40                 ld b,b
6BE7  90                 sub b
6BE8  44                 ld b,h
6BE9  09                 add hl,bc
6BEA  0F                 rrca
6BEB  E9                 jp (hl)
6BEC  1F                 rra
6BED  C5                 push bc
6BEE  10 E4              djnz $6BD4
6BF0  7C                 ld a,h
6BF1  98                 sbc a,b
6BF2  2C                 inc l
6BF3  91                 sub c
6BF4  02                 ld (bc),a
6BF5  22 04 44           ld ($4404),hl
6BF8  08                 ex af,af'
6BF9  E8                 ret pe
6BFA  12                 ld (de),a
6BFB  10 23              djnz $6C20
6BFD  A0                 and b
6BFE  48                 ld c,b
6BFF  23                 inc hl
6C00  E9                 jp (hl)
6C01  62                 ld h,d
6C02  11 F5 02           ld de,757
6C05  39                 add hl,sp
6C06  4F                 ld c,a
6C07  C8                 ret z
6C08  04                 inc b
6C09  51                 ld d,c
6C0A  1F                 rra
6C0B  C0                 ret nz
6C0C  30 88              jr nc,$6B96
6C0E  40                 ld b,b
6C0F  90                 sub b
6C10  81                 add a,c
6C11  21 02 22           ld hl,$2202
6C14  10 44              djnz $6C5A
6C16  09                 add hl,bc
6C17  FC 40 9F           call m,$9F40
6C1A  C4 09 FC           call nz,$FC09
6C1D  40                 ld b,b
6C1E  90                 sub b
6C1F  74                 ld (hl),h
6C20  09                 add hl,bc
6C21  07                 rlca
6C22  40                 ld b,b
6C23  8A                 adc a,d
6C24  54                 ld d,h
6C25  08                 ex af,af'
6C26  A5                 and l
6C27  40                 ld b,b
6C28  9F                 sbc a,a
6C29  C0                 ret nz
6C2A  88                 adc a,b
6C2B  27                 daa
6C2C  11 1F C5           ld de,$C51F
6C2F  11 04 10           ld de,$1004
6C32  42                 ld b,d
6C33  40                 ld b,b
6C34  A6                 and (hl)
6C35  8C                 adc a,h
6C36  12                 ld (de),a
6C37  10 22              djnz $6C5B
6C39  20 48              jr nz,$6C83
6C3B  40                 ld b,b
6C3C  90                 sub b
6C3D  81                 add a,c
6C3E  21 02 61           ld hl,$6102
6C41  1F                 rra
6C42  4B                 ld c,e
6C43  10 8F              djnz $6BD4
6C45  C4 08 A2           call nz,$A208
6C48  3E 40              ld a,64
6C4A  22 10 24           ld ($2410),hl
6C4D  1C                 inc e
6C4E  48                 ld c,b
6C4F  38 88              jr c,$6BD9
6C51  D6 91              sub -111
6C53  F7                 rst 30h
6C54  63                 ld h,e
6C55  F8                 ret m
6C56  81                 add a,c
6C57  3F                 ccf
6C58  88                 adc a,b
6C59  13                 inc de
6C5A  F8                 ret m
6C5B  81                 add a,c
6C5C  3F                 ccf
6C5D  88                 adc a,b
6C5E  13                 inc de
6C5F  F8                 ret m
6C60  11 04 42           ld de,$4204
6C63  23                 inc hl
6C64  F8                 ret m
6C65  A2                 and d
6C66  20 82              jr nz,$6BEA
6C68  0B                 dec bc
6C69  E4 C2 3E           call po,$3EC2
6C6C  93                 sub e
6C6D  23                 inc hl
6C6E  F8                 ret m
6C6F  66                 ld h,(hl)
6C70  18 CC              jr $6C3E
6C72  31 98 63           ld sp,$6398
6C75  30 C6              jr nc,$6C3D
6C77  61                 ld h,c
6C78  8C                 adc a,h
6C79  42                 ld b,d
6C7A  2D                 dec l
6C7B  FF                 rst 38h
6C7C  02                 ld (bc),a
6C7D  C8                 ret z
6C7E  A2                 and d
6C7F  3F                 ccf
6C80  80                 add a,b
6C81  61                 ld h,c
6C82  10 81              djnz $6C05
6C84  21 02 42           ld hl,$4202
6C87  04                 inc b
6C88  F9                 ld sp,hl
6C89  10 27              djnz $6CB2
6C8B  08                 ex af,af'
6C8C  82                 add a,d
6C8D  70                 ld (hl),b
6C8E  88                 adc a,b
6C8F  27                 daa
6C90  48                 ld c,b
6C91  92                 sub d
6C92  47                 ld b,a
6C93  CB 81              res 0,c
6C95  10 88              djnz $6C1F
6C97  FE 28              cp 40
6C99  88                 adc a,b
6C9A  20 85              jr nz,$6C21
6C9C  91                 sub c
6C9D  0B                 dec bc
6C9E  22 7D 2C           ld ($2C7D),hl
6CA1  41                 ld b,c
6CA2  98                 sbc a,b
6CA3  C3 31 86           jp $8631
6CA6  63                 ld h,e
6CA7  0C                 inc c
6CA8  C6 19              add a,25
6CAA  8C                 adc a,h
6CAB  33                 inc sp
6CAC  0C                 inc c
6CAD  84                 add a,h
6CAE  7F                 ld a,a
6CAF  00                 nop
6CB0  DB E1              in a,($E1)
6CB2  9E                 sbc a,(hl)
6CB3  8D                 adc a,l
6CB4  FC 03 C1           call m,$C103
6CB7  FB                 ei
6CB8  00                 nop
6CB9  8A                 adc a,d
6CBA  23                 inc hl
6CBB  E4 02 21           call po,$2102
6CBE  02                 ld (bc),a
6CBF  42                 ld b,d
6CC0  04                 inc b
6CC1  84                 add a,h
6CC2  08                 ex af,af'
6CC3  29                 add hl,hl
6CC4  08                 ex af,af'
6CC5  82                 add a,d
6CC6  90                 sub b
6CC7  88                 adc a,b
6CC8  29                 add hl,hl
6CC9  08                 ex af,af'
6CCA  82                 add a,d
6CCB  90                 sub b
6CCC  88                 adc a,b
6CCD  27                 daa
6CCE  08                 ex af,af'
6CCF  82                 add a,d
6CD0  70                 ld (hl),b
6CD1  88                 adc a,b
6CD2  27                 daa
6CD3  08                 ex af,af'
6CD4  82                 add a,d
6CD5  74                 ld (hl),h
6CD6  88                 adc a,b
6CD7  27                 daa
6CD8  08                 ex af,af'
6CD9  9F                 sbc a,a
6CDA  C0                 ret nz
6CDB  89                 adc a,c
6CDC  84                 add a,h
6CDD  7F                 ld a,a
6CDE  14                 inc d
6CDF  44                 ld b,h
6CE0  10 49              djnz $6D2B
6CE2  4C                 ld c,h
6CE3  92                 sub d
6CE4  99                 sbc a,c
6CE5  3E 96              ld a,-106
6CE7  23                 inc hl
6CE8  E9                 jp (hl)
6CE9  62                 ld h,d
6CEA  11 E0 BC           ld de,$BCE0
6CED  20 EF              jr nz,$6CDE
6CEF  7A                 ld a,d
6CF0  7F                 ld a,a
6CF1  D0                 ret nc
6CF2  1B                 dec de
6CF3  EA 70 2F           jp pe,$2F70
6CF6  08                 ex af,af'
6CF7  68                 ld l,b
6CF8  1B                 dec de
6CF9  64                 ld h,h
6CFA  51                 ld d,c
6CFB  00                 nop
6CFC  06 30              ld b,48
6CFE  88                 adc a,b
6CFF  40                 ld b,b
6D00  90                 sub b
6D01  81                 add a,c
6D02  21 02 7F           ld hl,$7F02
6D05  02                 ld (bc),a
6D06  27                 daa
6D07  F0                 ret p
6D08  22 7F 02           ld ($027F),hl
6D0B  24                 inc h
6D0C  84                 add a,h
6D0D  22 48 52           ld ($5248),hl
6D10  27                 daa
6D11  F0                 ret p
6D12  22 48 52           ld ($5248),hl
6D15  24                 inc h
6D16  84                 add a,h
6D17  22 7F 02           ld ($027F),hl
6D1A  27                 daa
6D1B  F0                 ret p
6D1C  22 11 1F           ld ($1F11),hl
6D1F  C5                 push bc
6D20  11 04 12           ld de,$1204
6D23  45                 ld b,l
6D24  21 64 46           ld hl,$4664
6D27  33                 inc sp
6D28  0C                 inc c
6D29  67                 ld h,a
6D2A  F1                 pop af
6D2B  8F                 adc a,a
6D2C  89                 adc a,c
6D2D  98                 sbc a,b
6D2E  62                 ld h,d
6D2F  0D                 dec c
6D30  F0                 ret p
6D31  B0                 or b
6D32  C8                 ret z
6D33  47                 ld b,a
6D34  C9                 ret
6D35  C7                 rst 00h
6D36  9D                 sbc a,l
6D37  7C                 ld a,h
6D38  DC 45 10           call c,$1045
6D3B  00                 nop
6D3C  20 48              jr nz,$6D86
6D3E  FE 31              cp 49
6D40  AD                 xor l
6D41  7C                 ld a,h
6D42  98                 sbc a,b
6D43  D6 BF              sub -65
6D45  CC 6B 5F           call z,$5F6B
6D48  4E                 ld c,(hl)
6D49  35                 dec (hl)
6D4A  AF                 xor a
6D4B  93                 sub e
6D4C  1A                 ld a,(de)
6D4D  D7                 rst 10h
6D4E  C9                 ret
6D4F  84                 add a,h
6D50  7F                 ld a,a
6D51  14                 inc d
6D52  44                 ld b,h
6D53  10 49              djnz $6D9E
6D55  C8                 ret z
6D56  92                 sub d
6D57  91                 sub c
6D58  3E A6              ld a,-90
6D5A  7C                 ld a,h
6D5B  98                 sbc a,b
6D5C  66                 ld h,(hl)
6D5D  11 F2 B1           ld de,$B1F2
6D60  08                 ex af,af'
6D61  FC 70 8A           call m,$8A70
6D64  23                 inc hl
6D65  18 42              jr $6DA9
6D67  3F                 ccf
6D68  FE 7C              cp 124
6D6A  BC                 cp h
6D6B  47                 ld b,a
6D6C  F1                 pop af
6D6D  44                 ld b,h
6D6E  69                 ld l,c
6D6F  1F                 rra
6D70  96                 sub (hl)
6D71  3E ED              ld a,-19
6D73  23                 inc hl
6D74  5F                 ld e,a
6D75  66                 ld h,(hl)
6D76  35                 dec (hl)
6D77  AF                 xor a
6D78  D3 1A              out ($1A),a
6D7A  D7                 rst 10h
6D7B  C9                 ret
6D7C  84                 add a,h
6D7D  7E                 ld a,(hl)
6D7E  00                 nop
6D7F  6C                 ld l,h
6D80  80                 add a,b
6D81  02                 ld (bc),a
6D82  6D                 ld l,l
6D83  5B                 ld e,e
6D84  17                 rla
6D85  6E                 ld l,(hl)
6D86  70                 ld (hl),b
6D87  6F                 ld l,a
6D88  5A                 ld e,d
6D89  10 01              djnz $6D8C
6D8B  38 34              jr c,$6DC1
6D8D  37                 scf
6D8E  36 78              ld (hl),120
6D90  76                 halt
6D91  73                 ld (hl),e
6D92  35                 dec (hl)
6D93  77                 ld (hl),a
6D94  75                 ld (hl),l
6D95  74                 ld (hl),h
6D96  72                 ld (hl),d
6D97  71                 ld (hl),c
6D98  59                 ld e,c
6D99  58                 ld e,b
6D9A  56                 ld d,(hl)
6D9B  11 91 47           ld de,$4791
6D9E  FF                 rst 38h
6D9F  0F                 rrca
6DA0  FE 1F              cp 31
6DA2  FC 3F 88           call m,$883F
6DA5  7F                 ld a,a
6DA6  E8                 ret pe
6DA7  47                 ld b,a
6DA8  FB                 ei
6DA9  0B                 dec bc
6DAA  65                 ld h,l
6DAB  C6 FF              add a,-1
6DAD  E1                 pop hl
6DAE  FA 42 D9           jp m,$D942
6DB1  71                 ld (hl),c
6DB2  BF                 cp a
6DB3  F8                 ret m
6DB4  7F                 ld a,a
6DB5  10 47              djnz $6DFE
6DB7  E2 04 7E           jp po,$7E04
6DBA  80                 add a,b
6DBB  FE 11              cp 17
6DBD  FF                 rst 38h
6DBE  C3 FF 87           jp $87FF
6DC1  FF                 rst 38h
6DC2  0F                 rrca
6DC3  D2 1F AE           jp nc,$AE1F
6DC6  11 FC 60           ld de,$60FC
6DC9  00                 nop
6DCA  FE 11              cp 17
6DCC  FC 42 D9           call m,$D942
6DCF  F3                 di
6DD0  8D                 adc a,l
6DD1  F4 C3 EC           call p,$ECC3
6DD4  C4 21 6C           call nz,$6C21
6DD7  F9                 ld sp,hl
6DD8  C6 FF              add a,-1
6DDA  21 67 7B           ld hl,$7B67
6DDD  C6 F8              add a,-8
6DDF  A1                 and c
6DE0  63                 ld h,e
6DE1  7F                 ld a,a
6DE2  F0                 ret p
6DE3  84                 add a,h
6DE4  2C                 inc l
6DE5  6F                 ld l,a
6DE6  D2 08 FF           jp nc,$FF08
6DE9  41                 ld b,c
6DEA  FC 23 EB           call m,$EB23
6DED  82                 add a,d
6DEE  3E 42              ld a,66
6DF0  23                 inc hl
6DF1  FF                 rst 38h
6DF2  87                 add a,a
6DF3  D3 0B              out ($0B),a
6DF5  67                 ld h,a
6DF6  CE 37              adc a,55
6DF8  FD 04              inc b
6DFA  7F                 ld a,a
6DFB  C0                 ret nz
6DFC  F8                 ret m
6DFD  B1                 or c
6DFE  FC 23 EE           call m,$EE23
6E01  85                 add a,l
6E02  9C                 sbc a,h
6E03  6F                 ld l,a
6E04  8A                 adc a,d
6E05  1A                 ld a,(de)
6E06  47                 ld b,a
6E07  C8                 ret z
6E08  44                 ld b,h
6E09  7F                 ld a,a
6E0A  F0                 ret p
6E0B  FF                 rst 38h
6E0C  E1                 pop hl
6E0D  F2 43 EE           jp p,$EE43
6E10  C0                 ret nz
6E11  01 FC 61           ld bc,$61FC
6E14  1F                 rra
6E15  E0                 ret po
6E16  3F                 ccf
6E17  84                 add a,h
6E18  7D                 ld a,l
6E19  90                 sub b
6E1A  1F                 rra
6E1B  52                 ld d,d
6E1C  3E A7              ld a,-89
6E1E  F4 80 FA           call p,$FA80
6E21  9F                 sbc a,a
6E22  B2                 or d
6E23  08                 ex af,af'
6E24  F9                 ld sp,hl
6E25  08                 ex af,af'
6E26  8F                 adc a,a
6E27  FE 1F              cp 31
6E29  FC 3F C8           call m,$C83F
6E2C  23                 inc hl
6E2D  E2 02 3F           jp po,$3F02
6E30  80                 add a,b
6E31  20 1F              jr nz,$6E52
6E33  E6 00              and $00
6E35  0F                 rrca
6E36  E1                 pop hl
6E37  1F                 rra
6E38  64                 ld h,h
6E39  1B                 dec de
6E3A  58                 ld e,b
6E3B  B6                 or (hl)
6E3C  B1                 or c
6E3D  FD 20 DA           jr nz,$6E1A
6E40  C7                 rst 00h
6E41  F4 86 91           call p,$9186
6E44  F2 11 1F           jp p,$1F11
6E47  AC                 xor h
6E48  3E 2C              ld a,44
6E4A  73                 ld (hl),e
6E4B  98                 sbc a,b
6E4C  C7                 rst 00h
6E4D  3E 4C              ld a,76
6E4F  73                 ld (hl),e
6E50  9F                 sbc a,a
6E51  C6 3E              add a,62
6E53  B8                 cp b
6E54  23                 inc hl
6E55  E2 02 00           jp po,$0002
6E58  0C                 inc c
6E59  20 00              jr nz,$6E5B
6E5B  47                 ld b,a
6E5C  C8                 ret z
6E5D  04                 inc b
6E5E  00                 nop
6E5F  1F                 rra
6E60  C2 3F 41           jp nz,$413F
6E63  28 A0              jr z,$6E05
6E65  09                 add hl,bc
6E66  40                 ld b,b
6E67  15                 dec d
6E68  7F                 ld a,a
6E69  02                 ld (bc),a
6E6A  51                 ld d,c
6E6B  40                 ld b,b
6E6C  15                 dec d
6E6D  7C                 ld a,h
6E6E  82                 add a,d
6E6F  52                 ld d,d
6E70  3E 42              ld a,66
6E72  21 5F A0           ld hl,$A05F
6E75  94                 sub h
6E76  73                 ld (hl),e
6E77  AB                 xor e
6E78  F2 12 8F           jp p,$8F12
6E7B  84                 add a,h
6E7C  31 80 47           ld sp,$4780
6E7F  F8                 ret m
6E80  44                 ld b,h
6E81  00                 nop
6E82  08                 ex af,af'
6E83  F9                 ld sp,hl
6E84  08                 ex af,af'
6E85  88                 adc a,b
6E86  0F                 rrca
6E87  A9                 xor c
6E88  FC A0 3E           call m,$3EA0
6E8B  A6                 and (hl)
6E8C  03                 inc bc
6E8D  EA 7F 88           jp pe,$887F
6E90  7C                 ld a,h
6E91  D8                 ret c
6E92  47                 ld b,a
6E93  F0                 ret p
6E94  04                 inc b
6E95  03                 inc bc
6E96  08                 ex af,af'
6E97  F9                 ld sp,hl
6E98  00                 nop
6E99  80                 add a,b
6E9A  03                 inc bc
6E9B  F8                 ret m
6E9C  47                 ld b,a
6E9D  D7                 rst 10h
6E9E  84                 add a,h
6E9F  7C                 ld a,h
6EA0  84                 add a,h
6EA1  47                 ld b,a
6EA2  F1                 pop af
6EA3  8E                 adc a,(hl)
6EA4  73                 ld (hl),e
6EA5  1F                 rra
6EA6  F5                 push af
6EA7  30 86              jr nc,$6E2F
6EA9  30 08              jr nc,$6EB3
6EAB  FF                 rst 38h
6EAC  08                 ex af,af'
6EAD  8C                 adc a,h
6EAE  01 1F 21           ld bc,$211F
6EB1  11 06 D6           ld de,$D606
6EB4  3F                 ccf
6EB5  94                 sub h
6EB6  1B                 dec de
6EB7  58                 ld e,b
6EB8  C1                 pop bc
6EB9  B5                 or l
6EBA  8F                 adc a,a
6EBB  D9                 exx
6EBC  04                 inc b
6EBD  7D                 ld a,l
6EBE  80                 add a,b
6EBF  40                 ld b,b
6EC0  10 0C              djnz $6ECE
6EC2  61                 ld h,c
6EC3  00                 nop
6EC4  02                 ld (bc),a
6EC5  3F                 ccf
6EC6  8C                 adc a,h
6EC7  01 00 07           ld bc,1792
6ECA  F0                 ret p
6ECB  8F                 adc a,a
6ECC  AC                 xor h
6ECD  88                 adc a,b
6ECE  FA 88 8F           jp m,$8F88
6ED1  EA 9F A6           jp pe,$A69F
6ED4  10 C6              djnz $6E9C
6ED6  01 1F E1           ld bc,$E11F
6ED9  10 00              djnz $6EDB
6EDB  23                 inc hl
6EDC  E4 22 14           call po,$1422
6EDF  50                 ld d,b
6EE0  05                 dec b
6EE1  42                 ld b,d
6EE2  12                 ld (de),a
6EE3  BE                 cp (hl)
6EE4  41                 ld b,c
6EE5  28 A0              jr z,$6E87
6EE7  0A                 ld a,(bc)
6EE8  8A                 adc a,d
6EE9  00                 nop
6EEA  A8                 xor b
6EEB  42                 ld b,d
6EEC  51                 ld d,c
6EED  E0                 ret po
6EEE  8F                 adc a,a
6EEF  B0                 or b
6EF0  08                 ex af,af'
6EF1  C2 00 8F           jp nz,$8F00
6EF4  E0                 ret po
6EF5  18 47              jr $6F3E
6EF7  C8                 ret z
6EF8  04                 inc b
6EF9  00                 nop
6EFA  1F                 rra
6EFB  7A                 ld a,d
6EFC  3E A2              ld a,-94
6EFE  23                 inc hl
6EFF  FA A2 3F           jp m,$3FA2
6F02  40                 ld b,b
6F03  31 80 FC           ld sp,$FC80
6F06  91                 sub c
6F07  80                 add a,b
6F08  7D                 ld a,l
6F09  88                 adc a,b
6F0A  F9                 ld sp,hl
6F0B  31 CE 7C           ld sp,$7CCE
6F0E  98                 sbc a,b
6F0F  E7                 rst 20h
6F10  3F                 ccf
6F11  EC 00 18           call pe,$1800
6F14  C0                 ret nz
6F15  00                 nop
6F16  80                 add a,b
6F17  20 08              jr nz,$6F21
6F19  FE 30              cp 48
6F1B  04                 inc b
6F1C  7F                 ld a,a
6F1D  18 02              jr $6F21
6F1F  00                 nop
6F20  0F                 rrca
6F21  E1                 pop hl
6F22  1F                 rra
6F23  59                 ld e,c
6F24  3E 64              ld a,100
6F26  63                 ld h,e
6F27  EA A2 31           jp pe,$31A2
6F2A  8F                 adc a,a
6F2B  E0                 ret po
6F2C  0C                 inc c
6F2D  67                 ld h,a
6F2E  CA 04 7D           jp z,$7D04
6F31  24                 inc h
6F32  47                 ld b,a
6F33  C8                 ret z
6F34  04                 inc b
6F35  00                 nop
6F36  08                 ex af,af'
6F37  FE 00              cp 0
6F39  8C                 adc a,h
6F3A  20 08              jr nz,$6F44
6F3C  02                 ld (bc),a
6F3D  01 8F D0           ld bc,$D08F
6F40  08                 ex af,af'
6F41  00                 nop
6F42  3F                 ccf
6F43  84                 add a,h
6F44  7D                 ld a,l
6F45  64                 ld h,h
6F46  40                 ld b,b
6F47  00                 nop
6F48  CF                 rst 08h
6F49  F0                 ret p
6F4A  08                 ex af,af'
6F4B  FE A8              cp -88
6F4D  8F                 adc a,a
6F4E  D0                 ret nc
6F4F  0C                 inc c
6F50  00                 nop
6F51  31 F4 81           ld sp,$81F4
6F54  1F                 rra
6F55  91                 sub c
6F56  3E A4              ld a,-92
6F58  7C                 ld a,h
6F59  80                 add a,b
6F5A  40                 ld b,b
6F5B  00                 nop
6F5C  8F                 adc a,a
6F5D  D0                 ret nc
6F5E  08                 ex af,af'
6F5F  02                 ld (bc),a
6F60  00                 nop
6F61  8F                 adc a,a
6F62  90                 sub b
6F63  18 C0              jr $6F25
6F65  10 00              djnz $6F67
6F67  7F                 ld a,a
6F68  08                 ex af,af'
6F69  FA C8 80           jp m,$80C8
6F6C  01 9F E0           ld bc,$E09F
6F6F  11 FD 51           ld de,$51FD
6F72  1F                 rra
6F73  A0                 and b
6F74  18 00              jr $6F76
6F76  23                 inc hl
6F77  E2 C0 01           jp po,$01C0
6F7A  8C                 adc a,h
6F7B  73                 ld (hl),e
6F7C  98                 sbc a,b
6F7D  C0                 ret nz
6F7E  00                 nop
6F7F  82                 add a,d
6F80  13                 inc de
6F81  E2 47 EC           jp po,$EC47
6F84  04                 inc b
6F85  00                 nop
6F86  08                 ex af,af'
6F87  00                 nop
6F88  3F                 ccf
6F89  8C                 adc a,h
6F8A  23                 inc hl
6F8B  F8                 ret m
6F8C  02                 ld (bc),a
6F8D  00                 nop
6F8E  0C                 inc c
6F8F  01 00 02           ld bc,512
6F92  00                 nop
6F93  0F                 rrca
6F94  D9                 exx
6F95  1F                 rra
6F96  B1                 or c
6F97  3E 44              ld a,68
6F99  00                 nop
6F9A  1F                 rra
6F9B  16 11              ld d,17
6F9D  8C                 adc a,h
6F9E  73                 ld (hl),e
6F9F  98                 sbc a,b
6FA0  FB                 ei
6FA1  28 8F              jr z,$6F32
6FA3  D0                 ret nc
6FA4  0C                 inc c
6FA5  00                 nop
6FA6  11 FE 11           ld de,$11FE
6FA9  18 02              jr $6FAD
6FAB  3E 42              ld a,66
6FAD  20 00              jr nz,$6FAF
6FAF  41                 ld b,c
6FB0  08                 ex af,af'
6FB1  80                 add a,b
6FB2  01 1F 70           ld bc,$701F
6FB5  3F                 ccf
6FB6  8C                 adc a,h
6FB7  23                 inc hl
6FB8  F4 02 00           call p,$0002
6FBB  04                 inc b
6FBC  03                 inc bc
6FBD  00                 nop
6FBE  47                 ld b,a
6FBF  C8                 ret z
6FC0  04                 inc b
6FC1  03                 inc bc
6FC2  08                 ex af,af'
6FC3  00                 nop
6FC4  3F                 ccf
6FC5  84                 add a,h
6FC6  31 80 03           ld sp,$0380
6FC9  00                 nop
6FCA  C0                 ret nz
6FCB  01 1F B1           ld bc,$B11F
6FCE  10 00              djnz $6FD0
6FD0  30 00              jr nc,$6FD2
6FD2  60                 ld h,b
6FD3  00                 nop
6FD4  C6 00              add a,0
6FD6  18 47              jr $701F
6FD8  F5                 push af
6FD9  44                 ld b,h
6FDA  7D                 ld a,l
6FDB  80                 add a,b
6FDC  C6 11              add a,17
6FDE  FE 11              cp 17
6FE0  00                 nop
6FE1  02                 ld (bc),a
6FE2  3E 42              ld a,66
6FE4  23                 inc hl
6FE5  9C                 sbc a,h
6FE6  41                 ld b,c
6FE7  08                 ex af,af'
6FE8  8F                 adc a,a
6FE9  C8                 ret z
6FEA  1F                 rra
6FEB  26 00              ld h,0
6FED  0C                 inc c
6FEE  01 00 02           ld bc,512
6FF1  00                 nop
6FF2  0C                 inc c
6FF3  00                 nop
6FF4  08                 ex af,af'
6FF5  C6 10              add a,16
6FF7  04                 inc b
6FF8  03                 inc bc
6FF9  F4 C2 31           call p,$31C2
6FFC  84                 add a,h
6FFD  00                 nop
6FFE  1F                 rra
6FFF  C2 18 C0           jp nz,$C018
7002  01 80 60           ld bc,$6080
7005  00                 nop
7006  8F                 adc a,a
7007  D8                 ret c
7008  88                 adc a,b
7009  FD 00              nop
700B  C0                 ret nz
700C  01 9F C0           ld bc,$C09F
700F  30 8F              jr nc,$6FA0
7011  EA 88 FB           jp pe,$FB88
7014  00                 nop
7015  C6 20              add a,32
7017  84                 add a,h
7018  FB                 ei
7019  11 80 7D           ld de,$7D80
701C  88                 adc a,b
701D  00                 nop
701E  3E 44              ld a,68
7020  00                 nop
7021  1F                 rra
7022  66                 ld h,(hl)
7023  00                 nop
7024  82                 add a,d
7025  11 00 C2           ld de,$C200
7028  00                 nop
7029  80                 add a,b
702A  61                 ld h,c
702B  1F                 rra
702C  A0                 and b
702D  10 00              djnz $702F
702F  20 08              jr nz,$7039
7031  FE 00              cp 0
7033  8F                 adc a,a
7034  90                 sub b
7035  08                 ex af,af'
7036  FE 01              cp 1
7038  FC 21 8C           call m,$8C21
703B  C6 18              add a,24
703D  06 00              ld b,0
703F  1F                 rra
7040  F2 00 0C           jp p,$0C00
7043  00                 nop
7044  0C                 inc c
7045  00                 nop
7046  19                 add hl,de
7047  F2 01 1F           jp p,$1F01
704A  D5                 push de
704B  11 F6 01           ld de,502
704E  8C                 adc a,h
704F  41                 ld b,c
7050  08                 ex af,af'
7051  8F                 adc a,a
7052  94                 sub h
7053  18 47              jr $709C
7055  F1                 pop af
7056  44                 ld b,h
7057  00                 nop
7058  08                 ex af,af'
7059  21 10 00           ld hl,16
705C  20 08              jr nz,$7066
705E  00                 nop
705F  11 F1 61           ld de,$61F1
7062  00                 nop
7063  40                 ld b,b
7064  3F                 ccf
7065  CC 23 00           call z,$0023
7068  07                 rlca
7069  F0                 ret p
706A  86                 add a,(hl)
706B  30 00              jr nc,$706D
706D  60                 ld h,b
706E  3F                 ccf
706F  8C                 adc a,h
7070  20 0C              jr nz,$707E
7072  02                 ld (bc),a
7073  00                 nop
7074  06 7E              ld b,126
7076  40                 ld b,b
7077  47                 ld b,a
7078  E8                 ret pe
7079  06 7C              ld b,124
707B  98                 sbc a,b
707C  47                 ld b,a
707D  F5                 push af
707E  44                 ld b,h
707F  63                 ld h,e
7080  1F                 rra
7081  20 18              jr nz,$709B
7083  C4 10 88           call nz,$8810
7086  FD 01 8F F8        ld bc,$F88F
708A  08                 ex af,af'
708B  FE 28              cp 40
708D  8C                 adc a,h
708E  03                 inc bc
708F  E4 40 30           call po,$3040
7092  80                 add a,b
7093  23                 inc hl
7094  18 47              jr $70DD
7096  D8                 ret c
7097  04                 inc b
7098  7F                 ld a,a
7099  00                 nop
709A  47                 ld b,a
709B  F0                 ret p
709C  04                 inc b
709D  7C                 ld a,h
709E  80                 add a,b
709F  47                 ld b,a
70A0  F0                 ret p
70A1  0F                 rrca
70A2  E1                 pop hl
70A3  0C                 inc c
70A4  67                 ld h,a
70A5  C8                 ret z
70A6  06 01              ld b,1
70A8  80                 add a,b
70A9  03                 inc bc
70AA  00                 nop
70AB  8F                 adc a,a
70AC  98                 sbc a,b
70AD  08                 ex af,af'
70AE  FE 01              cp 1
70B0  8F                 adc a,a
70B1  E0                 ret po
70B2  0C                 inc c
70B3  00                 nop
70B4  11 FD 51           ld de,$51FD
70B7  0C                 inc c
70B8  67                 ld h,a
70B9  E8                 ret pe
70BA  06 20              ld b,32
70BC  84                 add a,h
70BD  47                 ld b,a
70BE  E5                 push hl
70BF  8E                 adc a,(hl)
70C0  73                 ld (hl),e
70C1  18 00              jr $70C3
70C3  31 8E 73           ld sp,$738E
70C6  F8                 ret m
70C7  C7                 rst 00h
70C8  F1                 pop af
70C9  44                 ld b,h
70CA  00                 nop
70CB  1F                 rra
70CC  21 00 04           ld hl,1024
70CF  01 00 02           ld bc,512
70D2  01 80 03           ld bc,896
70D5  18 47              jr $711E
70D7  F1                 pop af
70D8  84                 add a,h
70D9  7F                 ld a,a
70DA  18 46              jr $7122
70DC  00                 nop
70DD  0F                 rrca
70DE  D3 1F              out ($1F),a
70E0  C2 18 C0           jp nz,$C018
70E3  63                 ld h,e
70E4  00                 nop
70E5  60                 ld h,b
70E6  18 00              jr $70E8
70E8  33                 inc sp
70E9  E4 07 C5           call po,$C507
70EC  80                 add a,b
70ED  01 18 00           ld bc,24
70F0  19                 add hl,de
70F1  FC 01 80           call m,$8001
70F4  07                 rlca
70F5  CD 44 31           call $3144
70F8  9F                 sbc a,a
70F9  A0                 and b
70FA  18 82              jr $707E
70FC  11 18 FE           ld de,-488
70FF  A8                 xor b
7100  80                 add a,b
7101  01 1F 95           ld bc,$951F
7104  11 80 7C           ld de,$7C80
7107  98                 sbc a,b
7108  06 10              ld b,16
710A  04                 inc b
710B  03                 inc bc
710C  08                 ex af,af'
710D  FF                 rst 38h
710E  41                 ld b,c
710F  FC 20 0C           call m,$0C20
7112  02                 ld (bc),a
7113  3E 60              ld a,96
7115  23                 inc hl
7116  EC 22 00           call pe,$0022
7119  04                 inc b
711A  7F                 ld a,a
711B  80                 add a,b
711C  60                 ld h,b
711D  01 F3 51           ld bc,$51F3
7120  0C                 inc c
7121  67                 ld h,a
7122  D8                 ret c
7123  04                 inc b
7124  10 88              djnz $70AE
7126  FB                 ei
7127  29                 add hl,hl
7128  8F                 adc a,a
7129  D2 88 E7           jp nc,$E788
712C  11 FA 53           ld de,$53FA
712F  1F                 rra
7130  C5                 push bc
7131  10 00              djnz $7133
7133  20 84              jr nz,$70B9
7135  40                 ld b,b
7136  00                 nop
7137  80                 add a,b
7138  20 00              jr nz,$713A
713A  40                 ld b,b
713B  3F                 ccf
713C  8C                 adc a,h
713D  00                 nop
713E  1F                 rra
713F  C6 3F              add a,63
7141  80                 add a,b
7142  7C                 ld a,h
7143  98                 sbc a,b
7144  00                 nop
7145  3E 4C              ld a,76
7147  03                 inc bc
7148  F8                 ret m
7149  46                 ld b,(hl)
714A  18 04              jr $7150
714C  7F                 ld a,a
714D  58                 ld e,b
714E  02                 ld (bc),a
714F  3E C2              ld a,-62
7151  20 00              jr nz,$7153
7153  47                 ld b,a
7154  DB 84              in a,($84)
7156  33                 inc sp
7157  F4 06 30           call p,$3006
715A  82                 add a,d
715B  11 18 F9           ld de,-1768
715E  29                 add hl,hl
715F  82                 add a,d
7160  62                 ld h,d
7161  98                 sbc a,b
7162  A6                 and (hl)
7163  3E 44              ld a,68
7165  62                 ld h,d
7166  98                 sbc a,b
7167  A6                 and (hl)
7168  09                 add hl,bc
7169  8A                 adc a,d
716A  51                 ld d,c
716B  18 02              jr $716F
716D  08                 ex af,af'
716E  44                 ld b,h
716F  03                 inc bc
7170  08                 ex af,af'
7171  02                 ld (bc),a
7172  31 84 01           ld sp,$0184
7175  00                 nop
7176  47                 ld b,a
7177  C8                 ret z
7178  04                 inc b
7179  00                 nop
717A  18 FB              jr $7177
717C  00                 nop
717D  8F                 adc a,a
717E  90                 sub b
717F  1F                 rra
7180  C2 00 C0           jp nz,$C000
7183  23                 inc hl
7184  F2 22 00           jp p,$0022
7187  0F                 rrca
7188  89                 adc a,c
7189  1F                 rra
718A  C0                 ret nz
718B  19                 add hl,de
718C  FA 81 0C           jp m,$0C81
718F  FD 00              nop
7191  C6 20              add a,32
7193  84                 add a,h
7194  47                 ld b,a
7195  D9                 exx
7196  4C                 ld c,h
7197  7D                 ld a,l
7198  D4 C7 F1           call nc,$F1C7
719B  44                 ld b,h
719C  00                 nop
719D  08                 ex af,af'
719E  21 10 00           ld hl,16
71A1  20 08              jr nz,$71AB
71A3  00                 nop
71A4  10 04              djnz $71AA
71A6  01 1F 20           ld bc,$201F
71A9  10 00              djnz $71AB
71AB  23                 inc hl
71AC  00                 nop
71AD  06 30              ld b,48
71AF  04                 inc b
71B0  7F                 ld a,a
71B1  18 07              jr $71BA
71B3  F0                 ret p
71B4  80                 add a,b
71B5  33                 inc sp
71B6  08                 ex af,af'
71B7  FC 88 8C           call m,$8C88
71BA  7C                 ld a,h
71BB  C0                 ret nz
71BC  67                 ld h,a
71BD  EA 04 7D           jp pe,$7D04
71C0  80                 add a,b
71C1  63                 ld h,e
71C2  10 42              djnz $7206
71C4  23                 inc hl
71C5  F3                 di
71C6  A2                 and d
71C7  30 04              jr nc,$71CD
71C9  10 88              djnz $7153
71CB  06 10              ld b,16
71CD  04                 inc b
71CE  7E                 ld a,(hl)
71CF  98                 sbc a,b
71D0  02                 ld (bc),a
71D1  3F                 ccf
71D2  C0                 ret nz
71D3  23                 inc hl
71D4  EC 02 3E           call pe,$3E02
71D7  40                 ld b,b
71D8  7F                 ld a,a
71D9  08                 ex af,af'
71DA  03                 inc bc
71DB  00                 nop
71DC  8F                 adc a,a
71DD  C8                 ret z
71DE  88                 adc a,b
71DF  FF                 rst 38h
71E0  81                 add a,c
71E1  F8                 ret m
71E2  E0                 ret po
71E3  00                 nop
71E4  D2 35 AF           jp nc,$AF35
71E7  90                 sub b
71E8  0C                 inc c
71E9  62                 ld h,d
71EA  08                 ex af,af'
71EB  44                 ld b,h
71EC  7F                 ld a,a
71ED  14                 inc d
71EE  F9                 ld sp,hl
71EF  31 F4 D1           ld sp,$D1F4
71F2  00                 nop
71F3  07                 rlca
71F4  C8                 ret z
71F5  80                 add a,b
71F6  01 1F 20           ld bc,$201F
71F9  11 FC 01           ld de,508
71FC  1F                 rra
71FD  A6                 and (hl)
71FE  00                 nop
71FF  04                 inc b
7200  63                 ld h,e
7201  1F                 rra
7202  20 11              jr nz,$7215
7204  FC 60 1F           call m,$1F60
7207  C2 30 C0           jp nz,$C030
720A  23                 inc hl
720B  F2 22 3E           jp p,$3E22
720E  4C                 ld c,h
720F  00                 nop
7210  1F                 rra
7211  96                 sub (hl)
7212  11 FF 11           ld de,$11FF
7215  1F                 rra
7216  C0                 ret nz
7217  10 C6              djnz $71DF
7219  7E                 ld a,(hl)
721A  80                 add a,b
721B  62                 ld h,d
721C  08                 ex af,af'
721D  44                 ld b,h
721E  7D                 ld a,l
721F  94                 sub h
7220  46                 ld b,(hl)
7221  3F                 ccf
7222  8A                 adc a,d
7223  63                 ld h,e
7224  E4 A6 3E           call po,$3EA6
7227  4A                 ld c,d
7228  63                 ld h,e
7229  F8                 ret m
722A  A2                 and d
722B  31 F6 03           ld sp,$03F6
722E  08                 ex af,af'
722F  F9                 ld sp,hl
7230  00                 nop
7231  8F                 adc a,a
7232  88                 adc a,b
7233  08                 ex af,af'
7234  F9                 ld sp,hl
7235  00                 nop
7236  80                 add a,b
7237  03                 inc bc
7238  18 47              jr $7281
723A  C8                 ret z
723B  0F                 rrca
723C  E1                 pop hl
723D  0C                 inc c
723E  60                 ld h,b
723F  3F                 ccf
7240  E4 00 1F           call po,$1F00
7243  AA                 xor d
7244  3E 42              ld a,66
7246  20 00              jr nz,$7248
7248  D2 19 F6           jp nc,$F619
724B  01 88 21           ld bc,$2188
724E  11 F6 51           ld de,$51F6
7251  1F                 rra
7252  25                 dec h
7253  11 F2 51           ld de,$51F2
7256  1F                 rra
7257  25                 dec h
7258  11 FC 51           ld de,$51FC
725B  1F                 rra
725C  B0                 or b
725D  11 FB 01           ld de,507
7260  1F                 rra
7261  90                 sub b
7262  3F                 ccf
7263  84                 add a,h
7264  31 98 40           ld sp,$4098
7267  18 06              jr $726F
7269  7F                 ld a,a
726A  80                 add a,b
726B  63                 ld h,e
726C  10 00              djnz $726E
726E  23                 inc hl
726F  E4 03 00           call po,$0003
7272  06 00              ld b,0
7274  0C                 inc c
7275  03                 inc bc
7276  3E C0              ld a,-64
7278  23                 inc hl
7279  E4 22 3F           call po,$3F22
727C  80                 add a,b
727D  21 9F E0           ld hl,$E09F
7280  10 42              djnz $72C4
7282  23                 inc hl
7283  E4 C5 28           call po,$28C5
7286  8F                 adc a,a
7287  93                 sub e
7288  08                 ex af,af'
7289  F9                 ld sp,hl
728A  30 8F              jr nc,$721B
728C  93                 sub e
728D  08                 ex af,af'
728E  FE 30              cp 48
7290  8F                 adc a,a
7291  DB 08              in a,($08)
7293  FF                 rst 38h
7294  71                 ld (hl),c
7295  FC 20 0C           call m,$0C20
7298  62                 ld h,d
7299  3F                 ccf
729A  A0                 and b
729B  31 88 E7           ld sp,$E788
729E  11 F2 01           ld de,498
72A1  80                 add a,b
72A2  03                 inc bc
72A3  00                 nop
72A4  06 01              ld b,1
72A6  9F                 sbc a,a
72A7  60                 ld h,b
72A8  11 F2 11           ld de,$11F2
72AB  00                 nop
72AC  06 90              ld b,-112
72AE  CF                 rst 08h
72AF  F0                 ret p
72B0  08                 ex af,af'
72B1  21 11 F6           ld hl,$F611
72B4  51                 ld d,c
72B5  1F                 rra
72B6  FC 3F B8           call m,$B83F
72B9  23                 inc hl
72BA  0C                 inc c
72BB  62                 ld h,d
72BC  3F                 ccf
72BD  2C                 inc l
72BE  7F                 ld a,a
72BF  00                 nop
72C0  62                 ld h,d
72C1  31 84 63           ld sp,$6384
72C4  F8                 ret m
72C5  03                 inc bc
72C6  00                 nop
72C7  06 00              ld b,0
72C9  0C                 inc c
72CA  03                 inc bc
72CB  00                 nop
72CC  0F                 rrca
72CD  E3                 ex (sp),hl
72CE  00                 nop
72CF  FB                 ei
72D0  11 FC 01           ld de,508
72D3  0C                 inc c
72D4  FB                 ei
72D5  01 A4 10           ld bc,$10A4
72D8  88                 adc a,b
72D9  FB                 ei
72DA  28 8F              jr z,$726B
72DC  FE 1F              cp 31
72DE  DC 10 C6           call c,$C610
72E1  01 1F 11           ld bc,$111F
72E4  11 80 01           ld de,384
72E7  80                 add a,b
72E8  63                 ld h,e
72E9  11 F2 01           ld de,498
72EC  9F                 sbc a,a
72ED  A0                 and b
72EE  19                 add hl,de
72EF  FA 01 1F           jp m,$1F01
72F2  E0                 ret po
72F3  10 00              djnz $72F5
72F5  69                 ld l,c
72F6  1F                 rra
72F7  67                 ld h,a
72F8  00                 nop
72F9  04                 inc b
72FA  10 88              djnz $7284
72FC  A5                 and l
72FD  3E 4C              ld a,76
72FF  22 10 B6           ld ($B610),hl
7302  7D                 ld a,l
7303  E7                 rst 20h
7304  1B                 dec de
7305  FF                 rst 38h
7306  87                 add a,a
7307  E5                 push hl
7308  0B                 dec bc
7309  67                 ld h,a
730A  D9                 exx
730B  37                 scf
730C  E9                 jp (hl)
730D  04                 inc b
730E  01 98 47           ld bc,$4798
7311  C4 44 7E           call nz,$7E44
7314  80                 add a,b
7315  63                 ld h,e
7316  11 F5 01           ld de,501
7319  9F                 sbc a,a
731A  A0                 and b
731B  11 8F B0           ld de,$B08F
731E  08                 ex af,af'
731F  FE 00              cp 0
7321  8F                 adc a,a
7322  B3                 or e
7323  60                 ld h,b
7324  D2 08 44           jp nc,$4408
7327  7D                 ld a,l
7328  94                 sub h
7329  47                 ld b,a
732A  D3 0B              out ($0B),a
732C  38 DF              jr c,$730D
732E  14                 inc d
732F  2D                 dec l
7330  E7                 rst 20h
7331  1B                 dec de
7332  E4 85 9C           call po,$9C85
7335  6F                 ld l,a
7336  A6                 and (hl)
7337  08                 ex af,af'
7338  00                 nop
7339  18 8F              jr $72CA
733B  88                 adc a,b
733C  88                 adc a,b
733D  FD 30 C0           jr nc,$7300
7340  23                 inc hl
7341  E1                 pop hl
7342  02                 ld (bc),a
7343  3F                 ccf
7344  AC                 xor h
7345  23                 inc hl
7346  E2 02 08           jp po,$0802
7349  44                 ld b,h
734A  7D                 ld a,l
734B  94                 sub h
734C  47                 ld b,a
734D  FF                 rst 38h
734E  0F                 rrca
734F  E6 0F              and $0F
7351  91                 sub c
7352  7A                 ld a,d
7353  8B                 adc a,e
7354  F4 60 3E           call p,$3E60
7357  A5                 and l
7358  28 00              jr z,$735A
735A  18 8F              jr $72EB
735C  D8                 ret c
735D  88                 adc a,b
735E  60                 ld h,b
735F  31 8F FB           ld sp,$FB8F
7362  88                 adc a,b
7363  60                 ld h,b
7364  18 06              jr $736C
7366  30 0C              jr nc,$7374
7368  00                 nop
7369  19                 add hl,de
736A  FC 03 EC           call m,$EC03
736D  C2 08 44           jp nz,$4408
7370  7C                 ld a,h
7371  98                 sbc a,b
7372  A5                 and l
7373  11 FF C3           ld de,$C3FF
7376  F9                 ld sp,hl
7377  83                 add a,e
7378  CC 5F 62           call z,$625F
737B  F7                 rst 30h
737C  18 36              jr $73B4
737E  B1                 or c
737F  8A                 adc a,d
7380  00                 nop
7381  06 7F              ld b,127
7383  C8                 ret z
7384  06 3E              ld b,62
7386  94                 sub h
7387  30 0C              jr nc,$7395
7389  00                 nop
738A  18 06              jr $7392
738C  00                 nop
738D  0C                 inc c
738E  F9                 ld sp,hl
738F  00                 nop
7390  8F                 adc a,a
7391  F0                 ret p
7392  88                 adc a,b
7393  FB                 ei
7394  28 8F              jr z,$7325
7396  9A                 sbc a,d
7397  0B                 dec bc
7398  91                 sub c
7399  27                 daa
739A  F2 41 72           jp p,$7241
739D  24                 inc h
739E  FE B8              cp -72
73A0  3F                 ccf
73A1  C5                 push bc
73A2  E1                 pop hl
73A3  37                 scf
73A4  81                 add a,c
73A5  40                 ld b,b
73A6  10 A3              djnz $734B
73A8  F8                 ret m
73A9  03                 inc bc
73AA  3F                 ccf
73AB  80                 add a,b
73AC  33                 inc sp
73AD  EF                 rst 28h
73AE  03                 inc bc
73AF  3E 40              ld a,64
73B1  68                 ld l,b
73B2  0C                 inc c
73B3  D0                 ret nc
73B4  19                 add hl,de
73B5  A0                 and b
73B6  03                 inc bc
73B7  48                 ld c,b
73B8  FF                 rst 38h
73B9  08                 ex af,af'
73BA  8F                 adc a,a
73BB  B2                 or d
73BC  88                 adc a,b
73BD  FE A0              cp -96
73BF  BF                 cp a
73C0  E2 44 9C           jp po,$9C44
73C3  20 BF              jr nz,$7384
73C5  E2 44 9F           jp po,$9F44
73C8  FB                 ei
73C9  0F                 rrca
73CA  8B                 adc a,e
73CB  1F                 rra
73CC  C2 3F 90           jp nz,$903F
73CF  63                 ld h,e
73D0  E5                 push hl
73D1  06 80              ld b,-128
73D3  06 7E              ld b,126
73D5  40                 ld b,b
73D6  47                 ld b,a
73D7  F8                 ret m
73D8  44                 ld b,h
73D9  7D                 ld a,l
73DA  94                 sub h
73DB  47                 ld b,a
73DC  D5                 push de
73DD  05                 dec b
73DE  FE 92              cp -110
73E0  24                 inc h
73E1  D7                 rst 10h
73E2  FA 48 93           jp m,$9348
73E5  FB                 ei
73E6  61                 ld h,c
73E7  84                 add a,h
73E8  7D                 ld a,l
73E9  84                 add a,h
73EA  FE 11              cp 17
73EC  FC 63 9C           call m,$9C63
73EF  FE 31              cp 49
73F1  CE 63              adc a,99
73F3  1C                 inc e
73F4  E7                 rst 20h
73F5  F7                 rst 30h
73F6  8E                 adc a,(hl)
73F7  73                 ld (hl),e
73F8  1C                 inc e
73F9  E6 39              and $39
73FB  CC 73 98           call z,$9873
73FE  47                 ld b,a
73FF  F8                 ret m
7400  44                 ld b,h
7401  7F                 ld a,a
7402  F8                 ret m
7403  FC F0 8F           call m,$8FF0
7406  B0                 or b
7407  88                 adc a,b
7408  47                 ld b,a
7409  E0                 ret po
740A  00                 nop
740B  00                 nop
740C  00                 nop
740D  00                 nop
740E  00                 nop
740F  00                 nop
7410  00                 nop
7411  00                 nop
7412  00                 nop
7413  00                 nop
7414  00                 nop
7415  00                 nop
7416  00                 nop
7417  00                 nop
7418  00                 nop
7419  00                 nop
741A  00                 nop
741B  00                 nop
741C  00                 nop
741D  00                 nop
741E  00                 nop
741F  00                 nop
7420  00                 nop
7421  00                 nop
7422  21 6A B0           ld hl,$B06A
7425  AF                 xor a
7426  0E 08              ld c,8
7428  06 05              ld b,5
742A  56                 ld d,(hl)
742B  CB 3A              srl d
742D  30 07              jr nc,$7436
742F  3C                 inc a
7430  10 F9              djnz $742B
7432  23                 inc hl
7433  0D                 dec c
7434  20 F2              jr nz,$7428
7436  A7                 and a
7437  28 F6              jr z,$742F
7439  FE 24              cp 36
743B  28 F2              jr z,$742F
743D  5F                 ld e,a
743E  16 00              ld d,0
7440  21 46 74           ld hl,$7446
7443  19                 add hl,de
7444  7E                 ld a,(hl)
7445  C9                 ret
7446  2A 5A 58           ld hl,($585A)
7449  43                 ld b,e
744A  56                 ld d,(hl)
744B  41                 ld b,c
744C  53                 ld d,e
744D  44                 ld b,h
744E  46                 ld b,(hl)
744F  47                 ld b,a
7450  51                 ld d,c
7451  57                 ld d,a
7452  45                 ld b,l
7453  52                 ld d,d
7454  54                 ld d,h
7455  31 32 33           ld sp,$3332
7458  34                 inc (hl)
7459  35                 dec (hl)
745A  30 39              jr nc,$7495
745C  38 37              jr c,$7495
745E  36 50              ld (hl),80
7460  4F                 ld c,a
7461  49                 ld c,c
7462  55                 ld d,l
7463  59                 ld e,c
7464  0D                 dec c
7465  4C                 ld c,h
7466  4B                 ld c,e
7467  4A                 ld c,d
7468  48                 ld c,b
7469  20 2A              jr nz,$7495
746B  4D                 ld c,l
746C  4E                 ld c,(hl)
746D  42                 ld b,d
746E  2A 01 14           ld hl,($1401)
7471  00                 nop
7472  21 9D B0           ld hl,$B09D
7475  34                 inc (hl)
7476  7E                 ld a,(hl)
7477  E6 03              and $03
7479  21 9E B0           ld hl,$B09E
747C  7E                 ld a,(hl)
747D  CC A1 74           call z,$74A1
7480  77                 ld (hl),a
7481  3A 9D B0           ld a,($B09D)
7484  68                 ld l,b
7485  26 00              ld h,0
7487  29                 add hl,hl
7488  29                 add hl,hl
7489  29                 add hl,hl
748A  29                 add hl,hl
748B  29                 add hl,hl
748C  ED 5B 66 B0        ld de,($B066)
7490  19                 add hl,de
7491  3A 9E B0           ld a,($B09E)
7494  06 20              ld b,32
7496  77                 ld (hl),a
7497  23                 inc hl
7498  CD A1 74           call $74A1
749B  10 F9              djnz $7496
749D  0D                 dec c
749E  20 F4              jr nz,$7494
74A0  C9                 ret
74A1  3C                 inc a
74A2  FE 48              cp 72
74A4  C0                 ret nz
74A5  3E 41              ld a,65
74A7  C9                 ret
74A8  FD E5              push iy
74AA  E1                 pop hl
74AB  11 16 76           ld de,$7616
74AE  01 03 00           ld bc,3
74B1  ED B0              ldir
74B3  32 D0 75           ld ($75D0),a
74B6  21 2E 2E           ld hl,$2E2E
74B9  22 12 76           ld ($7612),hl
74BC  22 14 76           ld ($7614),hl
74BF  22 D3 75           ld ($75D3),hl
74C2  22 D5 75           ld ($75D5),hl
74C5  11 12 76           ld de,$7612
74C8  ED 53 A0 B0        ld ($B0A0),de
74CC  3E 08              ld a,8
74CE  32 9F B0           ld ($B09F),a
74D1  ED 5B A0 B0        ld de,($B0A0)
74D5  13                 inc de
74D6  13                 inc de
74D7  13                 inc de
74D8  13                 inc de
74D9  21 F9 FF           ld hl,-7
74DC  19                 add hl,de
74DD  06 03              ld b,3
74DF  1A                 ld a,(de)
74E0  BE                 cp (hl)
74E1  38 23              jr c,$7506
74E3  20 04              jr nz,$74E9
74E5  23                 inc hl
74E6  13                 inc de
74E7  10 F6              djnz $74DF
74E9  ED 5B A0 B0        ld de,($B0A0)
74ED  21 F9 FF           ld hl,-7
74F0  19                 add hl,de
74F1  22 A0 B0           ld ($B0A0),hl
74F4  06 07              ld b,7
74F6  1A                 ld a,(de)
74F7  4F                 ld c,a
74F8  7E                 ld a,(hl)
74F9  12                 ld (de),a
74FA  71                 ld (hl),c
74FB  23                 inc hl
74FC  13                 inc de
74FD  10 F7              djnz $74F6
74FF  21 9F B0           ld hl,$B09F
7502  35                 dec (hl)
7503  F2 D1 74           jp p,$74D1
7506  3A 9F B0           ld a,($B09F)
7509  FE 08              cp 8
750B  C8                 ret z
750C  01 02 01           ld bc,258
750F  CD 29 B5           call $B529
7512  CD A0 79           call $79A0
7515  21 B1 75           ld hl,$75B1
7518  CD CE B2           call $B2CE
751B  CD 17 B5           call $B517
751E  21 00 00           ld hl,0
7521  22 D8 75           ld ($75D8),hl
7524  3E 02              ld a,2
7526  32 FF E9           ld ($E9FF),a
7529  2A 9F B0           ld hl,($B09F)
752C  26 00              ld h,0
752E  29                 add hl,hl
752F  11 03 0B           ld de,$0B03
7532  19                 add hl,de
7533  CD C4 B2           call $B2C4
7536  21 D3 75           ld hl,$75D3
7539  CD CE B2           call $B2CE
753C  CD 9B 75           call $759B
753F  CD 22 74           call $7422
7542  FE 2A              cp 42
7544  20 F6              jr nz,$753C
7546  CD 9B 75           call $759B
7549  CD 22 74           call $7422
754C  FE 2A              cp 42
754E  28 F6              jr z,$7546
7550  5F                 ld e,a
7551  21 D3 75           ld hl,$75D3
7554  ED 4B D8 75        ld bc,($75D8)
7558  09                 add hl,bc
7559  FE 0D              cp 13
755B  28 31              jr z,$758E
755D  FE 30              cp 48
755F  20 0A              jr nz,$756B
7561  DD 21 6A B0        ld ix,$B06A
7565  DD CB 00 46        bit 0,(ix+0)
7569  28 15              jr z,$7580
756B  FE 30              cp 48
756D  38 04              jr c,$7573
756F  FE 3A              cp 58
7571  38 D3              jr c,$7546
7573  79                 ld a,c
7574  FE 04              cp 4
7576  28 CE              jr z,$7546
7578  73                 ld (hl),e
7579  0C                 inc c
757A  ED 43 D8 75        ld ($75D8),bc
757E  18 A4              jr $7524
7580  79                 ld a,c
7581  A7                 and a
7582  28 C2              jr z,$7546
7584  0B                 dec bc
7585  ED 43 D8 75        ld ($75D8),bc
7589  2B                 dec hl
758A  36 2E              ld (hl),46
758C  18 96              jr $7524
758E  21 D3 75           ld hl,$75D3
7591  ED 5B A0 B0        ld de,($B0A0)
7595  01 04 00           ld bc,4
7598  ED B0              ldir
759A  C9                 ret
759B  76                 halt
759C  01 02 00           ld bc,2
759F  CD 72 74           call $7472
75A2  CD C3 CF           call $CFC3
75A5  3A 9F B0           ld a,($B09F)
75A8  87                 add a,a
75A9  C6 03              add a,3
75AB  47                 ld b,a
75AC  0E 01              ld c,1
75AE  C3 81 74           jp $7481
75B1  02                 ld (bc),a
75B2  00                 nop
75B3  07                 rlca
75B4  41                 ld b,c
75B5  20 4E              jr nz,$7605
75B7  45                 ld b,l
75B8  57                 ld d,a
75B9  20 48              jr nz,$7603
75BB  49                 ld c,c
75BC  47                 ld b,a
75BD  48                 ld c,b
75BE  20 53              jr nz,$7613
75C0  43                 ld b,e
75C1  4F                 ld c,a
75C2  52                 ld d,d
75C3  45                 ld b,l
75C4  20 21              jr nz,$75E7
75C6  02                 ld (bc),a
75C7  01 0B 50           ld bc,$500B
75CA  4C                 ld c,h
75CB  41                 ld b,c
75CC  59                 ld e,c
75CD  45                 ld b,l
75CE  52                 ld d,d
75CF  20 31              jr nz,$7602
75D1  20 A1              jr nz,$7574
75D3  2E 2E              ld l,46
75D5  2E 2E              ld l,46
75D7  A0                 and b
75D8  00                 nop
75D9  00                 nop
75DA  46                 ld b,(hl)
75DB  52                 ld d,d
75DC  45                 ld b,l
75DD  44                 ld b,h
75DE  01 00 00           ld bc,0
75E1  50                 ld d,b
75E2  45                 ld b,l
75E3  54                 ld d,h
75E4  45                 ld b,l
75E5  00                 nop
75E6  50                 ld d,b
75E7  00                 nop
75E8  50                 ld d,b
75E9  41                 ld b,c
75EA  55                 ld d,l
75EB  4C                 ld c,h
75EC  00                 nop
75ED  20 00              jr nz,$75EF
75EF  2D                 dec l
75F0  41                 ld b,c
75F1  4E                 ld c,(hl)
75F2  44                 ld b,h
75F3  00                 nop
75F4  10 00              djnz $75F6
75F6  20 54              jr nz,$764C
75F8  48                 ld c,b
75F9  45                 ld b,l
75FA  00                 nop
75FB  05                 dec b
75FC  00                 nop
75FD  45                 ld b,l
75FE  56                 ld d,(hl)
75FF  49                 ld c,c
7600  4C                 ld c,h
7601  00                 nop
7602  02                 ld (bc),a
7603  00                 nop
7604  50                 ld d,b
7605  49                 ld c,c
7606  4E                 ld c,(hl)
7607  4B                 ld c,e
7608  00                 nop
7609  01 00 20           ld bc,$2000
760C  45                 ld b,l
760D  4C                 ld c,h
760E  46                 ld b,(hl)
760F  00                 nop
7610  00                 nop
7611  50                 ld d,b
7612  2D                 dec l
7613  2D                 dec l
7614  2D                 dec l
7615  2D                 dec l
7616  FF                 rst 38h
7617  FF                 rst 38h
7618  FF                 rst 38h
7619  3E 09              ld a,9
761B  32 15 D2           ld ($D215),a
761E  3A AC B0           ld a,($B0AC)
7621  A7                 and a
7622  3E 09              ld a,9
7624  28 01              jr z,$7627
7626  AF                 xor a
7627  32 3D D2           ld ($D23D),a
762A  AF                 xor a
762B  32 CA 7F           ld ($7FCA),a
762E  32 D6 7F           ld ($7FD6),a
7631  32 72 B0           ld ($B072),a
7634  67                 ld h,a
7635  6F                 ld l,a
7636  32 CD 7F           ld ($7FCD),a
7639  22 CE 7F           ld ($7FCE),hl
763C  32 D9 7F           ld ($7FD9),a
763F  22 DA 7F           ld ($7FDA),hl
7642  3D                 dec a
7643  32 14 B3           ld ($B314),a
7646  C9                 ret
7647  3A 73 B0           ld a,($B073)
764A  ED 4B 72 B0        ld bc,($B072)
764E  B9                 cp c
764F  CA D8 76           jp z,$76D8
7652  21 22 60           ld hl,$6022
7655  11 00 D3           ld de,$D300
7658  01 00 14           ld bc,$1400
765B  ED B0              ldir
765D  DD 21 22 60        ld ix,$6022
7661  11 32 60           ld de,$6032
7664  ED 53 A4 B0        ld ($B0A4),de
7668  2E 00              ld l,0
766A  ED 5B A4 B0        ld de,($B0A4)
766E  E5                 push hl
766F  D5                 push de
7670  3A 72 B0           ld a,($B072)
7673  BD                 cp l
7674  28 3A              jr z,$76B0
7676  3A 73 B0           ld a,($B073)
7679  BD                 cp l
767A  28 1C              jr z,$7698
767C  DD 4E 08           ld c,(ix+8)
767F  DD 46 09           ld b,(ix+9)
7682  DD 6E 00           ld l,(ix+0)
7685  DD 66 01           ld h,(ix+1)
7688  11 DE 72           ld de,$72DE
768B  19                 add hl,de
768C  ED 5B A4 B0        ld de,($B0A4)
7690  ED B0              ldir
7692  ED 53 A4 B0        ld ($B0A4),de
7696  18 18              jr $76B0
7698  CD D2 7B           call $7BD2
769B  ED 5B A4 B0        ld de,($B0A4)
769F  D5                 push de
76A0  CD 50 7A           call $7A50
76A3  D1                 pop de
76A4  2A A4 B0           ld hl,($B0A4)
76A7  A7                 and a
76A8  ED 52              sbc hl,de
76AA  DD 75 08           ld (ix+8),l
76AD  DD 74 09           ld (ix+9),h
76B0  D1                 pop de
76B1  E1                 pop hl
76B2  DD 73 00           ld (ix+0),e
76B5  DD 72 01           ld (ix+1),d
76B8  DD 23              inc ix
76BA  DD 23              inc ix
76BC  2C                 inc l
76BD  7D                 ld a,l
76BE  FE 04              cp 4
76C0  20 A8              jr nz,$766A
76C2  3A 72 B0           ld a,($B072)
76C5  32 73 B0           ld ($B073),a
76C8  11 00 D3           ld de,$D300
76CB  CD 0A B3           call $B30A
76CE  11 DE 72           ld de,$72DE
76D1  19                 add hl,de
76D2  11 00 EB           ld de,$EB00
76D5  CD D2 79           call $79D2
76D8  3E FF              ld a,-1
76DA  32 74 B0           ld ($B074),a
76DD  21 00 20           ld hl,$2000
76E0  22 76 B0           ld ($B076),hl
76E3  3E 44              ld a,68
76E5  32 88 B0           ld ($B088),a
76E8  3A 72 B0           ld a,($B072)
76EB  32 87 B0           ld ($B087),a
76EE  32 0F 77           ld ($770F),a
76F1  11 F5 C8           ld de,$C8F5
76F4  CD 0A B3           call $B30A
76F7  7D                 ld a,l
76F8  32 2B C9           ld ($C92B),a
76FB  7C                 ld a,h
76FC  32 2D C9           ld ($C92D),a
76FF  DD 21 9F B2        ld ix,$B29F
7703  11 05 00           ld de,5
7706  06 06              ld b,6
7708  26 EA              ld h,-22
770A  DD 6E 04           ld l,(ix+4)
770D  DD 7E 03           ld a,(ix+3)
7710  77                 ld (hl),a
7711  DD 19              add ix,de
7713  10 F5              djnz $770A
7715  21 F4 CF           ld hl,$CFF4
7718  11 F5 CF           ld de,$CFF5
771B  01 E3 02           ld bc,739
771E  36 FF              ld (hl),-1
7720  ED B0              ldir
7722  11 F4 FF           ld de,-12
7725  3A 72 B0           ld a,($B072)
7728  CD 0A B3           call $B30A
772B  11 80 B0           ld de,$B080
772E  01 04 00           ld bc,4
7731  ED B0              ldir
7733  3E FF              ld a,-1
7735  32 84 B0           ld ($B084),a
7738  E5                 push hl
7739  DD E1              pop ix
773B  FD 21 F4 CF        ld iy,$CFF4
773F  DD 7E 00           ld a,(ix+0)
7742  FE FF              cp -1
7744  C8                 ret z
7745  5F                 ld e,a
7746  DD 56 01           ld d,(ix+1)
7749  DD 23              inc ix
774B  DD 23              inc ix
774D  D5                 push de
774E  7B                 ld a,e
774F  A7                 and a
7750  20 30              jr nz,$7782
7752  FD E5              push iy
7754  FD 21 10 D2        ld iy,$D210
7758  3A AE B0           ld a,($B0AE)
775B  A7                 and a
775C  C4 95 77           call nz,$7795
775F  FD 36 FE 40        ld (iy-2),64
7763  FD 21 38 D2        ld iy,$D238
7767  1E 00              ld e,0
7769  3A AF B0           ld a,($B0AF)
776C  A7                 and a
776D  C4 95 77           call nz,$7795
7770  FD 36 FE 40        ld (iy-2),64
7774  2A 3A D2           ld hl,($D23A)
7777  11 08 00           ld de,8
777A  19                 add hl,de
777B  22 3A D2           ld ($D23A),hl
777E  FD E1              pop iy
7780  18 03              jr $7785
7782  CD 95 77           call $7795
7785  DD CB 00 7E        bit 7,(ix+0)
7789  20 02              jr nz,$778D
778B  DD 23              inc ix
778D  DD 23              inc ix
778F  D1                 pop de
7790  15                 dec d
7791  20 BA              jr nz,$774D
7793  18 AA              jr $773F
7795  FD 73 09           ld (iy+9),e
7798  16 00              ld d,0
779A  21 BF FF           ld hl,-65
779D  19                 add hl,de
779E  7E                 ld a,(hl)
779F  FD 77 05           ld (iy+5),a
77A2  11 42 FF           ld de,-190
77A5  19                 add hl,de
77A6  7E                 ld a,(hl)
77A7  FD 77 04           ld (iy+4),a
77AA  DD CB 00 7E        bit 7,(ix+0)
77AE  20 13              jr nz,$77C3
77B0  DD 6E 01           ld l,(ix+1)
77B3  26 00              ld h,0
77B5  29                 add hl,hl
77B6  29                 add hl,hl
77B7  29                 add hl,hl
77B8  29                 add hl,hl
77B9  11 0F 00           ld de,15
77BC  19                 add hl,de
77BD  22 FA 77           ld ($77FA),hl
77C0  C3 C6 77           jp $77C6
77C3  2A FA 77           ld hl,($77FA)
77C6  FD 75 00           ld (iy+0),l
77C9  FD 74 01           ld (iy+1),h
77CC  DD 7E 00           ld a,(ix+0)
77CF  E6 7F              and $7F
77D1  6F                 ld l,a
77D2  26 00              ld h,0
77D4  29                 add hl,hl
77D5  29                 add hl,hl
77D6  2B                 dec hl
77D7  FD 75 02           ld (iy+2),l
77DA  FD 74 03           ld (iy+3),h
77DD  FD 36 08 00        ld (iy+8),0
77E1  FD 36 06 00        ld (iy+6),0
77E5  FD 7E 09           ld a,(iy+9)
77E8  FE 1F              cp 31
77EA  20 08              jr nz,$77F4
77EC  FD 36 06 03        ld (iy+6),3
77F0  FD 36 08 30        ld (iy+8),48
77F4  11 0A 00           ld de,10
77F7  FD 19              add iy,de
77F9  C9                 ret
77FA  00                 nop
77FB  00                 nop
77FC  21 01 00           ld hl,1
77FF  22 49 78           ld ($7849),hl
7802  11 1B 78           ld de,$781B
7805  D5                 push de
7806  2A 49 78           ld hl,($7849)
7809  2B                 dec hl
780A  22 49 78           ld ($7849),hl
780D  7D                 ld a,l
780E  A7                 and a
780F  C0                 ret nz
7810  7C                 ld a,h
7811  E6 03              and $03
7813  A7                 and a
7814  28 56              jr z,$786C
7816  FE 02              cp 2
7818  28 31              jr z,$784B
781A  C9                 ret
781B  76                 halt
781C  CD 6F 74           call $746F
781F  3A 14 B3           ld a,($B314)
7822  A7                 and a
7823  CC 4B FB           call z,$FB4B
7826  CD 22 74           call $7422
7829  D6 31              sub 49
782B  38 14              jr c,$7841
782D  FE 05              cp 5
782F  30 10              jr nc,$7841
7831  32 FE EA           ld ($EAFE),a
7834  FE 03              cp 3
7836  3E FF              ld a,-1
7838  38 01              jr c,$783B
783A  AF                 xor a
783B  32 AC B0           ld ($B0AC),a
783E  CD 6C 78           call $786C
7841  3A D5 7F           ld a,($7FD5)
7844  E6 10              and $10
7846  20 BA              jr nz,$7802
7848  C9                 ret
7849  01 00 01           ld bc,256
784C  00                 nop
784D  04                 inc b
784E  CD 29 B5           call $B529
7851  21 60 78           ld hl,$7860
7854  CD CE B2           call $B2CE
7857  CD A0 79           call $79A0
785A  CD 91 79           call $7991
785D  C3 17 B5           jp $B517
7860  02                 ld (bc),a
7861  01 0C 48           ld bc,$480C
7864  49                 ld c,c
7865  20 53              jr nz,$78BA
7867  43                 ld b,e
7868  4F                 ld c,a
7869  52                 ld d,d
786A  45                 ld b,l
786B  D3 01              out ($01),a
786D  00                 nop
786E  00                 nop
786F  CD 29 B5           call $B529
7872  21 8D 78           ld hl,$788D
7875  CD CE B2           call $B2CE
7878  CD 91 79           call $7991
787B  11 2B 79           ld de,$792B
787E  3A FE EA           ld a,($EAFE)
7881  CD 0A B3           call $B30A
7884  CD CE B2           call $B2CE
7887  CD 6F 74           call $746F
788A  C3 17 B5           jp $B517
788D  02                 ld (bc),a
788E  00                 nop
788F  06 43              ld b,67
7891  4A                 ld c,d
7892  27                 daa
7893  53                 ld d,e
7894  20 45              jr nz,$78DB
7896  4C                 ld c,h
7897  45                 ld b,l
7898  50                 ld d,b
7899  48                 ld c,b
789A  41                 ld b,c
789B  4E                 ld c,(hl)
789C  54                 ld d,h
789D  20 41              jr nz,$78E0
789F  4E                 ld c,(hl)
78A0  54                 ld d,h
78A1  49                 ld c,c
78A2  43                 ld b,e
78A3  53                 ld d,e
78A4  02                 ld (bc),a
78A5  02                 ld (bc),a
78A6  09                 add hl,bc
78A7  42                 ld b,d
78A8  59                 ld e,c
78A9  3A 20 44           ld a,($4420)
78AC  41                 ld b,c
78AD  56                 ld d,(hl)
78AE  45                 ld b,l
78AF  20 43              jr nz,$78F4
78B1  4C                 ld c,h
78B2  41                 ld b,c
78B3  52                 ld d,d
78B4  4B                 ld c,e
78B5  01 43 4F           ld bc,$4F43
78B8  4E                 ld c,(hl)
78B9  56                 ld d,(hl)
78BA  45                 ld b,l
78BB  52                 ld d,d
78BC  53                 ld d,e
78BD  49                 ld c,c
78BE  4F                 ld c,a
78BF  4E                 ld c,(hl)
78C0  20 42              jr nz,$7904
78C2  59                 ld e,c
78C3  3A 01 20           ld a,($2001)
78C6  52                 ld d,d
78C7  20 2D              jr nz,$78F6
78C9  20 46              jr nz,$7911
78CB  52                 ld d,d
78CC  45                 ld b,l
78CD  44                 ld b,h
78CE  20 2D              jr nz,$78FD
78D0  20 57              jr nz,$7929
78D2  01 20 20           ld bc,$2020
78D5  4F                 ld c,a
78D6  46                 ld b,(hl)
78D7  20 42              jr nz,$791B
78D9  49                 ld c,c
78DA  47                 ld b,a
78DB  20 52              jr nz,$792F
78DD  45                 ld b,l
78DE  44                 ld b,h
78DF  01 20 20           ld bc,$2020
78E2  20 53              jr nz,$7937
78E4  4F                 ld c,a
78E5  46                 ld b,(hl)
78E6  54                 ld d,h
78E7  57                 ld d,a
78E8  41                 ld b,c
78E9  52                 ld d,d
78EA  45                 ld b,l
78EB  02                 ld (bc),a
78EC  0C                 inc c
78ED  08                 ex af,af'
78EE  46                 ld b,(hl)
78EF  4F                 ld c,a
78F0  52                 ld d,d
78F1  20 20              jr nz,$7913
78F3  43                 ld b,e
78F4  4F                 ld c,a
78F5  44                 ld b,h
78F6  45                 ld b,l
78F7  4D                 ld c,l
78F8  41                 ld b,c
78F9  53                 ld d,e
78FA  54                 ld d,h
78FB  45                 ld b,l
78FC  52                 ld d,d
78FD  53                 ld d,e
78FE  02                 ld (bc),a
78FF  13                 inc de
7900  08                 ex af,af'
7901  48                 ld c,b
7902  49                 ld c,c
7903  54                 ld d,h
7904  20 46              jr nz,$794C
7906  49                 ld c,c
7907  52                 ld d,d
7908  45                 ld b,l
7909  20 54              jr nz,$795F
790B  4F                 ld c,a
790C  20 53              jr nz,$7961
790E  54                 ld d,h
790F  41                 ld b,c
7910  52                 ld d,d
7911  54                 ld d,h
7912  02                 ld (bc),a
7913  0E 0C              ld c,12
7915  20 4D              jr nz,$7964
7917  43                 ld b,e
7918  4D                 ld c,l
7919  58                 ld e,b
791A  43                 ld b,e
791B  49                 ld c,c
791C  81                 add a,c
791D  02                 ld (bc),a
791E  FF                 rst 38h
791F  00                 nop
7920  FF                 rst 38h
7921  03                 inc bc
7922  FF                 rst 38h
7923  06 07              ld b,7
7925  08                 ex af,af'
7926  00                 nop
7927  01 05 00           ld bc,5
792A  04                 inc b
792B  39                 add hl,sp
792C  79                 ld a,c
792D  41                 ld b,c
792E  79                 ld a,c
792F  49                 ld c,c
7930  79                 ld a,c
7931  51                 ld d,c
7932  79                 ld a,c
7933  67                 ld h,a
7934  79                 ld a,c
7935  7B                 ld a,e
7936  79                 ld a,c
7937  88                 adc a,b
7938  79                 ld a,c
7939  51                 ld d,c
793A  41                 ld b,c
793B  4F                 ld c,a
793C  50                 ld d,b
793D  20 53              jr nz,$7992
793F  50                 ld d,b
7940  C3 53 49           jp $4953
7943  4E                 ld c,(hl)
7944  43                 ld b,e
7945  4C                 ld c,h
7946  41                 ld b,c
7947  49                 ld c,c
7948  D2 4B 45           jp nc,$454B
794B  4D                 ld c,l
794C  50                 ld d,b
794D  53                 ld d,e
794E  54                 ld d,h
794F  4F                 ld c,a
7950  CE 51              adc a,81
7952  41                 ld b,c
7953  43                 ld b,e
7954  56                 ld d,(hl)
7955  20 53              jr nz,$79AA
7957  50                 ld d,b
7958  43                 ld b,e
7959  00                 nop
795A  50                 ld d,b
795B  20 45              jr nz,$79A2
795D  4E                 ld c,(hl)
795E  54                 ld d,h
795F  20 49              jr nz,$79AA
7961  4F                 ld c,a
7962  20 43              jr nz,$79A7
7964  41                 ld b,c
7965  50                 ld d,b
7966  D3 51              out ($51),a
7968  41                 ld b,c
7969  50                 ld d,b
796A  20 45              jr nz,$79B1
796C  4E                 ld c,(hl)
796D  54                 ld d,h
796E  20 53              jr nz,$79C3
7970  50                 ld d,b
7971  43                 ld b,e
7972  00                 nop
7973  53                 ld d,e
7974  49                 ld c,c
7975  4E                 ld c,(hl)
7976  43                 ld b,e
7977  4C                 ld c,h
7978  41                 ld b,c
7979  49                 ld c,c
797A  D2 51 41           jp nc,$4151
797D  45                 ld b,l
797E  52                 ld d,d
797F  20 5A              jr nz,$79DB
7981  00                 nop
7982  55                 ld d,l
7983  4A                 ld c,d
7984  4F                 ld c,a
7985  50                 ld d,b
7986  20 CD              jr nz,$7955
7988  53                 ld d,e
7989  49                 ld c,c
798A  4E                 ld c,(hl)
798B  43                 ld b,e
798C  4C                 ld c,h
798D  41                 ld b,c
798E  49                 ld c,c
798F  52                 ld d,d
7990  D3 3A              out ($3A),a
7992  FF                 rst 38h
7993  FF                 rst 38h
7994  A7                 and a
7995  C0                 ret nz
7996  21 9C 79           ld hl,$799C
7999  C3 CE B2           jp $B2CE
799C  02                 ld (bc),a
799D  00                 nop
799E  00                 nop
799F  C5                 push bc
79A0  21 03 0B           ld hl,$0B03
79A3  CD C4 B2           call $B2C4
79A6  21 DA 75           ld hl,$75DA
79A9  06 08              ld b,8
79AB  C5                 push bc
79AC  06 04              ld b,4
79AE  7E                 ld a,(hl)
79AF  C5                 push bc
79B0  CD 4E B2           call $B24E
79B3  C1                 pop bc
79B4  23                 inc hl
79B5  10 F7              djnz $79AE
79B7  3E 20              ld a,32
79B9  CD 4E B2           call $B24E
79BC  06 03              ld b,3
79BE  7E                 ld a,(hl)
79BF  23                 inc hl
79C0  C5                 push bc
79C1  CD 1F B2           call $B21F
79C4  C1                 pop bc
79C5  10 F7              djnz $79BE
79C7  E5                 push hl
79C8  CD BD B2           call $B2BD
79CB  E1                 pop hl
79CC  C1                 pop bc
79CD  10 DC              djnz $79AB
79CF  C3 6F 74           jp $746F
79D2  ED 53 A4 B0        ld ($B0A4),de
79D6  22 A2 B0           ld ($B0A2),hl
79D9  11 1E 00           ld de,30
79DC  19                 add hl,de
79DD  22 A6 B0           ld ($B0A6),hl
79E0  11 00 01           ld de,256
79E3  ED 53 A8 B0        ld ($B0A8),de
79E7  CD 21 7A           call $7A21
79EA  FE 1F              cp 31
79EC  28 08              jr z,$79F6
79EE  CD 0B 7A           call $7A0B
79F1  CD 18 7A           call $7A18
79F4  18 F1              jr $79E7
79F6  CD 21 7A           call $7A21
79F9  3D                 dec a
79FA  C8                 ret z
79FB  3C                 inc a
79FC  F5                 push af
79FD  CD 21 7A           call $7A21
7A00  CD 0B 7A           call $7A0B
7A03  C1                 pop bc
7A04  CD 18 7A           call $7A18
7A07  10 FB              djnz $7A04
7A09  18 DC              jr $79E7
7A0B  FE 1E              cp 30
7A0D  28 20              jr z,$7A2F
7A0F  2A A2 B0           ld hl,($B0A2)
7A12  5F                 ld e,a
7A13  16 00              ld d,0
7A15  19                 add hl,de
7A16  7E                 ld a,(hl)
7A17  C9                 ret
7A18  2A A4 B0           ld hl,($B0A4)
7A1B  77                 ld (hl),a
7A1C  23                 inc hl
7A1D  22 A4 B0           ld ($B0A4),hl
7A20  C9                 ret
7A21  06 05              ld b,5
7A23  CD 37 7A           call $7A37
7A26  10 FB              djnz $7A23
7A28  CB 3F              srl a
7A2A  CB 3F              srl a
7A2C  CB 3F              srl a
7A2E  C9                 ret
7A2F  06 08              ld b,8
7A31  CD 37 7A           call $7A37
7A34  10 FB              djnz $7A31
7A36  C9                 ret
7A37  ED 5B A8 B0        ld de,($B0A8)
7A3B  15                 dec d
7A3C  20 0A              jr nz,$7A48
7A3E  2A A6 B0           ld hl,($B0A6)
7A41  5E                 ld e,(hl)
7A42  23                 inc hl
7A43  22 A6 B0           ld ($B0A6),hl
7A46  16 08              ld d,8
7A48  CB 13              rl e
7A4A  1F                 rra
7A4B  ED 53 A8 B0        ld ($B0A8),de
7A4F  C9                 ret
7A50  21 00 EB           ld hl,$EB00
7A53  22 A6 B0           ld ($B0A6),hl
7A56  2A A4 B0           ld hl,($B0A4)
7A59  22 A2 B0           ld ($B0A2),hl
7A5C  2A 7F B0           ld hl,($B07F)
7A5F  2E 00              ld l,0
7A61  CB 3C              srl h
7A63  CB 1D              rr l
7A65  22 F2 FF           ld ($FFF2),hl
7A68  CD 14 7B           call $7B14
7A6B  3E 00              ld a,0
7A6D  F5                 push af
7A6E  CD 39 7B           call $7B39
7A71  2A A4 B0           ld hl,($B0A4)
7A74  3A F1 FF           ld a,($FFF1)
7A77  77                 ld (hl),a
7A78  23                 inc hl
7A79  22 A4 B0           ld ($B0A4),hl
7A7C  21 00 72           ld hl,$7200
7A7F  6F                 ld l,a
7A80  36 00              ld (hl),0
7A82  24                 inc h
7A83  36 00              ld (hl),0
7A85  F1                 pop af
7A86  3C                 inc a
7A87  FE 1E              cp 30
7A89  20 E2              jr nz,$7A6D
7A8B  CD D2 7A           call $7AD2
7A8E  CD 94 7B           call $7B94
7A91  ED 4B F2 FF        ld bc,($FFF2)
7A95  C5                 push bc
7A96  2A A6 B0           ld hl,($B0A6)
7A99  CD BB 7A           call $7ABB
7A9C  2A A6 B0           ld hl,($B0A6)
7A9F  23                 inc hl
7AA0  22 A6 B0           ld ($B0A6),hl
7AA3  C1                 pop bc
7AA4  0B                 dec bc
7AA5  78                 ld a,b
7AA6  B1                 or c
7AA7  20 EC              jr nz,$7A95
7AA9  CD DA 7A           call $7ADA
7AAC  3E 1F              ld a,31
7AAE  CD 7C 7B           call $7B7C
7AB1  3E 01              ld a,1
7AB3  CD 7C 7B           call $7B7C
7AB6  AF                 xor a
7AB7  CD 74 7B           call $7B74
7ABA  C9                 ret
7ABB  46                 ld b,(hl)
7ABC  78                 ld a,b
7ABD  32 EE FF           ld ($FFEE),a
7AC0  3A EF FF           ld a,($FFEF)
7AC3  B8                 cp b
7AC4  20 14              jr nz,$7ADA
7AC6  3A F0 FF           ld a,($FFF0)
7AC9  FE 1F              cp 31
7ACB  28 0D              jr z,$7ADA
7ACD  3C                 inc a
7ACE  32 F0 FF           ld ($FFF0),a
7AD1  C9                 ret
7AD2  AF                 xor a
7AD3  32 F0 FF           ld ($FFF0),a
7AD6  32 EF FF           ld ($FFEF),a
7AD9  C9                 ret
7ADA  3A F0 FF           ld a,($FFF0)
7ADD  FE 03              cp 3
7ADF  38 13              jr c,$7AF4
7AE1  3E 1F              ld a,31
7AE3  CD 7C 7B           call $7B7C
7AE6  3A F0 FF           ld a,($FFF0)
7AE9  CD 7C 7B           call $7B7C
7AEC  3A EF FF           ld a,($FFEF)
7AEF  CD 58 7B           call $7B58
7AF2  18 14              jr $7B08
7AF4  3A F0 FF           ld a,($FFF0)
7AF7  A7                 and a
7AF8  28 0E              jr z,$7B08
7AFA  47                 ld b,a
7AFB  C5                 push bc
7AFC  3A EF FF           ld a,($FFEF)
7AFF  CD 58 7B           call $7B58
7B02  C1                 pop bc
7B03  10 F6              djnz $7AFB
7B05  C3 08 7B           jp $7B08
7B08  3A EE FF           ld a,($FFEE)
7B0B  32 EF FF           ld ($FFEF),a
7B0E  3E 01              ld a,1
7B10  32 F0 FF           ld ($FFF0),a
7B13  C9                 ret
7B14  21 00 72           ld hl,$7200
7B17  11 01 72           ld de,$7201
7B1A  01 FF 01           ld bc,511
7B1D  AF                 xor a
7B1E  77                 ld (hl),a
7B1F  ED B0              ldir
7B21  11 00 EB           ld de,$EB00
7B24  ED 4B F2 FF        ld bc,($FFF2)
7B28  1A                 ld a,(de)
7B29  6F                 ld l,a
7B2A  26 72              ld h,114
7B2C  34                 inc (hl)
7B2D  20 03              jr nz,$7B32
7B2F  26 73              ld h,115
7B31  34                 inc (hl)
7B32  13                 inc de
7B33  0B                 dec bc
7B34  78                 ld a,b
7B35  B1                 or c
7B36  20 F0              jr nz,$7B28
7B38  C9                 ret
7B39  2E 00              ld l,0
7B3B  01 00 00           ld bc,0
7B3E  26 72              ld h,114
7B40  5E                 ld e,(hl)
7B41  26 73              ld h,115
7B43  56                 ld d,(hl)
7B44  7A                 ld a,d
7B45  B8                 cp b
7B46  38 0C              jr c,$7B54
7B48  20 04              jr nz,$7B4E
7B4A  7B                 ld a,e
7B4B  B9                 cp c
7B4C  38 06              jr c,$7B54
7B4E  7D                 ld a,l
7B4F  32 F1 FF           ld ($FFF1),a
7B52  42                 ld b,d
7B53  4B                 ld c,e
7B54  2C                 inc l
7B55  20 E7              jr nz,$7B3E
7B57  C9                 ret
7B58  2A A2 B0           ld hl,($B0A2)
7B5B  06 1E              ld b,30
7B5D  BE                 cp (hl)
7B5E  28 0E              jr z,$7B6E
7B60  23                 inc hl
7B61  10 FA              djnz $7B5D
7B63  F5                 push af
7B64  3E 1E              ld a,30
7B66  CD 7C 7B           call $7B7C
7B69  F1                 pop af
7B6A  CD 74 7B           call $7B74
7B6D  C9                 ret
7B6E  3E 1E              ld a,30
7B70  90                 sub b
7B71  C3 7C 7B           jp $7B7C
7B74  06 08              ld b,8
7B76  CD 84 7B           call $7B84
7B79  10 FB              djnz $7B76
7B7B  C9                 ret
7B7C  06 05              ld b,5
7B7E  CD 84 7B           call $7B84
7B81  10 FB              djnz $7B7E
7B83  C9                 ret
7B84  ED 5B A8 B0        ld de,($B0A8)
7B88  1F                 rra
7B89  CB 13              rl e
7B8B  15                 dec d
7B8C  CC 9C 7B           call z,$7B9C
7B8F  ED 53 A8 B0        ld ($B0A8),de
7B93  C9                 ret
7B94  11 00 08           ld de,2048
7B97  ED 53 A8 B0        ld ($B0A8),de
7B9B  C9                 ret
7B9C  2A A4 B0           ld hl,($B0A4)
7B9F  73                 ld (hl),e
7BA0  23                 inc hl
7BA1  22 A4 B0           ld ($B0A4),hl
7BA4  16 08              ld d,8
7BA6  C9                 ret
7BA7  3A 72 B0           ld a,($B072)
7BAA  FE 01              cp 1
7BAC  20 0C              jr nz,$7BBA
7BAE  3E 02              ld a,2
7BB0  32 6E F9           ld ($F96E),a
7BB3  32 EE F9           ld ($F9EE),a
7BB6  32 6E FA           ld ($FA6E),a
7BB9  C9                 ret
7BBA  FE 03              cp 3
7BBC  C0                 ret nz
7BBD  3E 02              ld a,2
7BBF  32 CF FD           ld ($FDCF),a
7BC2  32 4F FE           ld ($FE4F),a
7BC5  32 CF FE           ld ($FECF),a
7BC8  32 E0 FD           ld ($FDE0),a
7BCB  32 60 FE           ld ($FE60),a
7BCE  32 E0 FE           ld ($FEE0),a
7BD1  C9                 ret
7BD2  3A 73 B0           ld a,($B073)
7BD5  FE 01              cp 1
7BD7  20 0E              jr nz,$7BE7
7BD9  3E 80              ld a,-128
7BDB  32 6E F9           ld ($F96E),a
7BDE  32 EE F9           ld ($F9EE),a
7BE1  3E 08              ld a,8
7BE3  32 6E FA           ld ($FA6E),a
7BE6  C9                 ret
7BE7  FE 03              cp 3
7BE9  C0                 ret nz
7BEA  3E 80              ld a,-128
7BEC  32 CF FD           ld ($FDCF),a
7BEF  32 4F FE           ld ($FE4F),a
7BF2  32 E0 FD           ld ($FDE0),a
7BF5  32 60 FE           ld ($FE60),a
7BF8  3E 6D              ld a,109
7BFA  32 CF FE           ld ($FECF),a
7BFD  32 E0 FE           ld ($FEE0),a
7C00  C9                 ret
7C01  3A 84 B0           ld a,($B084)
7C04  A7                 and a
7C05  C0                 ret nz
7C06  3A 72 B0           ld a,($B072)
7C09  FE 02              cp 2
7C0B  20 0A              jr nz,$7C17
7C0D  CD DE E9           call $E9DE
7C10  E6 1F              and $1F
7C12  20 03              jr nz,$7C17
7C14  CD 3D CA           call $CA3D
7C17  21 DF D2           ld hl,$D2DF
7C1A  7E                 ld a,(hl)
7C1B  A7                 and a
7C1C  F2 22 7C           jp p,$7C22
7C1F  AF                 xor a
7C20  36 00              ld (hl),0
7C22  A7                 and a
7C23  20 79              jr nz,$7C9E
7C25  CD 9D C5           call $C59D
7C28  11 BE 01           ld de,446
7C2B  01 E8 01           ld bc,488
7C2E  3A 72 B0           ld a,($B072)
7C31  FE 02              cp 2
7C33  20 06              jr nz,$7C3B
7C35  11 B6 01           ld de,438
7C38  01 F0 01           ld bc,496
7C3B  2A DA D2           ld hl,($D2DA)
7C3E  A7                 and a
7C3F  ED 52              sbc hl,de
7C41  19                 add hl,de
7C42  38 0F              jr c,$7C53
7C44  11 E8 01           ld de,488
7C47  A7                 and a
7C48  ED 42              sbc hl,bc
7C4A  09                 add hl,bc
7C4B  30 0D              jr nc,$7C5A
7C4D  3A DC D2           ld a,($D2DC)
7C50  17                 rla
7C51  30 07              jr nc,$7C5A
7C53  3E 80              ld a,-128
7C55  01 01 00           ld bc,1
7C58  18 04              jr $7C5E
7C5A  AF                 xor a
7C5B  01 FF FF           ld bc,-1
7C5E  09                 add hl,bc
7C5F  22 DA D2           ld ($D2DA),hl
7C62  CD 17 C6           call $C617
7C65  11 DE D2           ld de,$D2DE
7C68  1A                 ld a,(de)
7C69  3C                 inc a
7C6A  3C                 inc a
7C6B  12                 ld (de),a
7C6C  2A D8 D2           ld hl,($D2D8)
7C6F  4F                 ld c,a
7C70  06 00              ld b,0
7C72  17                 rla
7C73  30 01              jr nc,$7C76
7C75  05                 dec b
7C76  2A D8 D2           ld hl,($D2D8)
7C79  09                 add hl,bc
7C7A  22 D8 D2           ld ($D2D8),hl
7C7D  01 DF 01           ld bc,479
7C80  3A 72 B0           ld a,($B072)
7C83  FE 01              cp 1
7C85  28 03              jr z,$7C8A
7C87  01 6F 02           ld bc,623
7C8A  A7                 and a
7C8B  ED 42              sbc hl,bc
7C8D  09                 add hl,bc
7C8E  D8                 ret c
7C8F  ED 43 D8 D2        ld ($D2D8),bc
7C93  CD DE E9           call $E9DE
7C96  E6 07              and $07
7C98  C6 03              add a,3
7C9A  32 DF D2           ld ($D2DF),a
7C9D  C9                 ret
7C9E  35                 dec (hl)
7C9F  CD C4 C5           call $C5C4
7CA2  21 DC D2           ld hl,$D2DC
7CA5  7E                 ld a,(hl)
7CA6  EE 80              xor $80
7CA8  77                 ld (hl),a
7CA9  3A DF D2           ld a,($D2DF)
7CAC  FE 03              cp 3
7CAE  D0                 ret nc
7CAF  CD B7 C5           call $C5B7
7CB2  CD DE E9           call $E9DE
7CB5  E6 03              and $03
7CB7  C6 0F              add a,15
7CB9  ED 44              neg
7CBB  32 DE D2           ld ($D2DE),a
7CBE  C3 20 C6           jp $C620
7CC1  3E FF              ld a,-1
7CC3  21 6A B0           ld hl,$B06A
7CC6  06 08              ld b,8
7CC8  A6                 and (hl)
7CC9  23                 inc hl
7CCA  10 FC              djnz $7CC8
7CCC  E6 1F              and $1F
7CCE  EE 1F              xor $1F
7CD0  C9                 ret
7CD1  E5                 push hl
7CD2  7C                 ld a,h
7CD3  F9                 ld sp,hl
7CD4  7C                 ld a,h
7CD5  13                 inc de
7CD6  7D                 ld a,l
7CD7  27                 daa
7CD8  7D                 ld a,l
7CD9  41                 ld b,c
7CDA  7D                 ld a,l
7CDB  67                 ld h,a
7CDC  7D                 ld a,l
7CDD  8D                 adc a,l
7CDE  7D                 ld a,l
7CDF  B3                 or e
7CE0  7D                 ld a,l
7CE1  D9                 exx
7CE2  7D                 ld a,l
7CE3  FF                 rst 38h
7CE4  7D                 ld a,l
7CE5  00                 nop
7CE6  5B                 ld e,e
7CE7  1E 04              ld e,4
7CE9  10 02              djnz $7CED
7CEB  78                 ld a,b
7CEC  5B                 ld e,e
7CED  10 04              djnz $7CF3
7CEF  00                 nop
7CF0  02                 ld (bc),a
7CF1  38 5C              jr c,$7D4F
7CF3  18 02              jr $7CF7
7CF5  18 00              jr $7CF7
7CF7  00                 nop
7CF8  00                 nop
7CF9  00                 nop
7CFA  5B                 ld e,e
7CFB  1E 04              ld e,4
7CFD  10 02              djnz $7D01
7CFF  B8                 cp b
7D00  5B                 ld e,e
7D01  10 04              djnz $7D07
7D03  00                 nop
7D04  02                 ld (bc),a
7D05  68                 ld l,b
7D06  5C                 ld e,h
7D07  10 02              djnz $7D0B
7D09  18 00              jr $7D0B
7D0B  88                 adc a,b
7D0C  5C                 ld e,h
7D0D  05                 dec b
7D0E  02                 ld (bc),a
7D0F  20 04              jr nz,$7D15
7D11  00                 nop
7D12  00                 nop
7D13  00                 nop
7D14  5B                 ld e,e
7D15  1E 04              ld e,4
7D17  10 02              djnz $7D1B
7D19  F8                 ret m
7D1A  5B                 ld e,e
7D1B  10 04              djnz $7D21
7D1D  00                 nop
7D1E  02                 ld (bc),a
7D1F  38 5C              jr c,$7D7D
7D21  18 02              jr $7D25
7D23  18 00              jr $7D25
7D25  00                 nop
7D26  00                 nop
7D27  00                 nop
7D28  5B                 ld e,e
7D29  1E 04              ld e,4
7D2B  10 02              djnz $7D2F
7D2D  B8                 cp b
7D2E  5B                 ld e,e
7D2F  10 04              djnz $7D35
7D31  00                 nop
7D32  02                 ld (bc),a
7D33  38 5C              jr c,$7D91
7D35  18 02              jr $7D39
7D37  18 00              jr $7D39
7D39  92                 sub d
7D3A  5C                 ld e,h
7D3B  04                 inc b
7D3C  02                 ld (bc),a
7D3D  18 04              jr $7D43
7D3F  00                 nop
7D40  00                 nop
7D41  9A                 sbc a,d
7D42  5C                 ld e,h
7D43  0C                 inc c
7D44  02                 ld (bc),a
7D45  20 04              jr nz,$7D4B
7D47  30 5D              jr nc,$7DA6
7D49  10 02              djnz $7D4D
7D4B  10 04              djnz $7D51
7D4D  0A                 ld a,(bc)
7D4E  5D                 ld e,l
7D4F  13                 inc de
7D50  02                 ld (bc),a
7D51  10 00              djnz $7D53
7D53  50                 ld d,b
7D54  5D                 ld e,l
7D55  13                 inc de
7D56  02                 ld (bc),a
7D57  10 08              djnz $7D61
7D59  CE 5D              adc a,93
7D5B  10 02              djnz $7D5F
7D5D  00                 nop
7D5E  02                 ld (bc),a
7D5F  0E 5E              ld c,94
7D61  10 02              djnz $7D65
7D63  00                 nop
7D64  06 00              ld b,0
7D66  00                 nop
7D67  9A                 sbc a,d
7D68  5C                 ld e,h
7D69  0C                 inc c
7D6A  02                 ld (bc),a
7D6B  20 04              jr nz,$7D71
7D6D  30 5D              jr nc,$7DCC
7D6F  10 02              djnz $7D73
7D71  10 04              djnz $7D77
7D73  B2                 or d
7D74  5C                 ld e,h
7D75  18 02              jr $7D79
7D77  10 00              djnz $7D79
7D79  9E                 sbc a,(hl)
7D7A  5D                 ld e,l
7D7B  18 02              jr $7D7F
7D7D  10 08              djnz $7D87
7D7F  CE 5D              adc a,93
7D81  10 02              djnz $7D85
7D83  00                 nop
7D84  02                 ld (bc),a
7D85  0E 5E              ld c,94
7D87  10 02              djnz $7D8B
7D89  00                 nop
7D8A  06 00              ld b,0
7D8C  00                 nop
7D8D  9A                 sbc a,d
7D8E  5C                 ld e,h
7D8F  0C                 inc c
7D90  02                 ld (bc),a
7D91  20 04              jr nz,$7D97
7D93  30 5D              jr nc,$7DF2
7D95  10 02              djnz $7D99
7D97  10 04              djnz $7D9D
7D99  E2 5C 14           jp po,$145C
7D9C  02                 ld (bc),a
7D9D  10 00              djnz $7D9F
7D9F  76                 halt
7DA0  5D                 ld e,l
7DA1  14                 inc d
7DA2  02                 ld (bc),a
7DA3  10 08              djnz $7DAD
7DA5  EE 5D              xor $5D
7DA7  10 02              djnz $7DAB
7DA9  00                 nop
7DAA  02                 ld (bc),a
7DAB  2E 5E              ld l,94
7DAD  10 02              djnz $7DB1
7DAF  00                 nop
7DB0  06 00              ld b,0
7DB2  00                 nop
7DB3  4E                 ld c,(hl)
7DB4  5E                 ld e,(hl)
7DB5  10 02              djnz $7DB9
7DB7  20 04              jr nz,$7DBD
7DB9  22 5F 18           ld ($185F),hl
7DBC  02                 ld (bc),a
7DBD  08                 ex af,af'
7DBE  04                 inc b
7DBF  E8                 ret pe
7DC0  5E                 ld e,(hl)
7DC1  1D                 dec e
7DC2  02                 ld (bc),a
7DC3  10 00              djnz $7DC5
7DC5  8C                 adc a,h
7DC6  5F                 ld e,a
7DC7  1D                 dec e
7DC8  02                 ld (bc),a
7DC9  10 08              djnz $7DD3
7DCB  D4 5F 10           call nc,$105F
7DCE  02                 ld (bc),a
7DCF  00                 nop
7DD0  00                 nop
7DD1  14                 inc d
7DD2  60                 ld h,b
7DD3  07                 rlca
7DD4  02                 ld (bc),a
7DD5  00                 nop
7DD6  06 00              ld b,0
7DD8  00                 nop
7DD9  6E                 ld l,(hl)
7DDA  5E                 ld e,(hl)
7DDB  10 02              djnz $7DDF
7DDD  20 04              jr nz,$7DE3
7DDF  22 5F 18           ld ($185F),hl
7DE2  02                 ld (bc),a
7DE3  08                 ex af,af'
7DE4  04                 inc b
7DE5  AE                 xor (hl)
7DE6  5E                 ld e,(hl)
7DE7  1D                 dec e
7DE8  02                 ld (bc),a
7DE9  10 00              djnz $7DEB
7DEB  8C                 adc a,h
7DEC  5F                 ld e,a
7DED  1D                 dec e
7DEE  02                 ld (bc),a
7DEF  10 08              djnz $7DF9
7DF1  C6 5F              add a,95
7DF3  07                 rlca
7DF4  02                 ld (bc),a
7DF5  00                 nop
7DF6  02                 ld (bc),a
7DF7  14                 inc d
7DF8  60                 ld h,b
7DF9  07                 rlca
7DFA  02                 ld (bc),a
7DFB  00                 nop
7DFC  06 00              ld b,0
7DFE  00                 nop
7DFF  8E                 adc a,(hl)
7E00  5E                 ld e,(hl)
7E01  10 02              djnz $7E05
7E03  20 04              jr nz,$7E09
7E05  22 5F 18           ld ($185F),hl
7E08  02                 ld (bc),a
7E09  08                 ex af,af'
7E0A  04                 inc b
7E0B  AE                 xor (hl)
7E0C  5E                 ld e,(hl)
7E0D  1D                 dec e
7E0E  02                 ld (bc),a
7E0F  10 00              djnz $7E11
7E11  52                 ld d,d
7E12  5F                 ld e,a
7E13  1D                 dec e
7E14  02                 ld (bc),a
7E15  10 08              djnz $7E1F
7E17  C6 5F              add a,95
7E19  07                 rlca
7E1A  02                 ld (bc),a
7E1B  00                 nop
7E1C  02                 ld (bc),a
7E1D  F4 5F 10           call p,$105F
7E20  02                 ld (bc),a
7E21  00                 nop
7E22  08                 ex af,af'
7E23  00                 nop
7E24  00                 nop
7E25  14                 inc d
7E26  7E                 ld a,(hl)
7E27  00                 nop
7E28  6E                 ld l,(hl)
7E29  00                 nop
7E2A  01 0A 11           ld bc,$110A
7E2D  1F                 rra
7E2E  02                 ld (bc),a
7E2F  01 11 83           ld bc,$8311
7E32  23                 inc hl
7E33  01 7A 0F           ld bc,$0F7A
7E36  1A                 ld a,(de)
7E37  01 36 0E           ld bc,$0E36
7E3A  05                 dec b
7E3B  03                 inc bc
7E3C  31 0C BA           ld sp,$BA0C
7E3F  36 0A              ld (hl),10
7E41  0E 0D              ld c,13
7E43  06 06              ld b,6
7E45  0B                 dec bc
7E46  04                 inc b
7E47  0B                 dec bc
7E48  0C                 inc c
7E49  21 03 A3           ld hl,$A303
7E4C  D6 34              sub 52
7E4E  0B                 dec bc
7E4F  5D                 ld e,l
7E50  12                 ld (de),a
7E51  DC DB ED           call c,$EDDB
7E54  6C                 ld l,h
7E55  0A                 ld a,(bc)
7E56  27                 daa
7E57  11 0F 09           ld de,$090F
7E5A  1B                 dec de
7E5B  0C                 inc c
7E5C  1B                 dec de
7E5D  12                 ld (de),a
7E5E  A2                 and d
7E5F  CB 2B              sra e
7E61  02                 ld (bc),a
7E62  57                 ld d,a
7E63  0B                 dec bc
7E64  68                 ld l,b
7E65  0F                 rrca
7E66  67                 ld h,a
7E67  0F                 rrca
7E68  23                 inc hl
7E69  08                 ex af,af'
7E6A  10 06              djnz $7E72
7E6C  40                 ld b,b
7E6D  08                 ex af,af'
7E6E  DE D8              sbc a,-40
7E70  3F                 ccf
7E71  0C                 inc c
7E72  C2 48 12           jp nz,$1248
7E75  11 09 1C           ld de,$1C09
7E78  06 16              ld b,22
7E7A  03                 inc bc
7E7B  9A                 sbc a,d
7E7C  1C                 inc e
7E7D  09                 add hl,bc
7E7E  CA 24 0E           jp z,$0E24
7E81  5D                 ld e,l
7E82  05                 dec b
7E83  E0                 ret po
7E84  5C                 ld e,h
7E85  02                 ld (bc),a
7E86  FF                 rst 38h
7E87  1F                 rra
7E88  7E                 ld a,(hl)
7E89  15                 dec d
7E8A  6E                 ld l,(hl)
7E8B  00                 nop
7E8C  01 09 07           ld bc,1801
7E8F  24                 inc h
7E90  01 7A 1D           ld bc,$1D7A
7E93  1B                 dec de
7E94  03                 inc bc
7E95  70                 ld (hl),b
7E96  16 F8              ld d,-8
7E98  FF                 rst 38h
7E99  06 02              ld b,2
7E9B  6E                 ld l,(hl)
7E9C  0A                 ld a,(bc)
7E9D  6C                 ld l,h
7E9E  0B                 dec bc
7E9F  1A                 ld a,(de)
7EA0  01 11 03           ld bc,785
7EA3  12                 ld (de),a
7EA4  18 02              jr $7EA8
7EA6  07                 rlca
7EA7  91                 sub c
7EA8  05                 dec b
7EA9  12                 ld (de),a
7EAA  E2 01 1D           jp po,$1D01
7EAD  83                 add a,e
7EAE  85                 add a,l
7EAF  87                 add a,a
7EB0  CE CF              adc a,-49
7EB2  D0                 ret nc
7EB3  D4 22 0C           call nc,$0C22
7EB6  21 0A 3C           ld hl,$3C0A
7EB9  18 B8              jr $7E73
7EBB  B3                 or e
7EBC  AF                 xor a
7EBD  AA                 xor d
7EBE  B5                 or l
7EBF  2A 1C B1           ld hl,($B11C)
7EC2  B8                 cp b
7EC3  BF                 cp a
7EC4  13                 inc de
7EC5  0B                 dec bc
7EC6  0A                 ld a,(bc)
7EC7  18 21              jr $7EEA
7EC9  06 4C              ld b,76
7ECB  09                 add hl,bc
7ECC  64                 ld h,h
7ECD  12                 ld (de),a
7ECE  2C                 inc l
7ECF  03                 inc bc
7ED0  AE                 xor (hl)
7ED1  B0                 or b
7ED2  B2                 or d
7ED3  BB                 cp e
7ED4  72                 ld (hl),d
7ED5  14                 inc d
7ED6  EF                 rst 28h
7ED7  14                 inc d
7ED8  09                 add hl,bc
7ED9  0A                 ld a,(bc)
7EDA  12                 ld (de),a
7EDB  22 10 26           ld ($2610),hl
7EDE  05                 dec b
7EDF  49                 ld c,c
7EE0  0D                 dec c
7EE1  4F                 ld c,a
7EE2  04                 inc b
7EE3  4D                 ld c,l
7EE4  17                 rla
7EE5  7B                 ld a,e
7EE6  09                 add hl,bc
7EE7  45                 ld b,l
7EE8  15                 dec d
7EE9  C9                 ret
7EEA  FF                 rst 38h
7EEB  28 80              jr z,$7E6D
7EED  1E 6C              ld e,108
7EEF  00                 nop
7EF0  01 03 26           ld bc,$2603
7EF3  25                 dec h
7EF4  01 7B 26           ld bc,$267B
7EF7  22 04 70           ld ($7004),hl
7EFA  14                 inc d
7EFB  F3                 di
7EFC  F6 F9              or $F9
7EFE  07                 rlca
7EFF  03                 inc bc
7F00  03                 inc bc
7F01  23                 inc hl
7F02  0A                 ld a,(bc)
7F03  0E 43              ld c,67
7F05  10 05              djnz $7F0C
7F07  05                 dec b
7F08  01 1F 11           ld bc,$111F
7F0B  16 01              ld d,1
7F0D  0C                 inc c
7F0E  2D                 dec l
7F0F  07                 rlca
7F10  6D                 ld l,l
7F11  12                 ld (de),a
7F12  1A                 ld a,(de)
7F13  01 57 06           ld bc,1623
7F16  15                 dec d
7F17  0B                 dec bc
7F18  07                 rlca
7F19  11 56 14           ld de,$1456
7F1C  59                 ld e,c
7F1D  18 60              jr $7F7F
7F1F  24                 inc h
7F20  1D                 dec e
7F21  1B                 dec de
7F22  A0                 and b
7F23  08                 ex af,af'
7F24  05                 dec b
7F25  1E 02              ld e,2
7F27  21 25 72           ld hl,$7225
7F2A  0E A0              ld c,-96
7F2C  16 0C              ld d,12
7F2E  20 02              jr nz,$7F32
7F30  1E 1E              ld e,30
7F32  26 25              ld h,37
7F34  39                 add hl,sp
7F35  22 62 14           ld ($1462),hl
7F38  77                 ld (hl),a
7F39  10 76              djnz $7FB1
7F3B  1D                 dec e
7F3C  6D                 ld l,l
7F3D  19                 add hl,de
7F3E  66                 ld h,(hl)
7F3F  21 AB 1E           ld hl,$1EAB
7F42  0E A2              ld c,-94
7F44  17                 rla
7F45  10 06              djnz $7F4D
7F47  1A                 ld a,(de)
7F48  14                 inc d
7F49  11 1D 16           ld de,$161D
7F4C  16 1C              ld d,28
7F4E  1E 08              ld e,8
7F50  66                 ld h,(hl)
7F51  21 54 1F           ld hl,$1F54
7F54  70                 ld (hl),b
7F55  0A                 ld a,(bc)
7F56  F6 2F              or $2F
7F58  14                 inc d
7F59  BE                 cp (hl)
7F5A  DA 66 25           jp c,$2566
7F5D  B7                 or a
7F5E  0F                 rrca
7F5F  26 8D              ld h,-115
7F61  FF                 rst 38h
7F62  28 7F              jr z,$7FE3
7F64  1E 50              ld e,80
7F66  00                 nop
7F67  01 03 08           ld bc,$0803
7F6A  21 02 53           ld hl,$5302
7F6D  23                 inc hl
7F6E  DB 07              in a,($07)
7F70  02                 ld (bc),a
7F71  1C                 inc e
7F72  1F                 rra
7F73  3A 15 05           ld a,($0515)
7F76  03                 inc bc
7F77  33                 inc sp
7F78  0E 3A              ld c,58
7F7A  1D                 dec e
7F7B  62                 ld h,d
7F7C  09                 add hl,bc
7F7D  0E 09              ld c,9
7F7F  20 0E              jr nz,$7F8F
7F81  12                 ld (de),a
7F82  14                 inc d
7F83  14                 inc d
7F84  23                 inc hl
7F85  35                 dec (hl)
7F86  26 59              ld h,89
7F88  16 29              ld d,41
7F8A  21 47 1C           ld hl,$1C47
7F8D  FB                 ei
7F8E  DC 15 0D           call c,$0D15
7F91  08                 ex af,af'
7F92  08                 ex af,af'
7F93  8B                 adc a,e
7F94  93                 sub e
7F95  07                 rlca
7F96  26 9E              ld h,-98
7F98  4E                 ld c,(hl)
7F99  0A                 ld a,(bc)
7F9A  DA DE 22           jp c,$22DE
7F9D  19                 add hl,de
7F9E  A8                 xor b
7F9F  76                 halt
7FA0  1C                 inc e
7FA1  40                 ld b,b
7FA2  13                 inc de
7FA3  3E 0D              ld a,13
7FA5  18 08              jr $7FAF
7FA7  24                 inc h
7FA8  16 17              ld d,23
7FAA  26 79              ld h,121
7FAC  1C                 inc e
7FAD  D1                 pop de
7FAE  CD 1E 02           call $021E
7FB1  37                 scf
7FB2  22 3A 06           ld ($063A),hl
7FB5  19                 add hl,de
7FB6  06 21              ld b,33
7FB8  11 07 15           ld de,$1507
7FBB  20 1F              jr nz,$7FDC
7FBD  1A                 ld a,(de)
7FBE  20 33              jr nz,$7FF3
7FC0  1C                 inc e
7FC1  46                 ld b,(hl)
7FC2  24                 inc h
7FC3  1D                 dec e
7FC4  03                 inc bc
7FC5  76                 halt
7FC6  24                 inc h
7FC7  F8                 ret m
7FC8  FC FF FF           call m,$FFFF
7FCB  FF                 rst 38h
7FCC  FF                 rst 38h
7FCD  FF                 rst 38h
7FCE  FF                 rst 38h
7FCF  FF                 rst 38h
7FD0  FF                 rst 38h
7FD1  FF                 rst 38h
7FD2  FF                 rst 38h
7FD3  FF                 rst 38h
7FD4  FF                 rst 38h
7FD5  FF                 rst 38h
7FD6  FF                 rst 38h
7FD7  FF                 rst 38h
7FD8  FF                 rst 38h
7FD9  FF                 rst 38h
7FDA  FF                 rst 38h
7FDB  FF                 rst 38h
7FDC  FF                 rst 38h
7FDD  FF                 rst 38h
7FDE  FF                 rst 38h
7FDF  FF                 rst 38h
7FE0  FF                 rst 38h
7FE1  FF                 rst 38h
7FE2  00                 nop
7FE3  00                 nop
7FE4  00                 nop
7FE5  00                 nop
7FE6  00                 nop
7FE7  00                 nop
7FE8  81                 add a,c
7FE9  11 82 29           ld de,$2982
7FEC  84                 add a,h
7FED  2E 00              ld l,0
7FEF  00                 nop
7FF0  00                 nop
7FF1  00                 nop
7FF2  00                 nop
7FF3  11 2B 87           ld de,$872B
7FF6  37                 scf
7FF7  00                 nop
7FF8  00                 nop
7FF9  00                 nop
7FFA  00                 nop
7FFB  00                 nop
7FFC  00                 nop
7FFD  2C                 inc l
7FFE  85                 add a,l
7FFF  35                 dec (hl)
8000  00                 nop
8001  00                 nop
8002  3E 3E              ld a,62
8004  7F                 ld a,a
8005  7F                 ld a,a
8006  2A 2A 00           ld hl,($002A)
8009  00                 nop
800A  97                 sub a
800B  F3                 di
800C  CF                 rst 08h
800D  49                 ld c,c
800E  CD A5 37           call $37A5
8011  54                 ld d,h
8012  77                 ld (hl),a
8013  AA                 xor d
8014  1B                 dec de
8015  F2 61 5A           jp p,$5A61
8018  F8                 ret m
8019  FA FA 52           jp m,$52FA
801C  F3                 di
801D  A2                 and d
801E  01 9C 00           ld bc,156
8021  00                 nop
8022  F5                 push af
8023  BE                 cp (hl)
8024  8A                 adc a,d
8025  FA 55 54           jp m,$5455
8028  00                 nop
8029  00                 nop
802A  7F                 ld a,a
802B  FC D0 22           call m,$22D0
802E  FC D8 BB           call m,$BBD8
8031  B2                 or d
8032  F7                 rst 30h
8033  E8                 ret pe
8034  BE                 cp (hl)
8035  D2 FD A8           jp nc,$A8FD
8038  B3                 or e
8039  4A                 ld c,d
803A  E0                 ret po
803B  90                 sub b
803C  80                 add a,b
803D  02                 ld (bc),a
803E  55                 ld d,l
803F  54                 ld d,h
8040  00                 nop
8041  00                 nop
8042  7F                 ld a,a
8043  FC DA 86           call m,$86DA
8046  FF                 rst 38h
8047  EA BF D8           jp pe,$D8BF
804A  F7                 rst 30h
804B  BA                 cp d
804C  AD                 xor l
804D  F0                 ret p
804E  DE EA              sbc a,-22
8050  BF                 cp a
8051  D8                 ret c
8052  FB                 ei
8053  B2                 or d
8054  B7                 or a
8055  20 EE              jr nz,$8045
8057  4A                 ld c,d
8058  9C                 sbc a,h
8059  90                 sub b
805A  C0                 ret nz
805B  02                 ld (bc),a
805C  55                 ld d,l
805D  54                 ld d,h
805E  00                 nop
805F  00                 nop
8060  00                 nop
8061  00                 nop
8062  44                 ld b,h
8063  44                 ld b,h
8064  44                 ld b,h
8065  44                 ld b,h
8066  44                 ld b,h
8067  44                 ld b,h
8068  44                 ld b,h
8069  44                 ld b,h
806A  EE EE              xor $EE
806C  EE EE              xor $EE
806E  55                 ld d,l
806F  55                 ld d,l
8070  BB                 cp e
8071  BB                 cp e
8072  00                 nop
8073  00                 nop
8074  FF                 rst 38h
8075  FF                 rst 38h
8076  00                 nop
8077  00                 nop
8078  E8                 ret pe
8079  FA DA 52           jp m,$52DA
807C  F3                 di
807D  A2                 and d
807E  01 9C 00           ld bc,156
8081  00                 nop
8082  00                 nop
8083  80                 add a,b
8084  20 82              jr nz,$8008
8086  10 84              djnz $800C
8088  0D                 dec c
8089  58                 ld e,b
808A  09                 add hl,bc
808B  48                 ld c,b
808C  03                 inc bc
808D  60                 ld h,b
808E  0E 38              ld c,56
8090  70                 ld (hl),b
8091  87                 add a,a
8092  0E 38              ld c,56
8094  03                 inc bc
8095  60                 ld h,b
8096  09                 add hl,bc
8097  48                 ld c,b
8098  0D                 dec c
8099  58                 ld e,b
809A  10 84              djnz $8020
809C  20 82              jr nz,$8020
809E  00                 nop
809F  80                 add a,b
80A0  00                 nop
80A1  80                 add a,b
80A2  00                 nop
80A3  00                 nop
80A4  01 40 00           ld bc,64
80A7  C0                 ret nz
80A8  02                 ld (bc),a
80A9  60                 ld h,b
80AA  00                 nop
80AB  A0                 and b
80AC  04                 inc b
80AD  70                 ld (hl),b
80AE  00                 nop
80AF  B0                 or b
80B0  08                 ex af,af'
80B1  58                 ld e,b
80B2  00                 nop
80B3  28 10              jr z,$80C5
80B5  9C                 sbc a,h
80B6  00                 nop
80B7  2C                 inc l
80B8  20 16              jr nz,$80D0
80BA  00                 nop
80BB  4A                 ld c,d
80BC  40                 ld b,b
80BD  17                 rla
80BE  00                 nop
80BF  AB                 xor e
80C0  00                 nop
80C1  00                 nop
80C2  FE FE              cp -2
80C4  FC FC AA           call m,$AAFC
80C7  AA                 xor d
80C8  00                 nop
80C9  00                 nop
80CA  EF                 rst 28h
80CB  EF                 rst 28h
80CC  CF                 rst 08h
80CD  CF                 rst 08h
80CE  AA                 xor d
80CF  AA                 xor d
80D0  00                 nop
80D1  00                 nop
80D2  FE FE              cp -2
80D4  FC FC AA           call m,$AAFC
80D7  AA                 xor d
80D8  00                 nop
80D9  00                 nop
80DA  EF                 rst 28h
80DB  EF                 rst 28h
80DC  CF                 rst 08h
80DD  CF                 rst 08h
80DE  AA                 xor d
80DF  AA                 xor d
80E0  AA                 xor d
80E1  AA                 xor d
80E2  05                 dec b
80E3  01 02 02           ld bc,514
80E6  55                 ld d,l
80E7  55                 ld d,l
80E8  AA                 xor d
80E9  AA                 xor d
80EA  10 10              djnz $80FC
80EC  20 20              jr nz,$810E
80EE  55                 ld d,l
80EF  55                 ld d,l
80F0  AA                 xor d
80F1  AA                 xor d
80F2  01 01 02           ld bc,513
80F5  02                 ld (bc),a
80F6  55                 ld d,l
80F7  55                 ld d,l
80F8  AA                 xor d
80F9  AA                 xor d
80FA  10 10              djnz $810C
80FC  20 20              jr nz,$811E
80FE  55                 ld d,l
80FF  55                 ld d,l
8100  7B                 ld a,e
8101  CC BF FE           call z,$FEBF
8104  FF                 rst 38h
8105  FE BF              cp -65
8107  FE BC              cp -68
8109  BE                 cp (hl)
810A  5F                 ld e,a
810B  FA A4 10           jp m,$10A4
810E  18 00              jr $8110
8110  81                 add a,c
8111  C0                 ret nz
8112  E3                 ex (sp),hl
8113  B2                 or d
8114  B7                 or a
8115  20 EE              jr nz,$8105
8117  4A                 ld c,d
8118  9C                 sbc a,h
8119  90                 sub b
811A  C0                 ret nz
811B  02                 ld (bc),a
811C  55                 ld d,l
811D  54                 ld d,h
811E  00                 nop
811F  00                 nop
8120  01 80 07           ld bc,1920
8123  E0                 ret po
8124  0E F8              ld c,-8
8126  3D                 dec a
8127  5C                 ld e,h
8128  FA AF D5           jp m,$D5AF
812B  57                 ld d,a
812C  A8                 xor b
812D  8A                 adc a,d
812E  52                 ld d,d
812F  25                 dec h
8130  08                 ex af,af'
8131  00                 nop
8132  20 82              jr nz,$80B6
8134  00                 nop
8135  00                 nop
8136  00                 nop
8137  00                 nop
8138  00                 nop
8139  00                 nop
813A  00                 nop
813B  00                 nop
813C  00                 nop
813D  00                 nop
813E  00                 nop
813F  00                 nop
8140  00                 nop
8141  00                 nop
8142  02                 ld (bc),a
8143  00                 nop
8144  10 00              djnz $8146
8146  42                 ld b,d
8147  A4                 and h
8148  52                 ld d,d
8149  80                 add a,b
814A  52                 ld d,d
814B  A4                 and h
814C  53                 ld d,e
814D  24                 inc h
814E  D3 24              out ($24),a
8150  D5                 push de
8151  2C                 inc l
8152  D5                 push de
8153  AC                 xor h
8154  EB                 ex de,hl
8155  AD                 xor l
8156  EB                 ex de,hl
8157  AD                 xor l
8158  EF                 rst 28h
8159  ED FF              ???
815B  FF                 rst 38h
815C  DA AA 48           jp c,$48AA
815F  88                 adc a,b
8160  3F                 ccf
8161  FC 0F FC           call m,$FC0F
8164  10 F8              djnz $815E
8166  1F                 rra
8167  08                 ex af,af'
8168  07                 rlca
8169  F0                 ret p
816A  08                 ex af,af'
816B  F0                 ret p
816C  0F                 rrca
816D  10 03              djnz $8172
816F  E0                 ret po
8170  04                 inc b
8171  E0                 ret po
8172  07                 rlca
8173  20 01              jr nz,$8176
8175  C0                 ret nz
8176  02                 ld (bc),a
8177  40                 ld b,b
8178  03                 inc bc
8179  80                 add a,b
817A  00                 nop
817B  80                 add a,b
817C  01 00 01           ld bc,256
817F  80                 add a,b
8180  C0                 ret nz
8181  00                 nop
8182  F0                 ret p
8183  00                 nop
8184  FC 00 FF           call m,$FF00
8187  00                 nop
8188  7F                 ld a,a
8189  C0                 ret nz
818A  4F                 ld c,a
818B  F0                 ret p
818C  87                 add a,a
818D  FC 87 FF           call m,$FF87
8190  32 FF 72           ld ($72FF),a
8193  FF                 rst 38h
8194  19                 add hl,de
8195  73                 ld (hl),e
8196  61                 ld h,c
8197  2B                 dec hl
8198  18 AA              jr $8144
819A  0A                 ld a,(bc)
819B  50                 ld d,b
819C  63                 ld h,e
819D  A2                 and d
819E  D1                 pop de
819F  9C                 sbc a,h
81A0  00                 nop
81A1  00                 nop
81A2  00                 nop
81A3  00                 nop
81A4  00                 nop
81A5  00                 nop
81A6  00                 nop
81A7  00                 nop
81A8  00                 nop
81A9  00                 nop
81AA  00                 nop
81AB  00                 nop
81AC  00                 nop
81AD  00                 nop
81AE  00                 nop
81AF  00                 nop
81B0  C0                 ret nz
81B1  00                 nop
81B2  F0                 ret p
81B3  00                 nop
81B4  FC 00 7F           call m,$7F00
81B7  00                 nop
81B8  3F                 ccf
81B9  C0                 ret nz
81BA  27                 daa
81BB  F0                 ret p
81BC  42                 ld b,d
81BD  FC D2 7F           call m,$7FD2
81C0  D0                 ret nc
81C1  BF                 cp a
81C2  61                 ld h,c
81C3  3F                 ccf
81C4  1D                 dec e
81C5  B7                 or a
81C6  3C                 inc a
81C7  16 B9              ld d,-71
81C9  E6 97              and $97
81CB  F1                 pop af
81CC  CF                 rst 08h
81CD  49                 ld c,c
81CE  CD A5 37           call $37A5
81D1  54                 ld d,h
81D2  77                 ld (hl),a
81D3  AA                 xor d
81D4  1B                 dec de
81D5  F2 61 5A           jp p,$5A61
81D8  18 FA              jr $81D4
81DA  0A                 ld a,(bc)
81DB  52                 ld d,d
81DC  63                 ld h,e
81DD  A2                 and d
81DE  D1                 pop de
81DF  9C                 sbc a,h
81E0  D0                 ret nc
81E1  E0                 ret po
81E2  61                 ld h,c
81E3  38 1D              jr c,$8202
81E5  DB 3C              in a,($3C)
81E7  06 B9              ld b,-71
81E9  E2 97 F3           jp po,$F397
81EC  CF                 rst 08h
81ED  49                 ld c,c
81EE  CD A5 37           call $37A5
81F1  54                 ld d,h
81F2  77                 ld (hl),a
81F3  AA                 xor d
81F4  1B                 dec de
81F5  F2 61 5A           jp p,$5A61
81F8  18 FA              jr $81F4
81FA  0A                 ld a,(bc)
81FB  52                 ld d,d
81FC  63                 ld h,e
81FD  A2                 and d
81FE  D1                 pop de
81FF  9C                 sbc a,h
8200  00                 nop
8201  00                 nop
8202  00                 nop
8203  20 04              jr nz,$8209
8205  80                 add a,b
8206  00                 nop
8207  20 42              jr nz,$824B
8209  A8                 xor b
820A  02                 ld (bc),a
820B  A8                 xor b
820C  6A                 ld l,d
820D  61                 ld h,c
820E  32 60 1B           ld ($1B60),a
8211  51                 ld d,c
8212  9B                 sbc a,e
8213  22 09 56           ld ($5609),hl
8216  6D                 ld l,l
8217  84                 add a,h
8218  35                 dec (hl)
8219  CC 36 F9           call z,$F936
821C  19                 add hl,de
821D  52                 ld d,d
821E  DD 53              ld d,e
8220  00                 nop
8221  00                 nop
8222  00                 nop
8223  00                 nop
8224  04                 inc b
8225  48                 ld c,b
8226  00                 nop
8227  40                 ld b,b
8228  06 40              ld b,64
822A  49                 ld c,c
822B  08                 ex af,af'
822C  01 52 18           ld bc,$1852
822F  90                 sub b
8230  36 A0              ld (hl),-96
8232  22 A4 0A           ld ($0AA4),hl
8235  A4                 and h
8236  E2 A8 35           jp po,$35A8
8239  22 5B 6A           ld ($6A5B),hl
823C  5A                 ld e,d
823D  88                 adc a,b
823E  1C                 inc e
823F  2C                 inc l
8240  00                 nop
8241  80                 add a,b
8242  00                 nop
8243  00                 nop
8244  02                 ld (bc),a
8245  C0                 ret nz
8246  00                 nop
8247  E0                 ret po
8248  05                 dec b
8249  60                 ld h,b
824A  10 F0              djnz $823C
824C  05                 dec b
824D  78                 ld a,b
824E  00                 nop
824F  FC 15 78           call m,$7815
8252  42                 ld b,d
8253  E8                 ret pe
8254  05                 dec b
8255  6E                 ld l,(hl)
8256  22 F7 05           ld ($05F7),hl
8259  FA 03 7B           jp m,$7B03
825C  05                 dec b
825D  EF                 rst 28h
825E  8B                 adc a,e
825F  FD 00              nop
8261  00                 nop
8262  00                 nop
8263  00                 nop
8264  00                 nop
8265  00                 nop
8266  00                 nop
8267  00                 nop
8268  00                 nop
8269  00                 nop
826A  00                 nop
826B  00                 nop
826C  00                 nop
826D  1A                 ld a,(de)
826E  00                 nop
826F  4D                 ld c,l
8270  00                 nop
8271  15                 dec d
8272  00                 nop
8273  45                 ld b,l
8274  00                 nop
8275  01 00 2A           ld bc,$2A00
8278  00                 nop
8279  00                 nop
827A  00                 nop
827B  06 00              ld b,0
827D  0B                 dec bc
827E  00                 nop
827F  27                 daa
8280  00                 nop
8281  00                 nop
8282  00                 nop
8283  00                 nop
8284  00                 nop
8285  00                 nop
8286  00                 nop
8287  00                 nop
8288  00                 nop
8289  00                 nop
828A  00                 nop
828B  00                 nop
828C  00                 nop
828D  00                 nop
828E  00                 nop
828F  00                 nop
8290  00                 nop
8291  03                 inc bc
8292  00                 nop
8293  0C                 inc c
8294  00                 nop
8295  30 00              jr nc,$8297
8297  C0                 ret nz
8298  03                 inc bc
8299  02                 ld (bc),a
829A  0C                 inc c
829B  08                 ex af,af'
829C  30 20              jr nc,$82BE
829E  C0                 ret nz
829F  80                 add a,b
82A0  00                 nop
82A1  03                 inc bc
82A2  00                 nop
82A3  0C                 inc c
82A4  00                 nop
82A5  30 00              jr nc,$82A7
82A7  C0                 ret nz
82A8  03                 inc bc
82A9  02                 ld (bc),a
82AA  0C                 inc c
82AB  08                 ex af,af'
82AC  30 20              jr nc,$82CE
82AE  C0                 ret nz
82AF  80                 add a,b
82B0  02                 ld (bc),a
82B1  00                 nop
82B2  08                 ex af,af'
82B3  00                 nop
82B4  20 00              jr nz,$82B6
82B6  80                 add a,b
82B7  00                 nop
82B8  00                 nop
82B9  00                 nop
82BA  00                 nop
82BB  00                 nop
82BC  00                 nop
82BD  00                 nop
82BE  00                 nop
82BF  00                 nop
82C0  00                 nop
82C1  00                 nop
82C2  50                 ld d,b
82C3  26 01              ld h,1
82C5  00                 nop
82C6  00                 nop
82C7  12                 ld (de),a
82C8  65                 ld h,l
82C9  40                 ld b,b
82CA  20 05              jr nz,$82D1
82CC  88                 adc a,b
82CD  25                 dec h
82CE  3B                 dec sp
82CF  08                 ex af,af'
82D0  A3                 and e
82D1  2B                 dec hl
82D2  3E FA              ld a,-6
82D4  43                 ld b,e
82D5  AA                 xor d
82D6  BA                 cp d
82D7  95                 sub l
82D8  53                 ld d,e
82D9  6F                 ld l,a
82DA  ED FF              ???
82DC  FF                 rst 38h
82DD  FF                 rst 38h
82DE  FF                 rst 38h
82DF  0F                 rrca
82E0  04                 inc b
82E1  4B                 ld c,e
82E2  11 40 10           ld de,$1040
82E5  40                 ld b,b
82E6  20 10              jr nz,$82F8
82E8  08                 ex af,af'
82E9  11 00 00           ld de,0
82EC  0D                 dec c
82ED  01 04 00           ld bc,4
82F0  20 80              jr nz,$8272
82F2  68                 ld l,b
82F3  6A                 ld l,d
82F4  12                 ld (de),a
82F5  08                 ex af,af'
82F6  00                 nop
82F7  00                 nop
82F8  00                 nop
82F9  11 60 00           ld de,96
82FC  02                 ld (bc),a
82FD  91                 sub c
82FE  0C                 inc c
82FF  51                 ld d,c
8300  00                 nop
8301  02                 ld (bc),a
8302  00                 nop
8303  00                 nop
8304  00                 nop
8305  08                 ex af,af'
8306  00                 nop
8307  03                 inc bc
8308  00                 nop
8309  22 00 8D           ld ($8D00),hl
830C  00                 nop
830D  07                 rlca
830E  00                 nop
830F  45                 ld b,l
8310  00                 nop
8311  0B                 dec bc
8312  00                 nop
8313  95                 sub l
8314  02                 ld (bc),a
8315  0B                 dec bc
8316  00                 nop
8317  8D                 adc a,l
8318  01 3A 00           ld bc,58
831B  27                 daa
831C  04                 inc b
831D  0D                 dec c
831E  00                 nop
831F  B7                 or a
8320  92                 sub d
8321  FF                 rst 38h
8322  A0                 and b
8323  2F                 cpl
8324  82                 add a,d
8325  2E 86              ld l,-122
8327  0B                 dec bc
8328  87                 add a,a
8329  75                 ld (hl),l
832A  BF                 cp a
832B  F5                 push af
832C  FF                 rst 38h
832D  FF                 rst 38h
832E  BF                 cp a
832F  FF                 rst 38h
8330  FF                 rst 38h
8331  FF                 rst 38h
8332  EF                 rst 28h
8333  F5                 push af
8334  FF                 rst 38h
8335  FE BF              cp -65
8337  FE FF              cp -1
8339  FB                 ei
833A  EF                 rst 28h
833B  FB                 ei
833C  EF                 rst 28h
833D  FD CF              rst 08h
833F  7F                 ld a,a
8340  80                 add a,b
8341  00                 nop
8342  C0                 ret nz
8343  00                 nop
8344  C0                 ret nz
8345  00                 nop
8346  50                 ld d,b
8347  00                 nop
8348  FC 00 FE           call m,$FE00
834B  00                 nop
834C  BF                 cp a
834D  00                 nop
834E  BC                 cp h
834F  00                 nop
8350  DC 00 D7           call c,$D700
8353  00                 nop
8354  D7                 rst 10h
8355  00                 nop
8356  DB 80              in a,($80)
8358  7F                 ld a,a
8359  80                 add a,b
835A  7F                 ld a,a
835B  C0                 ret nz
835C  FF                 rst 38h
835D  E0                 ret po
835E  FF                 rst 38h
835F  B0                 or b
8360  00                 nop
8361  0E 00              ld c,0
8363  27                 daa
8364  00                 nop
8365  0B                 dec bc
8366  00                 nop
8367  A7                 and a
8368  03                 inc bc
8369  0B                 dec bc
836A  0C                 inc c
836B  25                 dec h
836C  30 0B              jr nc,$8379
836E  C0                 ret nz
836F  A7                 and a
8370  02                 ld (bc),a
8371  0B                 dec bc
8372  08                 ex af,af'
8373  27                 daa
8374  20 0B              jr nz,$8381
8376  80                 add a,b
8377  27                 daa
8378  00                 nop
8379  0B                 dec bc
837A  00                 nop
837B  27                 daa
837C  00                 nop
837D  0B                 dec bc
837E  00                 nop
837F  27                 daa
8380  00                 nop
8381  00                 nop
8382  54                 ld d,h
8383  54                 ld d,h
8384  02                 ld (bc),a
8385  02                 ld (bc),a
8386  42                 ld b,d
8387  42                 ld b,d
8388  06 06              ld b,6
838A  5E                 ld e,(hl)
838B  5E                 ld e,(hl)
838C  3E 3E              ld a,62
838E  00                 nop
838F  00                 nop
8390  00                 nop
8391  00                 nop
8392  54                 ld d,h
8393  54                 ld d,h
8394  02                 ld (bc),a
8395  02                 ld (bc),a
8396  42                 ld b,d
8397  42                 ld b,d
8398  06 06              ld b,6
839A  5E                 ld e,(hl)
839B  5E                 ld e,(hl)
839C  3E 3E              ld a,62
839E  00                 nop
839F  00                 nop
83A0  6F                 ld l,a
83A1  00                 nop
83A2  7C                 ld a,h
83A3  B0                 or b
83A4  7B                 ld a,e
83A5  85                 add a,l
83A6  7C                 ld a,h
83A7  31 7B C0           ld sp,$C07B
83AA  35                 dec (hl)
83AB  04                 inc b
83AC  3B                 dec sp
83AD  A1                 and c
83AE  39                 add hl,sp
83AF  08                 ex af,af'
83B0  37                 scf
83B1  00                 nop
83B2  7A                 ld a,d
83B3  CA 6F C0           jp z,$C06F
83B6  71                 ld (hl),c
83B7  08                 ex af,af'
83B8  75                 ld (hl),l
83B9  60                 ld h,b
83BA  6D                 ld l,l
83BB  41                 ld b,c
83BC  75                 ld (hl),l
83BD  D8                 ret c
83BE  72                 ld (hl),d
83BF  08                 ex af,af'
83C0  00                 nop
83C1  7B                 ld a,e
83C2  06 9F              ld b,-97
83C4  50                 ld d,b
83C5  EF                 rst 28h
83C6  46                 ld b,(hl)
83C7  1F                 rra
83C8  01 EF 10           ld bc,$10EF
83CB  56                 ld d,(hl)
83CC  42                 ld b,d
83CD  EE 08              xor $08
83CF  4E                 ld c,(hl)
83D0  00                 nop
83D1  76                 halt
83D2  29                 add hl,hl
83D3  AF                 xor a
83D4  01 FB 08           ld bc,$08FB
83D7  47                 ld b,a
83D8  03                 inc bc
83D9  57                 ld d,a
83DA  41                 ld b,c
83DB  5B                 ld e,e
83DC  0D                 dec c
83DD  D7                 rst 10h
83DE  08                 ex af,af'
83DF  27                 daa
83E0  00                 nop
83E1  B7                 or a
83E2  00                 nop
83E3  02                 ld (bc),a
83E4  02                 ld (bc),a
83E5  00                 nop
83E6  08                 ex af,af'
83E7  28 20              jr z,$8409
83E9  51                 ld d,c
83EA  08                 ex af,af'
83EB  3E 10              ld a,16
83ED  5F                 ld e,a
83EE  40                 ld b,b
83EF  BB                 cp e
83F0  2A 17 00           ld hl,($0017)
83F3  BB                 cp e
83F4  28 7F              jr z,$8475
83F6  00                 nop
83F7  AF                 xor a
83F8  41                 ld b,c
83F9  57                 ld d,a
83FA  00                 nop
83FB  FA 03 0F           jp m,$0F03
83FE  54                 ld d,h
83FF  03                 inc bc
8400  0A                 ld a,(bc)
8401  6F                 ld l,a
8402  00                 nop
8403  F8                 ret m
8404  04                 inc b
8405  50                 ld d,b
8406  00                 nop
8407  D0                 ret nc
8408  02                 ld (bc),a
8409  50                 ld d,b
840A  00                 nop
840B  A0                 and b
840C  02                 ld (bc),a
840D  20 00              jr nz,$840F
840F  A0                 and b
8410  02                 ld (bc),a
8411  60                 ld h,b
8412  00                 nop
8413  E0                 ret po
8414  05                 dec b
8415  50                 ld d,b
8416  00                 nop
8417  D0                 ret nc
8418  05                 dec b
8419  D0                 ret nc
841A  12                 ld (de),a
841B  A8                 xor b
841C  04                 inc b
841D  B8                 cp b
841E  58                 ld e,b
841F  9F                 sbc a,a
8420  AF                 xor a
8421  A0                 and b
8422  00                 nop
8423  10 11              djnz $8436
8425  08                 ex af,af'
8426  1C                 inc e
8427  4C                 ld c,h
8428  F2 E8 F9           jp p,$F9E8
842B  FC BD FC           call m,$FCBD
842E  BC                 cp h
842F  DC DD DC           call c,$DCDD
8432  D7                 rst 10h
8433  C0                 ret nz
8434  D7                 rst 10h
8435  6C                 ld l,h
8436  DB F4              in a,($F4)
8438  7F                 ld a,a
8439  FA 7F FF           jp m,$FF7F
843C  CE 7D              adc a,125
843E  F2 E7 00           jp p,$00E7
8441  0B                 dec bc
8442  00                 nop
8443  27                 daa
8444  00                 nop
8445  0B                 dec bc
8446  00                 nop
8447  27                 daa
8448  00                 nop
8449  0B                 dec bc
844A  00                 nop
844B  27                 daa
844C  00                 nop
844D  0B                 dec bc
844E  00                 nop
844F  27                 daa
8450  00                 nop
8451  0B                 dec bc
8452  00                 nop
8453  27                 daa
8454  00                 nop
8455  0B                 dec bc
8456  00                 nop
8457  27                 daa
8458  00                 nop
8459  0B                 dec bc
845A  00                 nop
845B  27                 daa
845C  00                 nop
845D  0B                 dec bc
845E  00                 nop
845F  27                 daa
8460  02                 ld (bc),a
8461  00                 nop
8462  08                 ex af,af'
8463  00                 nop
8464  20 00              jr nz,$8466
8466  80                 add a,b
8467  00                 nop
8468  00                 nop
8469  00                 nop
846A  00                 nop
846B  00                 nop
846C  00                 nop
846D  00                 nop
846E  00                 nop
846F  00                 nop
8470  00                 nop
8471  00                 nop
8472  00                 nop
8473  00                 nop
8474  00                 nop
8475  00                 nop
8476  00                 nop
8477  00                 nop
8478  00                 nop
8479  00                 nop
847A  00                 nop
847B  00                 nop
847C  00                 nop
847D  00                 nop
847E  00                 nop
847F  00                 nop
8480  00                 nop
8481  00                 nop
8482  06 00              ld b,0
8484  06 80              ld b,-128
8486  0E 80              ld c,-128
8488  06 B8              ld b,-72
848A  0E A6              ld c,-90
848C  06 A9              ld b,-87
848E  0E 6E              ld c,110
8490  06 AF              ld b,-81
8492  0E AF              ld c,-81
8494  06 B0              ld b,-80
8496  0E 8F              ld c,-113
8498  06 80              ld b,-128
849A  06 00              ld b,0
849C  06 00              ld b,0
849E  00                 nop
849F  00                 nop
84A0  00                 nop
84A1  00                 nop
84A2  00                 nop
84A3  00                 nop
84A4  00                 nop
84A5  00                 nop
84A6  00                 nop
84A7  00                 nop
84A8  00                 nop
84A9  00                 nop
84AA  00                 nop
84AB  00                 nop
84AC  F8                 ret m
84AD  05                 dec b
84AE  07                 rlca
84AF  F0                 ret p
84B0  F8                 ret m
84B1  1D                 dec e
84B2  FE AE              cp -82
84B4  7C                 ld a,h
84B5  1F                 rra
84B6  86                 add a,(hl)
84B7  4B                 ld c,e
84B8  78                 ld a,b
84B9  DF                 rst 18h
84BA  06 50              ld b,80
84BC  00                 nop
84BD  C0                 ret nz
84BE  00                 nop
84BF  40                 ld b,b
84C0  00                 nop
84C1  00                 nop
84C2  00                 nop
84C3  10 00              djnz $84C5
84C5  B0                 or b
84C6  00                 nop
84C7  98                 sbc a,b
84C8  08                 ex af,af'
84C9  B8                 cp b
84CA  20 98              jr nz,$8464
84CC  00                 nop
84CD  BA                 cp d
84CE  00                 nop
84CF  1A                 ld a,(de)
84D0  40                 ld b,b
84D1  BA                 cp d
84D2  A0                 and b
84D3  9A                 sbc a,d
84D4  D0                 ret nc
84D5  B8                 cp b
84D6  F8                 ret m
84D7  98                 sbc a,b
84D8  00                 nop
84D9  B8                 cp b
84DA  00                 nop
84DB  10 00              djnz $84DD
84DD  30 00              jr nc,$84DF
84DF  00                 nop
84E0  00                 nop
84E1  00                 nop
84E2  00                 nop
84E3  00                 nop
84E4  00                 nop
84E5  00                 nop
84E6  00                 nop
84E7  00                 nop
84E8  00                 nop
84E9  00                 nop
84EA  00                 nop
84EB  01 00 00           ld bc,0
84EE  00                 nop
84EF  02                 ld (bc),a
84F0  00                 nop
84F1  00                 nop
84F2  00                 nop
84F3  04                 inc b
84F4  00                 nop
84F5  00                 nop
84F6  00                 nop
84F7  00                 nop
84F8  00                 nop
84F9  09                 add hl,bc
84FA  00                 nop
84FB  00                 nop
84FC  00                 nop
84FD  01 00 22           ld bc,$2200
8500  00                 nop
8501  D8                 ret c
8502  28 5E              jr z,$8562
8504  05                 dec b
8505  E3                 ex (sp),hl
8506  83                 add a,e
8507  61                 ld h,c
8508  17                 rla
8509  D0                 ret nc
850A  2F                 cpl
850B  E8                 ret pe
850C  5F                 ld e,a
850D  F4 3F F8           call p,$F83F
8510  7F                 ld a,a
8511  F4 BF FA           call p,$FABF
8514  7F                 ld a,a
8515  FD FF              rst 38h
8517  FE 7F              cp 127
8519  FF                 rst 38h
851A  FF                 rst 38h
851B  FE FF              cp -1
851D  FF                 rst 38h
851E  FF                 rst 38h
851F  FF                 rst 38h
8520  00                 nop
8521  E8                 ret pe
8522  00                 nop
8523  30 00              jr nc,$8525
8525  E8                 ret pe
8526  80                 add a,b
8527  30 C0              jr nc,$84E9
8529  E8                 ret pe
852A  C0                 ret nz
852B  30 E0              jr nc,$850D
852D  E8                 ret pe
852E  60                 ld h,b
852F  30 70              jr nc,$85A1
8531  E8                 ret pe
8532  30 30              jr nc,$8564
8534  30 E8              jr nc,$851E
8536  B8                 cp b
8537  30 18              jr nc,$8551
8539  E8                 ret pe
853A  9C                 sbc a,h
853B  30 4E              jr nc,$858B
853D  68                 ld l,b
853E  A6                 and (hl)
853F  30 00              jr nc,$8541
8541  05                 dec b
8542  00                 nop
8543  83                 add a,e
8544  00                 nop
8545  06 02              ld b,2
8547  03                 inc bc
8548  00                 nop
8549  10 08              djnz $8553
854B  2F                 cpl
854C  00                 nop
854D  1B                 dec de
854E  10 0E              djnz $855E
8550  00                 nop
8551  31 08 2A           ld sp,$2A08
8554  02                 ld (bc),a
8555  15                 dec d
8556  00                 nop
8557  A9                 xor c
8558  00                 nop
8559  02                 ld (bc),a
855A  00                 nop
855B  00                 nop
855C  00                 nop
855D  00                 nop
855E  00                 nop
855F  00                 nop
8560  FF                 rst 38h
8561  FF                 rst 38h
8562  FF                 rst 38h
8563  FF                 rst 38h
8564  FF                 rst 38h
8565  FF                 rst 38h
8566  FF                 rst 38h
8567  FF                 rst 38h
8568  0F                 rrca
8569  FE E0              cp -32
856B  01 FF FF           ld bc,-1
856E  FF                 rst 38h
856F  FF                 rst 38h
8570  FF                 rst 38h
8571  FF                 rst 38h
8572  07                 rlca
8573  F8                 ret m
8574  F8                 ret m
8575  07                 rlca
8576  57                 ld d,a
8577  FF                 rst 38h
8578  AA                 xor d
8579  BF                 cp a
857A  00                 nop
857B  00                 nop
857C  17                 rla
857D  F8                 ret m
857E  09                 add hl,bc
857F  F0                 ret p
8580  D6 08              sub 8
8582  AE                 xor (hl)
8583  80                 add a,b
8584  FD C0              ret nz
8586  F3                 di
8587  C0                 ret nz
8588  0F                 rrca
8589  D0                 ret nc
858A  FF                 rst 38h
858B  B8                 cp b
858C  FF                 rst 38h
858D  78                 ld a,b
858E  FC FC E3           call m,$E3FC
8591  F8                 ret m
8592  1F                 rra
8593  F0                 ret p
8594  FF                 rst 38h
8595  E0                 ret po
8596  FF                 rst 38h
8597  80                 add a,b
8598  E0                 ret po
8599  08                 ex af,af'
859A  00                 nop
859B  10 00              djnz $859D
859D  E8                 ret pe
859E  00                 nop
859F  30 00              jr nc,$85A1
85A1  E8                 ret pe
85A2  00                 nop
85A3  30 00              jr nc,$85A5
85A5  E8                 ret pe
85A6  00                 nop
85A7  30 00              jr nc,$85A9
85A9  E8                 ret pe
85AA  00                 nop
85AB  30 00              jr nc,$85AD
85AD  E8                 ret pe
85AE  00                 nop
85AF  30 00              jr nc,$85B1
85B1  E8                 ret pe
85B2  00                 nop
85B3  30 00              jr nc,$85B5
85B5  E8                 ret pe
85B6  00                 nop
85B7  30 00              jr nc,$85B9
85B9  E8                 ret pe
85BA  00                 nop
85BB  30 00              jr nc,$85BD
85BD  E8                 ret pe
85BE  00                 nop
85BF  30 00              jr nc,$85C1
85C1  00                 nop
85C2  78                 ld a,b
85C3  FE BC              cp -68
85C5  FC 3C AA           call m,$AA3C
85C8  BC                 cp h
85C9  00                 nop
85CA  3D                 dec a
85CB  EF                 rst 28h
85CC  8D                 adc a,l
85CD  CF                 rst 08h
85CE  50                 ld d,b
85CF  AA                 xor d
85D0  00                 nop
85D1  00                 nop
85D2  7F                 ld a,a
85D3  CE BF              adc a,-65
85D5  EC 3F EA           call pe,$EA3F
85D8  BF                 cp a
85D9  E0                 ret po
85DA  3F                 ccf
85DB  EF                 rst 28h
85DC  80                 add a,b
85DD  2F                 cpl
85DE  55                 ld d,l
85DF  4A                 ld c,d
85E0  00                 nop
85E1  00                 nop
85E2  FE 9E              cp -98
85E4  FC 3D AA           call m,$AA3D
85E7  BC                 cp h
85E8  00                 nop
85E9  3D                 dec a
85EA  EF                 rst 28h
85EB  BC                 cp h
85EC  CF                 rst 08h
85ED  B1                 or c
85EE  AA                 xor d
85EF  CA 00 00           jp z,$0000
85F2  FB                 ei
85F3  FE F7              cp -9
85F5  FD A7              and a
85F7  FC 17 FD           call m,$FD17
85FA  E7                 rst 20h
85FB  FC C4 01           call m,$01C4
85FE  AA                 xor d
85FF  AA                 xor d
8600  00                 nop
8601  00                 nop
8602  FE FE              cp -2
8604  FC FC AA           call m,$AAFC
8607  AA                 xor d
8608  00                 nop
8609  00                 nop
860A  EF                 rst 28h
860B  EC CF C8           call pe,$C8CF
860E  AA                 xor d
860F  A1                 and c
8610  00                 nop
8611  02                 ld (bc),a
8612  FE CB              cp -53
8614  FC 83 AA           call m,$AA83
8617  97                 sub a
8618  00                 nop
8619  07                 rlca
861A  EF                 rst 28h
861B  2F                 cpl
861C  CF                 rst 08h
861D  01 AA 95           ld bc,$95AA
8620  00                 nop
8621  00                 nop
8622  FC 83 C0           call m,$C083
8625  70                 ld (hl),b
8626  02                 ld (bc),a
8627  95                 sub l
8628  28 30              jr z,$865A
862A  8D                 adc a,l
862B  7A                 ld a,d
862C  3C                 inc a
862D  78                 ld a,b
862E  3C                 inc a
862F  81                 add a,c
8630  9E                 sbc a,(hl)
8631  00                 nop
8632  58                 ld e,b
8633  00                 nop
8634  A0                 and b
8635  00                 nop
8636  C0                 ret nz
8637  00                 nop
8638  C0                 ret nz
8639  00                 nop
863A  80                 add a,b
863B  00                 nop
863C  80                 add a,b
863D  00                 nop
863E  00                 nop
863F  00                 nop
8640  00                 nop
8641  00                 nop
8642  E0                 ret po
8643  FE F3              cp -13
8645  1C                 inc e
8646  F5                 push af
8647  E0                 ret po
8648  E1                 pop hl
8649  EE EB              xor $EB
864B  D7                 rst 10h
864C  E6 57              and $57
864E  41                 ld b,c
864F  C7                 rst 00h
8650  80                 add a,b
8651  33                 inc sp
8652  00                 nop
8653  0F                 rrca
8654  00                 nop
8655  02                 ld (bc),a
8656  00                 nop
8657  00                 nop
8658  00                 nop
8659  01 00 00           ld bc,0
865C  00                 nop
865D  00                 nop
865E  00                 nop
865F  00                 nop
8660  00                 nop
8661  00                 nop
8662  FE FE              cp -2
8664  FC FC AA           call m,$AAFC
8667  AA                 xor d
8668  00                 nop
8669  00                 nop
866A  AF                 xor a
866B  EF                 rst 28h
866C  CF                 rst 08h
866D  CF                 rst 08h
866E  CA AA B0           jp z,$B0AA
8671  00                 nop
8672  5A                 ld e,d
8673  FE 98              cp -104
8675  FC 3C AA           call m,$AA3C
8678  7C                 ld a,h
8679  00                 nop
867A  3E EF              ld a,-17
867C  9E                 sbc a,(hl)
867D  CF                 rst 08h
867E  7C                 ld a,h
867F  AA                 xor d
8680  00                 nop
8681  00                 nop
8682  00                 nop
8683  3C                 inc a
8684  00                 nop
8685  00                 nop
8686  00                 nop
8687  18 1C              jr $86A5
8689  3E 00              ld a,0
868B  FF                 rst 38h
868C  0F                 rrca
868D  FF                 rst 38h
868E  7F                 ld a,a
868F  FF                 rst 38h
8690  7F                 ld a,a
8691  FF                 rst 38h
8692  7C                 ld a,h
8693  00                 nop
8694  00                 nop
8695  00                 nop
8696  00                 nop
8697  00                 nop
8698  00                 nop
8699  3F                 ccf
869A  07                 rlca
869B  FF                 rst 38h
869C  00                 nop
869D  00                 nop
869E  00                 nop
869F  00                 nop
86A0  00                 nop
86A1  0C                 inc c
86A2  20 03              jr nz,$86A7
86A4  03                 inc bc
86A5  1F                 rra
86A6  6C                 ld l,h
86A7  FF                 rst 38h
86A8  FF                 rst 38h
86A9  FF                 rst 38h
86AA  7F                 ld a,a
86AB  FF                 rst 38h
86AC  FF                 rst 38h
86AD  FF                 rst 38h
86AE  FF                 rst 38h
86AF  7F                 ld a,a
86B0  E0                 ret po
86B1  FF                 rst 38h
86B2  00                 nop
86B3  0F                 rrca
86B4  1C                 inc e
86B5  01 00 00           ld bc,0
86B8  80                 add a,b
86B9  3E FF              ld a,-1
86BB  FE 70              cp 112
86BD  10 00              djnz $86BF
86BF  00                 nop
86C0  38 00              jr c,$86C2
86C2  FE 38              cp 56
86C4  FE 3C              cp 60
86C6  FF                 rst 38h
86C7  81                 add a,c
86C8  FF                 rst 38h
86C9  BF                 cp a
86CA  FF                 rst 38h
86CB  CF                 rst 08h
86CC  F1                 pop af
86CD  F7                 rst 30h
86CE  FE 7F              cp 127
86D0  FF                 rst 38h
86D1  BF                 cp a
86D2  F8                 ret m
86D3  1F                 rra
86D4  C0                 ret nz
86D5  06 03              ld b,3
86D7  80                 add a,b
86D8  00                 nop
86D9  01 08 07           ld bc,1800
86DC  00                 nop
86DD  07                 rlca
86DE  00                 nop
86DF  01 00 00           ld bc,0
86E2  00                 nop
86E3  00                 nop
86E4  43                 ld b,e
86E5  F8                 ret m
86E6  FC 00 FF           call m,$FF00
86E9  FF                 rst 38h
86EA  FF                 rst 38h
86EB  FE FF              cp -1
86ED  FF                 rst 38h
86EE  38 07              jr c,$86F7
86F0  C7                 rst 00h
86F1  C0                 ret nz
86F2  0F                 rrca
86F3  FC 0F FF           call m,$FF0F
86F6  00                 nop
86F7  3F                 ccf
86F8  38 00              jr c,$86FA
86FA  7F                 ld a,a
86FB  80                 add a,b
86FC  7F                 ld a,a
86FD  E0                 ret po
86FE  00                 nop
86FF  00                 nop
8700  00                 nop
8701  00                 nop
8702  00                 nop
8703  00                 nop
8704  06 08              ld b,8
8706  C0                 ret nz
8707  00                 nop
8708  7E                 ld a,(hl)
8709  F0                 ret p
870A  3F                 ccf
870B  EE F8              xor $F8
870D  00                 nop
870E  F0                 ret p
870F  00                 nop
8710  0F                 rrca
8711  90                 sub b
8712  00                 nop
8713  00                 nop
8714  F8                 ret m
8715  00                 nop
8716  80                 add a,b
8717  00                 nop
8718  00                 nop
8719  00                 nop
871A  00                 nop
871B  00                 nop
871C  00                 nop
871D  00                 nop
871E  00                 nop
871F  00                 nop
8720  07                 rlca
8721  FF                 rst 38h
8722  08                 ex af,af'
8723  00                 nop
8724  00                 nop
8725  55                 ld d,l
8726  02                 ld (bc),a
8727  2A 00 15           ld hl,($1500)
872A  00                 nop
872B  82                 add a,d
872C  00                 nop
872D  21 00 08           ld hl,2048
8730  00                 nop
8731  02                 ld (bc),a
8732  00                 nop
8733  00                 nop
8734  00                 nop
8735  35                 dec (hl)
8736  00                 nop
8737  17                 rla
8738  00                 nop
8739  78                 ld a,b
873A  00                 nop
873B  3B                 dec sp
873C  00                 nop
873D  F2 00 77           jp p,$7700
8740  FF                 rst 38h
8741  FF                 rst 38h
8742  00                 nop
8743  00                 nop
8744  FF                 rst 38h
8745  FF                 rst 38h
8746  FF                 rst 38h
8747  FF                 rst 38h
8748  FF                 rst 38h
8749  FF                 rst 38h
874A  FF                 rst 38h
874B  FF                 rst 38h
874C  FF                 rst 38h
874D  FF                 rst 38h
874E  00                 nop
874F  00                 nop
8750  FF                 rst 38h
8751  FF                 rst 38h
8752  00                 nop
8753  00                 nop
8754  55                 ld d,l
8755  55                 ld d,l
8756  77                 ld (hl),a
8757  77                 ld (hl),a
8758  00                 nop
8759  00                 nop
875A  FF                 rst 38h
875B  FF                 rst 38h
875C  AA                 xor d
875D  AA                 xor d
875E  FF                 rst 38h
875F  FF                 rst 38h
8760  FF                 rst 38h
8761  FE 00              cp 0
8763  02                 ld (bc),a
8764  FB                 ei
8765  42                 ld b,d
8766  FE 84              cp -124
8768  FB                 ei
8769  08                 ex af,af'
876A  FE 30              cp 48
876C  F0                 ret p
876D  C0                 ret nz
876E  07                 rlca
876F  00                 nop
8770  F8                 ret m
8771  00                 nop
8772  00                 nop
8773  00                 nop
8774  1C                 inc e
8775  00                 nop
8776  3C                 inc a
8777  00                 nop
8778  0E 00              ld c,0
877A  DE 00              sbc a,0
877C  A7                 and a
877D  00                 nop
877E  EF                 rst 28h
877F  00                 nop
8780  00                 nop
8781  00                 nop
8782  00                 nop
8783  F7                 rst 30h
8784  01 CF 01           ld bc,463
8787  E0                 ret po
8788  03                 inc bc
8789  8B                 adc a,e
878A  03                 inc bc
878B  DB 07              in a,($07)
878D  13                 inc de
878E  07                 rlca
878F  AB                 xor e
8790  0E 13              ld c,19
8792  0F                 rrca
8793  43                 ld b,e
8794  1C                 inc e
8795  63                 ld h,e
8796  1E B3              ld e,-77
8798  39                 add hl,sp
8799  5B                 ld e,e
879A  3C                 inc a
879B  00                 nop
879C  73                 ld (hl),e
879D  FF                 rst 38h
879E  7B                 ld a,e
879F  FF                 rst 38h
87A0  00                 nop
87A1  00                 nop
87A2  FF                 rst 38h
87A3  80                 add a,b
87A4  FF                 rst 38h
87A5  80                 add a,b
87A6  00                 nop
87A7  00                 nop
87A8  16 00              ld d,0
87AA  37                 scf
87AB  00                 nop
87AC  6C                 ld l,h
87AD  00                 nop
87AE  4E                 ld c,(hl)
87AF  00                 nop
87B0  18 00              jr $87B2
87B2  5C                 ld e,h
87B3  00                 nop
87B4  30 00              jr nc,$87B6
87B6  38 00              jr c,$87B8
87B8  60                 ld h,b
87B9  00                 nop
87BA  70                 ld (hl),b
87BB  00                 nop
87BC  40                 ld b,b
87BD  00                 nop
87BE  60                 ld h,b
87BF  00                 nop
87C0  00                 nop
87C1  00                 nop
87C2  00                 nop
87C3  00                 nop
87C4  00                 nop
87C5  01 00 01           ld bc,256
87C8  00                 nop
87C9  03                 inc bc
87CA  00                 nop
87CB  03                 inc bc
87CC  00                 nop
87CD  07                 rlca
87CE  00                 nop
87CF  07                 rlca
87D0  00                 nop
87D1  0E 00              ld c,0
87D3  0F                 rrca
87D4  00                 nop
87D5  1C                 inc e
87D6  00                 nop
87D7  1E 00              ld e,0
87D9  39                 add hl,sp
87DA  00                 nop
87DB  3C                 inc a
87DC  00                 nop
87DD  73                 ld (hl),e
87DE  00                 nop
87DF  7B                 ld a,e
87E0  00                 nop
87E1  00                 nop
87E2  F7                 rst 30h
87E3  FF                 rst 38h
87E4  CF                 rst 08h
87E5  FF                 rst 38h
87E6  E0                 ret po
87E7  00                 nop
87E8  8B                 adc a,e
87E9  16 DB              ld d,-37
87EB  37                 scf
87EC  13                 inc de
87ED  6C                 ld l,h
87EE  AB                 xor e
87EF  4E                 ld c,(hl)
87F0  13                 inc de
87F1  18 43              jr $8836
87F3  5C                 ld e,h
87F4  63                 ld h,e
87F5  30 B3              jr nc,$87AA
87F7  38 5B              jr c,$8854
87F9  60                 ld h,b
87FA  00                 nop
87FB  70                 ld (hl),b
87FC  FF                 rst 38h
87FD  40                 ld b,b
87FE  FF                 rst 38h
87FF  60                 ld h,b
8800  80                 add a,b
8801  07                 rlca
8802  FF                 rst 38h
8803  F7                 rst 30h
8804  FF                 rst 38h
8805  F9                 ld sp,hl
8806  00                 nop
8807  03                 inc bc
8808  34                 inc (hl)
8809  68                 ld l,b
880A  76                 halt
880B  6D                 ld l,l
880C  1B                 dec de
880D  64                 ld h,h
880E  39                 add hl,sp
880F  6A                 ld l,d
8810  0C                 inc c
8811  64                 ld h,h
8812  1D                 dec e
8813  61                 ld h,c
8814  06 63              ld b,99
8816  0E 66              ld c,102
8818  03                 inc bc
8819  6D                 ld l,l
881A  07                 rlca
881B  00                 nop
881C  01 7F 03           ld bc,895
881F  7F                 ld a,a
8820  80                 add a,b
8821  00                 nop
8822  80                 add a,b
8823  00                 nop
8824  C0                 ret nz
8825  00                 nop
8826  C0                 ret nz
8827  00                 nop
8828  E0                 ret po
8829  00                 nop
882A  E0                 ret po
882B  00                 nop
882C  70                 ld (hl),b
882D  00                 nop
882E  F0                 ret p
882F  00                 nop
8830  38 00              jr c,$8832
8832  78                 ld a,b
8833  00                 nop
8834  1C                 inc e
8835  00                 nop
8836  BC                 cp h
8837  00                 nop
8838  4E                 ld c,(hl)
8839  00                 nop
883A  1E 00              ld e,0
883C  E7                 rst 20h
883D  00                 nop
883E  EF                 rst 28h
883F  00                 nop
8840  00                 nop
8841  80                 add a,b
8842  00                 nop
8843  FF                 rst 38h
8844  00                 nop
8845  FF                 rst 38h
8846  00                 nop
8847  00                 nop
8848  00                 nop
8849  34                 inc (hl)
884A  00                 nop
884B  76                 halt
884C  00                 nop
884D  1B                 dec de
884E  00                 nop
884F  39                 add hl,sp
8850  00                 nop
8851  0C                 inc c
8852  00                 nop
8853  1D                 dec e
8854  00                 nop
8855  06 00              ld b,0
8857  0E 00              ld c,0
8859  03                 inc bc
885A  00                 nop
885B  07                 rlca
885C  00                 nop
885D  01 00 03           ld bc,768
8860  07                 rlca
8861  80                 add a,b
8862  F7                 rst 30h
8863  80                 add a,b
8864  F9                 ld sp,hl
8865  C0                 ret nz
8866  03                 inc bc
8867  C0                 ret nz
8868  68                 ld l,b
8869  E0                 ret po
886A  6D                 ld l,l
886B  E0                 ret po
886C  64                 ld h,h
886D  70                 ld (hl),b
886E  6A                 ld l,d
886F  F0                 ret p
8870  64                 ld h,h
8871  38 61              jr c,$88D4
8873  78                 ld a,b
8874  63                 ld h,e
8875  1C                 inc e
8876  66                 ld h,(hl)
8877  BC                 cp h
8878  6D                 ld l,l
8879  4E                 ld c,(hl)
887A  00                 nop
887B  1E 7F              ld e,127
887D  E7                 rst 20h
887E  7F                 ld a,a
887F  EF                 rst 28h
8880  00                 nop
8881  00                 nop
8882  7F                 ld a,a
8883  FF                 rst 38h
8884  7F                 ld a,a
8885  FF                 rst 38h
8886  00                 nop
8887  00                 nop
8888  69                 ld l,c
8889  4B                 ld c,e
888A  6C                 ld l,h
888B  9B                 sbc a,e
888C  66                 ld h,(hl)
888D  33                 inc sp
888E  6B                 ld l,e
888F  63                 ld h,e
8890  64                 ld h,h
8891  CB 61              bit 4,c
8893  A3                 and e
8894  63                 ld h,e
8895  33                 inc sp
8896  66                 ld h,(hl)
8897  9B                 sbc a,e
8898  6C                 ld l,h
8899  4B                 ld c,e
889A  00                 nop
889B  00                 nop
889C  7F                 ld a,a
889D  FF                 rst 38h
889E  7F                 ld a,a
889F  FF                 rst 38h
88A0  80                 add a,b
88A1  13                 inc de
88A2  27                 daa
88A3  26 1C              ld h,28
88A5  4D                 ld c,l
88A6  70                 ld (hl),b
88A7  99                 sbc a,c
88A8  D9                 exx
88A9  35                 dec (hl)
88AA  02                 ld (bc),a
88AB  65                 ld h,l
88AC  E4 D5 C9           call po,$C9D5
88AF  95                 sub l
88B0  13                 inc de
88B1  50                 ld d,b
88B2  26 40              ld h,64
88B4  4D                 ld c,l
88B5  00                 nop
88B6  98                 sbc a,b
88B7  00                 nop
88B8  30 00              jr nc,$88BA
88BA  60                 ld h,b
88BB  00                 nop
88BC  C0                 ret nz
88BD  00                 nop
88BE  80                 add a,b
88BF  00                 nop
88C0  00                 nop
88C1  00                 nop
88C2  55                 ld d,l
88C3  55                 ld d,l
88C4  FF                 rst 38h
88C5  FF                 rst 38h
88C6  55                 ld d,l
88C7  54                 ld d,h
88C8  55                 ld d,l
88C9  40                 ld b,b
88CA  54                 ld d,h
88CB  00                 nop
88CC  40                 ld b,b
88CD  00                 nop
88CE  00                 nop
88CF  00                 nop
88D0  00                 nop
88D1  00                 nop
88D2  00                 nop
88D3  00                 nop
88D4  00                 nop
88D5  00                 nop
88D6  00                 nop
88D7  00                 nop
88D8  00                 nop
88D9  00                 nop
88DA  00                 nop
88DB  00                 nop
88DC  00                 nop
88DD  00                 nop
88DE  00                 nop
88DF  00                 nop
88E0  00                 nop
88E1  00                 nop
88E2  55                 ld d,l
88E3  55                 ld d,l
88E4  7C                 ld a,h
88E5  00                 nop
88E6  00                 nop
88E7  00                 nop
88E8  00                 nop
88E9  00                 nop
88EA  00                 nop
88EB  00                 nop
88EC  00                 nop
88ED  00                 nop
88EE  00                 nop
88EF  00                 nop
88F0  00                 nop
88F1  00                 nop
88F2  00                 nop
88F3  00                 nop
88F4  00                 nop
88F5  00                 nop
88F6  00                 nop
88F7  00                 nop
88F8  00                 nop
88F9  00                 nop
88FA  00                 nop
88FB  00                 nop
88FC  00                 nop
88FD  00                 nop
88FE  00                 nop
88FF  00                 nop
8900  00                 nop
8901  00                 nop
8902  55                 ld d,l
8903  55                 ld d,l
8904  00                 nop
8905  1F                 rra
8906  00                 nop
8907  00                 nop
8908  00                 nop
8909  00                 nop
890A  00                 nop
890B  00                 nop
890C  00                 nop
890D  00                 nop
890E  00                 nop
890F  00                 nop
8910  00                 nop
8911  00                 nop
8912  00                 nop
8913  00                 nop
8914  00                 nop
8915  00                 nop
8916  00                 nop
8917  00                 nop
8918  00                 nop
8919  00                 nop
891A  00                 nop
891B  00                 nop
891C  00                 nop
891D  00                 nop
891E  00                 nop
891F  00                 nop
8920  00                 nop
8921  00                 nop
8922  55                 ld d,l
8923  55                 ld d,l
8924  7F                 ld a,a
8925  FF                 rst 38h
8926  15                 dec d
8927  55                 ld d,l
8928  01 55 00           ld bc,85
892B  15                 dec d
892C  00                 nop
892D  01 00 00           ld bc,0
8930  00                 nop
8931  00                 nop
8932  00                 nop
8933  00                 nop
8934  00                 nop
8935  00                 nop
8936  00                 nop
8937  00                 nop
8938  00                 nop
8939  00                 nop
893A  00                 nop
893B  00                 nop
893C  00                 nop
893D  00                 nop
893E  00                 nop
893F  00                 nop
8940  64                 ld h,h
8941  00                 nop
8942  32 72 59           ld ($5972),a
8945  1C                 inc e
8946  4C                 ld c,h
8947  87                 add a,a
8948  56                 ld d,(hl)
8949  4D                 ld c,l
894A  53                 ld d,e
894B  20 55              jr nz,$89A2
894D  93                 sub e
894E  54                 ld d,h
894F  C9                 ret
8950  05                 dec b
8951  64                 ld h,h
8952  01 32 00           ld bc,50
8955  59                 ld e,c
8956  00                 nop
8957  0C                 inc c
8958  00                 nop
8959  06 00              ld b,0
895B  03                 inc bc
895C  00                 nop
895D  01 00 00           ld bc,0
8960  00                 nop
8961  01 00 02           ld bc,512
8964  00                 nop
8965  05                 dec b
8966  00                 nop
8967  0B                 dec bc
8968  00                 nop
8969  17                 rla
896A  00                 nop
896B  2E 00              ld l,0
896D  5C                 ld e,h
896E  00                 nop
896F  B9                 cp c
8970  01 70 02           ld bc,624
8973  E4 05 C3           call po,$C305
8976  0B                 dec bc
8977  8E                 adc a,(hl)
8978  17                 rla
8979  39                 add hl,sp
897A  2E 00              ld l,0
897C  5C                 ld e,h
897D  FC F9 F9           call m,$F9F9
8980  70                 ld (hl),b
8981  03                 inc bc
8982  E4 E7 C3           call po,$C3E7
8985  8E                 adc a,(hl)
8986  8E                 adc a,(hl)
8987  1C                 inc e
8988  39                 add hl,sp
8989  39                 add hl,sp
898A  00                 nop
898B  70                 ld (hl),b
898C  FC E7 F9           call m,$F9E7
898F  CF                 rst 08h
8990  03                 inc bc
8991  80                 add a,b
8992  E7                 rst 20h
8993  27                 daa
8994  8E                 adc a,(hl)
8995  1C                 inc e
8996  1C                 inc e
8997  70                 ld (hl),b
8998  39                 add hl,sp
8999  C9                 ret
899A  70                 ld (hl),b
899B  02                 ld (bc),a
899C  E7                 rst 20h
899D  E4 CF C9           call po,$C9CF
89A0  80                 add a,b
89A1  13                 inc de
89A2  27                 daa
89A3  26 1C              ld h,28
89A5  4C                 ld c,h
89A6  70                 ld (hl),b
89A7  98                 sbc a,b
89A8  C9                 ret
89A9  30 02              jr nc,$89AD
89AB  60                 ld h,b
89AC  E4 C0 C9           call po,$C9C0
89AF  80                 add a,b
89B0  13                 inc de
89B1  00                 nop
89B2  26 00              ld h,0
89B4  4C                 ld c,h
89B5  00                 nop
89B6  98                 sbc a,b
89B7  00                 nop
89B8  30 00              jr nc,$89BA
89BA  60                 ld h,b
89BB  00                 nop
89BC  C0                 ret nz
89BD  00                 nop
89BE  80                 add a,b
89BF  00                 nop
89C0  64                 ld h,h
89C1  00                 nop
89C2  32 72 19           ld ($1972),a
89C5  1C                 inc e
89C6  0C                 inc c
89C7  87                 add a,a
89C8  06 49              ld b,73
89CA  03                 inc bc
89CB  20 01              jr nz,$89CE
89CD  93                 sub e
89CE  00                 nop
89CF  C9                 ret
89D0  00                 nop
89D1  64                 ld h,h
89D2  00                 nop
89D3  32 00 19           ld ($1900),a
89D6  00                 nop
89D7  0C                 inc c
89D8  00                 nop
89D9  06 00              ld b,0
89DB  03                 inc bc
89DC  00                 nop
89DD  01 00 00           ld bc,0
89E0  E0                 ret po
89E1  07                 rlca
89E2  73                 ld (hl),e
89E3  93                 sub e
89E4  38 E1              jr c,$89C7
89E6  1C                 inc e
89E7  38 CE              jr c,$89B7
89E9  4E                 ld c,(hl)
89EA  07                 rlca
89EB  00                 nop
89EC  F3                 di
89ED  9F                 sbc a,a
89EE  F9                 ld sp,hl
89EF  CF                 rst 08h
89F0  00                 nop
89F1  E0                 ret po
89F2  72                 ld (hl),d
89F3  73                 ld (hl),e
89F4  1C                 inc e
89F5  38 87              jr c,$897E
89F7  1C                 inc e
89F8  49                 ld c,c
89F9  CE 20              adc a,32
89FB  07                 rlca
89FC  93                 sub e
89FD  F3                 di
89FE  C9                 ret
89FF  F9                 ld sp,hl
8A00  40                 ld b,b
8A01  00                 nop
8A02  A0                 and b
8A03  00                 nop
8A04  D0                 ret nc
8A05  00                 nop
8A06  E8                 ret pe
8A07  00                 nop
8A08  74                 ld (hl),h
8A09  00                 nop
8A0A  3A 00 9D           ld a,($9D00)
8A0D  00                 nop
8A0E  CE 80              adc a,-128
8A10  07                 rlca
8A11  40                 ld b,b
8A12  93                 sub e
8A13  A0                 and b
8A14  E1                 pop hl
8A15  D0                 ret nc
8A16  38 E8              jr c,$8A00
8A18  4E                 ld c,(hl)
8A19  74                 ld (hl),h
8A1A  00                 nop
8A1B  3A 9F 9D           ld a,($9D9F)
8A1E  CF                 rst 08h
8A1F  CF                 rst 08h
8A20  00                 nop
8A21  00                 nop
8A22  00                 nop
8A23  FF                 rst 38h
8A24  00                 nop
8A25  FF                 rst 38h
8A26  00                 nop
8A27  00                 nop
8A28  00                 nop
8A29  69                 ld l,c
8A2A  00                 nop
8A2B  6C                 ld l,h
8A2C  00                 nop
8A2D  66                 ld h,(hl)
8A2E  00                 nop
8A2F  6B                 ld l,e
8A30  00                 nop
8A31  64                 ld h,h
8A32  00                 nop
8A33  61                 ld h,c
8A34  00                 nop
8A35  63                 ld h,e
8A36  00                 nop
8A37  66                 ld h,(hl)
8A38  00                 nop
8A39  6C                 ld l,h
8A3A  00                 nop
8A3B  00                 nop
8A3C  00                 nop
8A3D  FF                 rst 38h
8A3E  00                 nop
8A3F  FF                 rst 38h
8A40  00                 nop
8A41  00                 nop
8A42  7F                 ld a,a
8A43  80                 add a,b
8A44  7F                 ld a,a
8A45  80                 add a,b
8A46  00                 nop
8A47  00                 nop
8A48  4B                 ld c,e
8A49  00                 nop
8A4A  1B                 dec de
8A4B  00                 nop
8A4C  33                 inc sp
8A4D  00                 nop
8A4E  63                 ld h,e
8A4F  00                 nop
8A50  4B                 ld c,e
8A51  00                 nop
8A52  23                 inc hl
8A53  00                 nop
8A54  33                 inc sp
8A55  00                 nop
8A56  1B                 dec de
8A57  00                 nop
8A58  4B                 ld c,e
8A59  00                 nop
8A5A  00                 nop
8A5B  00                 nop
8A5C  7F                 ld a,a
8A5D  80                 add a,b
8A5E  7F                 ld a,a
8A5F  80                 add a,b
8A60  7F                 ld a,a
8A61  FC 80 02           call m,$0280
8A64  BF                 cp a
8A65  FA A0 0A           jp m,$0AA0
8A68  A0                 and b
8A69  0A                 ld a,(bc)
8A6A  BF                 cp a
8A6B  FA 80 02           jp m,$0280
8A6E  7F                 ld a,a
8A6F  FC 00 00           call m,$0000
8A72  00                 nop
8A73  00                 nop
8A74  00                 nop
8A75  00                 nop
8A76  00                 nop
8A77  00                 nop
8A78  00                 nop
8A79  00                 nop
8A7A  00                 nop
8A7B  00                 nop
8A7C  00                 nop
8A7D  00                 nop
8A7E  00                 nop
8A7F  00                 nop
8A80  7F                 ld a,a
8A81  FC 80 02           call m,$0280
8A84  BF                 cp a
8A85  FA A0 0A           jp m,$0AA0
8A88  A0                 and b
8A89  0A                 ld a,(bc)
8A8A  BF                 cp a
8A8B  FA 80 02           jp m,$0280
8A8E  7F                 ld a,a
8A8F  FC 00 00           call m,$0000
8A92  FF                 rst 38h
8A93  FF                 rst 38h
8A94  00                 nop
8A95  00                 nop
8A96  80                 add a,b
8A97  C0                 ret nz
8A98  E5                 push hl
8A99  72                 ld (hl),d
8A9A  C2 E1 81           jp nz,$81E1
8A9D  40                 ld b,b
8A9E  00                 nop
8A9F  80                 add a,b
8AA0  7F                 ld a,a
8AA1  FC 80 02           call m,$0280
8AA4  BF                 cp a
8AA5  FA A0 0A           jp m,$0AA0
8AA8  A0                 and b
8AA9  0A                 ld a,(bc)
8AAA  BF                 cp a
8AAB  FA 80 02           jp m,$0280
8AAE  7F                 ld a,a
8AAF  FC 00 00           call m,$0000
8AB2  7F                 ld a,a
8AB3  FC AA F6           call m,$F6AA
8AB6  1F                 rra
8AB7  DA AF FA           jp c,$FAAF
8ABA  15                 dec d
8ABB  BA                 cp d
8ABC  80                 add a,b
8ABD  66                 ld h,(hl)
8ABE  55                 ld d,l
8ABF  54                 ld d,h
8AC0  00                 nop
8AC1  00                 nop
8AC2  7F                 ld a,a
8AC3  FC AA F6           call m,$F6AA
8AC6  1F                 rra
8AC7  DA AF FA           jp c,$FAAF
8ACA  15                 dec d
8ACB  BA                 cp d
8ACC  80                 add a,b
8ACD  66                 ld h,(hl)
8ACE  55                 ld d,l
8ACF  54                 ld d,h
8AD0  00                 nop
8AD1  00                 nop
8AD2  00                 nop
8AD3  00                 nop
8AD4  00                 nop
8AD5  00                 nop
8AD6  00                 nop
8AD7  00                 nop
8AD8  00                 nop
8AD9  00                 nop
8ADA  00                 nop
8ADB  00                 nop
8ADC  00                 nop
8ADD  00                 nop
8ADE  00                 nop
8ADF  00                 nop
8AE0  00                 nop
8AE1  00                 nop
8AE2  00                 nop
8AE3  00                 nop
8AE4  00                 nop
8AE5  00                 nop
8AE6  00                 nop
8AE7  00                 nop
8AE8  00                 nop
8AE9  00                 nop
8AEA  00                 nop
8AEB  00                 nop
8AEC  00                 nop
8AED  00                 nop
8AEE  00                 nop
8AEF  00                 nop
8AF0  00                 nop
8AF1  00                 nop
8AF2  00                 nop
8AF3  80                 add a,b
8AF4  81                 add a,c
8AF5  40                 ld b,b
8AF6  C2 E1 E5           jp nz,$E5E1
8AF9  72                 ld (hl),d
8AFA  80                 add a,b
8AFB  C0                 ret nz
8AFC  00                 nop
8AFD  00                 nop
8AFE  FF                 rst 38h
8AFF  FF                 rst 38h
8B00  00                 nop
8B01  00                 nop
8B02  00                 nop
8B03  01 00 02           ld bc,512
8B06  00                 nop
8B07  05                 dec b
8B08  00                 nop
8B09  0B                 dec bc
8B0A  00                 nop
8B0B  1E 00              ld e,0
8B0D  19                 add hl,de
8B0E  00                 nop
8B0F  7C                 ld a,h
8B10  00                 nop
8B11  EB                 ex de,hl
8B12  00                 nop
8B13  7C                 ld a,h
8B14  03                 inc bc
8B15  96                 sub (hl)
8B16  07                 rlca
8B17  B6                 or (hl)
8B18  0A                 ld a,(bc)
8B19  FF                 rst 38h
8B1A  14                 inc d
8B1B  5A                 ld e,d
8B1C  28 BF              jr z,$8ADD
8B1E  50                 ld d,b
8B1F  0D                 dec c
8B20  80                 add a,b
8B21  00                 nop
8B22  00                 nop
8B23  00                 nop
8B24  A0                 and b
8B25  00                 nop
8B26  80                 add a,b
8B27  00                 nop
8B28  C8                 ret z
8B29  00                 nop
8B2A  A0                 and b
8B2B  00                 nop
8B2C  D2 00 F0           jp nc,$F000
8B2F  00                 nop
8B30  A8                 xor b
8B31  80                 add a,b
8B32  F2 00 F8           jp p,$F800
8B35  20 34              jr nz,$8B6B
8B37  80                 add a,b
8B38  FF                 rst 38h
8B39  08                 ex af,af'
8B3A  FA 80 BF           jp m,$BF80
8B3D  42                 ld b,d
8B3E  EA D0 B0           jp pe,$B0D0
8B41  EC 63 FA           call pe,$FA63
8B44  F7                 rst 30h
8B45  2D                 dec l
8B46  FF                 rst 38h
8B47  3A FF FF           ld a,($FFFF)
8B4A  FB                 ei
8B4B  ED E7              ???
8B4D  7E                 ld a,(hl)
8B4E  D3 DC              out ($DC),a
8B50  FF                 rst 38h
8B51  FD FF              rst 38h
8B53  FF                 rst 38h
8B54  FE 7B              cp 123
8B56  F7                 rst 30h
8B57  1F                 rra
8B58  FF                 rst 38h
8B59  FF                 rst 38h
8B5A  FF                 rst 38h
8B5B  FC DF FF           call m,$FFDF
8B5E  FF                 rst 38h
8B5F  FF                 rst 38h
8B60  00                 nop
8B61  00                 nop
8B62  A3                 and e
8B63  8A                 adc a,d
8B64  55                 ld d,l
8B65  6D                 ld l,l
8B66  FF                 rst 38h
8B67  BA                 cp d
8B68  F5                 push af
8B69  FF                 rst 38h
8B6A  FF                 rst 38h
8B6B  D7                 rst 10h
8B6C  FF                 rst 38h
8B6D  FF                 rst 38h
8B6E  BF                 cp a
8B6F  7F                 ld a,a
8B70  FF                 rst 38h
8B71  FB                 ei
8B72  FF                 rst 38h
8B73  FF                 rst 38h
8B74  EF                 rst 28h
8B75  FF                 rst 38h
8B76  FF                 rst 38h
8B77  DF                 rst 18h
8B78  FF                 rst 38h
8B79  FF                 rst 38h
8B7A  FF                 rst 38h
8B7B  FB                 ei
8B7C  FB                 ei
8B7D  FF                 rst 38h
8B7E  FF                 rst 38h
8B7F  FF                 rst 38h
8B80  00                 nop
8B81  00                 nop
8B82  87                 add a,a
8B83  0A                 ld a,(bc)
8B84  4F                 ld c,a
8B85  6D                 ld l,l
8B86  80                 add a,b
8B87  3A 8F 7F           ld a,($7F8F)
8B8A  16 D7              ld d,-41
8B8C  00                 nop
8B8D  FF                 rst 38h
8B8E  1E 7F              ld e,127
8B90  2E FB              ld l,-5
8B92  01 FF BD           ld bc,$BDFF
8B95  FF                 rst 38h
8B96  7D                 ld a,l
8B97  DF                 rst 18h
8B98  03                 inc bc
8B99  FF                 rst 38h
8B9A  7B                 ld a,e
8B9B  FB                 ei
8B9C  EB                 ex de,hl
8B9D  FF                 rst 38h
8B9E  07                 rlca
8B9F  FF                 rst 38h
8BA0  00                 nop
8BA1  BA                 cp d
8BA2  00                 nop
8BA3  68                 ld l,b
8BA4  03                 inc bc
8BA5  EC 01 80           call pe,$8001
8BA8  07                 rlca
8BA9  C0                 ret nz
8BAA  02                 ld (bc),a
8BAB  00                 nop
8BAC  07                 rlca
8BAD  80                 add a,b
8BAE  0C                 inc c
8BAF  00                 nop
8BB0  06 00              ld b,0
8BB2  08                 ex af,af'
8BB3  00                 nop
8BB4  0C                 inc c
8BB5  00                 nop
8BB6  00                 nop
8BB7  00                 nop
8BB8  08                 ex af,af'
8BB9  00                 nop
8BBA  00                 nop
8BBB  00                 nop
8BBC  00                 nop
8BBD  00                 nop
8BBE  00                 nop
8BBF  00                 nop
8BC0  03                 inc bc
8BC1  C0                 ret nz
8BC2  00                 nop
8BC3  00                 nop
8BC4  03                 inc bc
8BC5  C0                 ret nz
8BC6  01 80 00           ld bc,128
8BC9  00                 nop
8BCA  01 80 02           ld bc,640
8BCD  80                 add a,b
8BCE  00                 nop
8BCF  00                 nop
8BD0  03                 inc bc
8BD1  80                 add a,b
8BD2  05                 dec b
8BD3  80                 add a,b
8BD4  00                 nop
8BD5  00                 nop
8BD6  07                 rlca
8BD7  00                 nop
8BD8  03                 inc bc
8BD9  00                 nop
8BDA  00                 nop
8BDB  00                 nop
8BDC  0B                 dec bc
8BDD  00                 nop
8BDE  07                 rlca
8BDF  00                 nop
8BE0  26 10              ld h,16
8BE2  2F                 cpl
8BE3  C0                 ret nz
8BE4  29                 add hl,hl
8BE5  80                 add a,b
8BE6  13                 inc de
8BE7  C0                 ret nz
8BE8  16 60              ld d,96
8BEA  05                 dec b
8BEB  C0                 ret nz
8BEC  01 60 00           ld bc,96
8BEF  40                 ld b,b
8BF0  01 E0 00           ld bc,224
8BF3  40                 ld b,b
8BF4  00                 nop
8BF5  E0                 ret po
8BF6  01 C0 01           ld bc,448
8BF9  80                 add a,b
8BFA  00                 nop
8BFB  00                 nop
8BFC  00                 nop
8BFD  00                 nop
8BFE  00                 nop
8BFF  00                 nop
8C00  00                 nop
8C01  00                 nop
8C02  00                 nop
8C03  00                 nop
8C04  00                 nop
8C05  00                 nop
8C06  00                 nop
8C07  00                 nop
8C08  00                 nop
8C09  01 00 0F           ld bc,$0F00
8C0C  00                 nop
8C0D  18 00              jr $8C0F
8C0F  75                 ld (hl),l
8C10  00                 nop
8C11  E6 01              and $01
8C13  A8                 xor b
8C14  01 28 01           ld bc,296
8C17  82                 add a,d
8C18  03                 inc bc
8C19  0A                 ld a,(bc)
8C1A  02                 ld (bc),a
8C1B  0B                 dec bc
8C1C  06 2E              ld b,46
8C1E  00                 nop
8C1F  BA                 cp d
8C20  00                 nop
8C21  00                 nop
8C22  00                 nop
8C23  00                 nop
8C24  00                 nop
8C25  00                 nop
8C26  01 E0 FC           ld bc,-800
8C29  30 3E              jr nc,$8C69
8C2B  07                 rlca
8C2C  FF                 rst 38h
8C2D  9C                 sbc a,h
8C2E  1B                 dec de
8C2F  96                 sub (hl)
8C30  03                 inc bc
8C31  92                 sub d
8C32  00                 nop
8C33  00                 nop
8C34  95                 sub l
8C35  89                 adc a,c
8C36  B7                 or a
8C37  B3                 or e
8C38  BF                 cp a
8C39  BF                 cp a
8C3A  FA 7F AA           jp m,$AA7F
8C3D  2A A8 2A           ld hl,($2AA8)
8C40  00                 nop
8C41  00                 nop
8C42  00                 nop
8C43  00                 nop
8C44  00                 nop
8C45  00                 nop
8C46  00                 nop
8C47  00                 nop
8C48  00                 nop
8C49  00                 nop
8C4A  C0                 ret nz
8C4B  00                 nop
8C4C  F8                 ret m
8C4D  00                 nop
8C4E  4C                 ld c,h
8C4F  00                 nop
8C50  2A 00 0B           ld hl,($0B00)
8C53  00                 nop
8C54  80                 add a,b
8C55  80                 add a,b
8C56  39                 add hl,sp
8C57  C0                 ret nz
8C58  60                 ld h,b
8C59  60                 ld h,b
8C5A  FE 20              cp 32
8C5C  F8                 ret m
8C5D  20 9F              jr nz,$8BFE
8C5F  90                 sub b
8C60  AA                 xor d
8C61  6A                 ld l,d
8C62  54                 ld d,h
8C63  75                 ld (hl),l
8C64  AA                 xor d
8C65  6A                 ld l,d
8C66  54                 ld d,h
8C67  75                 ld (hl),l
8C68  AA                 xor d
8C69  6A                 ld l,d
8C6A  54                 ld d,h
8C6B  75                 ld (hl),l
8C6C  AA                 xor d
8C6D  6A                 ld l,d
8C6E  54                 ld d,h
8C6F  75                 ld (hl),l
8C70  AA                 xor d
8C71  6A                 ld l,d
8C72  54                 ld d,h
8C73  75                 ld (hl),l
8C74  AA                 xor d
8C75  6A                 ld l,d
8C76  54                 ld d,h
8C77  75                 ld (hl),l
8C78  AA                 xor d
8C79  6A                 ld l,d
8C7A  54                 ld d,h
8C7B  75                 ld (hl),l
8C7C  AA                 xor d
8C7D  6A                 ld l,d
8C7E  54                 ld d,h
8C7F  75                 ld (hl),l
8C80  AA                 xor d
8C81  AA                 xor d
8C82  55                 ld d,l
8C83  55                 ld d,l
8C84  B0                 or b
8C85  00                 nop
8C86  41                 ld b,c
8C87  45                 ld b,l
8C88  0B                 dec bc
8C89  20 44              jr nz,$8CCF
8C8B  40                 ld b,b
8C8C  A3                 and e
8C8D  C8                 ret z
8C8E  40                 ld b,b
8C8F  15                 dec d
8C90  A2                 and d
8C91  AA                 xor d
8C92  44                 ld b,h
8C93  15                 dec d
8C94  A4                 and h
8C95  8A                 adc a,d
8C96  44                 ld b,h
8C97  4D                 ld c,l
8C98  A2                 and d
8C99  8A                 adc a,d
8C9A  51                 ld d,c
8C9B  15                 dec d
8C9C  A8                 xor b
8C9D  2A 55 55           ld hl,($5555)
8CA0  A8                 xor b
8CA1  2A 51 15           ld hl,($1551)
8CA4  AB                 xor e
8CA5  AA                 xor d
8CA6  51                 ld d,c
8CA7  15                 dec d
8CA8  0C                 inc c
8CA9  60                 ld h,b
8CAA  00                 nop
8CAB  01 08 20           ld bc,$2008
8CAE  54                 ld d,h
8CAF  55                 ld d,l
8CB0  A8                 xor b
8CB1  2A 54 55           ld hl,($5554)
8CB4  A8                 xor b
8CB5  2A 54 55           ld hl,($5554)
8CB8  A8                 xor b
8CB9  2A 50 15           ld hl,($1550)
8CBC  A8                 xor b
8CBD  2A 50 15           ld hl,($1550)
8CC0  A8                 xor b
8CC1  AE                 xor (hl)
8CC2  5D                 ld e,l
8CC3  5C                 ld e,h
8CC4  98                 sbc a,b
8CC5  8E                 adc a,(hl)
8CC6  55                 ld d,l
8CC7  14                 inc d
8CC8  88                 adc a,b
8CC9  82                 add a,d
8CCA  75                 ld (hl),l
8CCB  44                 ld b,h
8CCC  B8                 cp b
8CCD  EA 1D D4           jp pe,$D41D
8CD0  88                 adc a,b
8CD1  8A                 adc a,d
8CD2  05                 dec b
8CD3  44                 ld b,h
8CD4  A8                 xor b
8CD5  A2                 and d
8CD6  15                 dec d
8CD7  54                 ld d,h
8CD8  A8                 xor b
8CD9  BA                 cp d
8CDA  D5                 push de
8CDB  0C                 inc c
8CDC  C8                 ret z
8CDD  8E                 adc a,(hl)
8CDE  45                 ld b,l
8CDF  44                 ld b,h
8CE0  AA                 xor d
8CE1  AA                 xor d
8CE2  44                 ld b,h
8CE3  44                 ld b,h
8CE4  8A                 adc a,d
8CE5  E2 44 74           jp po,$7444
8CE8  A2                 and d
8CE9  AA                 xor d
8CEA  54                 ld d,h
8CEB  54                 ld d,h
8CEC  8A                 adc a,d
8CED  A2                 and d
8CEE  5C                 ld e,h
8CEF  44                 ld b,h
8CF0  8E                 adc a,(hl)
8CF1  A2                 and d
8CF2  7C                 ld a,h
8CF3  74                 ld (hl),h
8CF4  A2                 and d
8CF5  EA 74 F0           jp pe,$F074
8CF8  AA                 xor d
8CF9  EA 44 50           jp pe,$5044
8CFC  AA                 xor d
8CFD  A2                 and d
8CFE  54                 ld d,h
8CFF  54                 ld d,h
8D00  AA                 xor d
8D01  AA                 xor d
8D02  55                 ld d,l
8D03  D5                 push de
8D04  AB                 xor e
8D05  AA                 xor d
8D06  55                 ld d,l
8D07  D5                 push de
8D08  AB                 xor e
8D09  6A                 ld l,d
8D0A  56                 ld d,(hl)
8D0B  35                 dec (hl)
8D0C  AB                 xor e
8D0D  6A                 ld l,d
8D0E  56                 ld d,(hl)
8D0F  35                 dec (hl)
8D10  AB                 xor e
8D11  EA 54 15           jp pe,$1554
8D14  80                 add a,b
8D15  00                 nop
8D16  44                 ld b,h
8D17  11 83 60           ld de,$6083
8D1A  40                 ld b,b
8D1B  01 A3 62           ld bc,$62A3
8D1E  50                 ld d,b
8D1F  05                 dec b
8D20  AA                 xor d
8D21  AA                 xor d
8D22  55                 ld d,l
8D23  55                 ld d,l
8D24  AA                 xor d
8D25  AA                 xor d
8D26  55                 ld d,l
8D27  55                 ld d,l
8D28  AA                 xor d
8D29  AA                 xor d
8D2A  55                 ld d,l
8D2B  55                 ld d,l
8D2C  AA                 xor d
8D2D  AA                 xor d
8D2E  55                 ld d,l
8D2F  55                 ld d,l
8D30  AA                 xor d
8D31  AA                 xor d
8D32  55                 ld d,l
8D33  55                 ld d,l
8D34  AA                 xor d
8D35  AA                 xor d
8D36  55                 ld d,l
8D37  55                 ld d,l
8D38  AA                 xor d
8D39  AA                 xor d
8D3A  55                 ld d,l
8D3B  55                 ld d,l
8D3C  AA                 xor d
8D3D  AA                 xor d
8D3E  55                 ld d,l
8D3F  55                 ld d,l
8D40  7F                 ld a,a
8D41  FC 80 02           call m,$0280
8D44  BF                 cp a
8D45  FA A0 0A           jp m,$0AA0
8D48  A0                 and b
8D49  0A                 ld a,(bc)
8D4A  BF                 cp a
8D4B  FA 80 02           jp m,$0280
8D4E  7F                 ld a,a
8D4F  FC 00 00           call m,$0000
8D52  55                 ld d,l
8D53  55                 ld d,l
8D54  AA                 xor d
8D55  AA                 xor d
8D56  55                 ld d,l
8D57  55                 ld d,l
8D58  AA                 xor d
8D59  AA                 xor d
8D5A  55                 ld d,l
8D5B  55                 ld d,l
8D5C  AA                 xor d
8D5D  AA                 xor d
8D5E  55                 ld d,l
8D5F  55                 ld d,l
8D60  AA                 xor d
8D61  AA                 xor d
8D62  55                 ld d,l
8D63  55                 ld d,l
8D64  AA                 xor d
8D65  AA                 xor d
8D66  55                 ld d,l
8D67  55                 ld d,l
8D68  AA                 xor d
8D69  AA                 xor d
8D6A  55                 ld d,l
8D6B  55                 ld d,l
8D6C  AA                 xor d
8D6D  AA                 xor d
8D6E  00                 nop
8D6F  01 7F FC           ld bc,-897
8D72  80                 add a,b
8D73  02                 ld (bc),a
8D74  BF                 cp a
8D75  FA A0 0A           jp m,$0AA0
8D78  A0                 and b
8D79  0A                 ld a,(bc)
8D7A  BF                 cp a
8D7B  FA 80 02           jp m,$0280
8D7E  7F                 ld a,a
8D7F  FC 00 00           call m,$0000
8D82  00                 nop
8D83  00                 nop
8D84  00                 nop
8D85  00                 nop
8D86  00                 nop
8D87  00                 nop
8D88  00                 nop
8D89  00                 nop
8D8A  00                 nop
8D8B  00                 nop
8D8C  00                 nop
8D8D  00                 nop
8D8E  00                 nop
8D8F  00                 nop
8D90  00                 nop
8D91  00                 nop
8D92  00                 nop
8D93  00                 nop
8D94  00                 nop
8D95  00                 nop
8D96  00                 nop
8D97  00                 nop
8D98  00                 nop
8D99  00                 nop
8D9A  00                 nop
8D9B  00                 nop
8D9C  00                 nop
8D9D  00                 nop
8D9E  00                 nop
8D9F  00                 nop
8DA0  00                 nop
8DA1  00                 nop
8DA2  1C                 inc e
8DA3  CC 72 F6           call z,$F672
8DA6  46                 ld b,(hl)
8DA7  3A D4 4C           ld a,($4CD4)
8DAA  BD                 cp l
8DAB  D4 B1 D4           call nc,$D4B1
8DAE  44                 ld b,h
8DAF  FA 1E 3A           jp m,$3A1E
8DB2  FB                 ei
8DB3  8E                 adc a,(hl)
8DB4  B7                 or a
8DB5  22 EE 48           ld ($48EE),hl
8DB8  9C                 sbc a,h
8DB9  90                 sub b
8DBA  C0                 ret nz
8DBB  02                 ld (bc),a
8DBC  55                 ld d,l
8DBD  54                 ld d,h
8DBE  00                 nop
8DBF  00                 nop
8DC0  01 00 00           ld bc,0
8DC3  00                 nop
8DC4  01 00 00           ld bc,0
8DC7  00                 nop
8DC8  01 00 00           ld bc,0
8DCB  00                 nop
8DCC  01 00 00           ld bc,0
8DCF  00                 nop
8DD0  01 00 00           ld bc,0
8DD3  00                 nop
8DD4  01 00 00           ld bc,0
8DD7  00                 nop
8DD8  01 00 00           ld bc,0
8DDB  00                 nop
8DDC  01 00 00           ld bc,0
8DDF  00                 nop
8DE0  00                 nop
8DE1  00                 nop
8DE2  48                 ld c,b
8DE3  80                 add a,b
8DE4  02                 ld (bc),a
8DE5  02                 ld (bc),a
8DE6  2A A0 22           ld hl,($22A0)
8DE9  A9                 xor c
8DEA  A6                 and (hl)
8DEB  A9                 xor c
8DEC  A6                 and (hl)
8DED  A9                 xor c
8DEE  A7                 and a
8DEF  29                 add hl,hl
8DF0  AA                 xor d
8DF1  A9                 xor c
8DF2  AE                 xor (hl)
8DF3  A9                 xor c
8DF4  F6 AD              or $AD
8DF6  56                 ld d,(hl)
8DF7  B5                 or l
8DF8  D6 BB              sub -69
8DFA  57                 ld d,a
8DFB  7B                 ld a,e
8DFC  5B                 ld e,e
8DFD  7B                 ld a,e
8DFE  FB                 ei
8DFF  7F                 ld a,a
8E00  00                 nop
8E01  00                 nop
8E02  00                 nop
8E03  00                 nop
8E04  00                 nop
8E05  00                 nop
8E06  00                 nop
8E07  00                 nop
8E08  00                 nop
8E09  00                 nop
8E0A  00                 nop
8E0B  00                 nop
8E0C  00                 nop
8E0D  00                 nop
8E0E  80                 add a,b
8E0F  80                 add a,b
8E10  C9                 ret
8E11  C9                 ret
8E12  9C                 sbc a,h
8E13  9C                 sbc a,h
8E14  BE                 cp (hl)
8E15  BE                 cp (hl)
8E16  88                 adc a,b
8E17  88                 adc a,b
8E18  88                 adc a,b
8E19  88                 adc a,b
8E1A  6B                 ld l,e
8E1B  6B                 ld l,e
8E1C  08                 ex af,af'
8E1D  08                 ex af,af'
8E1E  08                 ex af,af'
8E1F  08                 ex af,af'
8E20  00                 nop
8E21  00                 nop
8E22  00                 nop
8E23  00                 nop
8E24  00                 nop
8E25  00                 nop
8E26  00                 nop
8E27  00                 nop
8E28  00                 nop
8E29  00                 nop
8E2A  00                 nop
8E2B  00                 nop
8E2C  00                 nop
8E2D  00                 nop
8E2E  00                 nop
8E2F  03                 inc bc
8E30  00                 nop
8E31  07                 rlca
8E32  00                 nop
8E33  0F                 rrca
8E34  00                 nop
8E35  3F                 ccf
8E36  00                 nop
8E37  FF                 rst 38h
8E38  01 FF 03           ld bc,1023
8E3B  FF                 rst 38h
8E3C  0F                 rrca
8E3D  FF                 rst 38h
8E3E  3F                 ccf
8E3F  FF                 rst 38h
8E40  00                 nop
8E41  60                 ld h,b
8E42  00                 nop
8E43  90                 sub b
8E44  03                 inc bc
8E45  6C                 ld l,h
8E46  0F                 rrca
8E47  72                 ld (hl),d
8E48  1F                 rra
8E49  D9                 exx
8E4A  3F                 ccf
8E4B  B6                 or (hl)
8E4C  FF                 rst 38h
8E4D  69                 ld l,c
8E4E  FF                 rst 38h
8E4F  E4 FF 52           call po,$52FF
8E52  FE D1              cp -47
8E54  FF                 rst 38h
8E55  50                 ld d,b
8E56  FE C8              cp -56
8E58  FB                 ei
8E59  28 FE              jr z,$8E59
8E5B  A4                 and h
8E5C  FD 24              inc iyh
8E5E  FE 00              cp 0
8E60  00                 nop
8E61  00                 nop
8E62  00                 nop
8E63  00                 nop
8E64  00                 nop
8E65  00                 nop
8E66  00                 nop
8E67  00                 nop
8E68  80                 add a,b
8E69  00                 nop
8E6A  40                 ld b,b
8E6B  00                 nop
8E6C  30 00              jr nc,$8E6E
8E6E  C8                 ret z
8E6F  00                 nop
8E70  26 00              ld h,0
8E72  19                 add hl,de
8E73  00                 nop
8E74  84                 add a,h
8E75  C0                 ret nz
8E76  82                 add a,d
8E77  20 41              jr nz,$8EBA
8E79  98                 sbc a,b
8E7A  20 44              jr nz,$8EC0
8E7C  00                 nop
8E7D  03                 inc bc
8E7E  00                 nop
8E7F  00                 nop
8E80  A8                 xor b
8E81  8A                 adc a,d
8E82  00                 nop
8E83  00                 nop
8E84  20 04              jr nz,$8E8A
8E86  20 00              jr nz,$8E88
8E88  22 20 22           ld ($2220),hl
8E8B  01 20 00           ld bc,32
8E8E  20 00              jr nz,$8E90
8E90  20 01              jr nz,$8E93
8E92  22 01 22           ld ($2201),hl
8E95  01 20 02           ld bc,544
8E98  20 02              jr nz,$8E9C
8E9A  20 8A              jr nz,$8E26
8E9C  20 0A              jr nz,$8EA8
8E9E  20 04              jr nz,$8EA4
8EA0  A0                 and b
8EA1  00                 nop
8EA2  00                 nop
8EA3  00                 nop
8EA4  48                 ld c,b
8EA5  95                 sub l
8EA6  03                 inc bc
8EA7  E2 1C 1D           jp po,$1D1C
8EAA  20 02              jr nz,$8EAE
8EAC  40                 ld b,b
8EAD  01 80 00           ld bc,128
8EB0  00                 nop
8EB1  00                 nop
8EB2  00                 nop
8EB3  00                 nop
8EB4  00                 nop
8EB5  00                 nop
8EB6  00                 nop
8EB7  00                 nop
8EB8  00                 nop
8EB9  00                 nop
8EBA  00                 nop
8EBB  00                 nop
8EBC  00                 nop
8EBD  00                 nop
8EBE  00                 nop
8EBF  00                 nop
8EC0  00                 nop
8EC1  00                 nop
8EC2  00                 nop
8EC3  02                 ld (bc),a
8EC4  57                 ld d,a
8EC5  F4 FF AE           call p,$AEFF
8EC8  7F                 ld a,a
8EC9  F4 DF FC           call p,$FCDF
8ECC  7F                 ld a,a
8ECD  FC B7 FC           call m,$FCB7
8ED0  5F                 ld e,a
8ED1  FC 5F FC           call m,$FC5F
8ED4  5F                 ld e,a
8ED5  FE 2E              cp 46
8ED7  FE 2F              cp 47
8ED9  FE 2B              cp 43
8EDB  FC 2F FC           call m,$FC2F
8EDE  2F                 cpl
8EDF  FC 20 04           call m,$0420
8EE2  20 04              jr nz,$8EE8
8EE4  22 04 22           ld ($2204),hl
8EE7  04                 inc b
8EE8  20 44              jr nz,$8F2E
8EEA  20 64              jr nz,$8F50
8EEC  20 26              jr nz,$8F14
8EEE  24                 inc h
8EEF  22 04 12           ld ($1204),hl
8EF2  12                 ld (de),a
8EF3  94                 sub h
8EF4  4A                 ld c,d
8EF5  DC 2A 44           call c,$442A
8EF8  6A                 ld l,d
8EF9  64                 ld h,h
8EFA  77                 ld (hl),a
8EFB  E4 76 A4           call po,$A476
8EFE  22 04 2F           ld ($2F04),hl
8F01  FC 17 FE           call m,$FE17
8F04  17                 rla
8F05  FE 17              cp 23
8F07  FE 17              cp 23
8F09  FE 17              cp 23
8F0B  CE 17              adc a,23
8F0D  9E                 sbc a,(hl)
8F0E  27                 daa
8F0F  9E                 sbc a,(hl)
8F10  2E 1E              ld l,30
8F12  2E 18              ld l,24
8F14  2C                 inc l
8F15  52                 ld d,d
8F16  24                 inc h
8F17  12                 ld (de),a
8F18  24                 inc h
8F19  0A                 ld a,(bc)
8F1A  00                 nop
8F1B  10 02              djnz $8F1F
8F1D  00                 nop
8F1E  0E 72              ld c,114
8F20  00                 nop
8F21  00                 nop
8F22  00                 nop
8F23  00                 nop
8F24  00                 nop
8F25  00                 nop
8F26  00                 nop
8F27  00                 nop
8F28  00                 nop
8F29  00                 nop
8F2A  00                 nop
8F2B  00                 nop
8F2C  00                 nop
8F2D  00                 nop
8F2E  00                 nop
8F2F  00                 nop
8F30  00                 nop
8F31  01 00 07           ld bc,1792
8F34  00                 nop
8F35  0F                 rrca
8F36  00                 nop
8F37  7E                 ld a,(hl)
8F38  01 F7 0F           ld bc,$0FF7
8F3B  FE 3F              cp 63
8F3D  D5                 push de
8F3E  FD BA              cp d
8F40  00                 nop
8F41  0F                 rrca
8F42  00                 nop
8F43  7F                 ld a,a
8F44  00                 nop
8F45  FF                 rst 38h
8F46  03                 inc bc
8F47  F7                 rst 30h
8F48  07                 rlca
8F49  FB                 ei
8F4A  1F                 rra
8F4B  F5                 push af
8F4C  1F                 rra
8F4D  EF                 rst 28h
8F4E  7D                 ld a,l
8F4F  F5                 push af
8F50  FD AE FB           xor (iy-5)
8F53  DD FA EF 35        jp m,$35EF
8F57  55                 ld d,l
8F58  EB                 ex de,hl
8F59  AB                 xor e
8F5A  5D                 ld e,l
8F5B  5D                 ld e,l
8F5C  BA                 cp d
8F5D  AB                 xor e
8F5E  55                 ld d,l
8F5F  55                 ld d,l
8F60  F0                 ret p
8F61  00                 nop
8F62  7E                 ld a,(hl)
8F63  00                 nop
8F64  8F                 adc a,a
8F65  80                 add a,b
8F66  57                 ld d,a
8F67  80                 add a,b
8F68  A3                 and e
8F69  E0                 ret po
8F6A  75                 ld (hl),l
8F6B  F8                 ret m
8F6C  AB                 xor e
8F6D  FC 55 6F           call m,$6F55
8F70  AB                 xor e
8F71  AF                 xor a
8F72  54                 ld d,h
8F73  55                 ld d,l
8F74  AE                 xor (hl)
8F75  DE 75              sbc a,117
8F77  53                 ld d,e
8F78  BA                 cp d
8F79  AA                 xor d
8F7A  D7                 rst 10h
8F7B  75                 ld (hl),l
8F7C  EA EA F5           jp pe,$F5EA
8F7F  55                 ld d,l
8F80  00                 nop
8F81  00                 nop
8F82  00                 nop
8F83  00                 nop
8F84  00                 nop
8F85  00                 nop
8F86  00                 nop
8F87  00                 nop
8F88  00                 nop
8F89  00                 nop
8F8A  00                 nop
8F8B  00                 nop
8F8C  00                 nop
8F8D  00                 nop
8F8E  00                 nop
8F8F  00                 nop
8F90  C0                 ret nz
8F91  00                 nop
8F92  F0                 ret p
8F93  00                 nop
8F94  3C                 inc a
8F95  00                 nop
8F96  C7                 rst 00h
8F97  00                 nop
8F98  D5                 push de
8F99  E0                 ret po
8F9A  AA                 xor d
8F9B  30 55              jr nc,$8FF2
8F9D  5C                 ld e,h
8F9E  EE BF              xor $BF
8FA0  F0                 ret p
8FA1  01 7E 07           ld bc,1918
8FA4  8F                 adc a,a
8FA5  8E                 adc a,(hl)
8FA6  57                 ld d,a
8FA7  9D                 sbc a,l
8FA8  A3                 and e
8FA9  E2 75 F9           jp po,$F975
8FAC  AB                 xor e
8FAD  FC 55 6F           call m,$6F55
8FB0  AB                 xor e
8FB1  AF                 xor a
8FB2  54                 ld d,h
8FB3  55                 ld d,l
8FB4  AE                 xor (hl)
8FB5  DE 75              sbc a,117
8FB7  53                 ld d,e
8FB8  BA                 cp d
8FB9  AA                 xor d
8FBA  D7                 rst 10h
8FBB  75                 ld (hl),l
8FBC  EA EA F5           jp pe,$F5EA
8FBF  55                 ld d,l
8FC0  AA                 xor d
8FC1  AF                 xor a
8FC2  5D                 ld e,l
8FC3  57                 ld d,a
8FC4  AF                 xor a
8FC5  AA                 xor d
8FC6  57                 ld d,a
8FC7  55                 ld d,l
8FC8  EA BA D5           jp pe,$D5BA
8FCB  5D                 ld e,l
8FCC  AA                 xor d
8FCD  BE                 cp (hl)
8FCE  15                 dec d
8FCF  55                 ld d,l
8FD0  CB AB              res 5,e
8FD2  F1                 pop af
8FD3  77                 ld (hl),a
8FD4  3C                 inc a
8FD5  2B                 dec hl
8FD6  CF                 rst 08h
8FD7  15                 dec d
8FD8  D7                 rst 10h
8FD9  E2 AB F1           jp po,$F1AB
8FDC  55                 ld d,l
8FDD  5C                 ld e,h
8FDE  EE BF              xor $BF
8FE0  AA                 xor d
8FE1  AF                 xor a
8FE2  5D                 ld e,l
8FE3  57                 ld d,a
8FE4  AF                 xor a
8FE5  AA                 xor d
8FE6  57                 ld d,a
8FE7  55                 ld d,l
8FE8  EA BA D5           jp pe,$D5BA
8FEB  5D                 ld e,l
8FEC  AA                 xor d
8FED  BE                 cp (hl)
8FEE  55                 ld d,l
8FEF  55                 ld d,l
8FF0  AB                 xor e
8FF1  AB                 xor e
8FF2  DD 77 AE           ld (ix-82),a
8FF5  AB                 xor e
8FF6  5D                 ld e,l
8FF7  55                 ld d,l
8FF8  AA                 xor d
8FF9  AA                 xor d
8FFA  55                 ld d,l
8FFB  5D                 ld e,l
8FFC  AE                 xor (hl)
8FFD  BE                 cp (hl)
8FFE  57                 ld d,a
8FFF  5D                 ld e,l
9000  00                 nop
9001  00                 nop
9002  00                 nop
9003  00                 nop
9004  00                 nop
9005  00                 nop
9006  00                 nop
9007  00                 nop
9008  00                 nop
9009  00                 nop
900A  00                 nop
900B  00                 nop
900C  00                 nop
900D  00                 nop
900E  00                 nop
900F  00                 nop
9010  00                 nop
9011  00                 nop
9012  00                 nop
9013  00                 nop
9014  00                 nop
9015  00                 nop
9016  00                 nop
9017  00                 nop
9018  00                 nop
9019  00                 nop
901A  00                 nop
901B  00                 nop
901C  00                 nop
901D  00                 nop
901E  00                 nop
901F  00                 nop
9020  F0                 ret p
9021  00                 nop
9022  9F                 sbc a,a
9023  FE 9F              cp -97
9025  FE F0              cp -16
9027  01 0A AA           ld bc,$AA0A
902A  55                 ld d,l
902B  55                 ld d,l
902C  AA                 xor d
902D  AA                 xor d
902E  55                 ld d,l
902F  55                 ld d,l
9030  AA                 xor d
9031  AA                 xor d
9032  55                 ld d,l
9033  55                 ld d,l
9034  AA                 xor d
9035  AA                 xor d
9036  55                 ld d,l
9037  55                 ld d,l
9038  AA                 xor d
9039  AA                 xor d
903A  55                 ld d,l
903B  55                 ld d,l
903C  AA                 xor d
903D  AA                 xor d
903E  55                 ld d,l
903F  55                 ld d,l
9040  F2 AA 91           jp p,$91AA
9043  55                 ld d,l
9044  9C                 sbc a,h
9045  AA                 xor d
9046  FF                 rst 38h
9047  15                 dec d
9048  03                 inc bc
9049  CA 54 F5           jp z,$F554
904C  AA                 xor d
904D  32 55 45           ld ($4555),a
9050  AA                 xor d
9051  AA                 xor d
9052  55                 ld d,l
9053  55                 ld d,l
9054  AA                 xor d
9055  AA                 xor d
9056  55                 ld d,l
9057  55                 ld d,l
9058  AA                 xor d
9059  AA                 xor d
905A  55                 ld d,l
905B  55                 ld d,l
905C  AA                 xor d
905D  AA                 xor d
905E  55                 ld d,l
905F  55                 ld d,l
9060  F2 AA 95           jp p,$95AA
9063  55                 ld d,l
9064  92                 sub d
9065  AA                 xor d
9066  F5                 push af
9067  55                 ld d,l
9068  1A                 ld a,(de)
9069  AA                 xor d
906A  59                 ld e,c
906B  55                 ld d,l
906C  AC                 xor h
906D  AA                 xor d
906E  4D                 ld c,l
906F  55                 ld d,l
9070  A6                 and (hl)
9071  AA                 xor d
9072  56                 ld d,(hl)
9073  55                 ld d,l
9074  AB                 xor e
9075  2A 53 55           ld hl,($5553)
9078  A8                 xor b
9079  AA                 xor d
907A  55                 ld d,l
907B  55                 ld d,l
907C  AA                 xor d
907D  AA                 xor d
907E  55                 ld d,l
907F  55                 ld d,l
9080  F2 AA 95           jp p,$95AA
9083  55                 ld d,l
9084  92                 sub d
9085  AA                 xor d
9086  F5                 push af
9087  55                 ld d,l
9088  6A                 ld l,d
9089  AA                 xor d
908A  65                 ld h,l
908B  55                 ld d,l
908C  6A                 ld l,d
908D  AA                 xor d
908E  65                 ld h,l
908F  55                 ld d,l
9090  6A                 ld l,d
9091  AA                 xor d
9092  65                 ld h,l
9093  55                 ld d,l
9094  6A                 ld l,d
9095  AA                 xor d
9096  65                 ld h,l
9097  55                 ld d,l
9098  6A                 ld l,d
9099  AA                 xor d
909A  65                 ld h,l
909B  55                 ld d,l
909C  6A                 ld l,d
909D  AA                 xor d
909E  65                 ld h,l
909F  55                 ld d,l
90A0  80                 add a,b
90A1  0F                 rrca
90A2  7F                 ld a,a
90A3  F9                 ld sp,hl
90A4  7F                 ld a,a
90A5  F9                 ld sp,hl
90A6  00                 nop
90A7  0F                 rrca
90A8  AA                 xor d
90A9  A0                 and b
90AA  55                 ld d,l
90AB  55                 ld d,l
90AC  AA                 xor d
90AD  AA                 xor d
90AE  55                 ld d,l
90AF  55                 ld d,l
90B0  AA                 xor d
90B1  AA                 xor d
90B2  55                 ld d,l
90B3  55                 ld d,l
90B4  AA                 xor d
90B5  AA                 xor d
90B6  55                 ld d,l
90B7  55                 ld d,l
90B8  AA                 xor d
90B9  AA                 xor d
90BA  55                 ld d,l
90BB  55                 ld d,l
90BC  AA                 xor d
90BD  AA                 xor d
90BE  55                 ld d,l
90BF  55                 ld d,l
90C0  AA                 xor d
90C1  AF                 xor a
90C2  55                 ld d,l
90C3  49                 ld c,c
90C4  AA                 xor d
90C5  39                 add hl,sp
90C6  54                 ld d,h
90C7  FF                 rst 38h
90C8  A3                 and e
90C9  C0                 ret nz
90CA  4F                 ld c,a
90CB  15                 dec d
90CC  AC                 xor h
90CD  AA                 xor d
90CE  51                 ld d,c
90CF  55                 ld d,l
90D0  AA                 xor d
90D1  AA                 xor d
90D2  55                 ld d,l
90D3  55                 ld d,l
90D4  AA                 xor d
90D5  AA                 xor d
90D6  55                 ld d,l
90D7  55                 ld d,l
90D8  AA                 xor d
90D9  AA                 xor d
90DA  55                 ld d,l
90DB  55                 ld d,l
90DC  AA                 xor d
90DD  AA                 xor d
90DE  55                 ld d,l
90DF  55                 ld d,l
90E0  AA                 xor d
90E1  AF                 xor a
90E2  55                 ld d,l
90E3  49                 ld c,c
90E4  AA                 xor d
90E5  A9                 xor c
90E6  55                 ld d,l
90E7  4F                 ld c,a
90E8  AA                 xor d
90E9  98                 sbc a,b
90EA  55                 ld d,l
90EB  59                 ld e,c
90EC  AA                 xor d
90ED  B2                 or d
90EE  55                 ld d,l
90EF  35                 dec (hl)
90F0  AA                 xor d
90F1  6A                 ld l,d
90F2  55                 ld d,l
90F3  65                 ld h,l
90F4  AA                 xor d
90F5  CA 54 D5           jp z,$D554
90F8  AA                 xor d
90F9  2A 55 55           ld hl,($5555)
90FC  AA                 xor d
90FD  AA                 xor d
90FE  55                 ld d,l
90FF  55                 ld d,l
9100  AA                 xor d
9101  AF                 xor a
9102  55                 ld d,l
9103  49                 ld c,c
9104  AA                 xor d
9105  A9                 xor c
9106  55                 ld d,l
9107  4F                 ld c,a
9108  AA                 xor d
9109  A6                 and (hl)
910A  55                 ld d,l
910B  56                 ld d,(hl)
910C  AA                 xor d
910D  A6                 and (hl)
910E  55                 ld d,l
910F  56                 ld d,(hl)
9110  AA                 xor d
9111  A6                 and (hl)
9112  55                 ld d,l
9113  56                 ld d,(hl)
9114  AA                 xor d
9115  A6                 and (hl)
9116  55                 ld d,l
9117  56                 ld d,(hl)
9118  AA                 xor d
9119  A6                 and (hl)
911A  55                 ld d,l
911B  56                 ld d,(hl)
911C  AA                 xor d
911D  A6                 and (hl)
911E  55                 ld d,l
911F  56                 ld d,(hl)
9120  00                 nop
9121  00                 nop
9122  00                 nop
9123  00                 nop
9124  00                 nop
9125  00                 nop
9126  00                 nop
9127  00                 nop
9128  00                 nop
9129  00                 nop
912A  00                 nop
912B  01 00 01           ld bc,256
912E  00                 nop
912F  01 00 02           ld bc,512
9132  00                 nop
9133  02                 ld (bc),a
9134  00                 nop
9135  01 00 01           ld bc,256
9138  00                 nop
9139  03                 inc bc
913A  00                 nop
913B  03                 inc bc
913C  00                 nop
913D  03                 inc bc
913E  00                 nop
913F  03                 inc bc
9140  00                 nop
9141  00                 nop
9142  00                 nop
9143  00                 nop
9144  C0                 ret nz
9145  00                 nop
9146  E0                 ret po
9147  03                 inc bc
9148  F0                 ret p
9149  07                 rlca
914A  CF                 rst 08h
914B  C7                 rst 00h
914C  BF                 cp a
914D  F7                 rst 30h
914E  7F                 ld a,a
914F  EF                 rst 28h
9150  FF                 rst 38h
9151  FF                 rst 38h
9152  F9                 ld sp,hl
9153  7F                 ld a,a
9154  F6 BF              or $BF
9156  E3                 ex (sp),hl
9157  3F                 ccf
9158  C5                 push bc
9159  2F                 cpl
915A  A0                 and b
915B  F7                 rst 30h
915C  C9                 ret
915D  F7                 rst 30h
915E  9F                 sbc a,a
915F  F7                 rst 30h
9160  00                 nop
9161  00                 nop
9162  00                 nop
9163  00                 nop
9164  00                 nop
9165  00                 nop
9166  80                 add a,b
9167  00                 nop
9168  E0                 ret po
9169  00                 nop
916A  FE 00              cp 0
916C  FF                 rst 38h
916D  00                 nop
916E  FF                 rst 38h
916F  C0                 ret nz
9170  FF                 rst 38h
9171  C0                 ret nz
9172  5F                 ld e,a
9173  80                 add a,b
9174  AF                 xor a
9175  00                 nop
9176  D5                 push de
9177  00                 nop
9178  EA 00 F0           jp pe,$F000
917B  00                 nop
917C  EA 00 F6           jp pe,$F600
917F  00                 nop
9180  00                 nop
9181  07                 rlca
9182  00                 nop
9183  07                 rlca
9184  00                 nop
9185  07                 rlca
9186  00                 nop
9187  07                 rlca
9188  00                 nop
9189  07                 rlca
918A  01 D6 1F           ld bc,$1FD6
918D  A7                 and a
918E  3F                 ccf
918F  CE 3F              adc a,63
9191  AF                 xor a
9192  3F                 ccf
9193  4E                 ld c,(hl)
9194  3F                 ccf
9195  9F                 sbc a,a
9196  3F                 ccf
9197  3E 3E              ld a,62
9199  BD                 cp l
919A  3F                 ccf
919B  5E                 ld e,(hl)
919C  3A 9E 1C           ld a,($1C9E)
919F  05                 dec b
91A0  5D                 ld e,l
91A1  F3                 di
91A2  9D                 sbc a,l
91A3  EB                 ex de,hl
91A4  5E                 ld e,(hl)
91A5  71                 ld (hl),c
91A6  83                 add a,e
91A7  E9                 jp (hl)
91A8  5F                 ld e,a
91A9  E2 8F 15           jp po,$158F
91AC  00                 nop
91AD  AA                 xor d
91AE  8B                 adc a,e
91AF  47                 ld b,a
91B0  16 CF              ld d,-49
91B2  81                 add a,c
91B3  AF                 xor a
91B4  3F                 ccf
91B5  97                 sub a
91B6  5F                 ld e,a
91B7  B7                 or a
91B8  3F                 ccf
91B9  59                 ld e,c
91BA  5F                 ld e,a
91BB  3E BF              ld a,-65
91BD  5F                 ld e,a
91BE  7F                 ld a,a
91BF  3F                 ccf
91C0  FA 00 FE           jp m,$FE00
91C3  00                 nop
91C4  FE 00              cp 0
91C6  FE 00              cp 0
91C8  78                 ld a,b
91C9  00                 nop
91CA  00                 nop
91CB  00                 nop
91CC  A0                 and b
91CD  00                 nop
91CE  D8                 ret c
91CF  00                 nop
91D0  FC 00 FE           call m,$FE00
91D3  00                 nop
91D4  FF                 rst 38h
91D5  00                 nop
91D6  FF                 rst 38h
91D7  00                 nop
91D8  FF                 rst 38h
91D9  00                 nop
91DA  FF                 rst 38h
91DB  80                 add a,b
91DC  7F                 ld a,a
91DD  80                 add a,b
91DE  7F                 ld a,a
91DF  80                 add a,b
91E0  00                 nop
91E1  00                 nop
91E2  00                 nop
91E3  01 00 02           ld bc,512
91E6  00                 nop
91E7  02                 ld (bc),a
91E8  00                 nop
91E9  00                 nop
91EA  00                 nop
91EB  06 00              ld b,0
91ED  06 00              ld b,0
91EF  05                 dec b
91F0  00                 nop
91F1  04                 inc b
91F2  00                 nop
91F3  05                 dec b
91F4  00                 nop
91F5  00                 nop
91F6  00                 nop
91F7  00                 nop
91F8  00                 nop
91F9  01 00 01           ld bc,256
91FC  00                 nop
91FD  01 00 01           ld bc,256
9200  FF                 rst 38h
9201  5B                 ld e,e
9202  7F                 ld a,a
9203  A7                 and a
9204  FE 87              cp -121
9206  7D                 ld a,l
9207  4F                 ld c,a
9208  AA                 xor d
9209  A1                 and c
920A  00                 nop
920B  0B                 dec bc
920C  AA                 xor d
920D  E7                 rst 20h
920E  7F                 ld a,a
920F  EB                 ex de,hl
9210  FF                 rst 38h
9211  E7                 rst 20h
9212  7F                 ld a,a
9213  EB                 ex de,hl
9214  FF                 rst 38h
9215  E0                 ret po
9216  00                 nop
9217  1F                 rra
9218  FD 3F              ccf
921A  FE BF              cp -65
921C  FD 3F              ccf
921E  32 99 BF           ld ($BF99),a
9221  C0                 ret nz
9222  BF                 cp a
9223  C0                 ret nz
9224  BF                 cp a
9225  40                 ld b,b
9226  D2 00 C0           jp nc,$C000
9229  00                 nop
922A  E0                 ret po
922B  00                 nop
922C  F0                 ret p
922D  00                 nop
922E  F0                 ret p
922F  00                 nop
9230  F0                 ret p
9231  00                 nop
9232  F8                 ret m
9233  00                 nop
9234  7C                 ld a,h
9235  00                 nop
9236  80                 add a,b
9237  00                 nop
9238  F0                 ret p
9239  00                 nop
923A  F0                 ret p
923B  00                 nop
923C  F8                 ret m
923D  00                 nop
923E  98                 sbc a,b
923F  00                 nop
9240  FF                 rst 38h
9241  00                 nop
9242  FF                 rst 38h
9243  00                 nop
9244  FF                 rst 38h
9245  00                 nop
9246  FF                 rst 38h
9247  00                 nop
9248  8F                 adc a,a
9249  00                 nop
924A  9F                 sbc a,a
924B  00                 nop
924C  FF                 rst 38h
924D  00                 nop
924E  02                 ld (bc),a
924F  00                 nop
9250  0F                 rrca
9251  00                 nop
9252  FE 00              cp 0
9254  00                 nop
9255  70                 ld (hl),b
9256  07                 rlca
9257  60                 ld h,b
9258  FC 00 00           call m,$0000
925B  CD 03 90           call $9003
925E  FC 00 00           call m,$0000
9261  BF                 cp a
9262  01 78 F8           ld bc,-1928
9265  01 00 79           ld bc,$7900
9268  01 B8 F8           ld bc,-1864
926B  01 00 76           ld bc,$7600
926E  01 A8 FC           ld bc,-856
9271  00                 nop
9272  00                 nop
9273  EA 01 98           jp pe,$9801
9276  FC 00 00           call m,$0000
9279  E2 01 98           jp po,$9801
927C  FC 00 00           call m,$0000
927F  CD 01 B8           call $B801
9282  F8                 ret m
9283  01 00 DE           ld bc,$DE00
9286  03                 inc bc
9287  D0                 ret nc
9288  F8                 ret m
9289  01 00 DE           ld bc,$DE00
928C  03                 inc bc
928D  50                 ld d,b
928E  F8                 ret m
928F  01 00 CC           ld bc,$CC00
9292  03                 inc bc
9293  F0                 ret p
9294  F8                 ret m
9295  01 00 92           ld bc,$9200
9298  07                 rlca
9299  E0                 ret po
929A  F8                 ret m
929B  01 00 C7           ld bc,$C700
929E  0F                 rrca
929F  40                 ld b,b
92A0  FC 00 00           call m,$0000
92A3  E7                 rst 20h
92A4  0F                 rrca
92A5  80                 add a,b
92A6  FC 00 00           call m,$0000
92A9  EF                 rst 28h
92AA  0F                 rrca
92AB  40                 ld b,b
92AC  FE 00              cp 0
92AE  00                 nop
92AF  0F                 rrca
92B0  0F                 rrca
92B1  40                 ld b,b
92B2  FE 00              cp 0
92B4  00                 nop
92B5  4E                 ld c,(hl)
92B6  0F                 rrca
92B7  80                 add a,b
92B8  FC 00 00           call m,$0000
92BB  E1                 pop hl
92BC  07                 rlca
92BD  E0                 ret po
92BE  FC 00 00           call m,$0000
92C1  F3                 di
92C2  07                 rlca
92C3  E0                 ret po
92C4  F8                 ret m
92C5  01 00 FB           ld bc,-1280
92C8  03                 inc bc
92C9  F0                 ret p
92CA  FC 00 00           call m,$0000
92CD  F9                 ld sp,hl
92CE  07                 rlca
92CF  E0                 ret po
92D0  FF                 rst 38h
92D1  00                 nop
92D2  FF                 rst 38h
92D3  00                 nop
92D4  FF                 rst 38h
92D5  00                 nop
92D6  FF                 rst 38h
92D7  00                 nop
92D8  8F                 adc a,a
92D9  00                 nop
92DA  9F                 sbc a,a
92DB  00                 nop
92DC  FF                 rst 38h
92DD  00                 nop
92DE  02                 ld (bc),a
92DF  00                 nop
92E0  0F                 rrca
92E1  00                 nop
92E2  FE 00              cp 0
92E4  00                 nop
92E5  70                 ld (hl),b
92E6  07                 rlca
92E7  60                 ld h,b
92E8  FC 00 00           call m,$0000
92EB  CD 03 90           call $9003
92EE  FC 00 00           call m,$0000
92F1  BF                 cp a
92F2  01 78 F8           ld bc,-1928
92F5  01 00 79           ld bc,$7900
92F8  01 B8 F8           ld bc,-1864
92FB  01 00 76           ld bc,$7600
92FE  01 A8 FC           ld bc,-856
9301  00                 nop
9302  00                 nop
9303  6A                 ld l,d
9304  01 98 FC           ld bc,-872
9307  00                 nop
9308  00                 nop
9309  E2 01 98           jp po,$9801
930C  FC 00 00           call m,$0000
930F  ED 01              ???
9311  B8                 cp b
9312  FC 00 00           call m,$0000
9315  EE 03              xor $03
9317  D0                 ret nc
9318  FC 00 00           call m,$0000
931B  EE 03              xor $03
931D  50                 ld d,b
931E  FC 00 00           call m,$0000
9321  E0                 ret po
9322  03                 inc bc
9323  F0                 ret p
9324  FC 00 00           call m,$0000
9327  F2 07 E0           jp p,$E007
932A  FE 00              cp 0
932C  00                 nop
932D  77                 ld (hl),a
932E  0F                 rrca
932F  40                 ld b,b
9330  FE 00              cp 0
9332  00                 nop
9333  77                 ld (hl),a
9334  07                 rlca
9335  80                 add a,b
9336  FF                 rst 38h
9337  00                 nop
9338  00                 nop
9339  2F                 cpl
933A  07                 rlca
933B  A0                 and b
933C  FF                 rst 38h
933D  00                 nop
933E  00                 nop
933F  0F                 rrca
9340  0F                 rrca
9341  80                 add a,b
9342  FE 00              cp 0
9344  00                 nop
9345  43                 ld b,e
9346  1F                 rra
9347  00                 nop
9348  FE 00              cp 0
934A  00                 nop
934B  70                 ld (hl),b
934C  0F                 rrca
934D  40                 ld b,b
934E  FC 00 00           call m,$0000
9351  F7                 rst 30h
9352  0F                 rrca
9353  C0                 ret nz
9354  FC 00 00           call m,$0000
9357  FB                 ei
9358  1F                 rra
9359  80                 add a,b
935A  FE 00              cp 0
935C  00                 nop
935D  78                 ld a,b
935E  3F                 ccf
935F  00                 nop
9360  FF                 rst 38h
9361  00                 nop
9362  FF                 rst 38h
9363  00                 nop
9364  FF                 rst 38h
9365  00                 nop
9366  FF                 rst 38h
9367  00                 nop
9368  8F                 adc a,a
9369  00                 nop
936A  9F                 sbc a,a
936B  00                 nop
936C  FF                 rst 38h
936D  00                 nop
936E  02                 ld (bc),a
936F  00                 nop
9370  0F                 rrca
9371  00                 nop
9372  FE 00              cp 0
9374  00                 nop
9375  70                 ld (hl),b
9376  07                 rlca
9377  60                 ld h,b
9378  FC 00 00           call m,$0000
937B  CD 03 90           call $9003
937E  FC 00 00           call m,$0000
9381  BF                 cp a
9382  01 78 F8           ld bc,-1928
9385  01 00 79           ld bc,$7900
9388  01 B8 F8           ld bc,-1864
938B  01 00 76           ld bc,$7600
938E  01 A8 FC           ld bc,-856
9391  00                 nop
9392  00                 nop
9393  EA 01 98           jp pe,$9801
9396  FC 00 00           call m,$0000
9399  E2 01 98           jp po,$9801
939C  FC 00 00           call m,$0000
939F  CD 01 B8           call $B801
93A2  F8                 ret m
93A3  01 00 DE           ld bc,$DE00
93A6  03                 inc bc
93A7  D0                 ret nc
93A8  F8                 ret m
93A9  01 00 DE           ld bc,$DE00
93AC  03                 inc bc
93AD  50                 ld d,b
93AE  F8                 ret m
93AF  01 00 8C           ld bc,$8C00
93B2  03                 inc bc
93B3  F0                 ret p
93B4  F8                 ret m
93B5  01 00 D2           ld bc,$D200
93B8  07                 rlca
93B9  E0                 ret po
93BA  F8                 ret m
93BB  01 00 C7           ld bc,$C700
93BE  0F                 rrca
93BF  40                 ld b,b
93C0  F0                 ret p
93C1  03                 inc bc
93C2  00                 nop
93C3  87                 add a,a
93C4  1F                 rra
93C5  80                 add a,b
93C6  F0                 ret p
93C7  03                 inc bc
93C8  00                 nop
93C9  83                 add a,e
93CA  0F                 rrca
93CB  C0                 ret nz
93CC  F0                 ret p
93CD  03                 inc bc
93CE  00                 nop
93CF  1B                 dec de
93D0  0F                 rrca
93D1  C0                 ret nz
93D2  F8                 ret m
93D3  00                 nop
93D4  00                 nop
93D5  59                 ld e,c
93D6  1F                 rra
93D7  80                 add a,b
93D8  FC 00 00           call m,$0000
93DB  3C                 inc a
93DC  3F                 ccf
93DD  00                 nop
93DE  FE 00              cp 0
93E0  00                 nop
93E1  3E 7F              ld a,127
93E3  00                 nop
93E4  FE 00              cp 0
93E6  00                 nop
93E7  0C                 inc c
93E8  FF                 rst 38h
93E9  00                 nop
93EA  FF                 rst 38h
93EB  00                 nop
93EC  01 30 FF           ld bc,-208
93EF  00                 nop
93F0  FF                 rst 38h
93F1  00                 nop
93F2  FF                 rst 38h
93F3  00                 nop
93F4  FF                 rst 38h
93F5  00                 nop
93F6  FF                 rst 38h
93F7  00                 nop
93F8  8F                 adc a,a
93F9  00                 nop
93FA  9F                 sbc a,a
93FB  00                 nop
93FC  FF                 rst 38h
93FD  00                 nop
93FE  02                 ld (bc),a
93FF  00                 nop
9400  0F                 rrca
9401  00                 nop
9402  FE 00              cp 0
9404  00                 nop
9405  70                 ld (hl),b
9406  07                 rlca
9407  60                 ld h,b
9408  FC 00 00           call m,$0000
940B  CD 03 90           call $9003
940E  FC 00 00           call m,$0000
9411  BF                 cp a
9412  01 78 F8           ld bc,-1928
9415  01 00 79           ld bc,$7900
9418  01 B8 F8           ld bc,-1864
941B  01 00 76           ld bc,$7600
941E  01 A8 FC           ld bc,-856
9421  00                 nop
9422  00                 nop
9423  EA 01 98           jp pe,$9801
9426  FC 00 00           call m,$0000
9429  E2 01 98           jp po,$9801
942C  F0                 ret p
942D  00                 nop
942E  00                 nop
942F  CD 01 B8           call $B801
9432  E0                 ret po
9433  01 00 DE           ld bc,$DE00
9436  03                 inc bc
9437  D0                 ret nc
9438  C0                 ret nz
9439  0D                 dec c
943A  00                 nop
943B  DE 03              sbc a,3
943D  50                 ld d,b
943E  C0                 ret nz
943F  0B                 dec bc
9440  00                 nop
9441  8C                 adc a,h
9442  03                 inc bc
9443  F0                 ret p
9444  C0                 ret nz
9445  0F                 rrca
9446  00                 nop
9447  A2                 and d
9448  07                 rlca
9449  E0                 ret po
944A  E0                 ret po
944B  07                 rlca
944C  00                 nop
944D  3B                 dec sp
944E  0F                 rrca
944F  40                 ld b,b
9450  F0                 ret p
9451  00                 nop
9452  00                 nop
9453  2B                 dec hl
9454  07                 rlca
9455  A0                 and b
9456  F8                 ret m
9457  00                 nop
9458  80                 add a,b
9459  19                 add hl,de
945A  07                 rlca
945B  E0                 ret po
945C  FF                 rst 38h
945D  00                 nop
945E  80                 add a,b
945F  1D                 dec e
9460  07                 rlca
9461  E0                 ret po
9462  FF                 rst 38h
9463  00                 nop
9464  00                 nop
9465  3C                 inc a
9466  0F                 rrca
9467  C0                 ret nz
9468  FE 00              cp 0
946A  00                 nop
946B  7E                 ld a,(hl)
946C  1F                 rra
946D  00                 nop
946E  FC 00 00           call m,$0000
9471  FC 1F 80           call m,$801F
9474  FC 00 00           call m,$0000
9477  FD 1F              rra
9479  80                 add a,b
947A  FE 00              cp 0
947C  00                 nop
947D  79                 ld a,c
947E  3F                 ccf
947F  00                 nop
9480  FF                 rst 38h
9481  00                 nop
9482  FF                 rst 38h
9483  00                 nop
9484  FF                 rst 38h
9485  00                 nop
9486  FF                 rst 38h
9487  00                 nop
9488  8F                 adc a,a
9489  00                 nop
948A  9F                 sbc a,a
948B  00                 nop
948C  FF                 rst 38h
948D  00                 nop
948E  02                 ld (bc),a
948F  00                 nop
9490  0F                 rrca
9491  00                 nop
9492  FE 00              cp 0
9494  00                 nop
9495  70                 ld (hl),b
9496  07                 rlca
9497  60                 ld h,b
9498  FC 00 00           call m,$0000
949B  CD 03 90           call $9003
949E  FC 00 00           call m,$0000
94A1  BF                 cp a
94A2  01 78 F8           ld bc,-1928
94A5  01 00 79           ld bc,$7900
94A8  01 B8 F8           ld bc,-1864
94AB  01 00 76           ld bc,$7600
94AE  01 A8 F8           ld bc,-1880
94B1  01 00 6A           ld bc,$6A00
94B4  01 98 FC           ld bc,-872
94B7  00                 nop
94B8  00                 nop
94B9  E2 01 98           jp po,$9801
94BC  FC 00 00           call m,$0000
94BF  ED 01              ???
94C1  B8                 cp b
94C2  FC 00 00           call m,$0000
94C5  DE 03              sbc a,3
94C7  D0                 ret nc
94C8  F8                 ret m
94C9  01 00 DE           ld bc,$DE00
94CC  03                 inc bc
94CD  50                 ld d,b
94CE  F0                 ret p
94CF  01 00 CC           ld bc,$CC00
94D2  03                 inc bc
94D3  F0                 ret p
94D4  E0                 ret po
94D5  05                 dec b
94D6  00                 nop
94D7  A2                 and d
94D8  07                 rlca
94D9  E0                 ret po
94DA  E0                 ret po
94DB  07                 rlca
94DC  00                 nop
94DD  BF                 cp a
94DE  03                 inc bc
94DF  50                 ld d,b
94E0  F0                 ret p
94E1  03                 inc bc
94E2  00                 nop
94E3  5D                 ld e,l
94E4  01 B8 F8           ld bc,-1864
94E7  00                 nop
94E8  00                 nop
94E9  4C                 ld c,h
94EA  01 F8 FC           ld bc,-776
94ED  00                 nop
94EE  00                 nop
94EF  1E 03              ld e,3
94F1  30 FF              jr nc,$94F2
94F3  00                 nop
94F4  80                 add a,b
94F5  1E 07              ld e,7
94F7  00                 nop
94F8  FF                 rst 38h
94F9  00                 nop
94FA  00                 nop
94FB  3E 0F              ld a,15
94FD  80                 add a,b
94FE  FF                 rst 38h
94FF  00                 nop
9500  00                 nop
9501  3E 3F              ld a,63
9503  00                 nop
9504  FF                 rst 38h
9505  00                 nop
9506  00                 nop
9507  3E 7F              ld a,127
9509  00                 nop
950A  FF                 rst 38h
950B  00                 nop
950C  80                 add a,b
950D  1C                 inc e
950E  FF                 rst 38h
950F  00                 nop
9510  FF                 rst 38h
9511  00                 nop
9512  FF                 rst 38h
9513  00                 nop
9514  FF                 rst 38h
9515  00                 nop
9516  FF                 rst 38h
9517  00                 nop
9518  8F                 adc a,a
9519  00                 nop
951A  9F                 sbc a,a
951B  00                 nop
951C  FF                 rst 38h
951D  00                 nop
951E  02                 ld (bc),a
951F  00                 nop
9520  0F                 rrca
9521  00                 nop
9522  FE 00              cp 0
9524  00                 nop
9525  70                 ld (hl),b
9526  07                 rlca
9527  60                 ld h,b
9528  FC 00 00           call m,$0000
952B  CD 03 90           call $9003
952E  FC 00 00           call m,$0000
9531  BF                 cp a
9532  01 78 F8           ld bc,-1928
9535  01 00 79           ld bc,$7900
9538  01 B8 F8           ld bc,-1864
953B  01 00 76           ld bc,$7600
953E  01 A8 FC           ld bc,-856
9541  00                 nop
9542  00                 nop
9543  EA 01 98           jp pe,$9801
9546  FC 00 00           call m,$0000
9549  E2 01 98           jp po,$9801
954C  FC 00 00           call m,$0000
954F  CD 01 B8           call $B801
9552  F8                 ret m
9553  01 00 DE           ld bc,$DE00
9556  03                 inc bc
9557  D0                 ret nc
9558  F8                 ret m
9559  01 00 DE           ld bc,$DE00
955C  03                 inc bc
955D  50                 ld d,b
955E  F8                 ret m
955F  01 00 8C           ld bc,$8C00
9562  03                 inc bc
9563  F0                 ret p
9564  F8                 ret m
9565  01 00 A2           ld bc,$A200
9568  07                 rlca
9569  E0                 ret po
956A  F0                 ret p
956B  03                 inc bc
956C  00                 nop
956D  BF                 cp a
956E  0F                 rrca
956F  40                 ld b,b
9570  E0                 ret po
9571  07                 rlca
9572  00                 nop
9573  5B                 ld e,e
9574  07                 rlca
9575  A0                 and b
9576  E0                 ret po
9577  07                 rlca
9578  00                 nop
9579  59                 ld e,c
957A  07                 rlca
957B  E0                 ret po
957C  F0                 ret p
957D  00                 nop
957E  00                 nop
957F  0D                 dec c
9580  07                 rlca
9581  E0                 ret po
9582  F8                 ret m
9583  00                 nop
9584  00                 nop
9585  0C                 inc c
9586  0F                 rrca
9587  C0                 ret nz
9588  FE 00              cp 0
958A  00                 nop
958B  5E                 ld e,(hl)
958C  1F                 rra
958D  00                 nop
958E  FE 00              cp 0
9590  00                 nop
9591  5F                 ld e,a
9592  3F                 ccf
9593  00                 nop
9594  FF                 rst 38h
9595  00                 nop
9596  00                 nop
9597  1F                 rra
9598  3F                 ccf
9599  00                 nop
959A  FF                 rst 38h
959B  00                 nop
959C  80                 add a,b
959D  0F                 rrca
959E  3F                 ccf
959F  00                 nop
95A0  FF                 rst 38h
95A1  00                 nop
95A2  FF                 rst 38h
95A3  00                 nop
95A4  FF                 rst 38h
95A5  00                 nop
95A6  FF                 rst 38h
95A7  00                 nop
95A8  8F                 adc a,a
95A9  00                 nop
95AA  9F                 sbc a,a
95AB  00                 nop
95AC  FF                 rst 38h
95AD  00                 nop
95AE  02                 ld (bc),a
95AF  00                 nop
95B0  0F                 rrca
95B1  00                 nop
95B2  FE 00              cp 0
95B4  00                 nop
95B5  70                 ld (hl),b
95B6  07                 rlca
95B7  60                 ld h,b
95B8  FC 00 00           call m,$0000
95BB  CD 03 90           call $9003
95BE  FC 00 00           call m,$0000
95C1  BF                 cp a
95C2  01 78 F8           ld bc,-1928
95C5  01 00 79           ld bc,$7900
95C8  01 B8 F8           ld bc,-1864
95CB  01 00 76           ld bc,$7600
95CE  01 A8 FC           ld bc,-856
95D1  00                 nop
95D2  00                 nop
95D3  EA 01 98           jp pe,$9801
95D6  FC 00 00           call m,$0000
95D9  E2 01 98           jp po,$9801
95DC  FC 00 00           call m,$0000
95DF  CD 01 B8           call $B801
95E2  F8                 ret m
95E3  01 00 DE           ld bc,$DE00
95E6  03                 inc bc
95E7  D0                 ret nc
95E8  F8                 ret m
95E9  01 00 DE           ld bc,$DE00
95EC  03                 inc bc
95ED  50                 ld d,b
95EE  F8                 ret m
95EF  01 00 8C           ld bc,$8C00
95F2  03                 inc bc
95F3  F0                 ret p
95F4  F8                 ret m
95F5  01 00 A2           ld bc,$A200
95F8  07                 rlca
95F9  E0                 ret po
95FA  F8                 ret m
95FB  01 00 BF           ld bc,$BF00
95FE  0F                 rrca
95FF  40                 ld b,b
9600  F8                 ret m
9601  01 00 DB           ld bc,$DB00
9604  1F                 rra
9605  80                 add a,b
9606  F8                 ret m
9607  01 00 D3           ld bc,$D300
960A  0F                 rrca
960B  C0                 ret nz
960C  FC 00 00           call m,$0000
960F  13                 inc de
9610  0F                 rrca
9611  C0                 ret nz
9612  FE 00              cp 0
9614  00                 nop
9615  43                 ld b,e
9616  0F                 rrca
9617  C0                 ret nz
9618  FC 00 00           call m,$0000
961B  60                 ld h,b
961C  1F                 rra
961D  00                 nop
961E  FC 00 00           call m,$0000
9621  F7                 rst 30h
9622  0F                 rrca
9623  C0                 ret nz
9624  FE 00              cp 0
9626  00                 nop
9627  77                 ld (hl),a
9628  07                 rlca
9629  E0                 ret po
962A  FF                 rst 38h
962B  00                 nop
962C  00                 nop
962D  03                 inc bc
962E  0F                 rrca
962F  C0                 ret nz
9630  FF                 rst 38h
9631  00                 nop
9632  FF                 rst 38h
9633  00                 nop
9634  FF                 rst 38h
9635  00                 nop
9636  FC 00 7F           call m,$7F00
9639  00                 nop
963A  87                 add a,a
963B  00                 nop
963C  F8                 ret m
963D  00                 nop
963E  03                 inc bc
963F  00                 nop
9640  03                 inc bc
9641  00                 nop
9642  F0                 ret p
9643  03                 inc bc
9644  00                 nop
9645  00                 nop
9646  01 78 E0           ld bc,$E078
9649  07                 rlca
964A  00                 nop
964B  7C                 ld a,h
964C  00                 nop
964D  CC E0 06           call z,$06E0
9650  00                 nop
9651  FF                 rst 38h
9652  00                 nop
9653  B4                 or h
9654  C0                 ret nz
9655  0E 00              ld c,0
9657  79                 ld a,c
9658  00                 nop
9659  BE                 cp (hl)
965A  C0                 ret nz
965B  0D                 dec c
965C  00                 nop
965D  B6                 or (hl)
965E  00                 nop
965F  AE                 xor (hl)
9660  C0                 ret nz
9661  0D                 dec c
9662  00                 nop
9663  31 00 1E           ld sp,$1E00
9666  C0                 ret nz
9667  0C                 inc c
9668  00                 nop
9669  B5                 or l
966A  00                 nop
966B  1E C0              ld e,-64
966D  0C                 inc c
966E  00                 nop
966F  31 00 0C           ld sp,$0C00
9672  E0                 ret po
9673  05                 dec b
9674  00                 nop
9675  B6                 or (hl)
9676  00                 nop
9677  BC                 cp h
9678  E0                 ret po
9679  06 00              ld b,0
967B  B6                 or (hl)
967C  00                 nop
967D  3C                 inc a
967E  F0                 ret p
967F  02                 ld (bc),a
9680  00                 nop
9681  30 00              jr nc,$9683
9683  5C                 ld e,h
9684  F8                 ret m
9685  01 00 35           ld bc,$3500
9688  01 78 F0           ld bc,$F078
968B  02                 ld (bc),a
968C  00                 nop
968D  B5                 or l
968E  03                 inc bc
968F  B0                 or b
9690  F0                 ret p
9691  02                 ld (bc),a
9692  00                 nop
9693  B6                 or (hl)
9694  01 C0 E0           ld bc,$E0C0
9697  05                 dec b
9698  00                 nop
9699  87                 add a,a
969A  01 78 E0           ld bc,$E078
969D  05                 dec b
969E  00                 nop
969F  FF                 rst 38h
96A0  01 B8 E0           ld bc,$E0B8
96A3  06 00              ld b,0
96A5  CF                 rst 08h
96A6  03                 inc bc
96A7  30 F0              jr nc,$9699
96A9  01 00 39           ld bc,$3900
96AC  07                 rlca
96AD  C0                 ret nz
96AE  F8                 ret m
96AF  01 00 C7           ld bc,$C700
96B2  0F                 rrca
96B3  C0                 ret nz
96B4  F8                 ret m
96B5  01 00 EF           ld bc,$EF00
96B8  0F                 rrca
96B9  C0                 ret nz
96BA  F8                 ret m
96BB  01 00 EF           ld bc,$EF00
96BE  0F                 rrca
96BF  C0                 ret nz
96C0  FF                 rst 38h
96C1  00                 nop
96C2  FF                 rst 38h
96C3  00                 nop
96C4  FF                 rst 38h
96C5  00                 nop
96C6  FF                 rst 38h
96C7  00                 nop
96C8  8F                 adc a,a
96C9  00                 nop
96CA  9F                 sbc a,a
96CB  00                 nop
96CC  FF                 rst 38h
96CD  00                 nop
96CE  02                 ld (bc),a
96CF  00                 nop
96D0  0F                 rrca
96D1  00                 nop
96D2  FE 00              cp 0
96D4  00                 nop
96D5  70                 ld (hl),b
96D6  07                 rlca
96D7  60                 ld h,b
96D8  FC 00 00           call m,$0000
96DB  CD 03 90           call $9003
96DE  FC 00 00           call m,$0000
96E1  BF                 cp a
96E2  01 78 F8           ld bc,-1928
96E5  01 00 79           ld bc,$7900
96E8  01 B8 F8           ld bc,-1864
96EB  01 00 76           ld bc,$7600
96EE  01 A8 FC           ld bc,-856
96F1  00                 nop
96F2  00                 nop
96F3  EA 01 98           jp pe,$9801
96F6  FC 00 00           call m,$0000
96F9  E2 01 98           jp po,$9801
96FC  FC 00 00           call m,$0000
96FF  CD 01 B8           call $B801
9702  F8                 ret m
9703  01 00 DE           ld bc,$DE00
9706  03                 inc bc
9707  80                 add a,b
9708  F8                 ret m
9709  01 00 DE           ld bc,$DE00
970C  01 38 F8           ld bc,-1992
970F  01 00 8D           ld bc,$8D00
9712  01 F8 F8           ld bc,-1800
9715  01 00 A3           ld bc,$A300
9718  01 F8 F8           ld bc,-1800
971B  01 00 BB           ld bc,$BB00
971E  03                 inc bc
971F  F0                 ret p
9720  F8                 ret m
9721  01 00 CB           ld bc,$CB00
9724  07                 rlca
9725  80                 add a,b
9726  FC 00 00           call m,$0000
9729  3C                 inc a
972A  0F                 rrca
972B  00                 nop
972C  FC 00 00           call m,$0000
972F  FF                 rst 38h
9730  07                 rlca
9731  80                 add a,b
9732  F8                 ret m
9733  01 00 FF           ld bc,-256
9736  07                 rlca
9737  A0                 and b
9738  F8                 ret m
9739  01 00 F8           ld bc,-2048
973C  07                 rlca
973D  60                 ld h,b
973E  FC 00 00           call m,$0000
9741  E0                 ret po
9742  0F                 rrca
9743  C0                 ret nz
9744  FE 00              cp 0
9746  06 00              ld b,0
9748  1F                 rra
9749  00                 nop
974A  FF                 rst 38h
974B  00                 nop
974C  FF                 rst 38h
974D  00                 nop
974E  FF                 rst 38h
974F  00                 nop
9750  FF                 rst 38h
9751  00                 nop
9752  FF                 rst 38h
9753  00                 nop
9754  FF                 rst 38h
9755  00                 nop
9756  FF                 rst 38h
9757  00                 nop
9758  00                 nop
9759  00                 nop
975A  7F                 ld a,a
975B  00                 nop
975C  FE 00              cp 0
975E  00                 nop
975F  E7                 rst 20h
9760  33                 inc sp
9761  00                 nop
9762  FC 00 00           call m,$0000
9765  FB                 ei
9766  01 80 F8           ld bc,-1920
9769  03                 inc bc
976A  00                 nop
976B  60                 ld h,b
976C  00                 nop
976D  CC F0 05           call z,$05F0
9770  00                 nop
9771  8F                 adc a,a
9772  00                 nop
9773  5A                 ld e,d
9774  F0                 ret p
9775  06 00              ld b,0
9777  DF                 rst 18h
9778  00                 nop
9779  B6                 or (hl)
977A  F0                 ret p
977B  06 00              ld b,0
977D  CF                 rst 08h
977E  00                 nop
977F  36 F0              ld (hl),-16
9781  06 00              ld b,0
9783  B6                 or (hl)
9784  00                 nop
9785  D6 F0              sub -16
9787  05                 dec b
9788  00                 nop
9789  A6                 and (hl)
978A  00                 nop
978B  5A                 ld e,d
978C  F0                 ret p
978D  05                 dec b
978E  00                 nop
978F  56                 ld d,(hl)
9790  00                 nop
9791  AA                 xor d
9792  F8                 ret m
9793  02                 ld (bc),a
9794  00                 nop
9795  46                 ld b,(hl)
9796  00                 nop
9797  24                 inc h
9798  F8                 ret m
9799  03                 inc bc
979A  00                 nop
979B  36 00              ld (hl),0
979D  CC E0 01           call z,$01E0
97A0  00                 nop
97A1  96                 sub (hl)
97A2  00                 nop
97A3  9C                 sbc a,h
97A4  80                 add a,b
97A5  15                 dec d
97A6  00                 nop
97A7  C6 01              add a,1
97A9  38 00              jr c,$97AB
97AB  20 00              jr nz,$97AD
97AD  D6 03              sub 3
97AF  B0                 or b
97B0  00                 nop
97B1  47                 ld b,a
97B2  00                 nop
97B3  36 01              ld (hl),1
97B5  C8                 ret z
97B6  00                 nop
97B7  09                 add hl,bc
97B8  00                 nop
97B9  56                 ld d,(hl)
97BA  01 F8 00           ld bc,248
97BD  50                 ld d,b
97BE  00                 nop
97BF  A9                 xor c
97C0  03                 inc bc
97C1  F0                 ret p
97C2  00                 nop
97C3  10 00              djnz $97C5
97C5  BF                 cp a
97C6  03                 inc bc
97C7  F0                 ret p
97C8  80                 add a,b
97C9  11 00 FF           ld de,-256
97CC  07                 rlca
97CD  E0                 ret po
97CE  C0                 ret nz
97CF  0F                 rrca
97D0  00                 nop
97D1  FF                 rst 38h
97D2  0F                 rrca
97D3  C0                 ret nz
97D4  E0                 ret po
97D5  03                 inc bc
97D6  00                 nop
97D7  E7                 rst 20h
97D8  1F                 rra
97D9  00                 nop
97DA  F0                 ret p
97DB  00                 nop
97DC  00                 nop
97DD  08                 ex af,af'
97DE  3F                 ccf
97DF  00                 nop
97E0  FF                 rst 38h
97E1  00                 nop
97E2  FF                 rst 38h
97E3  00                 nop
97E4  FF                 rst 38h
97E5  00                 nop
97E6  FF                 rst 38h
97E7  00                 nop
97E8  FF                 rst 38h
97E9  00                 nop
97EA  FF                 rst 38h
97EB  00                 nop
97EC  FF                 rst 38h
97ED  00                 nop
97EE  E7                 rst 20h
97EF  00                 nop
97F0  FF                 rst 38h
97F1  00                 nop
97F2  FF                 rst 38h
97F3  00                 nop
97F4  C3 00 FF           jp $FF00
97F7  00                 nop
97F8  FF                 rst 38h
97F9  00                 nop
97FA  81                 add a,c
97FB  18 FF              jr $97FC
97FD  00                 nop
97FE  FF                 rst 38h
97FF  00                 nop
9800  00                 nop
9801  18 FF              jr $9802
9803  00                 nop
9804  FC 00 00           call m,$0000
9807  18 3F              jr $9848
9809  00                 nop
980A  F0                 ret p
980B  00                 nop
980C  00                 nop
980D  DB 0F              in a,($0F)
980F  00                 nop
9810  E0                 ret po
9811  03                 inc bc
9812  00                 nop
9813  99                 sbc a,c
9814  07                 rlca
9815  C0                 ret nz
9816  C0                 ret nz
9817  0D                 dec c
9818  00                 nop
9819  24                 inc h
981A  03                 inc bc
981B  B0                 or b
981C  80                 add a,b
981D  12                 ld (de),a
981E  00                 nop
981F  7E                 ld a,(hl)
9820  01 48 00           ld bc,72
9823  24                 inc h
9824  00                 nop
9825  FF                 rst 38h
9826  00                 nop
9827  24                 inc h
9828  00                 nop
9829  4A                 ld c,d
982A  00                 nop
982B  FF                 rst 38h
982C  00                 nop
982D  52                 ld d,d
982E  00                 nop
982F  54                 ld d,h
9830  00                 nop
9831  FF                 rst 38h
9832  00                 nop
9833  AA                 xor d
9834  00                 nop
9835  69                 ld l,c
9836  00                 nop
9837  FF                 rst 38h
9838  00                 nop
9839  96                 sub (hl)
983A  00                 nop
983B  35                 dec (hl)
983C  00                 nop
983D  FF                 rst 38h
983E  00                 nop
983F  AC                 xor h
9840  80                 add a,b
9841  0C                 inc c
9842  00                 nop
9843  FF                 rst 38h
9844  01 38 C0           ld bc,$C038
9847  03                 inc bc
9848  00                 nop
9849  FF                 rst 38h
984A  03                 inc bc
984B  E0                 ret po
984C  F0                 ret p
984D  00                 nop
984E  00                 nop
984F  00                 nop
9850  07                 rlca
9851  00                 nop
9852  FC 00 00           call m,$0000
9855  18 1F              jr $9876
9857  00                 nop
9858  FF                 rst 38h
9859  00                 nop
985A  81                 add a,c
985B  18 FF              jr $985C
985D  00                 nop
985E  FF                 rst 38h
985F  00                 nop
9860  81                 add a,c
9861  18 FF              jr $9862
9863  00                 nop
9864  FF                 rst 38h
9865  00                 nop
9866  81                 add a,c
9867  18 FF              jr $9868
9869  00                 nop
986A  FF                 rst 38h
986B  00                 nop
986C  81                 add a,c
986D  18 FF              jr $986E
986F  00                 nop
9870  FF                 rst 38h
9871  00                 nop
9872  FF                 rst 38h
9873  00                 nop
9874  FF                 rst 38h
9875  00                 nop
9876  FF                 rst 38h
9877  00                 nop
9878  FF                 rst 38h
9879  00                 nop
987A  FF                 rst 38h
987B  00                 nop
987C  FF                 rst 38h
987D  00                 nop
987E  FF                 rst 38h
987F  00                 nop
9880  FF                 rst 38h
9881  00                 nop
9882  FF                 rst 38h
9883  00                 nop
9884  FF                 rst 38h
9885  00                 nop
9886  FF                 rst 38h
9887  00                 nop
9888  FF                 rst 38h
9889  00                 nop
988A  FF                 rst 38h
988B  00                 nop
988C  FF                 rst 38h
988D  00                 nop
988E  FF                 rst 38h
988F  00                 nop
9890  FF                 rst 38h
9891  00                 nop
9892  FF                 rst 38h
9893  00                 nop
9894  FF                 rst 38h
9895  00                 nop
9896  FF                 rst 38h
9897  00                 nop
9898  FF                 rst 38h
9899  00                 nop
989A  FF                 rst 38h
989B  00                 nop
989C  FF                 rst 38h
989D  00                 nop
989E  FF                 rst 38h
989F  00                 nop
98A0  F1                 pop af
98A1  00                 nop
98A2  FF                 rst 38h
98A3  00                 nop
98A4  FF                 rst 38h
98A5  00                 nop
98A6  E0                 ret po
98A7  0E FF              ld c,-1
98A9  00                 nop
98AA  FF                 rst 38h
98AB  00                 nop
98AC  C0                 ret nz
98AD  1F                 rra
98AE  7F                 ld a,a
98AF  00                 nop
98B0  FF                 rst 38h
98B1  00                 nop
98B2  C0                 ret nz
98B3  13                 inc de
98B4  7F                 ld a,a
98B5  00                 nop
98B6  FF                 rst 38h
98B7  00                 nop
98B8  C0                 ret nz
98B9  0E 1F              ld c,31
98BB  80                 add a,b
98BC  FF                 rst 38h
98BD  00                 nop
98BE  80                 add a,b
98BF  31 03 E0           ld sp,$E003
98C2  FF                 rst 38h
98C3  00                 nop
98C4  00                 nop
98C5  7F                 ld a,a
98C6  01 FC FF           ld bc,-4
98C9  00                 nop
98CA  00                 nop
98CB  73                 ld (hl),e
98CC  00                 nop
98CD  CE FF              adc a,-1
98CF  00                 nop
98D0  80                 add a,b
98D1  2F                 cpl
98D2  00                 nop
98D3  BF                 cp a
98D4  7F                 ld a,a
98D5  00                 nop
98D6  C0                 ret nz
98D7  1D                 dec e
98D8  00                 nop
98D9  7F                 ld a,a
98DA  3F                 ccf
98DB  80                 add a,b
98DC  E0                 ret po
98DD  0A                 ld a,(bc)
98DE  00                 nop
98DF  FF                 rst 38h
98E0  3F                 ccf
98E1  80                 add a,b
98E2  F0                 ret p
98E3  06 00              ld b,0
98E5  FF                 rst 38h
98E6  3F                 ccf
98E7  80                 add a,b
98E8  F0                 ret p
98E9  05                 dec b
98EA  00                 nop
98EB  6F                 ld l,a
98EC  7F                 ld a,a
98ED  00                 nop
98EE  F0                 ret p
98EF  04                 inc b
98F0  00                 nop
98F1  90                 sub b
98F2  FF                 rst 38h
98F3  00                 nop
98F4  E0                 ret po
98F5  0E 00              ld c,0
98F7  0F                 rrca
98F8  7F                 ld a,a
98F9  00                 nop
98FA  C0                 ret nz
98FB  1A                 ld a,(de)
98FC  00                 nop
98FD  7C                 ld a,h
98FE  3F                 ccf
98FF  80                 add a,b
9900  FF                 rst 38h
9901  00                 nop
9902  FF                 rst 38h
9903  00                 nop
9904  FF                 rst 38h
9905  00                 nop
9906  FF                 rst 38h
9907  00                 nop
9908  FF                 rst 38h
9909  00                 nop
990A  FF                 rst 38h
990B  00                 nop
990C  FF                 rst 38h
990D  00                 nop
990E  FF                 rst 38h
990F  00                 nop
9910  FF                 rst 38h
9911  00                 nop
9912  F1                 pop af
9913  00                 nop
9914  FF                 rst 38h
9915  00                 nop
9916  FF                 rst 38h
9917  00                 nop
9918  E0                 ret po
9919  0E FF              ld c,-1
991B  00                 nop
991C  FF                 rst 38h
991D  00                 nop
991E  C0                 ret nz
991F  1F                 rra
9920  3F                 ccf
9921  00                 nop
9922  FF                 rst 38h
9923  00                 nop
9924  C0                 ret nz
9925  13                 inc de
9926  1F                 rra
9927  40                 ld b,b
9928  FF                 rst 38h
9929  00                 nop
992A  E0                 ret po
992B  0E 0F              ld c,15
992D  E0                 ret po
992E  FF                 rst 38h
992F  00                 nop
9930  80                 add a,b
9931  11 0F E0           ld de,$E00F
9934  FF                 rst 38h
9935  00                 nop
9936  00                 nop
9937  5F                 ld e,a
9938  07                 rlca
9939  F0                 ret p
993A  FF                 rst 38h
993B  00                 nop
993C  80                 add a,b
993D  19                 add hl,de
993E  03                 inc bc
993F  F8                 ret m
9940  FF                 rst 38h
9941  00                 nop
9942  80                 add a,b
9943  27                 daa
9944  01 E4 FF           ld bc,-28
9947  00                 nop
9948  00                 nop
9949  5D                 ld e,l
994A  01 F8 FF           ld bc,-8
994D  00                 nop
994E  00                 nop
994F  4B                 ld c,e
9950  00                 nop
9951  7C                 ld a,h
9952  FF                 rst 38h
9953  00                 nop
9954  B0                 or b
9955  03                 inc bc
9956  00                 nop
9957  7D                 ld a,l
9958  7F                 ld a,a
9959  00                 nop
995A  FC 01 00           call m,$0001
995D  BE                 cp (hl)
995E  1F                 rra
995F  80                 add a,b
9960  FE 00              cp 0
9962  40                 ld b,b
9963  0F                 rrca
9964  07                 rlca
9965  20 FF              jr nz,$9966
9967  00                 nop
9968  F0                 ret p
9969  01 03 C8           ld bc,$C803
996C  FF                 rst 38h
996D  00                 nop
996E  FE 00              cp 0
9970  01 74 FF           ld bc,-140
9973  00                 nop
9974  FF                 rst 38h
9975  00                 nop
9976  81                 add a,c
9977  18 FF              jr $9978
9979  00                 nop
997A  FF                 rst 38h
997B  00                 nop
997C  E1                 pop hl
997D  04                 inc b
997E  FF                 rst 38h
997F  00                 nop
9980  FF                 rst 38h
9981  00                 nop
9982  C1                 pop bc
9983  14                 inc d
9984  FF                 rst 38h
9985  00                 nop
9986  FF                 rst 38h
9987  00                 nop
9988  C0                 ret nz
9989  1A                 ld a,(de)
998A  FF                 rst 38h
998B  00                 nop
998C  FF                 rst 38h
998D  00                 nop
998E  E5                 push hl
998F  00                 nop
9990  FF                 rst 38h
9991  00                 nop
9992  FF                 rst 38h
9993  00                 nop
9994  FF                 rst 38h
9995  00                 nop
9996  FF                 rst 38h
9997  00                 nop
9998  FF                 rst 38h
9999  00                 nop
999A  FF                 rst 38h
999B  00                 nop
999C  FF                 rst 38h
999D  00                 nop
999E  FF                 rst 38h
999F  00                 nop
99A0  FF                 rst 38h
99A1  00                 nop
99A2  FF                 rst 38h
99A3  00                 nop
99A4  F0                 ret p
99A5  00                 nop
99A6  7F                 ld a,a
99A7  00                 nop
99A8  FF                 rst 38h
99A9  00                 nop
99AA  C0                 ret nz
99AB  0A                 ld a,(bc)
99AC  1F                 rra
99AD  80                 add a,b
99AE  FF                 rst 38h
99AF  00                 nop
99B0  00                 nop
99B1  20 0F              jr nz,$99C2
99B3  60                 ld h,b
99B4  FE 00              cp 0
99B6  00                 nop
99B7  8A                 adc a,d
99B8  07                 rlca
99B9  10 FE              djnz $99B9
99BB  00                 nop
99BC  00                 nop
99BD  25                 dec h
99BE  03                 inc bc
99BF  28 8C              jr z,$994D
99C1  01 00 18           ld bc,$1800
99C4  03                 inc bc
99C5  18 04              jr $99CB
99C7  70                 ld (hl),b
99C8  00                 nop
99C9  50                 ld d,b
99CA  01 8C 00           ld bc,140
99CD  FA 00 20           jp m,$2000
99D0  01 54 00           ld bc,84
99D3  98                 sbc a,b
99D4  00                 nop
99D5  60                 ld h,b
99D6  01 CC 00           ld bc,204
99D9  9A                 sbc a,d
99DA  00                 nop
99DB  25                 dec h
99DC  01 44 00           ld bc,68
99DF  74                 ld (hl),h
99E0  00                 nop
99E1  56                 ld d,(hl)
99E2  01 CC 80           ld bc,$80CC
99E5  0D                 dec c
99E6  00                 nop
99E7  11 03 18           ld de,$1803
99EA  C0                 ret nz
99EB  1C                 inc e
99EC  00                 nop
99ED  28 03              jr z,$99F2
99EF  28 E0              jr z,$99D1
99F1  0D                 dec c
99F2  00                 nop
99F3  06 07              ld b,7
99F5  90                 sub b
99F6  80                 add a,b
99F7  16 00              ld d,0
99F9  C3 0F E0           jp $E00F
99FC  00                 nop
99FD  77                 ld (hl),a
99FE  00                 nop
99FF  60                 ld h,b
9A00  1F                 rra
9A01  80                 add a,b
9A02  80                 add a,b
9A03  37                 scf
9A04  00                 nop
9A05  F8                 ret m
9A06  0F                 rrca
9A07  60                 ld h,b
9A08  80                 add a,b
9A09  3F                 ccf
9A0A  00                 nop
9A0B  FE 07              cp 7
9A0D  10 C0              djnz $99CF
9A0F  1F                 rra
9A10  00                 nop
9A11  FF                 rst 38h
9A12  0F                 rrca
9A13  80                 add a,b
9A14  C0                 ret nz
9A15  1F                 rra
9A16  00                 nop
9A17  FF                 rst 38h
9A18  03                 inc bc
9A19  F0                 ret p
9A1A  E0                 ret po
9A1B  0A                 ld a,(bc)
9A1C  00                 nop
9A1D  AA                 xor d
9A1E  01 AC FF           ld bc,-84
9A21  00                 nop
9A22  80                 add a,b
9A23  00                 nop
9A24  7F                 ld a,a
9A25  00                 nop
9A26  FF                 rst 38h
9A27  00                 nop
9A28  00                 nop
9A29  7F                 ld a,a
9A2A  3F                 ccf
9A2B  80                 add a,b
9A2C  FF                 rst 38h
9A2D  00                 nop
9A2E  00                 nop
9A2F  5F                 ld e,a
9A30  3F                 ccf
9A31  80                 add a,b
9A32  FF                 rst 38h
9A33  00                 nop
9A34  00                 nop
9A35  50                 ld d,b
9A36  3F                 ccf
9A37  00                 nop
9A38  FC 00 00           call m,$0000
9A3B  00                 nop
9A3C  3F                 ccf
9A3D  00                 nop
9A3E  F8                 ret m
9A3F  03                 inc bc
9A40  00                 nop
9A41  FF                 rst 38h
9A42  3F                 ccf
9A43  80                 add a,b
9A44  FC 00 00           call m,$0000
9A47  00                 nop
9A48  3F                 ccf
9A49  00                 nop
9A4A  FE 00              cp 0
9A4C  00                 nop
9A4D  98                 sbc a,b
9A4E  1F                 rra
9A4F  00                 nop
9A50  FE 00              cp 0
9A52  00                 nop
9A53  45                 ld b,l
9A54  1F                 rra
9A55  00                 nop
9A56  FE 00              cp 0
9A58  00                 nop
9A59  E5                 push hl
9A5A  0F                 rrca
9A5B  00                 nop
9A5C  FC 00 00           call m,$0000
9A5F  03                 inc bc
9A60  07                 rlca
9A61  00                 nop
9A62  FC 00 00           call m,$0000
9A65  0E 0F              ld c,15
9A67  00                 nop
9A68  F8                 ret m
9A69  00                 nop
9A6A  00                 nop
9A6B  67                 ld h,a
9A6C  1F                 rra
9A6D  00                 nop
9A6E  F8                 ret m
9A6F  01 00 3E           ld bc,$3E00
9A72  3F                 ccf
9A73  80                 add a,b
9A74  F8                 ret m
9A75  02                 ld (bc),a
9A76  00                 nop
9A77  01 1F C0           ld bc,$C01F
9A7A  F8                 ret m
9A7B  01 00 11           ld bc,$1100
9A7E  0F                 rrca
9A7F  E0                 ret po
9A80  F8                 ret m
9A81  01 00 B6           ld bc,$B600
9A84  0F                 rrca
9A85  E0                 ret po
9A86  C8                 ret z
9A87  01 00 B6           ld bc,$B600
9A8A  0F                 rrca
9A8B  A0                 and b
9A8C  84                 add a,h
9A8D  30 00              jr nc,$9A8F
9A8F  00                 nop
9A90  03                 inc bc
9A91  58                 ld e,b
9A92  00                 nop
9A93  7A                 ld a,d
9A94  00                 nop
9A95  BF                 cp a
9A96  03                 inc bc
9A97  D8                 ret c
9A98  00                 nop
9A99  7D                 ld a,l
9A9A  00                 nop
9A9B  1F                 rra
9A9C  03                 inc bc
9A9D  38 80              jr c,$9A1F
9A9F  3C                 inc a
9AA0  40                 ld b,b
9AA1  1C                 inc e
9AA2  03                 inc bc
9AA3  D8                 ret c
9AA4  C0                 ret nz
9AA5  1F                 rra
9AA6  20 81              jr nz,$9A29
9AA8  07                 rlca
9AA9  F0                 ret p
9AAA  E0                 ret po
9AAB  07                 rlca
9AAC  3C                 inc a
9AAD  81                 add a,c
9AAE  0F                 rrca
9AAF  E0                 ret po
9AB0  FF                 rst 38h
9AB1  00                 nop
9AB2  80                 add a,b
9AB3  00                 nop
9AB4  7F                 ld a,a
9AB5  00                 nop
9AB6  FF                 rst 38h
9AB7  00                 nop
9AB8  00                 nop
9AB9  7F                 ld a,a
9ABA  3F                 ccf
9ABB  80                 add a,b
9ABC  FF                 rst 38h
9ABD  00                 nop
9ABE  00                 nop
9ABF  5F                 ld e,a
9AC0  3F                 ccf
9AC1  80                 add a,b
9AC2  FF                 rst 38h
9AC3  00                 nop
9AC4  00                 nop
9AC5  50                 ld d,b
9AC6  3F                 ccf
9AC7  00                 nop
9AC8  FC 00 00           call m,$0000
9ACB  00                 nop
9ACC  3F                 ccf
9ACD  00                 nop
9ACE  F8                 ret m
9ACF  03                 inc bc
9AD0  00                 nop
9AD1  FF                 rst 38h
9AD2  3F                 ccf
9AD3  80                 add a,b
9AD4  FC 00 00           call m,$0000
9AD7  00                 nop
9AD8  3F                 ccf
9AD9  00                 nop
9ADA  FE 00              cp 0
9ADC  00                 nop
9ADD  98                 sbc a,b
9ADE  1F                 rra
9ADF  00                 nop
9AE0  FE 00              cp 0
9AE2  00                 nop
9AE3  45                 ld b,l
9AE4  1F                 rra
9AE5  00                 nop
9AE6  FE 00              cp 0
9AE8  00                 nop
9AE9  E5                 push hl
9AEA  0F                 rrca
9AEB  00                 nop
9AEC  FC 00 00           call m,$0000
9AEF  03                 inc bc
9AF0  07                 rlca
9AF1  00                 nop
9AF2  FC 00 00           call m,$0000
9AF5  0E 0F              ld c,15
9AF7  00                 nop
9AF8  F8                 ret m
9AF9  00                 nop
9AFA  00                 nop
9AFB  67                 ld h,a
9AFC  1F                 rra
9AFD  00                 nop
9AFE  F8                 ret m
9AFF  01 00 3E           ld bc,$3E00
9B02  3F                 ccf
9B03  80                 add a,b
9B04  F8                 ret m
9B05  02                 ld (bc),a
9B06  00                 nop
9B07  00                 nop
9B08  1F                 rra
9B09  C0                 ret nz
9B0A  F0                 ret p
9B0B  00                 nop
9B0C  00                 nop
9B0D  88                 adc a,b
9B0E  0F                 rrca
9B0F  E0                 ret po
9B10  F0                 ret p
9B11  04                 inc b
9B12  00                 nop
9B13  DB 0F              in a,($0F)
9B15  60                 ld h,b
9B16  F0                 ret p
9B17  02                 ld (bc),a
9B18  00                 nop
9B19  DB 07              in a,($07)
9B1B  70                 ld (hl),b
9B1C  F0                 ret p
9B1D  04                 inc b
9B1E  00                 nop
9B1F  00                 nop
9B20  07                 rlca
9B21  70                 ld (hl),b
9B22  F8                 ret m
9B23  02                 ld (bc),a
9B24  00                 nop
9B25  77                 ld (hl),a
9B26  0F                 rrca
9B27  60                 ld h,b
9B28  F8                 ret m
9B29  00                 nop
9B2A  00                 nop
9B2B  70                 ld (hl),b
9B2C  1F                 rra
9B2D  40                 ld b,b
9B2E  F0                 ret p
9B2F  07                 rlca
9B30  00                 nop
9B31  07                 rlca
9B32  3F                 ccf
9B33  80                 add a,b
9B34  E0                 ret po
9B35  0F                 rrca
9B36  00                 nop
9B37  EF                 rst 28h
9B38  1F                 rra
9B39  C0                 ret nz
9B3A  E0                 ret po
9B3B  0F                 rrca
9B3C  00                 nop
9B3D  EF                 rst 28h
9B3E  1F                 rra
9B3F  C0                 ret nz
9B40  FF                 rst 38h
9B41  00                 nop
9B42  80                 add a,b
9B43  00                 nop
9B44  7F                 ld a,a
9B45  00                 nop
9B46  FF                 rst 38h
9B47  00                 nop
9B48  00                 nop
9B49  7F                 ld a,a
9B4A  3F                 ccf
9B4B  80                 add a,b
9B4C  FF                 rst 38h
9B4D  00                 nop
9B4E  00                 nop
9B4F  5F                 ld e,a
9B50  3F                 ccf
9B51  80                 add a,b
9B52  FF                 rst 38h
9B53  00                 nop
9B54  00                 nop
9B55  50                 ld d,b
9B56  3F                 ccf
9B57  00                 nop
9B58  FC 00 00           call m,$0000
9B5B  00                 nop
9B5C  3F                 ccf
9B5D  00                 nop
9B5E  F8                 ret m
9B5F  03                 inc bc
9B60  00                 nop
9B61  FF                 rst 38h
9B62  3F                 ccf
9B63  80                 add a,b
9B64  FC 00 00           call m,$0000
9B67  00                 nop
9B68  3F                 ccf
9B69  00                 nop
9B6A  FE 00              cp 0
9B6C  00                 nop
9B6D  98                 sbc a,b
9B6E  1F                 rra
9B6F  00                 nop
9B70  FE 00              cp 0
9B72  00                 nop
9B73  45                 ld b,l
9B74  1F                 rra
9B75  00                 nop
9B76  FE 00              cp 0
9B78  00                 nop
9B79  E5                 push hl
9B7A  0F                 rrca
9B7B  00                 nop
9B7C  FC 00 00           call m,$0000
9B7F  03                 inc bc
9B80  07                 rlca
9B81  00                 nop
9B82  FC 00 00           call m,$0000
9B85  0E 0F              ld c,15
9B87  00                 nop
9B88  F8                 ret m
9B89  00                 nop
9B8A  00                 nop
9B8B  67                 ld h,a
9B8C  1F                 rra
9B8D  00                 nop
9B8E  F8                 ret m
9B8F  01 00 3E           ld bc,$3E00
9B92  1F                 rra
9B93  C0                 ret nz
9B94  F8                 ret m
9B95  02                 ld (bc),a
9B96  00                 nop
9B97  00                 nop
9B98  0F                 rrca
9B99  E0                 ret po
9B9A  FC 00 00           call m,$0000
9B9D  04                 inc b
9B9E  0F                 rrca
9B9F  E0                 ret po
9BA0  FC 01 00           call m,$0001
9BA3  CD 07 70           call $7007
9BA6  F8                 ret m
9BA7  00                 nop
9BA8  00                 nop
9BA9  ED 07              ???
9BAB  70                 ld (hl),b
9BAC  F8                 ret m
9BAD  03                 inc bc
9BAE  00                 nop
9BAF  E0                 ret po
9BB0  07                 rlca
9BB1  70                 ld (hl),b
9BB2  F8                 ret m
9BB3  00                 nop
9BB4  00                 nop
9BB5  6F                 ld l,a
9BB6  07                 rlca
9BB7  30 F8              jr nc,$9BB1
9BB9  03                 inc bc
9BBA  00                 nop
9BBB  EF                 rst 28h
9BBC  0F                 rrca
9BBD  20 F8              jr nz,$9BB7
9BBF  00                 nop
9BC0  00                 nop
9BC1  72                 ld (hl),d
9BC2  DF                 rst 18h
9BC3  00                 nop
9BC4  FC 01 01           call m,$0101
9BC7  38 FF              jr c,$9BC8
9BC9  00                 nop
9BCA  FE 00              cp 0
9BCC  07                 rlca
9BCD  E0                 ret po
9BCE  FF                 rst 38h
9BCF  00                 nop
9BD0  F9                 ld sp,hl
9BD1  00                 nop
9BD2  47                 ld b,a
9BD3  00                 nop
9BD4  FF                 rst 38h
9BD5  00                 nop
9BD6  F0                 ret p
9BD7  06 03              ld b,3
9BD9  B8                 cp b
9BDA  3F                 ccf
9BDB  00                 nop
9BDC  F0                 ret p
9BDD  05                 dec b
9BDE  02                 ld (bc),a
9BDF  B8                 cp b
9BE0  1F                 rra
9BE1  C0                 ret nz
9BE2  F8                 ret m
9BE3  00                 nop
9BE4  00                 nop
9BE5  19                 add hl,de
9BE6  0F                 rrca
9BE7  E0                 ret po
9BE8  F8                 ret m
9BE9  02                 ld (bc),a
9BEA  04                 inc b
9BEB  E1                 pop hl
9BEC  07                 rlca
9BED  E0                 ret po
9BEE  80                 add a,b
9BEF  07                 rlca
9BF0  06 F0              ld b,-16
9BF2  03                 inc bc
9BF3  C8                 ret z
9BF4  00                 nop
9BF5  7A                 ld a,d
9BF6  03                 inc bc
9BF7  F8                 ret m
9BF8  31 04 00           ld sp,$0004
9BFB  1E 07              ld e,7
9BFD  F0                 ret p
9BFE  F1                 pop af
9BFF  04                 inc b
9C00  00                 nop
9C01  3E 03              ld a,3
9C03  68                 ld l,b
9C04  88                 adc a,b
9C05  02                 ld (bc),a
9C06  00                 nop
9C07  7F                 ld a,a
9C08  00                 nop
9C09  9C                 sbc a,h
9C0A  00                 nop
9C0B  72                 ld (hl),d
9C0C  00                 nop
9C0D  7F                 ld a,a
9C0E  00                 nop
9C0F  3D                 dec a
9C10  00                 nop
9C11  FB                 ei
9C12  80                 add a,b
9C13  3C                 inc a
9C14  00                 nop
9C15  7D                 ld a,l
9C16  00                 nop
9C17  FB                 ei
9C18  C0                 ret nz
9C19  01 00 FD           ld bc,-768
9C1C  00                 nop
9C1D  FD C0              ret nz
9C1F  03                 inc bc
9C20  00                 nop
9C21  FD 00              nop
9C23  FC 80 29           call m,$2980
9C26  00                 nop
9C27  FD 03              inc bc
9C29  E0                 ret po
9C2A  00                 nop
9C2B  54                 ld d,h
9C2C  00                 nop
9C2D  7D                 ld a,l
9C2E  01 DC 80           ld bc,$80DC
9C31  2A 00 9A           ld hl,($9A00)
9C34  00                 nop
9C35  3E C0              ld a,-64
9C37  01 04 D1           ld bc,$D104
9C3A  00                 nop
9C3B  3E F8              ld a,-8
9C3D  03                 inc bc
9C3E  0E C0              ld c,-64
9C40  00                 nop
9C41  BE                 cp (hl)
9C42  E0                 ret po
9C43  07                 rlca
9C44  3F                 ccf
9C45  80                 add a,b
9C46  00                 nop
9C47  5E                 ld e,(hl)
9C48  C0                 ret nz
9C49  1F                 rra
9C4A  7F                 ld a,a
9C4B  00                 nop
9C4C  81                 add a,c
9C4D  0C                 inc c
9C4E  C0                 ret nz
9C4F  1F                 rra
9C50  7F                 ld a,a
9C51  00                 nop
9C52  80                 add a,b
9C53  32 E0 0E           ld ($0EE0),a
9C56  FF                 rst 38h
9C57  00                 nop
9C58  01 7C F1           ld bc,$F17C
9C5B  00                 nop
9C5C  FF                 rst 38h
9C5D  00                 nop
9C5E  03                 inc bc
9C5F  78                 ld a,b
9C60  FF                 rst 38h
9C61  00                 nop
9C62  FF                 rst 38h
9C63  00                 nop
9C64  FF                 rst 38h
9C65  00                 nop
9C66  FF                 rst 38h
9C67  00                 nop
9C68  FF                 rst 38h
9C69  00                 nop
9C6A  FF                 rst 38h
9C6B  00                 nop
9C6C  F9                 ld sp,hl
9C6D  00                 nop
9C6E  47                 ld b,a
9C6F  00                 nop
9C70  FF                 rst 38h
9C71  00                 nop
9C72  F0                 ret p
9C73  06 03              ld b,3
9C75  B8                 cp b
9C76  F9                 ld sp,hl
9C77  00                 nop
9C78  F0                 ret p
9C79  05                 dec b
9C7A  03                 inc bc
9C7B  B8                 cp b
9C7C  F0                 ret p
9C7D  06 F8              ld b,-8
9C7F  00                 nop
9C80  03                 inc bc
9C81  18 E0              jr $9C63
9C83  0F                 rrca
9C84  F8                 ret m
9C85  02                 ld (bc),a
9C86  07                 rlca
9C87  E0                 ret po
9C88  E0                 ret po
9C89  0F                 rrca
9C8A  80                 add a,b
9C8B  07                 rlca
9C8C  07                 rlca
9C8D  F0                 ret p
9C8E  F0                 ret p
9C8F  06 00              ld b,0
9C91  7A                 ld a,d
9C92  03                 inc bc
9C93  F8                 ret m
9C94  80                 add a,b
9C95  01 00 1E           ld bc,$1E00
9C98  00                 nop
9C99  F4 00 71           call p,$7100
9C9C  00                 nop
9C9D  3E 00              ld a,0
9C9F  6D                 ld l,l
9CA0  00                 nop
9CA1  F9                 ld sp,hl
9CA2  00                 nop
9CA3  7F                 ld a,a
9CA4  00                 nop
9CA5  9D                 sbc a,l
9CA6  00                 nop
9CA7  F9                 ld sp,hl
9CA8  80                 add a,b
9CA9  3F                 ccf
9CAA  00                 nop
9CAB  3C                 inc a
9CAC  00                 nop
9CAD  3D                 dec a
9CAE  80                 add a,b
9CAF  3C                 inc a
9CB0  00                 nop
9CB1  79                 ld a,c
9CB2  00                 nop
9CB3  BD                 cp l
9CB4  C0                 ret nz
9CB5  01 00 F3           ld bc,$F300
9CB8  00                 nop
9CB9  FA FE 00           jp m,$00FE
9CBC  00                 nop
9CBD  77                 ld (hl),a
9CBE  00                 nop
9CBF  FA FF 00           jp m,$00FF
9CC2  80                 add a,b
9CC3  17                 rla
9CC4  01 F4 FF           ld bc,-12
9CC7  00                 nop
9CC8  00                 nop
9CC9  17                 rla
9CCA  03                 inc bc
9CCB  80                 add a,b
9CCC  FE 00              cp 0
9CCE  00                 nop
9CCF  C3 03 08           jp $0803
9CD2  FC 01 00           call m,$0001
9CD5  08                 ex af,af'
9CD6  03                 inc bc
9CD7  68                 ld l,b
9CD8  FE 00              cp 0
9CDA  00                 nop
9CDB  7E                 ld a,(hl)
9CDC  07                 rlca
9CDD  E0                 ret po
9CDE  FE 00              cp 0
9CE0  01 FC 1F           ld bc,$1FFC
9CE3  00                 nop
9CE4  FF                 rst 38h
9CE5  00                 nop
9CE6  03                 inc bc
9CE7  78                 ld a,b
9CE8  FF                 rst 38h
9CE9  00                 nop
9CEA  FF                 rst 38h
9CEB  00                 nop
9CEC  87                 add a,a
9CED  00                 nop
9CEE  FF                 rst 38h
9CEF  00                 nop
9CF0  FF                 rst 38h
9CF1  00                 nop
9CF2  C1                 pop bc
9CF3  00                 nop
9CF4  FF                 rst 38h
9CF5  00                 nop
9CF6  E1                 pop hl
9CF7  00                 nop
9CF8  00                 nop
9CF9  3E FF              ld a,-1
9CFB  00                 nop
9CFC  C0                 ret nz
9CFD  1E 00              ld e,0
9CFF  FF                 rst 38h
9D00  67                 ld h,a
9D01  00                 nop
9D02  80                 add a,b
9D03  3F                 ccf
9D04  00                 nop
9D05  3F                 ccf
9D06  03                 inc bc
9D07  98                 sbc a,b
9D08  80                 add a,b
9D09  3F                 ccf
9D0A  00                 nop
9D0B  DF                 rst 18h
9D0C  01 48 80           ld bc,$8048
9D0F  33                 inc sp
9D10  00                 nop
9D11  F9                 ld sp,hl
9D12  00                 nop
9D13  E6 00              and $00
9D15  79                 ld a,c
9D16  00                 nop
9D17  F3                 di
9D18  00                 nop
9D19  F6 00              or $00
9D1B  71                 ld (hl),c
9D1C  00                 nop
9D1D  F1                 pop af
9D1E  00                 nop
9D1F  F2 00 65           jp p,$6500
9D22  00                 nop
9D23  B4                 or h
9D24  01 E8 00           ld bc,232
9D27  4E                 ld c,(hl)
9D28  00                 nop
9D29  4E                 ld c,(hl)
9D2A  01 6C 00           ld bc,108
9D2D  4F                 ld c,a
9D2E  00                 nop
9D2F  BE                 cp (hl)
9D30  03                 inc bc
9D31  60                 ld h,b
9D32  00                 nop
9D33  49                 ld c,c
9D34  00                 nop
9D35  8E                 adc a,(hl)
9D36  03                 inc bc
9D37  48                 ld c,b
9D38  00                 nop
9D39  61                 ld h,c
9D3A  00                 nop
9D3B  0C                 inc c
9D3C  01 5C 80           ld bc,$805C
9D3F  30 00              jr nc,$9D41
9D41  03                 inc bc
9D42  01 CC 00           ld bc,204
9D45  77                 ld (hl),a
9D46  00                 nop
9D47  F7                 rst 30h
9D48  03                 inc bc
9D49  E8                 ret pe
9D4A  00                 nop
9D4B  6F                 ld l,a
9D4C  00                 nop
9D4D  FB                 ei
9D4E  07                 rlca
9D4F  40                 ld b,b
9D50  00                 nop
9D51  6F                 ld l,a
9D52  00                 nop
9D53  FA 03 B8           jp m,$B803
9D56  80                 add a,b
9D57  3F                 ccf
9D58  00                 nop
9D59  FD 03              inc bc
9D5B  B8                 cp b
9D5C  C0                 ret nz
9D5D  1E 00              ld e,0
9D5F  FD 03              inc bc
9D61  B8                 cp b
9D62  E0                 ret po
9D63  0C                 inc c
9D64  00                 nop
9D65  7C                 ld a,h
9D66  45                 ld b,l
9D67  00                 nop
9D68  F0                 ret p
9D69  01 00 9E           ld bc,$9E00
9D6C  C8                 ret z
9D6D  02                 ld (bc),a
9D6E  FC 01 21           call m,$2101
9D71  80                 add a,b
9D72  85                 add a,l
9D73  30 FE              jr nc,$9D73
9D75  00                 nop
9D76  7F                 ld a,a
9D77  00                 nop
9D78  83                 add a,e
9D79  38 FF              jr c,$9D7A
9D7B  00                 nop
9D7C  FF                 rst 38h
9D7D  00                 nop
9D7E  C7                 rst 00h
9D7F  00                 nop
9D80  FF                 rst 38h
9D81  00                 nop
9D82  80                 add a,b
9D83  00                 nop
9D84  FF                 rst 38h
9D85  00                 nop
9D86  FF                 rst 38h
9D87  00                 nop
9D88  00                 nop
9D89  7D                 ld a,l
9D8A  7F                 ld a,a
9D8B  00                 nop
9D8C  FF                 rst 38h
9D8D  00                 nop
9D8E  00                 nop
9D8F  68                 ld l,b
9D90  7F                 ld a,a
9D91  00                 nop
9D92  FF                 rst 38h
9D93  00                 nop
9D94  00                 nop
9D95  69                 ld l,c
9D96  7F                 ld a,a
9D97  00                 nop
9D98  FE 00              cp 0
9D9A  00                 nop
9D9B  00                 nop
9D9C  3F                 ccf
9D9D  00                 nop
9D9E  FC 01 00           call m,$0001
9DA1  F5                 push af
9DA2  1F                 rra
9DA3  40                 ld b,b
9DA4  FE 00              cp 0
9DA6  00                 nop
9DA7  00                 nop
9DA8  3F                 ccf
9DA9  00                 nop
9DAA  FE 00              cp 0
9DAC  00                 nop
9DAD  B3                 or e
9DAE  3F                 ccf
9DAF  80                 add a,b
9DB0  FC 00 00           call m,$0000
9DB3  0F                 rrca
9DB4  3F                 ccf
9DB5  80                 add a,b
9DB6  FC 00 00           call m,$0000
9DB9  8F                 adc a,a
9DBA  3F                 ccf
9DBB  80                 add a,b
9DBC  F8                 ret m
9DBD  00                 nop
9DBE  00                 nop
9DBF  FD 3F              ccf
9DC1  00                 nop
9DC2  F0                 ret p
9DC3  06 00              ld b,0
9DC5  56                 ld d,(hl)
9DC6  1F                 rra
9DC7  C0                 ret nz
9DC8  E0                 ret po
9DC9  08                 ex af,af'
9DCA  00                 nop
9DCB  B9                 cp c
9DCC  0F                 rrca
9DCD  E0                 ret po
9DCE  C0                 ret nz
9DCF  11 00 C5           ld de,$C500
9DD2  07                 rlca
9DD3  F0                 ret p
9DD4  C0                 ret nz
9DD5  11 00 7E           ld de,$7E00
9DD8  03                 inc bc
9DD9  F8                 ret m
9DDA  E0                 ret po
9DDB  09                 add hl,bc
9DDC  00                 nop
9DDD  FF                 rst 38h
9DDE  03                 inc bc
9DDF  78                 ld a,b
9DE0  F6 00              or $00
9DE2  00                 nop
9DE3  FF                 rst 38h
9DE4  01 7C FE           ld bc,-388
9DE7  00                 nop
9DE8  00                 nop
9DE9  BF                 cp a
9DEA  01 BC FF           ld bc,-68
9DED  00                 nop
9DEE  00                 nop
9DEF  3F                 ccf
9DF0  03                 inc bc
9DF1  38 FF              jr c,$9DF2
9DF3  00                 nop
9DF4  00                 nop
9DF5  7E                 ld a,(hl)
9DF6  C7                 rst 00h
9DF7  00                 nop
9DF8  FF                 rst 38h
9DF9  00                 nop
9DFA  00                 nop
9DFB  7E                 ld a,(hl)
9DFC  FF                 rst 38h
9DFD  00                 nop
9DFE  FF                 rst 38h
9DFF  00                 nop
9E00  01 7C FF           ld bc,-132
9E03  00                 nop
9E04  FF                 rst 38h
9E05  00                 nop
9E06  83                 add a,e
9E07  38 FF              jr c,$9E08
9E09  00                 nop
9E0A  FF                 rst 38h
9E0B  00                 nop
9E0C  C7                 rst 00h
9E0D  00                 nop
9E0E  FF                 rst 38h
9E0F  00                 nop
9E10  FF                 rst 38h
9E11  00                 nop
9E12  01 00 FF           ld bc,-256
9E15  00                 nop
9E16  FE 00              cp 0
9E18  00                 nop
9E19  FA FF 00           jp m,$00FF
9E1C  FE 00              cp 0
9E1E  00                 nop
9E1F  D0                 ret nc
9E20  FF                 rst 38h
9E21  00                 nop
9E22  FE 00              cp 0
9E24  00                 nop
9E25  D2 FF 00           jp nc,$00FF
9E28  FC 00 00           call m,$0000
9E2B  00                 nop
9E2C  7F                 ld a,a
9E2D  00                 nop
9E2E  F8                 ret m
9E2F  03                 inc bc
9E30  00                 nop
9E31  EA 3F 80           jp pe,$803F
9E34  FC 00 00           call m,$0000
9E37  00                 nop
9E38  7F                 ld a,a
9E39  00                 nop
9E3A  FC 01 00           call m,$0001
9E3D  67                 ld h,a
9E3E  7F                 ld a,a
9E3F  00                 nop
9E40  F8                 ret m
9E41  00                 nop
9E42  00                 nop
9E43  1F                 rra
9E44  7F                 ld a,a
9E45  00                 nop
9E46  FC 01 00           call m,$0001
9E49  1F                 rra
9E4A  7F                 ld a,a
9E4B  00                 nop
9E4C  FC 01 00           call m,$0001
9E4F  FA FF 00           jp m,$00FF
9E52  FC 00 00           call m,$0000
9E55  AC                 xor h
9E56  7F                 ld a,a
9E57  00                 nop
9E58  F8                 ret m
9E59  02                 ld (bc),a
9E5A  00                 nop
9E5B  71                 ld (hl),c
9E5C  3F                 ccf
9E5D  80                 add a,b
9E5E  F8                 ret m
9E5F  03                 inc bc
9E60  00                 nop
9E61  83                 add a,e
9E62  1F                 rra
9E63  C0                 ret nz
9E64  F0                 ret p
9E65  05                 dec b
9E66  00                 nop
9E67  7B                 ld a,e
9E68  1F                 rra
9E69  C0                 ret nz
9E6A  F0                 ret p
9E6B  05                 dec b
9E6C  00                 nop
9E6D  FB                 ei
9E6E  0F                 rrca
9E6F  E0                 ret po
9E70  F0                 ret p
9E71  05                 dec b
9E72  00                 nop
9E73  FD 0F              rrca
9E75  E0                 ret po
9E76  F0                 ret p
9E77  05                 dec b
9E78  00                 nop
9E79  BD                 cp l
9E7A  0F                 rrca
9E7B  E0                 ret po
9E7C  F8                 ret m
9E7D  01 00 3D           ld bc,$3D00
9E80  0F                 rrca
9E81  E0                 ret po
9E82  FC 01 00           call m,$0001
9E85  1E 1F              ld e,31
9E87  C0                 ret nz
9E88  FC 01 00           call m,$0001
9E8B  0F                 rrca
9E8C  3F                 ccf
9E8D  00                 nop
9E8E  FC 01 00           call m,$0001
9E91  2F                 cpl
9E92  3F                 ccf
9E93  80                 add a,b
9E94  FE 00              cp 0
9E96  10 87              djnz $9E1F
9E98  7F                 ld a,a
9E99  00                 nop
9E9A  FF                 rst 38h
9E9B  00                 nop
9E9C  38 00              jr c,$9E9E
9E9E  FF                 rst 38h
9E9F  00                 nop
9EA0  FF                 rst 38h
9EA1  00                 nop
9EA2  01 00 FF           ld bc,-256
9EA5  00                 nop
9EA6  FE 00              cp 0
9EA8  00                 nop
9EA9  FA FF 00           jp m,$00FF
9EAC  FE 00              cp 0
9EAE  00                 nop
9EAF  D0                 ret nc
9EB0  FF                 rst 38h
9EB1  00                 nop
9EB2  FE 00              cp 0
9EB4  00                 nop
9EB5  D2 FF 00           jp nc,$00FF
9EB8  FC 00 00           call m,$0000
9EBB  00                 nop
9EBC  7F                 ld a,a
9EBD  00                 nop
9EBE  F8                 ret m
9EBF  03                 inc bc
9EC0  00                 nop
9EC1  EA 3F 80           jp pe,$803F
9EC4  FC 00 00           call m,$0000
9EC7  00                 nop
9EC8  7F                 ld a,a
9EC9  00                 nop
9ECA  FC 01 00           call m,$0001
9ECD  67                 ld h,a
9ECE  7F                 ld a,a
9ECF  00                 nop
9ED0  F8                 ret m
9ED1  00                 nop
9ED2  00                 nop
9ED3  1F                 rra
9ED4  7F                 ld a,a
9ED5  00                 nop
9ED6  FC 01 00           call m,$0001
9ED9  1F                 rra
9EDA  7F                 ld a,a
9EDB  00                 nop
9EDC  FC 01 00           call m,$0001
9EDF  FA FF 00           jp m,$00FF
9EE2  FE 00              cp 0
9EE4  00                 nop
9EE5  AC                 xor h
9EE6  7F                 ld a,a
9EE7  00                 nop
9EE8  FF                 rst 38h
9EE9  00                 nop
9EEA  00                 nop
9EEB  71                 ld (hl),c
9EEC  3F                 ccf
9EED  80                 add a,b
9EEE  FE 00              cp 0
9EF0  00                 nop
9EF1  83                 add a,e
9EF2  1F                 rra
9EF3  C0                 ret nz
9EF4  FC 01 00           call m,$0001
9EF7  7B                 ld a,e
9EF8  1F                 rra
9EF9  C0                 ret nz
9EFA  FC 01 00           call m,$0001
9EFD  E7                 rst 20h
9EFE  1F                 rra
9EFF  C0                 ret nz
9F00  FC 01 00           call m,$0001
9F03  DF                 rst 18h
9F04  3F                 ccf
9F05  80                 add a,b
9F06  F8                 ret m
9F07  03                 inc bc
9F08  00                 nop
9F09  DF                 rst 18h
9F0A  7F                 ld a,a
9F0B  00                 nop
9F0C  F8                 ret m
9F0D  03                 inc bc
9F0E  00                 nop
9F0F  0E 3F              ld c,63
9F11  80                 add a,b
9F12  F0                 ret p
9F13  06 10              ld b,16
9F15  01 1F C0           ld bc,$C01F
9F18  F0                 ret p
9F19  04                 inc b
9F1A  10 47              djnz $9F63
9F1C  1F                 rra
9F1D  C0                 ret nz
9F1E  F0                 ret p
9F1F  04                 inc b
9F20  38 03              jr c,$9F25
9F22  0F                 rrca
9F23  E0                 ret po
9F24  F8                 ret m
9F25  03                 inc bc
9F26  7C                 ld a,h
9F27  01 1F C0           ld bc,$C01F
9F2A  FC 00 FE           call m,$FE00
9F2D  00                 nop
9F2E  3F                 ccf
9F2F  00                 nop
9F30  FF                 rst 38h
9F31  00                 nop
9F32  FF                 rst 38h
9F33  00                 nop
9F34  FF                 rst 38h
9F35  00                 nop
9F36  FF                 rst 38h
9F37  00                 nop
9F38  FF                 rst 38h
9F39  00                 nop
9F3A  FF                 rst 38h
9F3B  00                 nop
9F3C  FF                 rst 38h
9F3D  00                 nop
9F3E  FF                 rst 38h
9F3F  00                 nop
9F40  FF                 rst 38h
9F41  00                 nop
9F42  FF                 rst 38h
9F43  00                 nop
9F44  FF                 rst 38h
9F45  00                 nop
9F46  FF                 rst 38h
9F47  00                 nop
9F48  FF                 rst 38h
9F49  00                 nop
9F4A  FF                 rst 38h
9F4B  00                 nop
9F4C  FF                 rst 38h
9F4D  00                 nop
9F4E  FF                 rst 38h
9F4F  00                 nop
9F50  FF                 rst 38h
9F51  00                 nop
9F52  FF                 rst 38h
9F53  00                 nop
9F54  FF                 rst 38h
9F55  00                 nop
9F56  FF                 rst 38h
9F57  00                 nop
9F58  FF                 rst 38h
9F59  00                 nop
9F5A  FF                 rst 38h
9F5B  00                 nop
9F5C  FF                 rst 38h
9F5D  00                 nop
9F5E  FF                 rst 38h
9F5F  00                 nop
9F60  FF                 rst 38h
9F61  00                 nop
9F62  FF                 rst 38h
9F63  00                 nop
9F64  FF                 rst 38h
9F65  00                 nop
9F66  FF                 rst 38h
9F67  00                 nop
9F68  FF                 rst 38h
9F69  00                 nop
9F6A  FF                 rst 38h
9F6B  00                 nop
9F6C  FF                 rst 38h
9F6D  00                 nop
9F6E  FF                 rst 38h
9F6F  00                 nop
9F70  FF                 rst 38h
9F71  00                 nop
9F72  FF                 rst 38h
9F73  00                 nop
9F74  FF                 rst 38h
9F75  00                 nop
9F76  FF                 rst 38h
9F77  00                 nop
9F78  FF                 rst 38h
9F79  00                 nop
9F7A  FF                 rst 38h
9F7B  00                 nop
9F7C  FF                 rst 38h
9F7D  00                 nop
9F7E  FF                 rst 38h
9F7F  00                 nop
9F80  E3                 ex (sp),hl
9F81  00                 nop
9F82  FF                 rst 38h
9F83  00                 nop
9F84  FF                 rst 38h
9F85  00                 nop
9F86  C1                 pop bc
9F87  1C                 inc e
9F88  FF                 rst 38h
9F89  00                 nop
9F8A  FF                 rst 38h
9F8B  00                 nop
9F8C  80                 add a,b
9F8D  2E FF              ld l,-1
9F8F  00                 nop
9F90  FF                 rst 38h
9F91  00                 nop
9F92  80                 add a,b
9F93  1E FF              ld e,-1
9F95  00                 nop
9F96  FF                 rst 38h
9F97  00                 nop
9F98  80                 add a,b
9F99  2E FF              ld l,-1
9F9B  00                 nop
9F9C  FF                 rst 38h
9F9D  00                 nop
9F9E  C1                 pop bc
9F9F  14                 inc d
9FA0  FF                 rst 38h
9FA1  00                 nop
9FA2  FF                 rst 38h
9FA3  00                 nop
9FA4  E3                 ex (sp),hl
9FA5  00                 nop
9FA6  FF                 rst 38h
9FA7  00                 nop
9FA8  FF                 rst 38h
9FA9  00                 nop
9FAA  FF                 rst 38h
9FAB  00                 nop
9FAC  FF                 rst 38h
9FAD  00                 nop
9FAE  FF                 rst 38h
9FAF  00                 nop
9FB0  FF                 rst 38h
9FB1  00                 nop
9FB2  FF                 rst 38h
9FB3  00                 nop
9FB4  FF                 rst 38h
9FB5  00                 nop
9FB6  FF                 rst 38h
9FB7  00                 nop
9FB8  FF                 rst 38h
9FB9  00                 nop
9FBA  FF                 rst 38h
9FBB  00                 nop
9FBC  FF                 rst 38h
9FBD  00                 nop
9FBE  FF                 rst 38h
9FBF  00                 nop
9FC0  FF                 rst 38h
9FC1  00                 nop
9FC2  FF                 rst 38h
9FC3  00                 nop
9FC4  FF                 rst 38h
9FC5  00                 nop
9FC6  FF                 rst 38h
9FC7  00                 nop
9FC8  FF                 rst 38h
9FC9  00                 nop
9FCA  FF                 rst 38h
9FCB  00                 nop
9FCC  FF                 rst 38h
9FCD  00                 nop
9FCE  FF                 rst 38h
9FCF  00                 nop
9FD0  FF                 rst 38h
9FD1  00                 nop
9FD2  FF                 rst 38h
9FD3  00                 nop
9FD4  FF                 rst 38h
9FD5  00                 nop
9FD6  FF                 rst 38h
9FD7  00                 nop
9FD8  FF                 rst 38h
9FD9  00                 nop
9FDA  FF                 rst 38h
9FDB  00                 nop
9FDC  DF                 rst 18h
9FDD  00                 nop
9FDE  FF                 rst 38h
9FDF  00                 nop
9FE0  FF                 rst 38h
9FE1  00                 nop
9FE2  0F                 rrca
9FE3  20 FF              jr nz,$9FE4
9FE5  00                 nop
9FE6  FC 00 0F           call m,$0F00
9FE9  E0                 ret po
9FEA  FF                 rst 38h
9FEB  00                 nop
9FEC  F0                 ret p
9FED  03                 inc bc
9FEE  1F                 rra
9FEF  C0                 ret nz
9FF0  FF                 rst 38h
9FF1  00                 nop
9FF2  C0                 ret nz
9FF3  0F                 rrca
9FF4  3F                 ccf
9FF5  80                 add a,b
9FF6  FF                 rst 38h
9FF7  00                 nop
9FF8  00                 nop
9FF9  3F                 ccf
9FFA  7F                 ld a,a
9FFB  00                 nop
9FFC  FE 00              cp 0
9FFE  00                 nop
9FFF  FE FF              cp -1
A001  00                 nop
A002  FF                 rst 38h
A003  00                 nop
A004  01 7C FF           ld bc,-132
A007  00                 nop
A008  FF                 rst 38h
A009  00                 nop
A00A  80                 add a,b
A00B  3E FF              ld a,-1
A00D  00                 nop
A00E  FF                 rst 38h
A00F  00                 nop
A010  C0                 ret nz
A011  1F                 rra
A012  7F                 ld a,a
A013  00                 nop
A014  FF                 rst 38h
A015  00                 nop
A016  80                 add a,b
A017  3F                 ccf
A018  3F                 ccf
A019  80                 add a,b
A01A  FF                 rst 38h
A01B  00                 nop
A01C  00                 nop
A01D  7E                 ld a,(hl)
A01E  7F                 ld a,a
A01F  00                 nop
A020  FE 00              cp 0
A022  01 F8 FF           ld bc,-8
A025  00                 nop
A026  FC 01 07           call m,$0701
A029  E0                 ret po
A02A  FF                 rst 38h
A02B  00                 nop
A02C  F8                 ret m
A02D  03                 inc bc
A02E  1F                 rra
A02F  80                 add a,b
A030  FF                 rst 38h
A031  00                 nop
A032  F0                 ret p
A033  06 7F              ld b,127
A035  00                 nop
A036  FF                 rst 38h
A037  00                 nop
A038  FF                 rst 38h
A039  00                 nop
A03A  FF                 rst 38h
A03B  00                 nop
A03C  FF                 rst 38h
A03D  00                 nop
A03E  FF                 rst 38h
A03F  00                 nop
A040  FF                 rst 38h
A041  00                 nop
A042  FF                 rst 38h
A043  00                 nop
A044  FF                 rst 38h
A045  00                 nop
A046  FF                 rst 38h
A047  00                 nop
A048  FF                 rst 38h
A049  00                 nop
A04A  FF                 rst 38h
A04B  00                 nop
A04C  FF                 rst 38h
A04D  00                 nop
A04E  FF                 rst 38h
A04F  00                 nop
A050  FF                 rst 38h
A051  00                 nop
A052  00                 nop
A053  00                 nop
A054  7F                 ld a,a
A055  00                 nop
A056  FE 00              cp 0
A058  00                 nop
A059  E3                 ex (sp),hl
A05A  3F                 ccf
A05B  80                 add a,b
A05C  FE 00              cp 0
A05E  00                 nop
A05F  D5                 push de
A060  3F                 ccf
A061  80                 add a,b
A062  FE 00              cp 0
A064  00                 nop
A065  E3                 ex (sp),hl
A066  3F                 ccf
A067  80                 add a,b
A068  FF                 rst 38h
A069  00                 nop
A06A  00                 nop
A06B  55                 ld d,l
A06C  7F                 ld a,a
A06D  00                 nop
A06E  FF                 rst 38h
A06F  00                 nop
A070  00                 nop
A071  63                 ld h,e
A072  7F                 ld a,a
A073  00                 nop
A074  FF                 rst 38h
A075  00                 nop
A076  00                 nop
A077  55                 ld d,l
A078  7F                 ld a,a
A079  00                 nop
A07A  FF                 rst 38h
A07B  00                 nop
A07C  80                 add a,b
A07D  22 FF 00           ld ($00FF),hl
A080  FF                 rst 38h
A081  00                 nop
A082  80                 add a,b
A083  36 FF              ld (hl),-1
A085  00                 nop
A086  FF                 rst 38h
A087  00                 nop
A088  80                 add a,b
A089  22 FF 00           ld ($00FF),hl
A08C  FF                 rst 38h
A08D  00                 nop
A08E  C1                 pop bc
A08F  14                 inc d
A090  FF                 rst 38h
A091  00                 nop
A092  FF                 rst 38h
A093  00                 nop
A094  C1                 pop bc
A095  14                 inc d
A096  FF                 rst 38h
A097  00                 nop
A098  FF                 rst 38h
A099  00                 nop
A09A  C1                 pop bc
A09B  14                 inc d
A09C  FF                 rst 38h
A09D  00                 nop
A09E  FF                 rst 38h
A09F  00                 nop
A0A0  E3                 ex (sp),hl
A0A1  08                 ex af,af'
A0A2  FF                 rst 38h
A0A3  00                 nop
A0A4  FF                 rst 38h
A0A5  00                 nop
A0A6  E3                 ex (sp),hl
A0A7  08                 ex af,af'
A0A8  FF                 rst 38h
A0A9  00                 nop
A0AA  FF                 rst 38h
A0AB  00                 nop
A0AC  E3                 ex (sp),hl
A0AD  08                 ex af,af'
A0AE  FF                 rst 38h
A0AF  00                 nop
A0B0  FF                 rst 38h
A0B1  00                 nop
A0B2  F7                 rst 30h
A0B3  00                 nop
A0B4  FF                 rst 38h
A0B5  00                 nop
A0B6  FF                 rst 38h
A0B7  00                 nop
A0B8  FF                 rst 38h
A0B9  00                 nop
A0BA  FF                 rst 38h
A0BB  00                 nop
A0BC  FF                 rst 38h
A0BD  00                 nop
A0BE  FF                 rst 38h
A0BF  00                 nop
A0C0  FF                 rst 38h
A0C1  00                 nop
A0C2  FF                 rst 38h
A0C3  00                 nop
A0C4  FF                 rst 38h
A0C5  00                 nop
A0C6  FF                 rst 38h
A0C7  00                 nop
A0C8  FF                 rst 38h
A0C9  00                 nop
A0CA  FF                 rst 38h
A0CB  00                 nop
A0CC  FF                 rst 38h
A0CD  00                 nop
A0CE  FF                 rst 38h
A0CF  00                 nop
A0D0  FF                 rst 38h
A0D1  00                 nop
A0D2  FF                 rst 38h
A0D3  00                 nop
A0D4  FF                 rst 38h
A0D5  00                 nop
A0D6  FF                 rst 38h
A0D7  00                 nop
A0D8  FF                 rst 38h
A0D9  00                 nop
A0DA  FF                 rst 38h
A0DB  00                 nop
A0DC  FF                 rst 38h
A0DD  00                 nop
A0DE  FF                 rst 38h
A0DF  00                 nop
A0E0  FF                 rst 38h
A0E1  00                 nop
A0E2  FF                 rst 38h
A0E3  00                 nop
A0E4  FF                 rst 38h
A0E5  00                 nop
A0E6  FF                 rst 38h
A0E7  00                 nop
A0E8  FF                 rst 38h
A0E9  00                 nop
A0EA  FF                 rst 38h
A0EB  00                 nop
A0EC  FF                 rst 38h
A0ED  00                 nop
A0EE  FF                 rst 38h
A0EF  00                 nop
A0F0  FF                 rst 38h
A0F1  00                 nop
A0F2  FF                 rst 38h
A0F3  00                 nop
A0F4  FF                 rst 38h
A0F5  00                 nop
A0F6  FF                 rst 38h
A0F7  00                 nop
A0F8  FF                 rst 38h
A0F9  00                 nop
A0FA  FF                 rst 38h
A0FB  00                 nop
A0FC  FF                 rst 38h
A0FD  00                 nop
A0FE  FF                 rst 38h
A0FF  00                 nop
A100  FF                 rst 38h
A101  00                 nop
A102  FF                 rst 38h
A103  00                 nop
A104  FF                 rst 38h
A105  00                 nop
A106  FF                 rst 38h
A107  00                 nop
A108  FF                 rst 38h
A109  00                 nop
A10A  FF                 rst 38h
A10B  00                 nop
A10C  FF                 rst 38h
A10D  00                 nop
A10E  FF                 rst 38h
A10F  00                 nop
A110  FF                 rst 38h
A111  00                 nop
A112  E0                 ret po
A113  00                 nop
A114  FF                 rst 38h
A115  00                 nop
A116  FF                 rst 38h
A117  00                 nop
A118  80                 add a,b
A119  1F                 rra
A11A  7F                 ld a,a
A11B  00                 nop
A11C  FF                 rst 38h
A11D  00                 nop
A11E  00                 nop
A11F  7D                 ld a,l
A120  3F                 ccf
A121  80                 add a,b
A122  FE 00              cp 0
A124  00                 nop
A125  FF                 rst 38h
A126  1F                 rra
A127  C0                 ret nz
A128  FC 01 00           call m,$0001
A12B  FA 1F 80           jp m,$801F
A12E  FC 01 00           call m,$0001
A131  DF                 rst 18h
A132  1F                 rra
A133  00                 nop
A134  FC 01 00           call m,$0001
A137  FA 1F 00           jp m,$001F
A13A  FC 01 00           call m,$0001
A13D  6D                 ld l,l
A13E  1F                 rra
A13F  00                 nop
A140  FC 01 00           call m,$0001
A143  FA 3F 00           jp m,$003F
A146  FC 01 00           call m,$0001
A149  54                 ld d,h
A14A  7F                 ld a,a
A14B  00                 nop
A14C  FE 00              cp 0
A14E  00                 nop
A14F  A0                 and b
A150  FF                 rst 38h
A151  00                 nop
A152  FF                 rst 38h
A153  00                 nop
A154  01 40 FF           ld bc,-192
A157  00                 nop
A158  FF                 rst 38h
A159  00                 nop
A15A  83                 add a,e
A15B  00                 nop
A15C  FF                 rst 38h
A15D  00                 nop
A15E  FF                 rst 38h
A15F  00                 nop
A160  FF                 rst 38h
A161  00                 nop
A162  FF                 rst 38h
A163  00                 nop
A164  FF                 rst 38h
A165  00                 nop
A166  FF                 rst 38h
A167  00                 nop
A168  FF                 rst 38h
A169  00                 nop
A16A  FF                 rst 38h
A16B  00                 nop
A16C  FF                 rst 38h
A16D  00                 nop
A16E  FF                 rst 38h
A16F  00                 nop
A170  FF                 rst 38h
A171  00                 nop
A172  FF                 rst 38h
A173  00                 nop
A174  FF                 rst 38h
A175  00                 nop
A176  FF                 rst 38h
A177  00                 nop
A178  FF                 rst 38h
A179  00                 nop
A17A  FF                 rst 38h
A17B  00                 nop
A17C  FF                 rst 38h
A17D  00                 nop
A17E  F7                 rst 30h
A17F  00                 nop
A180  7F                 ld a,a
A181  00                 nop
A182  F8                 ret m
A183  00                 nop
A184  E2 08 37           jp po,$3708
A187  80                 add a,b
A188  F0                 ret p
A189  07                 rlca
A18A  77                 ld (hl),a
A18B  00                 nop
A18C  63                 ld h,e
A18D  08                 ex af,af'
A18E  D8                 ret c
A18F  01 37 00           ld bc,55
A192  F7                 rst 30h
A193  00                 nop
A194  8C                 adc a,h
A195  21 03 C8           ld hl,$C803
A198  DF                 rst 18h
A199  00                 nop
A19A  C2 10 01           jp nz,$0110
A19D  CC 8F 20           call z,$208F
A1A0  E1                 pop hl
A1A1  0C                 inc c
A1A2  21 0C 0D           ld hl,$0D0C
A1A5  60                 ld h,b
A1A6  F0                 ret p
A1A7  04                 inc b
A1A8  21 00 08           ld hl,2048
A1AB  02                 ld (bc),a
A1AC  F8                 ret m
A1AD  03                 inc bc
A1AE  00                 nop
A1AF  DE 01              sbc a,1
A1B1  F4 F0 02           call p,$02F0
A1B4  00                 nop
A1B5  52                 ld d,d
A1B6  01 94 E0           ld bc,$E094
A1B9  0A                 ld a,(bc)
A1BA  00                 nop
A1BB  52                 ld d,d
A1BC  01 94 F0           ld bc,$F094
A1BF  03                 inc bc
A1C0  00                 nop
A1C1  D2 01 F4           jp nc,$F401
A1C4  F0                 ret p
A1C5  02                 ld (bc),a
A1C6  00                 nop
A1C7  12                 ld (de),a
A1C8  03                 inc bc
A1C9  80                 add a,b
A1CA  E0                 ret po
A1CB  0A                 ld a,(bc)
A1CC  00                 nop
A1CD  1E 01              ld e,1
A1CF  84                 add a,h
A1D0  F4 00 00           call p,$0000
A1D3  E1                 pop hl
A1D4  03                 inc bc
A1D5  30 F8              jr nc,$A1CF
A1D7  01 0C E1           ld bc,$E10C
A1DA  01 98 E0           ld bc,$E098
A1DD  07                 rlca
A1DE  04                 inc b
A1DF  11 20 8E           ld de,$8E20
A1E2  C0                 ret nz
A1E3  1C                 inc e
A1E4  66                 ld h,(hl)
A1E5  00                 nop
A1E6  30 82              jr nc,$A16A
A1E8  E2 00 02           jp po,$0200
A1EB  98                 sbc a,b
A1EC  2D                 dec l
A1ED  80                 add a,b
A1EE  F7                 rst 30h
A1EF  00                 nop
A1F0  43                 ld b,e
A1F1  18 07              jr $A1FA
A1F3  50                 ld d,b
A1F4  E3                 ex (sp),hl
A1F5  08                 ex af,af'
A1F6  C7                 rst 00h
A1F7  10 AF              djnz $A1A8
A1F9  00                 nop
A1FA  F7                 rst 30h
A1FB  00                 nop
A1FC  E7                 rst 20h
A1FD  10 FF              djnz $A1FE
A1FF  00                 nop
A200  FF                 rst 38h
A201  00                 nop
A202  FF                 rst 38h
A203  00                 nop
A204  FF                 rst 38h
A205  00                 nop
A206  80                 add a,b
A207  00                 nop
A208  00                 nop
A209  00                 nop
A20A  01 00 00           ld bc,0
A20D  7F                 ld a,a
A20E  00                 nop
A20F  FF                 rst 38h
A210  00                 nop
A211  FE 00              cp 0
A213  4A                 ld c,d
A214  00                 nop
A215  52                 ld d,d
A216  00                 nop
A217  86                 add a,(hl)
A218  00                 nop
A219  44                 ld b,h
A21A  00                 nop
A21B  73                 ld (hl),e
A21C  00                 nop
A21D  92                 sub d
A21E  00                 nop
A21F  40                 ld b,b
A220  00                 nop
A221  61                 ld h,c
A222  00                 nop
A223  92                 sub d
A224  00                 nop
A225  4A                 ld c,d
A226  00                 nop
A227  6D                 ld l,l
A228  00                 nop
A229  86                 add a,(hl)
A22A  00                 nop
A22B  4E                 ld c,(hl)
A22C  00                 nop
A22D  40                 ld b,b
A22E  00                 nop
A22F  9E                 sbc a,(hl)
A230  00                 nop
A231  4A                 ld c,d
A232  00                 nop
A233  4C                 ld c,h
A234  00                 nop
A235  92                 sub d
A236  00                 nop
A237  7F                 ld a,a
A238  00                 nop
A239  FF                 rst 38h
A23A  00                 nop
A23B  FE 00              cp 0
A23D  40                 ld b,b
A23E  00                 nop
A23F  14                 inc d
A240  00                 nop
A241  02                 ld (bc),a
A242  00                 nop
A243  5F                 ld e,a
A244  00                 nop
A245  CA 00 FA           jp z,$FA00
A248  00                 nop
A249  4F                 ld c,a
A24A  00                 nop
A24B  E4 00 F2           call po,$F200
A24E  00                 nop
A24F  67                 ld h,a
A250  00                 nop
A251  F1                 pop af
A252  00                 nop
A253  CA 00 53           jp z,$5300
A256  00                 nop
A257  FF                 rst 38h
A258  00                 nop
A259  96                 sub (hl)
A25A  00                 nop
A25B  69                 ld l,c
A25C  00                 nop
A25D  FF                 rst 38h
A25E  00                 nop
A25F  AE                 xor (hl)
A260  00                 nop
A261  55                 ld d,l
A262  00                 nop
A263  FF                 rst 38h
A264  01 98 00           ld bc,152
A267  6A                 ld l,d
A268  00                 nop
A269  FF                 rst 38h
A26A  01 84 00           ld bc,132
A26D  7C                 ld a,h
A26E  00                 nop
A26F  FF                 rst 38h
A270  01 BC 00           ld bc,188
A273  6A                 ld l,d
A274  00                 nop
A275  FF                 rst 38h
A276  03                 inc bc
A277  38 00              jr c,$A279
A279  7D                 ld a,l
A27A  00                 nop
A27B  00                 nop
A27C  07                 rlca
A27D  30 00              jr nc,$A27F
A27F  76                 halt
A280  00                 nop
A281  AA                 xor d
A282  0F                 rrca
A283  60                 ld h,b
A284  00                 nop
A285  7F                 ld a,a
A286  00                 nop
A287  FE 1F              cp 31
A289  00                 nop
A28A  80                 add a,b
A28B  00                 nop
A28C  00                 nop
A28D  00                 nop
A28E  7F                 ld a,a
A28F  00                 nop
A290  FF                 rst 38h
A291  00                 nop
A292  FF                 rst 38h
A293  00                 nop
A294  FF                 rst 38h
A295  00                 nop
A296  FF                 rst 38h
A297  00                 nop
A298  FF                 rst 38h
A299  00                 nop
A29A  FF                 rst 38h
A29B  00                 nop
A29C  FF                 rst 38h
A29D  00                 nop
A29E  FF                 rst 38h
A29F  00                 nop
A2A0  FF                 rst 38h
A2A1  00                 nop
A2A2  FF                 rst 38h
A2A3  00                 nop
A2A4  FF                 rst 38h
A2A5  00                 nop
A2A6  FF                 rst 38h
A2A7  00                 nop
A2A8  FF                 rst 38h
A2A9  00                 nop
A2AA  FF                 rst 38h
A2AB  00                 nop
A2AC  FF                 rst 38h
A2AD  00                 nop
A2AE  FF                 rst 38h
A2AF  00                 nop
A2B0  F7                 rst 30h
A2B1  00                 nop
A2B2  FF                 rst 38h
A2B3  00                 nop
A2B4  FF                 rst 38h
A2B5  00                 nop
A2B6  A2                 and d
A2B7  08                 ex af,af'
A2B8  FF                 rst 38h
A2B9  00                 nop
A2BA  FF                 rst 38h
A2BB  00                 nop
A2BC  00                 nop
A2BD  41                 ld b,c
A2BE  7F                 ld a,a
A2BF  00                 nop
A2C0  FE 00              cp 0
A2C2  80                 add a,b
A2C3  1C                 inc e
A2C4  BF                 cp a
A2C5  00                 nop
A2C6  FC 01 00           call m,$0001
A2C9  77                 ld (hl),a
A2CA  1F                 rra
A2CB  40                 ld b,b
A2CC  FE 00              cp 0
A2CE  00                 nop
A2CF  43                 ld b,e
A2D0  3F                 ccf
A2D1  00                 nop
A2D2  FC 00 00           call m,$0000
A2D5  EF                 rst 28h
A2D6  1F                 rra
A2D7  80                 add a,b
A2D8  F8                 ret m
A2D9  02                 ld (bc),a
A2DA  00                 nop
A2DB  9F                 sbc a,a
A2DC  0F                 rrca
A2DD  A0                 and b
A2DE  FC 00 00           call m,$0000
A2E1  DF                 rst 18h
A2E2  1F                 rra
A2E3  80                 add a,b
A2E4  FE 00              cp 0
A2E6  00                 nop
A2E7  7F                 ld a,a
A2E8  3F                 ccf
A2E9  00                 nop
A2EA  FC 01 00           call m,$0001
A2ED  7F                 ld a,a
A2EE  1F                 rra
A2EF  40                 ld b,b
A2F0  FE 00              cp 0
A2F2  80                 add a,b
A2F3  1C                 inc e
A2F4  BF                 cp a
A2F5  00                 nop
A2F6  FF                 rst 38h
A2F7  00                 nop
A2F8  00                 nop
A2F9  41                 ld b,c
A2FA  7F                 ld a,a
A2FB  00                 nop
A2FC  FF                 rst 38h
A2FD  00                 nop
A2FE  A2                 and d
A2FF  08                 ex af,af'
A300  FF                 rst 38h
A301  00                 nop
A302  FF                 rst 38h
A303  00                 nop
A304  F7                 rst 30h
A305  00                 nop
A306  FF                 rst 38h
A307  00                 nop
A308  FF                 rst 38h
A309  00                 nop
A30A  FF                 rst 38h
A30B  00                 nop
A30C  FF                 rst 38h
A30D  00                 nop
A30E  FF                 rst 38h
A30F  00                 nop
A310  FF                 rst 38h
A311  00                 nop
A312  FF                 rst 38h
A313  00                 nop
A314  FF                 rst 38h
A315  00                 nop
A316  FF                 rst 38h
A317  00                 nop
A318  FF                 rst 38h
A319  00                 nop
A31A  FF                 rst 38h
A31B  00                 nop
A31C  FF                 rst 38h
A31D  00                 nop
A31E  FF                 rst 38h
A31F  00                 nop
A320  FF                 rst 38h
A321  00                 nop
A322  FF                 rst 38h
A323  00                 nop
A324  FF                 rst 38h
A325  00                 nop
A326  FF                 rst 38h
A327  00                 nop
A328  FF                 rst 38h
A329  00                 nop
A32A  FF                 rst 38h
A32B  00                 nop
A32C  FF                 rst 38h
A32D  00                 nop
A32E  FF                 rst 38h
A32F  00                 nop
A330  FF                 rst 38h
A331  00                 nop
A332  FF                 rst 38h
A333  00                 nop
A334  FF                 rst 38h
A335  00                 nop
A336  FF                 rst 38h
A337  00                 nop
A338  FF                 rst 38h
A339  00                 nop
A33A  FF                 rst 38h
A33B  00                 nop
A33C  FF                 rst 38h
A33D  00                 nop
A33E  FF                 rst 38h
A33F  00                 nop
A340  FF                 rst 38h
A341  00                 nop
A342  FF                 rst 38h
A343  00                 nop
A344  FF                 rst 38h
A345  00                 nop
A346  F3                 di
A347  00                 nop
A348  FF                 rst 38h
A349  00                 nop
A34A  FF                 rst 38h
A34B  00                 nop
A34C  E1                 pop hl
A34D  0C                 inc c
A34E  FF                 rst 38h
A34F  00                 nop
A350  FF                 rst 38h
A351  00                 nop
A352  C0                 ret nz
A353  12                 ld (de),a
A354  FF                 rst 38h
A355  00                 nop
A356  FF                 rst 38h
A357  00                 nop
A358  80                 add a,b
A359  25                 dec h
A35A  7F                 ld a,a
A35B  00                 nop
A35C  FF                 rst 38h
A35D  00                 nop
A35E  00                 nop
A35F  21 3F 00           ld hl,63
A362  FE 00              cp 0
A364  00                 nop
A365  F3                 di
A366  1F                 rra
A367  C0                 ret nz
A368  FC 01 00           call m,$0001
A36B  7E                 ld a,(hl)
A36C  0F                 rrca
A36D  E0                 ret po
A36E  FC 01 00           call m,$0001
A371  62                 ld h,d
A372  0F                 rrca
A373  20 F8              jr nz,$A36D
A375  02                 ld (bc),a
A376  00                 nop
A377  00                 nop
A378  07                 rlca
A379  10 F8              djnz $A373
A37B  02                 ld (bc),a
A37C  00                 nop
A37D  2D                 dec l
A37E  07                 rlca
A37F  F0                 ret p
A380  FC 01 00           call m,$0001
A383  17                 rla
A384  0F                 rrca
A385  E0                 ret po
A386  FC 01 00           call m,$0001
A389  0B                 dec bc
A38A  0F                 rrca
A38B  60                 ld h,b
A38C  FE 00              cp 0
A38E  00                 nop
A38F  97                 sub a
A390  1F                 rra
A391  C0                 ret nz
A392  FE 00              cp 0
A394  00                 nop
A395  8D                 adc a,l
A396  1F                 rra
A397  C0                 ret nz
A398  FF                 rst 38h
A399  00                 nop
A39A  00                 nop
A39B  57                 ld d,a
A39C  3F                 ccf
A39D  80                 add a,b
A39E  FF                 rst 38h
A39F  00                 nop
A3A0  80                 add a,b
A3A1  00                 nop
A3A2  7F                 ld a,a
A3A3  00                 nop
A3A4  FF                 rst 38h
A3A5  00                 nop
A3A6  FF                 rst 38h
A3A7  00                 nop
A3A8  FF                 rst 38h
A3A9  00                 nop
A3AA  FF                 rst 38h
A3AB  00                 nop
A3AC  FF                 rst 38h
A3AD  00                 nop
A3AE  FF                 rst 38h
A3AF  00                 nop
A3B0  FF                 rst 38h
A3B1  00                 nop
A3B2  FF                 rst 38h
A3B3  00                 nop
A3B4  FF                 rst 38h
A3B5  00                 nop
A3B6  F3                 di
A3B7  00                 nop
A3B8  FF                 rst 38h
A3B9  00                 nop
A3BA  FF                 rst 38h
A3BB  00                 nop
A3BC  E1                 pop hl
A3BD  0C                 inc c
A3BE  FF                 rst 38h
A3BF  00                 nop
A3C0  FF                 rst 38h
A3C1  00                 nop
A3C2  C0                 ret nz
A3C3  1E FF              ld e,-1
A3C5  00                 nop
A3C6  FF                 rst 38h
A3C7  00                 nop
A3C8  C0                 ret nz
A3C9  0E FF              ld c,-1
A3CB  00                 nop
A3CC  FF                 rst 38h
A3CD  00                 nop
A3CE  C0                 ret nz
A3CF  17                 rla
A3D0  7F                 ld a,a
A3D1  00                 nop
A3D2  FF                 rst 38h
A3D3  00                 nop
A3D4  C0                 ret nz
A3D5  0F                 rrca
A3D6  63                 ld h,e
A3D7  00                 nop
A3D8  FF                 rst 38h
A3D9  00                 nop
A3DA  C0                 ret nz
A3DB  17                 rla
A3DC  00                 nop
A3DD  9C                 sbc a,h
A3DE  FF                 rst 38h
A3DF  00                 nop
A3E0  E0                 ret po
A3E1  0B                 dec bc
A3E2  00                 nop
A3E3  A3                 and e
A3E4  7F                 ld a,a
A3E5  00                 nop
A3E6  C0                 ret nz
A3E7  05                 dec b
A3E8  00                 nop
A3E9  C4 3F 80           call nz,$803F
A3EC  80                 add a,b
A3ED  2B                 dec hl
A3EE  00                 nop
A3EF  E9                 jp (hl)
A3F0  3F                 ccf
A3F1  80                 add a,b
A3F2  80                 add a,b
A3F3  25                 dec h
A3F4  00                 nop
A3F5  EC 7F 00           call pe,$007F
A3F8  C0                 ret nz
A3F9  02                 ld (bc),a
A3FA  00                 nop
A3FB  DE FF              sbc a,-1
A3FD  00                 nop
A3FE  C0                 ret nz
A3FF  10 00              djnz $A401
A401  3F                 ccf
A402  3F                 ccf
A403  00                 nop
A404  80                 add a,b
A405  2A 00 1F           ld hl,($1F00)
A408  0F                 rrca
A409  40                 ld b,b
A40A  80                 add a,b
A40B  2B                 dec hl
A40C  00                 nop
A40D  4F                 ld c,a
A40E  03                 inc bc
A40F  B0                 or b
A410  80                 add a,b
A411  2A 00 77           ld hl,($7700)
A414  01 8C 80           ld bc,$808C
A417  2A 80 1B           ld hl,($1B80)
A41A  00                 nop
A41B  B2                 or d
A41C  80                 add a,b
A41D  2A E0 03           ld hl,($03E0)
A420  01 8C C0           ld bc,$C08C
A423  12                 ld (de),a
A424  FC 01 03           call m,$0301
A427  B0                 or b
A428  C0                 ret nz
A429  12                 ld (de),a
A42A  FE 00              cp 0
A42C  0F                 rrca
A42D  00                 nop
A42E  E0                 ret po
A42F  0E FF              ld c,-1
A431  00                 nop
A432  FF                 rst 38h
A433  00                 nop
A434  F0                 ret p
A435  07                 rlca
A436  7F                 ld a,a
A437  00                 nop
A438  FF                 rst 38h
A439  00                 nop
A43A  F8                 ret m
A43B  00                 nop
A43C  FF                 rst 38h
A43D  00                 nop
A43E  FF                 rst 38h
A43F  00                 nop
A440  FF                 rst 38h
A441  00                 nop
A442  FF                 rst 38h
A443  00                 nop
A444  FF                 rst 38h
A445  00                 nop
A446  FF                 rst 38h
A447  00                 nop
A448  FF                 rst 38h
A449  00                 nop
A44A  FF                 rst 38h
A44B  00                 nop
A44C  FF                 rst 38h
A44D  00                 nop
A44E  FF                 rst 38h
A44F  00                 nop
A450  FF                 rst 38h
A451  00                 nop
A452  FF                 rst 38h
A453  00                 nop
A454  FF                 rst 38h
A455  00                 nop
A456  FF                 rst 38h
A457  00                 nop
A458  FF                 rst 38h
A459  00                 nop
A45A  FF                 rst 38h
A45B  00                 nop
A45C  FF                 rst 38h
A45D  00                 nop
A45E  FF                 rst 38h
A45F  00                 nop
A460  FF                 rst 38h
A461  00                 nop
A462  FF                 rst 38h
A463  00                 nop
A464  FF                 rst 38h
A465  00                 nop
A466  FF                 rst 38h
A467  00                 nop
A468  FF                 rst 38h
A469  00                 nop
A46A  FF                 rst 38h
A46B  00                 nop
A46C  FF                 rst 38h
A46D  00                 nop
A46E  FF                 rst 38h
A46F  00                 nop
A470  FF                 rst 38h
A471  00                 nop
A472  FE 00              cp 0
A474  07                 rlca
A475  00                 nop
A476  FF                 rst 38h
A477  00                 nop
A478  FC 00 03           call m,$0300
A47B  B0                 or b
A47C  FF                 rst 38h
A47D  00                 nop
A47E  00                 nop
A47F  01 03 C8           ld bc,$C803
A482  FC 00 00           call m,$0000
A485  7C                 ld a,h
A486  03                 inc bc
A487  E8                 ret pe
A488  FC 01 00           call m,$0001
A48B  C7                 rst 00h
A48C  03                 inc bc
A48D  48                 ld c,b
A48E  F8                 ret m
A48F  01 00 1D           ld bc,$1D00
A492  03                 inc bc
A493  08                 ex af,af'
A494  F8                 ret m
A495  03                 inc bc
A496  00                 nop
A497  7F                 ld a,a
A498  03                 inc bc
A499  90                 sub b
A49A  F8                 ret m
A49B  02                 ld (bc),a
A49C  00                 nop
A49D  7F                 ld a,a
A49E  07                 rlca
A49F  90                 sub b
A4A0  F8                 ret m
A4A1  02                 ld (bc),a
A4A2  00                 nop
A4A3  FF                 rst 38h
A4A4  07                 rlca
A4A5  A0                 and b
A4A6  F8                 ret m
A4A7  02                 ld (bc),a
A4A8  00                 nop
A4A9  FF                 rst 38h
A4AA  0F                 rrca
A4AB  A0                 and b
A4AC  F8                 ret m
A4AD  03                 inc bc
A4AE  00                 nop
A4AF  FF                 rst 38h
A4B0  0F                 rrca
A4B1  A0                 and b
A4B2  F8                 ret m
A4B3  01 00 7F           ld bc,$7F00
A4B6  0F                 rrca
A4B7  20 FC              jr nz,$A4B5
A4B9  01 00 FF           ld bc,-256
A4BC  0F                 rrca
A4BD  00                 nop
A4BE  FC 00 00           call m,$0000
A4C1  7C                 ld a,h
A4C2  7F                 ld a,a
A4C3  00                 nop
A4C4  FF                 rst 38h
A4C5  00                 nop
A4C6  01 00 FF           ld bc,-256
A4C9  00                 nop
A4CA  FF                 rst 38h
A4CB  00                 nop
A4CC  FF                 rst 38h
A4CD  00                 nop
A4CE  FF                 rst 38h
A4CF  00                 nop
A4D0  FF                 rst 38h
A4D1  00                 nop
A4D2  C9                 ret
A4D3  00                 nop
A4D4  FF                 rst 38h
A4D5  00                 nop
A4D6  FF                 rst 38h
A4D7  00                 nop
A4D8  00                 nop
A4D9  36 7F              ld (hl),127
A4DB  00                 nop
A4DC  EE 00              xor $00
A4DE  00                 nop
A4DF  CB 3F              srl a
A4E1  80                 add a,b
A4E2  C4 11 00           call nz,$0011
A4E5  BD                 cp l
A4E6  1D                 dec e
A4E7  C0                 ret nz
A4E8  8C                 adc a,h
A4E9  21 00 7E           ld hl,$7E00
A4EC  18 C2              jr $A4B0
A4EE  80                 add a,b
A4EF  13                 inc de
A4F0  00                 nop
A4F1  13                 inc de
A4F2  08                 ex af,af'
A4F3  61                 ld h,c
A4F4  88                 adc a,b
A4F5  27                 daa
A4F6  00                 nop
A4F7  7E                 ld a,(hl)
A4F8  08                 ex af,af'
A4F9  E3                 ex (sp),hl
A4FA  C0                 ret nz
A4FB  12                 ld (de),a
A4FC  00                 nop
A4FD  FF                 rst 38h
A4FE  00                 nop
A4FF  70                 ld (hl),b
A500  C0                 ret nz
A501  0C                 inc c
A502  00                 nop
A503  F9                 ld sp,hl
A504  01 3C 80           ld bc,$803C
A507  31 00 63           ld sp,$6300
A50A  03                 inc bc
A50B  00                 nop
A50C  00                 nop
A50D  6F                 ld l,a
A50E  00                 nop
A50F  1E 1F              ld e,31
A511  C0                 ret nz
A512  00                 nop
A513  6E                 ld l,(hl)
A514  00                 nop
A515  E0                 ret po
A516  0F                 rrca
A517  E0                 ret po
A518  80                 add a,b
A519  1E 00              ld e,0
A51B  FE 01              cp 1
A51D  F0                 ret p
A51E  E0                 ret po
A51F  00                 nop
A520  00                 nop
A521  3E 00              ld a,0
A523  6E                 ld l,(hl)
A524  F8                 ret m
A525  03                 inc bc
A526  00                 nop
A527  C1                 pop bc
A528  00                 nop
A529  0E F0              ld c,-16
A52B  07                 rlca
A52C  00                 nop
A52D  FF                 rst 38h
A52E  70                 ld (hl),b
A52F  06 E0              ld b,-32
A531  08                 ex af,af'
A532  00                 nop
A533  EF                 rst 28h
A534  39                 add hl,sp
A535  80                 add a,b
A536  F0                 ret p
A537  06 00              ld b,0
A539  57                 ld d,a
A53A  3F                 ccf
A53B  80                 add a,b
A53C  E0                 ret po
A53D  00                 nop
A53E  00                 nop
A53F  CF                 rst 08h
A540  7F                 ld a,a
A541  00                 nop
A542  C1                 pop bc
A543  1C                 inc e
A544  30 00              jr nc,$A546
A546  7F                 ld a,a
A547  00                 nop
A548  E1                 pop hl
A549  0C                 inc c
A54A  F8                 ret m
A54B  03                 inc bc
A54C  3F                 ccf
A54D  80                 add a,b
A54E  F3                 di
A54F  00                 nop
A550  F8                 ret m
A551  03                 inc bc
A552  1F                 rra
A553  C0                 ret nz
A554  FF                 rst 38h
A555  00                 nop
A556  FC 01 1F           call m,$1F01
A559  C0                 ret nz
A55A  FF                 rst 38h
A55B  00                 nop
A55C  FC 01 3F           call m,$3F01
A55F  80                 add a,b
A560  FF                 rst 38h
A561  00                 nop
A562  FF                 rst 38h
A563  00                 nop
A564  FF                 rst 38h
A565  00                 nop
A566  FF                 rst 38h
A567  00                 nop
A568  C9                 ret
A569  00                 nop
A56A  FF                 rst 38h
A56B  00                 nop
A56C  FF                 rst 38h
A56D  00                 nop
A56E  00                 nop
A56F  36 7F              ld (hl),127
A571  00                 nop
A572  F6 00              or $00
A574  00                 nop
A575  CB 3B              srl e
A577  80                 add a,b
A578  E0                 ret po
A579  09                 add hl,bc
A57A  00                 nop
A57B  BD                 cp l
A57C  11 C4 C0           ld de,$C0C4
A57F  11 00 7E           ld de,$7E00
A582  18 C2              jr $A546
A584  C8                 ret z
A585  03                 inc bc
A586  00                 nop
A587  13                 inc de
A588  00                 nop
A589  66                 ld h,(hl)
A58A  C0                 ret nz
A58B  17                 rla
A58C  00                 nop
A58D  7E                 ld a,(hl)
A58E  01 E0 E0           ld bc,$E0E0
A591  0A                 ld a,(bc)
A592  00                 nop
A593  FF                 rst 38h
A594  01 7C F0           ld bc,$F07C
A597  04                 inc b
A598  00                 nop
A599  C1                 pop bc
A59A  03                 inc bc
A59B  38 C8              jr c,$A565
A59D  00                 nop
A59E  00                 nop
A59F  63                 ld h,e
A5A0  47                 ld b,a
A5A1  00                 nop
A5A2  80                 add a,b
A5A3  31 00 1E           ld sp,$1E00
A5A6  71                 ld (hl),c
A5A7  00                 nop
A5A8  00                 nop
A5A9  6F                 ld l,a
A5AA  00                 nop
A5AB  40                 ld b,b
A5AC  00                 nop
A5AD  8E                 adc a,(hl)
A5AE  00                 nop
A5AF  6E                 ld l,(hl)
A5B0  00                 nop
A5B1  E2 00 EE           jp po,$EE00
A5B4  80                 add a,b
A5B5  1E 00              ld e,0
A5B7  FE 00              cp 0
A5B9  F6 E1              or $E1
A5BB  00                 nop
A5BC  00                 nop
A5BD  20 01              jr nz,$A5C0
A5BF  70                 ld (hl),b
A5C0  FE 00              cp 0
A5C2  00                 nop
A5C3  DF                 rst 18h
A5C4  01 1C FC           ld bc,-996
A5C7  01 00 A3           ld bc,$A300
A5CA  23                 inc hl
A5CB  80                 add a,b
A5CC  FC 01 00           call m,$0001
A5CF  DD 3F              ccf
A5D1  80                 add a,b
A5D2  F8                 ret m
A5D3  03                 inc bc
A5D4  00                 nop
A5D5  DD 7F              ld a,a
A5D7  00                 nop
A5D8  FC 00 00           call m,$0000
A5DB  00                 nop
A5DC  FF                 rst 38h
A5DD  00                 nop
A5DE  FE 00              cp 0
A5E0  01 1C FF           ld bc,-228
A5E3  00                 nop
A5E4  FC 01 23           call m,$2301
A5E7  80                 add a,b
A5E8  FF                 rst 38h
A5E9  00                 nop
A5EA  F8                 ret m
A5EB  03                 inc bc
A5EC  1F                 rra
A5ED  C0                 ret nz
A5EE  FF                 rst 38h
A5EF  00                 nop
A5F0  FF                 rst 38h
A5F1  00                 nop
A5F2  FF                 rst 38h
A5F3  00                 nop
A5F4  FF                 rst 38h
A5F5  00                 nop
A5F6  FF                 rst 38h
A5F7  00                 nop
A5F8  F0                 ret p
A5F9  00                 nop
A5FA  FF                 rst 38h
A5FB  00                 nop
A5FC  FF                 rst 38h
A5FD  00                 nop
A5FE  E0                 ret po
A5FF  0F                 rrca
A600  7F                 ld a,a
A601  00                 nop
A602  FF                 rst 38h
A603  00                 nop
A604  C0                 ret nz
A605  10 3F              djnz $A646
A607  80                 add a,b
A608  FF                 rst 38h
A609  00                 nop
A60A  C0                 ret nz
A60B  00                 nop
A60C  1F                 rra
A60D  40                 ld b,b
A60E  FF                 rst 38h
A60F  00                 nop
A610  80                 add a,b
A611  2C                 inc l
A612  1F                 rra
A613  40                 ld b,b
A614  FF                 rst 38h
A615  00                 nop
A616  80                 add a,b
A617  06 0F              ld b,15
A619  20 FF              jr nz,$A61A
A61B  00                 nop
A61C  00                 nop
A61D  6C                 ld l,h
A61E  0F                 rrca
A61F  20 FE              jr nz,$A61F
A621  00                 nop
A622  00                 nop
A623  A0                 and b
A624  0F                 rrca
A625  60                 ld h,b
A626  FC 01 00           call m,$0001
A629  F8                 ret m
A62A  07                 rlca
A62B  90                 sub b
A62C  FE 00              cp 0
A62E  00                 nop
A62F  00                 nop
A630  07                 rlca
A631  10 FF              djnz $A632
A633  00                 nop
A634  00                 nop
A635  49                 ld c,c
A636  07                 rlca
A637  10 FF              djnz $A638
A639  00                 nop
A63A  00                 nop
A63B  55                 ld d,l
A63C  03                 inc bc
A63D  08                 ex af,af'
A63E  FE 00              cp 0
A640  00                 nop
A641  AD                 xor l
A642  03                 inc bc
A643  08                 ex af,af'
A644  FE 00              cp 0
A646  00                 nop
A647  9D                 sbc a,l
A648  03                 inc bc
A649  08                 ex af,af'
A64A  FE 00              cp 0
A64C  00                 nop
A64D  AD                 xor l
A64E  03                 inc bc
A64F  08                 ex af,af'
A650  FE 00              cp 0
A652  00                 nop
A653  9C                 sbc a,h
A654  03                 inc bc
A655  88                 adc a,b
A656  FE 00              cp 0
A658  00                 nop
A659  AC                 xor h
A65A  07                 rlca
A65B  90                 sub b
A65C  FF                 rst 38h
A65D  00                 nop
A65E  00                 nop
A65F  5C                 ld e,h
A660  07                 rlca
A661  50                 ld d,b
A662  FF                 rst 38h
A663  00                 nop
A664  00                 nop
A665  4E                 ld c,(hl)
A666  0F                 rrca
A667  20 FF              jr nz,$A668
A669  00                 nop
A66A  80                 add a,b
A66B  26 0F              ld h,15
A66D  00                 nop
A66E  FF                 rst 38h
A66F  00                 nop
A670  C0                 ret nz
A671  0E 0F              ld c,15
A673  20 FF              jr nz,$A674
A675  00                 nop
A676  80                 add a,b
A677  11 1F C0           ld de,$C01F
A67A  FF                 rst 38h
A67B  00                 nop
A67C  00                 nop
A67D  77                 ld (hl),a
A67E  0F                 rrca
A67F  E0                 ret po
A680  F9                 ld sp,hl
A681  00                 nop
A682  FF                 rst 38h
A683  00                 nop
A684  FF                 rst 38h
A685  00                 nop
A686  F0                 ret p
A687  06 FF              ld b,-1
A689  00                 nop
A68A  FF                 rst 38h
A68B  00                 nop
A68C  E0                 ret po
A68D  09                 add hl,bc
A68E  7F                 ld a,a
A68F  00                 nop
A690  FF                 rst 38h
A691  00                 nop
A692  E0                 ret po
A693  08                 ex af,af'
A694  0F                 rrca
A695  00                 nop
A696  F3                 di
A697  00                 nop
A698  E0                 ret po
A699  09                 add hl,bc
A69A  07                 rlca
A69B  F0                 ret p
A69C  E1                 pop hl
A69D  0C                 inc c
A69E  E0                 ret po
A69F  0A                 ld a,(bc)
A6A0  03                 inc bc
A6A1  08                 ex af,af'
A6A2  C0                 ret nz
A6A3  0A                 ld a,(bc)
A6A4  C0                 ret nz
A6A5  10 01              djnz $A6A8
A6A7  04                 inc b
A6A8  00                 nop
A6A9  32 C0 15           ld ($15C0),a
A6AC  00                 nop
A6AD  84                 add a,h
A6AE  00                 nop
A6AF  C2 E0 00           jp nz,$00E0
A6B2  00                 nop
A6B3  C3 01 04           jp $0401
A6B6  E0                 ret po
A6B7  0D                 dec c
A6B8  00                 nop
A6B9  80                 add a,b
A6BA  01 04 C0           ld bc,$C004
A6BD  1C                 inc e
A6BE  00                 nop
A6BF  00                 nop
A6C0  03                 inc bc
A6C1  08                 ex af,af'
A6C2  C0                 ret nz
A6C3  16 00              ld d,0
A6C5  00                 nop
A6C6  07                 rlca
A6C7  30 80              jr nc,$A649
A6C9  38 00              jr c,$A6CB
A6CB  03                 inc bc
A6CC  0F                 rrca
A6CD  C0                 ret nz
A6CE  C0                 ret nz
A6CF  02                 ld (bc),a
A6D0  00                 nop
A6D1  80                 add a,b
A6D2  3F                 ccf
A6D3  00                 nop
A6D4  FC 01 00           call m,$0001
A6D7  00                 nop
A6D8  3F                 ccf
A6D9  80                 add a,b
A6DA  FC 00 00           call m,$0000
A6DD  C0                 ret nz
A6DE  3F                 ccf
A6DF  80                 add a,b
A6E0  FC 01 00           call m,$0001
A6E3  40                 ld b,b
A6E4  1F                 rra
A6E5  40                 ld b,b
A6E6  FE 00              cp 0
A6E8  00                 nop
A6E9  A0                 and b
A6EA  1F                 rra
A6EB  40                 ld b,b
A6EC  FF                 rst 38h
A6ED  00                 nop
A6EE  00                 nop
A6EF  60                 ld h,b
A6F0  0F                 rrca
A6F1  20 FF              jr nz,$A6F2
A6F3  00                 nop
A6F4  00                 nop
A6F5  30 0F              jr nc,$A706
A6F7  00                 nop
A6F8  FE 00              cp 0
A6FA  00                 nop
A6FB  9C                 sbc a,h
A6FC  07                 rlca
A6FD  70                 ld (hl),b
A6FE  FE 00              cp 0
A700  00                 nop
A701  E6 03              and $03
A703  D8                 ret c
A704  FF                 rst 38h
A705  00                 nop
A706  00                 nop
A707  70                 ld (hl),b
A708  01 2C FF           ld bc,-212
A70B  00                 nop
A70C  81                 add a,c
A70D  1C                 inc e
A70E  01 7C FF           ld bc,-132
A711  00                 nop
A712  F8                 ret m
A713  03                 inc bc
A714  3F                 ccf
A715  80                 add a,b
A716  FF                 rst 38h
A717  00                 nop
A718  F0                 ret p
A719  04                 inc b
A71A  1F                 rra
A71B  40                 ld b,b
A71C  FF                 rst 38h
A71D  00                 nop
A71E  E0                 ret po
A71F  08                 ex af,af'
A720  0F                 rrca
A721  20 CB              jr nz,$A6EE
A723  00                 nop
A724  C0                 ret nz
A725  10 07              djnz $A72E
A727  10 81              djnz $A6AA
A729  34                 inc (hl)
A72A  80                 add a,b
A72B  26 03              ld h,3
A72D  C8                 ret z
A72E  01 74 80           ld bc,$8074
A731  29                 add hl,hl
A732  03                 inc bc
A733  28 01              jr z,$A736
A735  74                 ld (hl),h
A736  00                 nop
A737  56                 ld d,(hl)
A738  01 94 00           ld bc,148
A73B  78                 ld a,b
A73C  00                 nop
A73D  52                 ld d,d
A73E  01 D4 80           ld bc,$80D4
A741  30 00              jr nc,$A743
A743  68                 ld l,b
A744  01 AC C8           ld bc,$C8AC
A747  00                 nop
A748  00                 nop
A749  5C                 ld e,h
A74A  01 74 FC           ld bc,-908
A74D  00                 nop
A74E  00                 nop
A74F  1F                 rra
A750  03                 inc bc
A751  F0                 ret p
A752  F8                 ret m
A753  03                 inc bc
A754  00                 nop
A755  0E 07              ld c,7
A757  E0                 ret po
A758  F0                 ret p
A759  04                 inc b
A75A  00                 nop
A75B  83                 add a,e
A75C  03                 inc bc
A75D  88                 adc a,b
A75E  F0                 ret p
A75F  04                 inc b
A760  00                 nop
A761  4C                 ld c,h
A762  03                 inc bc
A763  08                 ex af,af'
A764  F0                 ret p
A765  04                 inc b
A766  00                 nop
A767  2F                 cpl
A768  01 84 E0           ld bc,$E084
A76B  09                 add hl,bc
A76C  00                 nop
A76D  17                 rla
A76E  01 84 80           ld bc,$8084
A771  09                 add hl,bc
A772  00                 nop
A773  17                 rla
A774  01 44 00           ld bc,68
A777  64                 ld h,h
A778  00                 nop
A779  8B                 adc a,e
A77A  03                 inc bc
A77B  20 00              jr nz,$A77D
A77D  60                 ld h,b
A77E  00                 nop
A77F  84                 add a,h
A780  01 0C 00           ld bc,12
A783  7C                 ld a,h
A784  00                 nop
A785  80                 add a,b
A786  00                 nop
A787  FE 80              cp -128
A789  3C                 inc a
A78A  00                 nop
A78B  40                 ld b,b
A78C  00                 nop
A78D  BE                 cp (hl)
A78E  C0                 ret nz
A78F  1E 00              ld e,0
A791  A0                 and b
A792  00                 nop
A793  2A E0 07           ld hl,($07E0)
A796  40                 ld b,b
A797  1E 80              ld e,-128
A799  2A F8 00           ld hl,($00F8)
A79C  E1                 pop hl
A79D  00                 nop
A79E  D5                 push de
A79F  00                 nop
A7A0  FF                 rst 38h
A7A1  00                 nop
A7A2  00                 nop
A7A3  00                 nop
A7A4  FF                 rst 38h
A7A5  00                 nop
A7A6  FA 00 00           jp m,$0000
A7A9  B7                 or a
A7AA  3F                 ccf
A7AB  00                 nop
A7AC  F0                 ret p
A7AD  04                 inc b
A7AE  00                 nop
A7AF  81                 add a,c
A7B0  1F                 rra
A7B1  C0                 ret nz
A7B2  E0                 ret po
A7B3  0C                 inc c
A7B4  00                 nop
A7B5  37                 scf
A7B6  0F                 rrca
A7B7  60                 ld h,b
A7B8  E0                 ret po
A7B9  0C                 inc c
A7BA  00                 nop
A7BB  CB 0F              rrc a
A7BD  E0                 ret po
A7BE  C0                 ret nz
A7BF  13                 inc de
A7C0  00                 nop
A7C1  8B                 adc a,e
A7C2  07                 rlca
A7C3  F0                 ret p
A7C4  C0                 ret nz
A7C5  1F                 rra
A7C6  00                 nop
A7C7  FD 07              rlca
A7C9  F0                 ret p
A7CA  C0                 ret nz
A7CB  1F                 rra
A7CC  00                 nop
A7CD  FD 07              rlca
A7CF  B0                 or b
A7D0  E0                 ret po
A7D1  0F                 rrca
A7D2  00                 nop
A7D3  39                 add hl,sp
A7D4  07                 rlca
A7D5  B0                 or b
A7D6  F0                 ret p
A7D7  01 00 E4           ld bc,$E400
A7DA  0F                 rrca
A7DB  00                 nop
A7DC  FE 00              cp 0
A7DE  00                 nop
A7DF  1D                 dec e
A7E0  1F                 rra
A7E1  C0                 ret nz
A7E2  FE 00              cp 0
A7E4  00                 nop
A7E5  FB                 ei
A7E6  0F                 rrca
A7E7  E0                 ret po
A7E8  FE 00              cp 0
A7EA  00                 nop
A7EB  FB                 ei
A7EC  07                 rlca
A7ED  F0                 ret p
A7EE  FC 01 00           call m,$0001
A7F1  E9                 jp (hl)
A7F2  07                 rlca
A7F3  F0                 ret p
A7F4  F8                 ret m
A7F5  03                 inc bc
A7F6  00                 nop
A7F7  F6 03              or $03
A7F9  F8                 ret m
A7FA  F8                 ret m
A7FB  03                 inc bc
A7FC  00                 nop
A7FD  EF                 rst 28h
A7FE  03                 inc bc
A7FF  F8                 ret m
A800  F8                 ret m
A801  02                 ld (bc),a
A802  00                 nop
A803  6B                 ld l,e
A804  07                 rlca
A805  F0                 ret p
A806  FC 01 00           call m,$0001
A809  E6 0F              and $0F
A80B  C0                 ret nz
A80C  F8                 ret m
A80D  02                 ld (bc),a
A80E  00                 nop
A80F  00                 nop
A810  3F                 ccf
A811  00                 nop
A812  F0                 ret p
A813  07                 rlca
A814  00                 nop
A815  07                 rlca
A816  1F                 rra
A817  C0                 ret nz
A818  F0                 ret p
A819  07                 rlca
A81A  00                 nop
A81B  C7                 rst 00h
A81C  0F                 rrca
A81D  E0                 ret po
A81E  E0                 ret po
A81F  0F                 rrca
A820  00                 nop
A821  87                 add a,a
A822  0F                 rrca
A823  E0                 ret po
A824  E0                 ret po
A825  0B                 dec bc
A826  30 87              jr nc,$A7AF
A828  0F                 rrca
A829  E0                 ret po
A82A  F0                 ret p
A82B  00                 nop
A82C  70                 ld (hl),b
A82D  05                 dec b
A82E  0F                 rrca
A82F  60                 ld h,b
A830  FF                 rst 38h
A831  00                 nop
A832  00                 nop
A833  00                 nop
A834  FF                 rst 38h
A835  00                 nop
A836  FA 00 00           jp m,$0000
A839  B7                 or a
A83A  3F                 ccf
A83B  00                 nop
A83C  F0                 ret p
A83D  04                 inc b
A83E  00                 nop
A83F  81                 add a,c
A840  1F                 rra
A841  C0                 ret nz
A842  E0                 ret po
A843  0C                 inc c
A844  00                 nop
A845  37                 scf
A846  0F                 rrca
A847  60                 ld h,b
A848  E0                 ret po
A849  0C                 inc c
A84A  00                 nop
A84B  CB 0F              rrc a
A84D  E0                 ret po
A84E  C0                 ret nz
A84F  13                 inc de
A850  00                 nop
A851  8B                 adc a,e
A852  07                 rlca
A853  F0                 ret p
A854  C0                 ret nz
A855  1F                 rra
A856  00                 nop
A857  FD 07              rlca
A859  F0                 ret p
A85A  C0                 ret nz
A85B  1F                 rra
A85C  00                 nop
A85D  FD 07              rlca
A85F  B0                 or b
A860  E0                 ret po
A861  0F                 rrca
A862  00                 nop
A863  39                 add hl,sp
A864  07                 rlca
A865  B0                 or b
A866  F0                 ret p
A867  01 00 E4           ld bc,$E400
A86A  0F                 rrca
A86B  00                 nop
A86C  FC 00 00           call m,$0000
A86F  1C                 inc e
A870  0F                 rrca
A871  E0                 ret po
A872  F8                 ret m
A873  02                 ld (bc),a
A874  00                 nop
A875  FD 07              rlca
A877  F0                 ret p
A878  F8                 ret m
A879  02                 ld (bc),a
A87A  00                 nop
A87B  FD 03              inc bc
A87D  F8                 ret m
A87E  F0                 ret p
A87F  05                 dec b
A880  00                 nop
A881  EE 01              xor $01
A883  7C                 ld a,h
A884  E0                 ret po
A885  0B                 dec bc
A886  00                 nop
A887  FE 01              cp 1
A889  FC C0 1B           call m,$1BC0
A88C  00                 nop
A88D  FD 01 FC C0        ld bc,$C0FC
A891  1A                 ld a,(de)
A892  00                 nop
A893  79                 ld a,c
A894  03                 inc bc
A895  78                 ld a,b
A896  E0                 ret po
A897  0D                 dec c
A898  00                 nop
A899  E0                 ret po
A89A  07                 rlca
A89B  D0                 ret nc
A89C  F0                 ret p
A89D  02                 ld (bc),a
A89E  00                 nop
A89F  03                 inc bc
A8A0  0F                 rrca
A8A1  00                 nop
A8A2  F8                 ret m
A8A3  03                 inc bc
A8A4  00                 nop
A8A5  07                 rlca
A8A6  1F                 rra
A8A7  C0                 ret nz
A8A8  FC 01 00           call m,$0001
A8AB  0F                 rrca
A8AC  1F                 rra
A8AD  C0                 ret nz
A8AE  FE 00              cp 0
A8B0  00                 nop
A8B1  8F                 adc a,a
A8B2  1F                 rra
A8B3  C0                 ret nz
A8B4  FF                 rst 38h
A8B5  00                 nop
A8B6  00                 nop
A8B7  1F                 rra
A8B8  3F                 ccf
A8B9  80                 add a,b
A8BA  FF                 rst 38h
A8BB  00                 nop
A8BC  C0                 ret nz
A8BD  15                 dec d
A8BE  3F                 ccf
A8BF  80                 add a,b
A8C0  FF                 rst 38h
A8C1  00                 nop
A8C2  00                 nop
A8C3  00                 nop
A8C4  FF                 rst 38h
A8C5  00                 nop
A8C6  FA 00 00           jp m,$0000
A8C9  B7                 or a
A8CA  3F                 ccf
A8CB  00                 nop
A8CC  F0                 ret p
A8CD  04                 inc b
A8CE  00                 nop
A8CF  81                 add a,c
A8D0  1F                 rra
A8D1  C0                 ret nz
A8D2  E0                 ret po
A8D3  0C                 inc c
A8D4  00                 nop
A8D5  37                 scf
A8D6  0F                 rrca
A8D7  60                 ld h,b
A8D8  E0                 ret po
A8D9  0C                 inc c
A8DA  00                 nop
A8DB  CB 0F              rrc a
A8DD  E0                 ret po
A8DE  C0                 ret nz
A8DF  13                 inc de
A8E0  00                 nop
A8E1  8B                 adc a,e
A8E2  07                 rlca
A8E3  F0                 ret p
A8E4  C0                 ret nz
A8E5  1F                 rra
A8E6  00                 nop
A8E7  FD 07              rlca
A8E9  F0                 ret p
A8EA  C0                 ret nz
A8EB  1F                 rra
A8EC  00                 nop
A8ED  FD 07              rlca
A8EF  B0                 or b
A8F0  E0                 ret po
A8F1  0F                 rrca
A8F2  00                 nop
A8F3  39                 add hl,sp
A8F4  07                 rlca
A8F5  B0                 or b
A8F6  F0                 ret p
A8F7  01 00 E4           ld bc,$E400
A8FA  0F                 rrca
A8FB  00                 nop
A8FC  FC 00 00           call m,$0000
A8FF  1D                 dec e
A900  1F                 rra
A901  C0                 ret nz
A902  F8                 ret m
A903  02                 ld (bc),a
A904  00                 nop
A905  FB                 ei
A906  0F                 rrca
A907  E0                 ret po
A908  FC 00 00           call m,$0000
A90B  FB                 ei
A90C  07                 rlca
A90D  F0                 ret p
A90E  FC 01 00           call m,$0001
A911  ED 07              ???
A913  F0                 ret p
A914  F8                 ret m
A915  03                 inc bc
A916  00                 nop
A917  FE 03              cp 3
A919  F8                 ret m
A91A  F8                 ret m
A91B  03                 inc bc
A91C  00                 nop
A91D  F9                 ld sp,hl
A91E  03                 inc bc
A91F  F8                 ret m
A920  F8                 ret m
A921  02                 ld (bc),a
A922  00                 nop
A923  77                 ld (hl),a
A924  07                 rlca
A925  F0                 ret p
A926  FC 01 00           call m,$0001
A929  E5                 push hl
A92A  0F                 rrca
A92B  E0                 ret po
A92C  F8                 ret m
A92D  02                 ld (bc),a
A92E  00                 nop
A92F  0B                 dec bc
A930  1F                 rra
A931  40                 ld b,b
A932  F8                 ret m
A933  03                 inc bc
A934  00                 nop
A935  1C                 inc e
A936  3F                 ccf
A937  00                 nop
A938  F8                 ret m
A939  03                 inc bc
A93A  00                 nop
A93B  DF                 rst 18h
A93C  1F                 rra
A93D  C0                 ret nz
A93E  F8                 ret m
A93F  03                 inc bc
A940  00                 nop
A941  CE 3F              adc a,63
A943  80                 add a,b
A944  F8                 ret m
A945  03                 inc bc
A946  00                 nop
A947  EB                 ex de,hl
A948  7F                 ld a,a
A949  00                 nop
A94A  F8                 ret m
A94B  02                 ld (bc),a
A94C  14                 inc d
A94D  E0                 ret po
A94E  FF                 rst 38h
A94F  00                 nop
A950  FF                 rst 38h
A951  00                 nop
A952  FF                 rst 38h
A953  00                 nop
A954  FF                 rst 38h
A955  00                 nop
A956  FF                 rst 38h
A957  00                 nop
A958  FF                 rst 38h
A959  00                 nop
A95A  FF                 rst 38h
A95B  00                 nop
A95C  FF                 rst 38h
A95D  00                 nop
A95E  FF                 rst 38h
A95F  00                 nop
A960  FF                 rst 38h
A961  00                 nop
A962  FF                 rst 38h
A963  00                 nop
A964  F0                 ret p
A965  00                 nop
A966  7F                 ld a,a
A967  00                 nop
A968  FF                 rst 38h
A969  00                 nop
A96A  E0                 ret po
A96B  0D                 dec c
A96C  1F                 rra
A96D  80                 add a,b
A96E  FF                 rst 38h
A96F  00                 nop
A970  80                 add a,b
A971  03                 inc bc
A972  0F                 rrca
A973  60                 ld h,b
A974  F0                 ret p
A975  00                 nop
A976  00                 nop
A977  42                 ld b,d
A978  07                 rlca
A979  00                 nop
A97A  E0                 ret po
A97B  0D                 dec c
A97C  00                 nop
A97D  34                 inc (hl)
A97E  03                 inc bc
A97F  78                 ld a,b
A980  E0                 ret po
A981  01 00 5A           ld bc,$5A00
A984  07                 rlca
A985  00                 nop
A986  E0                 ret po
A987  0D                 dec c
A988  00                 nop
A989  6D                 ld l,l
A98A  07                 rlca
A98B  70                 ld (hl),b
A98C  F0                 ret p
A98D  05                 dec b
A98E  00                 nop
A98F  41                 ld b,c
A990  0F                 rrca
A991  00                 nop
A992  F0                 ret p
A993  05                 dec b
A994  00                 nop
A995  42                 ld b,d
A996  0F                 rrca
A997  60                 ld h,b
A998  F8                 ret m
A999  01 00 1C           ld bc,$1C00
A99C  1F                 rra
A99D  80                 add a,b
A99E  FE 00              cp 0
A9A0  80                 add a,b
A9A1  01 0F A0           ld bc,$A00F
A9A4  FF                 rst 38h
A9A5  00                 nop
A9A6  00                 nop
A9A7  75                 ld (hl),l
A9A8  07                 rlca
A9A9  10 FF              djnz $A9AA
A9AB  00                 nop
A9AC  00                 nop
A9AD  01 07 B0           ld bc,$B007
A9B0  FF                 rst 38h
A9B1  00                 nop
A9B2  00                 nop
A9B3  6C                 ld l,h
A9B4  0F                 rrca
A9B5  E0                 ret po
A9B6  FF                 rst 38h
A9B7  00                 nop
A9B8  80                 add a,b
A9B9  18 07              jr $A9C2
A9BB  00                 nop
A9BC  FF                 rst 38h
A9BD  00                 nop
A9BE  E0                 ret po
A9BF  07                 rlca
A9C0  01 D8 FF           ld bc,-40
A9C3  00                 nop
A9C4  F8                 ret m
A9C5  00                 nop
A9C6  00                 nop
A9C7  16 FE              ld d,-2
A9C9  00                 nop
A9CA  38 02              jr c,$A9CE
A9CC  C0                 ret nz
A9CD  06 FC              ld b,-4
A9CF  01 00 40           ld bc,$4000
A9D2  80                 add a,b
A9D3  36 F8              ld (hl),-8
A9D5  02                 ld (bc),a
A9D6  00                 nop
A9D7  14                 inc d
A9D8  01 58 F8           ld bc,-1960
A9DB  02                 ld (bc),a
A9DC  00                 nop
A9DD  00                 nop
A9DE  03                 inc bc
A9DF  68                 ld l,b
A9E0  FF                 rst 38h
A9E1  00                 nop
A9E2  FF                 rst 38h
A9E3  00                 nop
A9E4  FF                 rst 38h
A9E5  00                 nop
A9E6  FF                 rst 38h
A9E7  00                 nop
A9E8  FF                 rst 38h
A9E9  00                 nop
A9EA  FF                 rst 38h
A9EB  00                 nop
A9EC  FF                 rst 38h
A9ED  00                 nop
A9EE  E0                 ret po
A9EF  00                 nop
A9F0  FF                 rst 38h
A9F1  00                 nop
A9F2  FF                 rst 38h
A9F3  00                 nop
A9F4  C0                 ret nz
A9F5  1B                 dec de
A9F6  3F                 ccf
A9F7  00                 nop
A9F8  FF                 rst 38h
A9F9  00                 nop
A9FA  00                 nop
A9FB  06 1F              ld b,31
A9FD  C0                 ret nz
A9FE  E0                 ret po
A9FF  00                 nop
AA00  00                 nop
AA01  84                 add a,h
AA02  0F                 rrca
AA03  00                 nop
AA04  C0                 ret nz
AA05  1A                 ld a,(de)
AA06  00                 nop
AA07  68                 ld l,b
AA08  07                 rlca
AA09  F0                 ret p
AA0A  C0                 ret nz
AA0B  02                 ld (bc),a
AA0C  00                 nop
AA0D  B4                 or h
AA0E  0F                 rrca
AA0F  00                 nop
AA10  C0                 ret nz
AA11  1A                 ld a,(de)
AA12  00                 nop
AA13  DA 0F E0           jp c,$E00F
AA16  E0                 ret po
AA17  0A                 ld a,(bc)
AA18  00                 nop
AA19  82                 add a,d
AA1A  1F                 rra
AA1B  00                 nop
AA1C  E0                 ret po
AA1D  0A                 ld a,(bc)
AA1E  30 84              jr nc,$A9A4
AA20  1F                 rra
AA21  C0                 ret nz
AA22  F0                 ret p
AA23  02                 ld (bc),a
AA24  60                 ld h,b
AA25  00                 nop
AA26  3F                 ccf
AA27  00                 nop
AA28  F9                 ld sp,hl
AA29  00                 nop
AA2A  C0                 ret nz
AA2B  1C                 inc e
AA2C  1F                 rra
AA2D  80                 add a,b
AA2E  FF                 rst 38h
AA2F  00                 nop
AA30  80                 add a,b
AA31  01 0F A0           ld bc,$A00F
AA34  FF                 rst 38h
AA35  00                 nop
AA36  00                 nop
AA37  75                 ld (hl),l
AA38  07                 rlca
AA39  10 FF              djnz $AA3A
AA3B  00                 nop
AA3C  00                 nop
AA3D  01 07 B0           ld bc,$B007
AA40  FF                 rst 38h
AA41  00                 nop
AA42  00                 nop
AA43  6C                 ld l,h
AA44  0F                 rrca
AA45  E0                 ret po
AA46  FF                 rst 38h
AA47  00                 nop
AA48  80                 add a,b
AA49  16 1F              ld d,31
AA4B  00                 nop
AA4C  FF                 rst 38h
AA4D  00                 nop
AA4E  E0                 ret po
AA4F  0C                 inc c
AA50  3F                 ccf
AA51  80                 add a,b
AA52  FF                 rst 38h
AA53  00                 nop
AA54  F0                 ret p
AA55  02                 ld (bc),a
AA56  3F                 ccf
AA57  00                 nop
AA58  FF                 rst 38h
AA59  00                 nop
AA5A  80                 add a,b
AA5B  06 0F              ld b,15
AA5D  40                 ld b,b
AA5E  FF                 rst 38h
AA5F  00                 nop
AA60  00                 nop
AA61  76                 halt
AA62  07                 rlca
AA63  10 FE              djnz $AA63
AA65  00                 nop
AA66  00                 nop
AA67  BB                 cp e
AA68  0F                 rrca
AA69  00                 nop
AA6A  FE 00              cp 0
AA6C  00                 nop
AA6D  DD 0F              rrca
AA6F  20 FF              jr nz,$AA70
AA71  00                 nop
AA72  FF                 rst 38h
AA73  00                 nop
AA74  FF                 rst 38h
AA75  00                 nop
AA76  FF                 rst 38h
AA77  00                 nop
AA78  FF                 rst 38h
AA79  00                 nop
AA7A  FF                 rst 38h
AA7B  00                 nop
AA7C  FF                 rst 38h
AA7D  00                 nop
AA7E  FF                 rst 38h
AA7F  00                 nop
AA80  FF                 rst 38h
AA81  00                 nop
AA82  FF                 rst 38h
AA83  00                 nop
AA84  F0                 ret p
AA85  00                 nop
AA86  7F                 ld a,a
AA87  00                 nop
AA88  FF                 rst 38h
AA89  00                 nop
AA8A  E0                 ret po
AA8B  0D                 dec c
AA8C  1F                 rra
AA8D  80                 add a,b
AA8E  FF                 rst 38h
AA8F  00                 nop
AA90  80                 add a,b
AA91  03                 inc bc
AA92  0F                 rrca
AA93  60                 ld h,b
AA94  F0                 ret p
AA95  00                 nop
AA96  00                 nop
AA97  42                 ld b,d
AA98  07                 rlca
AA99  00                 nop
AA9A  E0                 ret po
AA9B  0D                 dec c
AA9C  00                 nop
AA9D  34                 inc (hl)
AA9E  03                 inc bc
AA9F  78                 ld a,b
AAA0  E0                 ret po
AAA1  01 00 5A           ld bc,$5A00
AAA4  07                 rlca
AAA5  00                 nop
AAA6  E0                 ret po
AAA7  0D                 dec c
AAA8  00                 nop
AAA9  6D                 ld l,l
AAAA  07                 rlca
AAAB  70                 ld (hl),b
AAAC  F0                 ret p
AAAD  05                 dec b
AAAE  00                 nop
AAAF  41                 ld b,c
AAB0  0F                 rrca
AAB1  00                 nop
AAB2  F0                 ret p
AAB3  05                 dec b
AAB4  00                 nop
AAB5  42                 ld b,d
AAB6  0F                 rrca
AAB7  60                 ld h,b
AAB8  F8                 ret m
AAB9  01 00 1C           ld bc,$1C00
AABC  1F                 rra
AABD  80                 add a,b
AABE  FE 00              cp 0
AAC0  80                 add a,b
AAC1  01 0F A0           ld bc,$A00F
AAC4  FF                 rst 38h
AAC5  00                 nop
AAC6  00                 nop
AAC7  75                 ld (hl),l
AAC8  07                 rlca
AAC9  10 FF              djnz $AACA
AACB  00                 nop
AACC  00                 nop
AACD  01 07 B0           ld bc,$B007
AAD0  FF                 rst 38h
AAD1  00                 nop
AAD2  00                 nop
AAD3  6C                 ld l,h
AAD4  0F                 rrca
AAD5  E0                 ret po
AAD6  FC 00 80           call m,$8000
AAD9  02                 ld (bc),a
AADA  1F                 rra
AADB  00                 nop
AADC  F8                 ret m
AADD  03                 inc bc
AADE  40                 ld b,b
AADF  18 0F              jr $AAF0
AAE1  20 F0              jr nz,$AAD3
AAE3  06 07              ld b,7
AAE5  A0                 and b
AAE6  07                 rlca
AAE7  80                 add a,b
AAE8  F0                 ret p
AAE9  05                 dec b
AAEA  07                 rlca
AAEB  90                 sub b
AAEC  C3 08 F8           jp $F808
AAEF  03                 inc bc
AAF0  0F                 rrca
AAF1  60                 ld h,b
AAF2  07                 rlca
AAF3  00                 nop
AAF4  F8                 ret m
AAF5  02                 ld (bc),a
AAF6  1E 80              ld e,-128
AAF8  07                 rlca
AAF9  00                 nop
AAFA  FC 01 0E           call m,$0E01
AAFD  E1                 pop hl
AAFE  07                 rlca
AAFF  50                 ld d,b
AB00  FF                 rst 38h
AB01  00                 nop
AB02  FF                 rst 38h
AB03  00                 nop
AB04  FF                 rst 38h
AB05  00                 nop
AB06  FF                 rst 38h
AB07  00                 nop
AB08  FF                 rst 38h
AB09  00                 nop
AB0A  FF                 rst 38h
AB0B  00                 nop
AB0C  FF                 rst 38h
AB0D  00                 nop
AB0E  FF                 rst 38h
AB0F  00                 nop
AB10  FF                 rst 38h
AB11  00                 nop
AB12  FF                 rst 38h
AB13  00                 nop
AB14  FF                 rst 38h
AB15  00                 nop
AB16  FF                 rst 38h
AB17  00                 nop
AB18  FF                 rst 38h
AB19  00                 nop
AB1A  FF                 rst 38h
AB1B  00                 nop
AB1C  FF                 rst 38h
AB1D  00                 nop
AB1E  FF                 rst 38h
AB1F  00                 nop
AB20  FF                 rst 38h
AB21  00                 nop
AB22  FF                 rst 38h
AB23  00                 nop
AB24  FF                 rst 38h
AB25  00                 nop
AB26  FF                 rst 38h
AB27  00                 nop
AB28  FF                 rst 38h
AB29  00                 nop
AB2A  FF                 rst 38h
AB2B  00                 nop
AB2C  FF                 rst 38h
AB2D  00                 nop
AB2E  FF                 rst 38h
AB2F  00                 nop
AB30  FF                 rst 38h
AB31  00                 nop
AB32  FF                 rst 38h
AB33  00                 nop
AB34  FF                 rst 38h
AB35  00                 nop
AB36  F7                 rst 30h
AB37  00                 nop
AB38  FF                 rst 38h
AB39  00                 nop
AB3A  FF                 rst 38h
AB3B  00                 nop
AB3C  C3 08 FF           jp $FF08
AB3F  00                 nop
AB40  FF                 rst 38h
AB41  00                 nop
AB42  00                 nop
AB43  3C                 inc a
AB44  00                 nop
AB45  00                 nop
AB46  00                 nop
AB47  00                 nop
AB48  00                 nop
AB49  FD 00              nop
AB4B  FF                 rst 38h
AB4C  00                 nop
AB4D  FF                 rst 38h
AB4E  00                 nop
AB4F  3C                 inc a
AB50  00                 nop
AB51  00                 nop
AB52  00                 nop
AB53  00                 nop
AB54  C3 08 FF           jp $FF08
AB57  00                 nop
AB58  FF                 rst 38h
AB59  00                 nop
AB5A  F7                 rst 30h
AB5B  00                 nop
AB5C  FF                 rst 38h
AB5D  00                 nop
AB5E  FF                 rst 38h
AB5F  00                 nop
AB60  FF                 rst 38h
AB61  00                 nop
AB62  FF                 rst 38h
AB63  00                 nop
AB64  FF                 rst 38h
AB65  00                 nop
AB66  FF                 rst 38h
AB67  00                 nop
AB68  FF                 rst 38h
AB69  00                 nop
AB6A  FF                 rst 38h
AB6B  00                 nop
AB6C  FF                 rst 38h
AB6D  00                 nop
AB6E  FF                 rst 38h
AB6F  00                 nop
AB70  FF                 rst 38h
AB71  00                 nop
AB72  FF                 rst 38h
AB73  00                 nop
AB74  FF                 rst 38h
AB75  00                 nop
AB76  FF                 rst 38h
AB77  00                 nop
AB78  FF                 rst 38h
AB79  00                 nop
AB7A  FF                 rst 38h
AB7B  00                 nop
AB7C  FF                 rst 38h
AB7D  00                 nop
AB7E  FF                 rst 38h
AB7F  00                 nop
AB80  FF                 rst 38h
AB81  00                 nop
AB82  FF                 rst 38h
AB83  00                 nop
AB84  FF                 rst 38h
AB85  00                 nop
AB86  FF                 rst 38h
AB87  00                 nop
AB88  FF                 rst 38h
AB89  00                 nop
AB8A  FF                 rst 38h
AB8B  00                 nop
AB8C  FF                 rst 38h
AB8D  00                 nop
AB8E  FF                 rst 38h
AB8F  00                 nop
AB90  FF                 rst 38h
AB91  00                 nop
AB92  FF                 rst 38h
AB93  00                 nop
AB94  FF                 rst 38h
AB95  00                 nop
AB96  FF                 rst 38h
AB97  00                 nop
AB98  FF                 rst 38h
AB99  00                 nop
AB9A  FF                 rst 38h
AB9B  00                 nop
AB9C  FF                 rst 38h
AB9D  00                 nop
AB9E  FF                 rst 38h
AB9F  00                 nop
ABA0  FF                 rst 38h
ABA1  00                 nop
ABA2  FF                 rst 38h
ABA3  00                 nop
ABA4  FF                 rst 38h
ABA5  00                 nop
ABA6  FF                 rst 38h
ABA7  00                 nop
ABA8  FF                 rst 38h
ABA9  00                 nop
ABAA  FF                 rst 38h
ABAB  00                 nop
ABAC  FF                 rst 38h
ABAD  00                 nop
ABAE  FF                 rst 38h
ABAF  00                 nop
ABB0  FF                 rst 38h
ABB1  00                 nop
ABB2  FF                 rst 38h
ABB3  00                 nop
ABB4  FF                 rst 38h
ABB5  00                 nop
ABB6  FF                 rst 38h
ABB7  00                 nop
ABB8  FF                 rst 38h
ABB9  00                 nop
ABBA  FF                 rst 38h
ABBB  00                 nop
ABBC  FF                 rst 38h
ABBD  00                 nop
ABBE  FF                 rst 38h
ABBF  00                 nop
ABC0  FF                 rst 38h
ABC1  00                 nop
ABC2  FF                 rst 38h
ABC3  00                 nop
ABC4  FF                 rst 38h
ABC5  00                 nop
ABC6  FE 00              cp 0
ABC8  1B                 dec de
ABC9  00                 nop
ABCA  BF                 cp a
ABCB  00                 nop
ABCC  FC 01 01           call m,$0101
ABCF  E4 0F 40           call po,$400F
ABD2  F8                 ret m
ABD3  03                 inc bc
ABD4  00                 nop
ABD5  FE 87              cp -121
ABD7  10 FC              djnz $ABD5
ABD9  01 00 C1           ld bc,$C100
ABDC  0F                 rrca
ABDD  40                 ld b,b
ABDE  FE 00              cp 0
ABE0  3E 00              ld a,0
ABE2  BF                 cp a
ABE3  00                 nop
ABE4  FF                 rst 38h
ABE5  00                 nop
ABE6  FF                 rst 38h
ABE7  00                 nop
ABE8  FF                 rst 38h
ABE9  00                 nop
ABEA  FF                 rst 38h
ABEB  00                 nop
ABEC  FF                 rst 38h
ABED  00                 nop
ABEE  FF                 rst 38h
ABEF  00                 nop
ABF0  FF                 rst 38h
ABF1  00                 nop
ABF2  FF                 rst 38h
ABF3  00                 nop
ABF4  FF                 rst 38h
ABF5  00                 nop
ABF6  FF                 rst 38h
ABF7  00                 nop
ABF8  FF                 rst 38h
ABF9  00                 nop
ABFA  FF                 rst 38h
ABFB  00                 nop
ABFC  FF                 rst 38h
ABFD  00                 nop
ABFE  FF                 rst 38h
ABFF  00                 nop
AC00  FF                 rst 38h
AC01  00                 nop
AC02  FF                 rst 38h
AC03  00                 nop
AC04  FF                 rst 38h
AC05  00                 nop
AC06  FF                 rst 38h
AC07  00                 nop
AC08  FF                 rst 38h
AC09  00                 nop
AC0A  FF                 rst 38h
AC0B  00                 nop
AC0C  FF                 rst 38h
AC0D  00                 nop
AC0E  FF                 rst 38h
AC0F  00                 nop
AC10  FF                 rst 38h
AC11  00                 nop
AC12  FF                 rst 38h
AC13  00                 nop
AC14  FF                 rst 38h
AC15  00                 nop
AC16  FF                 rst 38h
AC17  00                 nop
AC18  FF                 rst 38h
AC19  00                 nop
AC1A  FF                 rst 38h
AC1B  00                 nop
AC1C  FF                 rst 38h
AC1D  00                 nop
AC1E  FF                 rst 38h
AC1F  00                 nop
AC20  FF                 rst 38h
AC21  00                 nop
AC22  FF                 rst 38h
AC23  00                 nop
AC24  FF                 rst 38h
AC25  00                 nop
AC26  FF                 rst 38h
AC27  00                 nop
AC28  FF                 rst 38h
AC29  00                 nop
AC2A  FF                 rst 38h
AC2B  00                 nop
AC2C  FF                 rst 38h
AC2D  00                 nop
AC2E  FF                 rst 38h
AC2F  00                 nop
AC30  FF                 rst 38h
AC31  00                 nop
AC32  FF                 rst 38h
AC33  00                 nop
AC34  FF                 rst 38h
AC35  00                 nop
AC36  FF                 rst 38h
AC37  00                 nop
AC38  FF                 rst 38h
AC39  00                 nop
AC3A  FF                 rst 38h
AC3B  00                 nop
AC3C  FF                 rst 38h
AC3D  00                 nop
AC3E  FF                 rst 38h
AC3F  00                 nop
AC40  FF                 rst 38h
AC41  00                 nop
AC42  FF                 rst 38h
AC43  00                 nop
AC44  FF                 rst 38h
AC45  00                 nop
AC46  FF                 rst 38h
AC47  00                 nop
AC48  FF                 rst 38h
AC49  00                 nop
AC4A  FF                 rst 38h
AC4B  00                 nop
AC4C  FF                 rst 38h
AC4D  00                 nop
AC4E  FF                 rst 38h
AC4F  00                 nop
AC50  C1                 pop bc
AC51  3E C1              ld a,-63
AC53  3E C3              ld a,-61
AC55  3E 80              ld a,-128
AC57  7F                 ld a,a
AC58  00                 nop
AC59  7F                 ld a,a
AC5A  01 7F 00           ld bc,127
AC5D  7E                 ld a,(hl)
AC5E  00                 nop
AC5F  7E                 ld a,(hl)
AC60  00                 nop
AC61  7E                 ld a,(hl)
AC62  00                 nop
AC63  55                 ld d,l
AC64  00                 nop
AC65  55                 ld d,l
AC66  00                 nop
AC67  55                 ld d,l
AC68  00                 nop
AC69  00                 nop
AC6A  00                 nop
AC6B  00                 nop
AC6C  00                 nop
AC6D  00                 nop
AC6E  80                 add a,b
AC6F  17                 rla
AC70  00                 nop
AC71  F3                 di
AC72  01 97 80           ld bc,$8097
AC75  4F                 ld c,a
AC76  00                 nop
AC77  49                 ld c,c
AC78  00                 nop
AC79  CF                 rst 08h
AC7A  00                 nop
AC7B  4D                 ld c,l
AC7C  00                 nop
AC7D  A5                 and l
AC7E  00                 nop
AC7F  CD 00 37           call $3700
AC82  00                 nop
AC83  54                 ld d,h
AC84  01 37 80           ld bc,$8037
AC87  77                 ld (hl),a
AC88  00                 nop
AC89  AA                 xor d
AC8A  00                 nop
AC8B  77                 ld (hl),a
AC8C  00                 nop
AC8D  1B                 dec de
AC8E  00                 nop
AC8F  F2 00 1B           jp p,$1B00
AC92  80                 add a,b
AC93  61                 ld h,c
AC94  00                 nop
AC95  5A                 ld e,d
AC96  00                 nop
AC97  61                 ld h,c
AC98  00                 nop
AC99  78                 ld a,b
AC9A  00                 nop
AC9B  FA 01 F8           jp m,$F801
AC9E  00                 nop
AC9F  7A                 ld a,d
ACA0  00                 nop
ACA1  52                 ld d,d
ACA2  01 FA 00           ld bc,250
ACA5  73                 ld (hl),e
ACA6  00                 nop
ACA7  A2                 and d
ACA8  00                 nop
ACA9  F3                 di
ACAA  00                 nop
ACAB  01 00 9C           ld bc,$9C00
ACAE  00                 nop
ACAF  01 FF 00           ld bc,255
ACB2  FF                 rst 38h
ACB3  00                 nop
ACB4  FF                 rst 38h
ACB5  00                 nop
ACB6  FF                 rst 38h
ACB7  00                 nop
ACB8  FF                 rst 38h
ACB9  00                 nop
ACBA  C3 00 FF           jp $FF00
ACBD  00                 nop
ACBE  FE 00              cp 0
ACC0  03                 inc bc
ACC1  10 FF              djnz $ACC2
ACC3  00                 nop
ACC4  FC 01 03           call m,$0301
ACC7  C8                 ret z
ACC8  FF                 rst 38h
ACC9  00                 nop
ACCA  E4 00 07           call po,$0700
ACCD  20 FF              jr nz,$ACCE
ACCF  00                 nop
ACD0  C0                 ret nz
ACD1  1B                 dec de
ACD2  07                 rlca
ACD3  10 FF              djnz $ACD4
ACD5  00                 nop
ACD6  E0                 ret po
ACD7  09                 add hl,bc
ACD8  07                 rlca
ACD9  50                 ld d,b
ACDA  FF                 rst 38h
ACDB  00                 nop
ACDC  C0                 ret nz
ACDD  00                 nop
ACDE  07                 rlca
ACDF  D0                 ret nc
ACE0  FF                 rst 38h
ACE1  00                 nop
ACE2  80                 add a,b
ACE3  3E 0F              ld a,15
ACE5  E0                 ret po
ACE6  FF                 rst 38h
ACE7  00                 nop
ACE8  00                 nop
ACE9  7F                 ld a,a
ACEA  1F                 rra
ACEB  00                 nop
ACEC  FF                 rst 38h
ACED  00                 nop
ACEE  00                 nop
ACEF  63                 ld h,e
ACF0  1F                 rra
ACF1  40                 ld b,b
ACF2  FF                 rst 38h
ACF3  00                 nop
ACF4  80                 add a,b
ACF5  3E 0F              ld a,15
ACF7  E0                 ret po
ACF8  FC 00 00           call m,$0000
ACFB  01 07 F0           ld bc,$F007
ACFE  E0                 ret po
ACFF  00                 nop
AD00  00                 nop
AD01  03                 inc bc
AD02  04                 inc b
AD03  F0                 ret p
AD04  C0                 ret nz
AD05  00                 nop
AD06  00                 nop
AD07  FF                 rst 38h
AD08  00                 nop
AD09  C0                 ret nz
AD0A  80                 add a,b
AD0B  01 00 FF           ld bc,-256
AD0E  00                 nop
AD0F  A0                 and b
AD10  00                 nop
AD11  03                 inc bc
AD12  00                 nop
AD13  38 00              jr c,$AD15
AD15  70                 ld (hl),b
AD16  80                 add a,b
AD17  03                 inc bc
AD18  00                 nop
AD19  80                 add a,b
AD1A  00                 nop
AD1B  0E C4              ld c,-60
AD1D  38 00              jr c,$AD1F
AD1F  2F                 cpl
AD20  00                 nop
AD21  EE E8              xor $E8
AD23  3C                 inc a
AD24  00                 nop
AD25  11 03 CE           ld de,$CE03
AD28  FC 3E 00           call m,$003E
AD2B  60                 ld h,b
AD2C  03                 inc bc
AD2D  3E FC              ld a,-4
AD2F  1E 00              ld e,0
AD31  00                 nop
AD32  03                 inc bc
AD33  7E                 ld a,(hl)
AD34  FE 0F              cp 15
AD36  00                 nop
AD37  C0                 ret nz
AD38  07                 rlca
AD39  FC FF 03           call m,$03FF
AD3C  80                 add a,b
AD3D  C0                 ret nz
AD3E  0F                 rrca
AD3F  F8                 ret m
AD40  FF                 rst 38h
AD41  00                 nop
AD42  FF                 rst 38h
AD43  00                 nop
AD44  FF                 rst 38h
AD45  00                 nop
AD46  FF                 rst 38h
AD47  00                 nop
AD48  FF                 rst 38h
AD49  00                 nop
AD4A  E3                 ex (sp),hl
AD4B  00                 nop
AD4C  FF                 rst 38h
AD4D  00                 nop
AD4E  FE 00              cp 0
AD50  03                 inc bc
AD51  10 FF              djnz $AD52
AD53  00                 nop
AD54  FC 01 03           call m,$0301
AD57  C8                 ret z
AD58  FF                 rst 38h
AD59  00                 nop
AD5A  E4 00 07           call po,$0700
AD5D  20 FF              jr nz,$AD5E
AD5F  00                 nop
AD60  C0                 ret nz
AD61  1B                 dec de
AD62  07                 rlca
AD63  10 FF              djnz $AD64
AD65  00                 nop
AD66  C0                 ret nz
AD67  09                 add hl,bc
AD68  07                 rlca
AD69  50                 ld d,b
AD6A  FF                 rst 38h
AD6B  00                 nop
AD6C  C0                 ret nz
AD6D  00                 nop
AD6E  07                 rlca
AD6F  D0                 ret nc
AD70  FF                 rst 38h
AD71  00                 nop
AD72  80                 add a,b
AD73  3E 0F              ld a,15
AD75  E0                 ret po
AD76  FF                 rst 38h
AD77  00                 nop
AD78  00                 nop
AD79  7F                 ld a,a
AD7A  1F                 rra
AD7B  40                 ld b,b
AD7C  FF                 rst 38h
AD7D  00                 nop
AD7E  00                 nop
AD7F  63                 ld h,e
AD80  1F                 rra
AD81  00                 nop
AD82  FF                 rst 38h
AD83  00                 nop
AD84  80                 add a,b
AD85  3E 0F              ld a,15
AD87  E0                 ret po
AD88  FC 00 00           call m,$0000
AD8B  41                 ld b,c
AD8C  07                 rlca
AD8D  F0                 ret p
AD8E  E0                 ret po
AD8F  03                 inc bc
AD90  00                 nop
AD91  DD 04              inc b
AD93  70                 ld (hl),b
AD94  C0                 ret nz
AD95  1F                 rra
AD96  00                 nop
AD97  9F                 sbc a,a
AD98  00                 nop
AD99  7B                 ld a,e
AD9A  80                 add a,b
AD9B  3F                 ccf
AD9C  00                 nop
AD9D  3F                 ccf
AD9E  00                 nop
AD9F  B9                 cp c
ADA0  00                 nop
ADA1  70                 ld (hl),b
ADA2  00                 nop
ADA3  2F                 cpl
ADA4  00                 nop
ADA5  DF                 rst 18h
ADA6  80                 add a,b
ADA7  39                 add hl,sp
ADA8  00                 nop
ADA9  80                 add a,b
ADAA  00                 nop
ADAB  07                 rlca
ADAC  C4 11 00           call nz,$0011
ADAF  E1                 pop hl
ADB0  00                 nop
ADB1  D8                 ret c
ADB2  E8                 ret pe
ADB3  03                 inc bc
ADB4  00                 nop
ADB5  F7                 rst 30h
ADB6  03                 inc bc
ADB7  98                 sbc a,b
ADB8  FC 01 00           call m,$0001
ADBB  F2 03 38           jp p,$3803
ADBE  FC 01 00           call m,$0001
ADC1  FC 03 F8           call m,$F803
ADC4  FE 00              cp 0
ADC6  00                 nop
ADC7  7D                 ld a,l
ADC8  07                 rlca
ADC9  F0                 ret p
ADCA  FF                 rst 38h
ADCB  00                 nop
ADCC  80                 add a,b
ADCD  1D                 dec e
ADCE  0F                 rrca
ADCF  E0                 ret po
ADD0  FE 00              cp 0
ADD2  1F                 rra
ADD3  00                 nop
ADD4  FF                 rst 38h
ADD5  00                 nop
ADD6  F8                 ret m
ADD7  01 0F E0           ld bc,$E00F
ADDA  FF                 rst 38h
ADDB  00                 nop
ADDC  F0                 ret p
ADDD  06 07              ld b,7
ADDF  10 FF              djnz $ADE0
ADE1  00                 nop
ADE2  F0                 ret p
ADE3  04                 inc b
ADE4  03                 inc bc
ADE5  08                 ex af,af'
ADE6  FF                 rst 38h
ADE7  00                 nop
ADE8  F8                 ret m
ADE9  02                 ld (bc),a
ADEA  03                 inc bc
ADEB  08                 ex af,af'
ADEC  F7                 rst 30h
ADED  00                 nop
ADEE  80                 add a,b
ADEF  06 03              ld b,3
ADF1  04                 inc b
ADF2  E3                 ex (sp),hl
ADF3  08                 ex af,af'
ADF4  00                 nop
ADF5  7A                 ld a,d
ADF6  03                 inc bc
ADF7  04                 inc b
ADF8  F1                 pop af
ADF9  04                 inc b
ADFA  00                 nop
ADFB  1E 03              ld e,3
ADFD  44                 ld b,h
ADFE  F1                 pop af
ADFF  04                 inc b
AE00  80                 add a,b
AE01  3C                 inc a
AE02  03                 inc bc
AE03  C4 88 02           call nz,$0288
AE06  00                 nop
AE07  7F                 ld a,a
AE08  00                 nop
AE09  84                 add a,h
AE0A  00                 nop
AE0B  72                 ld (hl),d
AE0C  00                 nop
AE0D  7F                 ld a,a
AE0E  00                 nop
AE0F  0F                 rrca
AE10  00                 nop
AE11  FB                 ei
AE12  80                 add a,b
AE13  3C                 inc a
AE14  00                 nop
AE15  1F                 rra
AE16  00                 nop
AE17  FB                 ei
AE18  C0                 ret nz
AE19  00                 nop
AE1A  00                 nop
AE1B  3F                 ccf
AE1C  00                 nop
AE1D  FD C0              ret nz
AE1F  04                 inc b
AE20  00                 nop
AE21  7F                 ld a,a
AE22  00                 nop
AE23  FD 80              add a,b
AE25  2C                 inc l
AE26  00                 nop
AE27  3E 03              ld a,3
AE29  FC 00 56           call m,$5600
AE2C  00                 nop
AE2D  7E                 ld a,(hl)
AE2E  01 FE 80           ld bc,$80FE
AE31  2A 00 FE           ld hl,($FE00)
AE34  00                 nop
AE35  7E                 ld a,(hl)
AE36  C0                 ret nz
AE37  01 04 F0           ld bc,$F004
AE3A  00                 nop
AE3B  3E F8              ld a,-8
AE3D  03                 inc bc
AE3E  0E C0              ld c,-64
AE40  00                 nop
AE41  BC                 cp h
AE42  E0                 ret po
AE43  07                 rlca
AE44  3F                 ccf
AE45  80                 add a,b
AE46  00                 nop
AE47  5C                 ld e,h
AE48  C0                 ret nz
AE49  1F                 rra
AE4A  7F                 ld a,a
AE4B  00                 nop
AE4C  81                 add a,c
AE4D  0C                 inc c
AE4E  C0                 ret nz
AE4F  1F                 rra
AE50  7F                 ld a,a
AE51  00                 nop
AE52  80                 add a,b
AE53  3E E0              ld a,-32
AE55  0E FF              ld c,-1
AE57  00                 nop
AE58  01 7C F1           ld bc,$F17C
AE5B  00                 nop
AE5C  FF                 rst 38h
AE5D  00                 nop
AE5E  03                 inc bc
AE5F  78                 ld a,b
AE60  FE 00              cp 0
AE62  1F                 rra
AE63  00                 nop
AE64  FF                 rst 38h
AE65  00                 nop
AE66  F8                 ret m
AE67  01 0F E0           ld bc,$E00F
AE6A  FF                 rst 38h
AE6B  00                 nop
AE6C  F0                 ret p
AE6D  06 07              ld b,7
AE6F  10 FF              djnz $AE70
AE71  00                 nop
AE72  F0                 ret p
AE73  04                 inc b
AE74  03                 inc bc
AE75  08                 ex af,af'
AE76  FF                 rst 38h
AE77  00                 nop
AE78  F8                 ret m
AE79  02                 ld (bc),a
AE7A  03                 inc bc
AE7B  08                 ex af,af'
AE7C  FE 00              cp 0
AE7E  80                 add a,b
AE7F  06 01              ld b,1
AE81  04                 inc b
AE82  FE 00              cp 0
AE84  00                 nop
AE85  7A                 ld a,d
AE86  01 04 FE           ld bc,-508
AE89  00                 nop
AE8A  00                 nop
AE8B  1E 01              ld e,1
AE8D  44                 ld b,h
AE8E  FE 00              cp 0
AE90  80                 add a,b
AE91  3C                 inc a
AE92  01 C4 8C           ld bc,$8CC4
AE95  01 00 7F           ld bc,$7F00
AE98  00                 nop
AE99  84                 add a,h
AE9A  04                 inc b
AE9B  71                 ld (hl),c
AE9C  00                 nop
AE9D  7F                 ld a,a
AE9E  00                 nop
AE9F  0F                 rrca
AEA0  00                 nop
AEA1  F9                 ld sp,hl
AEA2  80                 add a,b
AEA3  3C                 inc a
AEA4  00                 nop
AEA5  1F                 rra
AEA6  00                 nop
AEA7  F9                 ld sp,hl
AEA8  C0                 ret nz
AEA9  00                 nop
AEAA  00                 nop
AEAB  3F                 ccf
AEAC  00                 nop
AEAD  BD                 cp l
AEAE  F0                 ret p
AEAF  04                 inc b
AEB0  00                 nop
AEB1  7F                 ld a,a
AEB2  00                 nop
AEB3  3D                 dec a
AEB4  F0                 ret p
AEB5  04                 inc b
AEB6  00                 nop
AEB7  3E 00              ld a,0
AEB9  FA F8 02           jp m,$02F8
AEBC  00                 nop
AEBD  7E                 ld a,(hl)
AEBE  00                 nop
AEBF  FA F8 02           jp m,$02F8
AEC2  00                 nop
AEC3  1C                 inc e
AEC4  01 F4 FC           ld bc,-780
AEC7  00                 nop
AEC8  00                 nop
AEC9  03                 inc bc
AECA  03                 inc bc
AECB  E0                 ret po
AECC  FE 00              cp 0
AECE  00                 nop
AECF  E3                 ex (sp),hl
AED0  03                 inc bc
AED1  18 FC              jr $AECF
AED3  01 00 0E           ld bc,$0E00
AED6  03                 inc bc
AED7  78                 ld a,b
AED8  FE 00              cp 0
AEDA  00                 nop
AEDB  7E                 ld a,(hl)
AEDC  07                 rlca
AEDD  E0                 ret po
AEDE  FE 00              cp 0
AEE0  01 FC 1F           ld bc,$1FFC
AEE3  00                 nop
AEE4  FF                 rst 38h
AEE5  00                 nop
AEE6  03                 inc bc
AEE7  78                 ld a,b
AEE8  FF                 rst 38h
AEE9  00                 nop
AEEA  FF                 rst 38h
AEEB  00                 nop
AEEC  87                 add a,a
AEED  00                 nop
AEEE  FF                 rst 38h
AEEF  00                 nop
AEF0  3D                 dec a
AEF1  3D                 dec a
AEF2  3D                 dec a
AEF3  3D                 dec a
AEF4  3D                 dec a
AEF5  3D                 dec a
AEF6  3D                 dec a
AEF7  3D                 dec a
AEF8  3D                 dec a
AEF9  3D                 dec a
AEFA  3D                 dec a
AEFB  3D                 dec a
AEFC  3D                 dec a
AEFD  3D                 dec a
AEFE  3D                 dec a
AEFF  3D                 dec a
AF00  B0                 or b
AF01  B0                 or b
AF02  B0                 or b
AF03  B0                 or b
AF04  B0                 or b
AF05  B0                 or b
AF06  B0                 or b
AF07  B0                 or b
AF08  B0                 or b
AF09  B0                 or b
AF0A  B0                 or b
AF0B  B0                 or b
AF0C  B0                 or b
AF0D  B0                 or b
AF0E  B0                 or b
AF0F  B0                 or b
AF10  B0                 or b
AF11  B0                 or b
AF12  B0                 or b
AF13  B0                 or b
AF14  B0                 or b
AF15  B0                 or b
AF16  B0                 or b
AF17  B0                 or b
AF18  B0                 or b
AF19  B0                 or b
AF1A  B0                 or b
AF1B  B0                 or b
AF1C  B0                 or b
AF1D  B0                 or b
AF1E  B0                 or b
AF1F  B0                 or b
AF20  B0                 or b
AF21  B0                 or b
AF22  B0                 or b
AF23  B0                 or b
AF24  B0                 or b
AF25  B0                 or b
AF26  B0                 or b
AF27  B0                 or b
AF28  B0                 or b
AF29  B0                 or b
AF2A  B0                 or b
AF2B  B0                 or b
AF2C  B0                 or b
AF2D  B0                 or b
AF2E  B0                 or b
AF2F  B0                 or b
AF30  B0                 or b
AF31  B0                 or b
AF32  B0                 or b
AF33  B0                 or b
AF34  B0                 or b
AF35  B0                 or b
AF36  B0                 or b
AF37  B0                 or b
AF38  B0                 or b
AF39  B0                 or b
AF3A  B0                 or b
AF3B  B0                 or b
AF3C  B0                 or b
AF3D  B0                 or b
AF3E  B0                 or b
AF3F  B0                 or b
AF40  B0                 or b
AF41  B0                 or b
AF42  B0                 or b
AF43  B0                 or b
AF44  B0                 or b
AF45  B0                 or b
AF46  B0                 or b
AF47  B0                 or b
AF48  B0                 or b
AF49  B0                 or b
AF4A  B0                 or b
AF4B  B0                 or b
AF4C  B0                 or b
AF4D  B0                 or b
AF4E  B0                 or b
AF4F  B0                 or b
AF50  B0                 or b
AF51  B0                 or b
AF52  B0                 or b
AF53  B0                 or b
AF54  B0                 or b
AF55  B0                 or b
AF56  B0                 or b
AF57  B0                 or b
AF58  B0                 or b
AF59  B0                 or b
AF5A  B0                 or b
AF5B  B0                 or b
AF5C  B0                 or b
AF5D  B0                 or b
AF5E  B0                 or b
AF5F  B0                 or b
AF60  B0                 or b
AF61  B0                 or b
AF62  B0                 or b
AF63  B0                 or b
AF64  B0                 or b
AF65  B0                 or b
AF66  B0                 or b
AF67  B0                 or b
AF68  B0                 or b
AF69  B0                 or b
AF6A  B0                 or b
AF6B  B0                 or b
AF6C  B0                 or b
AF6D  B0                 or b
AF6E  B0                 or b
AF6F  B0                 or b
AF70  B0                 or b
AF71  B0                 or b
AF72  B0                 or b
AF73  B0                 or b
AF74  B0                 or b
AF75  B0                 or b
AF76  B0                 or b
AF77  B0                 or b
AF78  B0                 or b
AF79  B0                 or b
AF7A  B0                 or b
AF7B  B0                 or b
AF7C  B0                 or b
AF7D  B0                 or b
AF7E  B0                 or b
AF7F  B0                 or b
AF80  B0                 or b
AF81  B0                 or b
AF82  B0                 or b
AF83  B0                 or b
AF84  B0                 or b
AF85  B0                 or b
AF86  B0                 or b
AF87  B0                 or b
AF88  B0                 or b
AF89  B0                 or b
AF8A  B0                 or b
AF8B  B0                 or b
AF8C  B0                 or b
AF8D  B0                 or b
AF8E  B0                 or b
AF8F  B0                 or b
AF90  B0                 or b
AF91  B0                 or b
AF92  B0                 or b
AF93  B0                 or b
AF94  B0                 or b
AF95  B0                 or b
AF96  B0                 or b
AF97  B0                 or b
AF98  B0                 or b
AF99  B0                 or b
AF9A  B0                 or b
AF9B  B0                 or b
AF9C  B0                 or b
AF9D  B0                 or b
AF9E  B0                 or b
AF9F  B0                 or b
AFA0  B0                 or b
AFA1  B0                 or b
AFA2  B0                 or b
AFA3  B0                 or b
AFA4  B0                 or b
AFA5  B0                 or b
AFA6  B0                 or b
AFA7  B0                 or b
AFA8  B0                 or b
AFA9  B0                 or b
AFAA  B0                 or b
AFAB  B0                 or b
AFAC  B0                 or b
AFAD  B0                 or b
AFAE  B0                 or b
AFAF  B0                 or b
AFB0  B0                 or b
AFB1  B0                 or b
AFB2  B0                 or b
AFB3  B0                 or b
AFB4  B0                 or b
AFB5  B0                 or b
AFB6  B0                 or b
AFB7  B0                 or b
AFB8  B0                 or b
AFB9  B0                 or b
AFBA  B0                 or b
AFBB  B0                 or b
AFBC  B0                 or b
AFBD  B0                 or b
AFBE  B0                 or b
AFBF  B0                 or b
AFC0  B0                 or b
AFC1  B0                 or b
AFC2  B0                 or b
AFC3  B0                 or b
AFC4  B0                 or b
AFC5  B0                 or b
AFC6  B0                 or b
AFC7  B0                 or b
AFC8  B0                 or b
AFC9  B0                 or b
AFCA  B0                 or b
AFCB  B0                 or b
AFCC  B0                 or b
AFCD  B0                 or b
AFCE  B0                 or b
AFCF  B0                 or b
AFD0  B0                 or b
AFD1  B0                 or b
AFD2  B0                 or b
AFD3  B0                 or b
AFD4  B0                 or b
AFD5  B0                 or b
AFD6  B0                 or b
AFD7  B0                 or b
AFD8  B0                 or b
AFD9  B0                 or b
AFDA  B0                 or b
AFDB  B0                 or b
AFDC  B0                 or b
AFDD  B0                 or b
AFDE  B0                 or b
AFDF  B0                 or b
AFE0  B0                 or b
AFE1  B0                 or b
AFE2  B0                 or b
AFE3  B0                 or b
AFE4  B0                 or b
AFE5  B0                 or b
AFE6  B0                 or b
AFE7  B0                 or b
AFE8  B0                 or b
AFE9  B0                 or b
AFEA  B0                 or b
AFEB  B0                 or b
AFEC  B0                 or b
AFED  B0                 or b
AFEE  B0                 or b
AFEF  B0                 or b
AFF0  B0                 or b
AFF1  B0                 or b
AFF2  B0                 or b
AFF3  B0                 or b
AFF4  B0                 or b
AFF5  B0                 or b
AFF6  B0                 or b
AFF7  B0                 or b
AFF8  B0                 or b
AFF9  B0                 or b
AFFA  B0                 or b
AFFB  B0                 or b
AFFC  B0                 or b
AFFD  B0                 or b
AFFE  B0                 or b
AFFF  B0                 or b
B000  B0                 or b
B001  2A 2A 2A           ld hl,($2A2A)
B004  2A 2A 2A           ld hl,($2A2A)
B007  2A 2A 2A           ld hl,($2A2A)
B00A  2A 2A 2A           ld hl,($2A2A)
B00D  2A 2A 2A           ld hl,($2A2A)
B010  2A 2A 2A           ld hl,($2A2A)
B013  2A 2A 2A           ld hl,($2A2A)
B016  2A 2A 2A           ld hl,($2A2A)
B019  2A 2A 2A           ld hl,($2A2A)
B01C  2A 2A 2A           ld hl,($2A2A)
B01F  2A 2A 2A           ld hl,($2A2A)
B022  2A 2A 2A           ld hl,($2A2A)
B025  2A 2A 2A           ld hl,($2A2A)
B028  2A 2A 2A           ld hl,($2A2A)
B02B  2A 2A 2A           ld hl,($2A2A)
B02E  2A 2A 2A           ld hl,($2A2A)
B031  2A 2A 2A           ld hl,($2A2A)
B034  2A 2A 2A           ld hl,($2A2A)
B037  2A 2A 2A           ld hl,($2A2A)
B03A  2A 2A 2A           ld hl,($2A2A)
B03D  2A 2A 2A           ld hl,($2A2A)
B040  2A 2A 2A           ld hl,($2A2A)
B043  2A 2A 2A           ld hl,($2A2A)
B046  2A 2A 2A           ld hl,($2A2A)
B049  2A 2A 2A           ld hl,($2A2A)
B04C  2A 2A 2A           ld hl,($2A2A)
B04F  2A 2A 2A           ld hl,($2A2A)
B052  2A 2A 2A           ld hl,($2A2A)
B055  2A 2A 2A           ld hl,($2A2A)
B058  2A 2A 2A           ld hl,($2A2A)
B05B  2A 2A 2A           ld hl,($2A2A)
B05E  2A 2A 00           ld hl,($002A)
B061  00                 nop
B062  00                 nop
B063  00                 nop
B064  00                 nop
B065  00                 nop
B066  00                 nop
B067  58                 ld e,b
B068  00                 nop
B069  00                 nop
B06A  FF                 rst 38h
B06B  FF                 rst 38h
B06C  FF                 rst 38h
B06D  FF                 rst 38h
B06E  FF                 rst 38h
B06F  FF                 rst 38h
B070  FF                 rst 38h
B071  FF                 rst 38h
B072  00                 nop
B073  00                 nop
B074  FF                 rst 38h
B075  00                 nop
B076  00                 nop
B077  00                 nop
B078  00                 nop
B079  00                 nop
B07A  00                 nop
B07B  00                 nop
B07C  00                 nop
B07D  00                 nop
B07E  00                 nop
B07F  00                 nop
B080  28 80              jr z,$B002
B082  00                 nop
B083  00                 nop
B084  FF                 rst 38h
B085  00                 nop
B086  FF                 rst 38h
B087  FF                 rst 38h
B088  00                 nop
B089  00                 nop
B08A  00                 nop
B08B  00                 nop
B08C  00                 nop
B08D  00                 nop
B08E  00                 nop
B08F  00                 nop
B090  00                 nop
B091  00                 nop
B092  00                 nop
B093  00                 nop
B094  00                 nop
B095  00                 nop
B096  00                 nop
B097  10 00              djnz $B099
B099  00                 nop
B09A  00                 nop
B09B  00                 nop
B09C  FF                 rst 38h
B09D  00                 nop
B09E  41                 ld b,c
B09F  00                 nop
B0A0  00                 nop
B0A1  00                 nop
B0A2  00                 nop
B0A3  00                 nop
B0A4  00                 nop
B0A5  00                 nop
B0A6  00                 nop
B0A7  00                 nop
B0A8  00                 nop
B0A9  00                 nop
B0AA  00                 nop
B0AB  00                 nop
B0AC  FF                 rst 38h
B0AD  00                 nop
B0AE  00                 nop
B0AF  00                 nop
B0B0  22 DF B0           ld ($B0DF),hl
B0B3  21 00 00           ld hl,0
B0B6  E3                 ex (sp),hl
B0B7  22 E8 B0           ld ($B0E8),hl
B0BA  ED 73 E2 B0        ld ($B0E2),sp
B0BE  31 24 B0           ld sp,$B024
B0C1  F5                 push af
B0C2  C5                 push bc
B0C3  D5                 push de
B0C4  CD EA B0           call $B0EA
B0C7  CD 60 B1           call $B160
B0CA  CD 6F B1           call $B16F
B0CD  3A 96 B0           ld a,($B096)
B0D0  A7                 and a
B0D1  20 08              jr nz,$B0DB
B0D3  3E 14              ld a,20
B0D5  CD 19 B2           call $B219
B0D8  CD 00 D5           call $D500
B0DB  D1                 pop de
B0DC  C1                 pop bc
B0DD  F1                 pop af
B0DE  21 FF FF           ld hl,-1
B0E1  31 FF FF           ld sp,$FFFF
B0E4  33                 inc sp
B0E5  33                 inc sp
B0E6  FB                 ei
B0E7  C3 00 00           jp $0000
B0EA  ED 5B E2 B0        ld de,($B0E2)
B0EE  7A                 ld a,d
B0EF  FE 80              cp -128
B0F1  D8                 ret c
B0F2  FE AF              cp -81
B0F4  D0                 ret nc
B0F5  CD F9 B0           call $B0F9
B0F8  13                 inc de
B0F9  6B                 ld l,e
B0FA  26 80              ld h,-128
B0FC  AF                 xor a
B0FD  AE                 xor (hl)
B0FE  24                 inc h
B0FF  AE                 xor (hl)
B100  24                 inc h
B101  AE                 xor (hl)
B102  24                 inc h
B103  AE                 xor (hl)
B104  24                 inc h
B105  AE                 xor (hl)
B106  24                 inc h
B107  AE                 xor (hl)
B108  24                 inc h
B109  AE                 xor (hl)
B10A  24                 inc h
B10B  AE                 xor (hl)
B10C  24                 inc h
B10D  AE                 xor (hl)
B10E  24                 inc h
B10F  AE                 xor (hl)
B110  24                 inc h
B111  AE                 xor (hl)
B112  24                 inc h
B113  AE                 xor (hl)
B114  24                 inc h
B115  AE                 xor (hl)
B116  24                 inc h
B117  AE                 xor (hl)
B118  24                 inc h
B119  AE                 xor (hl)
B11A  24                 inc h
B11B  AE                 xor (hl)
B11C  24                 inc h
B11D  AE                 xor (hl)
B11E  24                 inc h
B11F  AE                 xor (hl)
B120  24                 inc h
B121  AE                 xor (hl)
B122  24                 inc h
B123  AE                 xor (hl)
B124  24                 inc h
B125  AE                 xor (hl)
B126  24                 inc h
B127  AE                 xor (hl)
B128  24                 inc h
B129  AE                 xor (hl)
B12A  24                 inc h
B12B  AE                 xor (hl)
B12C  24                 inc h
B12D  AE                 xor (hl)
B12E  24                 inc h
B12F  AE                 xor (hl)
B130  24                 inc h
B131  AE                 xor (hl)
B132  24                 inc h
B133  AE                 xor (hl)
B134  24                 inc h
B135  AE                 xor (hl)
B136  24                 inc h
B137  AE                 xor (hl)
B138  24                 inc h
B139  AE                 xor (hl)
B13A  24                 inc h
B13B  AE                 xor (hl)
B13C  24                 inc h
B13D  AE                 xor (hl)
B13E  24                 inc h
B13F  AE                 xor (hl)
B140  24                 inc h
B141  AE                 xor (hl)
B142  24                 inc h
B143  AE                 xor (hl)
B144  24                 inc h
B145  AE                 xor (hl)
B146  24                 inc h
B147  AE                 xor (hl)
B148  24                 inc h
B149  AE                 xor (hl)
B14A  24                 inc h
B14B  AE                 xor (hl)
B14C  24                 inc h
B14D  AE                 xor (hl)
B14E  24                 inc h
B14F  AE                 xor (hl)
B150  24                 inc h
B151  AE                 xor (hl)
B152  24                 inc h
B153  AE                 xor (hl)
B154  24                 inc h
B155  AE                 xor (hl)
B156  24                 inc h
B157  AE                 xor (hl)
B158  24                 inc h
B159  AE                 xor (hl)
B15A  24                 inc h
B15B  26 E7              ld h,-25
B15D  AE                 xor (hl)
B15E  12                 ld (de),a
B15F  C9                 ret
B160  01 FE FE           ld bc,-258
B163  21 6A B0           ld hl,$B06A
B166  ED 78              in a,(c)
B168  77                 ld (hl),a
B169  23                 inc hl
B16A  CB 00              rlc b
B16C  38 F8              jr c,$B166
B16E  C9                 ret
B16F  11 1D 79           ld de,$791D
B172  3A FE EA           ld a,($EAFE)
B175  CD 0A B3           call $B30A
B178  E5                 push hl
B179  7C                 ld a,h
B17A  FE FF              cp -1
B17C  C4 9D B1           call nz,$B19D
B17F  32 E1 7F           ld ($7FE1),a
B182  E1                 pop hl
B183  7D                 ld a,l
B184  CD 9D B1           call $B19D
B187  32 D5 7F           ld ($7FD5),a
B18A  C9                 ret
B18B  CE B1              adc a,-79
B18D  D8                 ret c
B18E  B1                 or c
B18F  C4 B1 00           call nz,$00B1
B192  00                 nop
B193  EC B1 E2           call pe,$E2B1
B196  B1                 or c
B197  F6 B1              or $B1
B199  00                 nop
B19A  B2                 or d
B19B  0A                 ld a,(bc)
B19C  B2                 or d
B19D  11 8B B1           ld de,$B18B
B1A0  CD 0A B3           call $B30A
B1A3  7C                 ld a,h
B1A4  B5                 or l
B1A5  28 11              jr z,$B1B8
B1A7  06 08              ld b,8
B1A9  7E                 ld a,(hl)
B1AA  23                 inc hl
B1AB  DB FE              in a,($FE)
B1AD  A6                 and (hl)
B1AE  23                 inc hl
B1AF  D6 01              sub 1
B1B1  CB 10              rl b
B1B3  30 F4              jr nc,$B1A9
B1B5  78                 ld a,b
B1B6  2F                 cpl
B1B7  C9                 ret
B1B8  3A 86 B0           ld a,($B086)
B1BB  A7                 and a
B1BC  C0                 ret nz
B1BD  AF                 xor a
B1BE  DB 1F              in a,($1F)
B1C0  2F                 cpl
B1C1  F6 E0              or $E0
B1C3  C9                 ret
B1C4  7F                 ld a,a
B1C5  01 FB 01           ld bc,507
B1C8  FD 01 DF 02        ld bc,$02DF
B1CC  DF                 rst 18h
B1CD  01 EF 01           ld bc,495
B1D0  EF                 rst 28h
B1D1  02                 ld (bc),a
B1D2  EF                 rst 28h
B1D3  04                 inc b
B1D4  EF                 rst 28h
B1D5  10 EF              djnz $B1C6
B1D7  08                 ex af,af'
B1D8  FE 02              cp 2
B1DA  FB                 ei
B1DB  01 FD 01           ld bc,509
B1DE  FB                 ei
B1DF  04                 inc b
B1E0  FB                 ei
B1E1  08                 ex af,af'
B1E2  7F                 ld a,a
B1E3  04                 inc b
B1E4  DF                 rst 18h
B1E5  08                 ex af,af'
B1E6  BF                 cp a
B1E7  08                 ex af,af'
B1E8  DF                 rst 18h
B1E9  02                 ld (bc),a
B1EA  DF                 rst 18h
B1EB  01 F7 10           ld bc,$10F7
B1EE  F7                 rst 30h
B1EF  08                 ex af,af'
B1F0  F7                 rst 30h
B1F1  04                 inc b
B1F2  F7                 rst 30h
B1F3  01 F7 02           ld bc,759
B1F6  7F                 ld a,a
B1F7  01 FB 01           ld bc,507
B1FA  FD 01 FE 08        ld bc,$08FE
B1FE  FE 10              cp 16
B200  FE 01              cp 1
B202  DF                 rst 18h
B203  01 BF 01           ld bc,447
B206  DF                 rst 18h
B207  04                 inc b
B208  DF                 rst 18h
B209  02                 ld (bc),a
B20A  7F                 ld a,a
B20B  01 FB 01           ld bc,507
B20E  FD 01 DF 01        ld bc,$01DF
B212  BF                 cp a
B213  01 3E 14           ld bc,$143E
B216  32 97 B0           ld ($B097),a
B219  01 FD 7F           ld bc,$7FFD
B21C  ED 79              out (c),a
B21E  C9                 ret
B21F  F5                 push af
B220  CB 3F              srl a
B222  CB 3F              srl a
B224  CB 3F              srl a
B226  CB 3F              srl a
B228  CD 2E B2           call $B22E
B22B  F1                 pop af
B22C  E6 0F              and $0F
B22E  FE 0A              cp 10
B230  38 02              jr c,$B234
B232  C6 07              add a,7
B234  C6 30              add a,48
B236  C3 4E B2           jp $B24E
B239  21 00 40           ld hl,$4000
B23C  0E A0              ld c,-96
B23E  E5                 push hl
B23F  06 20              ld b,32
B241  36 00              ld (hl),0
B243  23                 inc hl
B244  10 FB              djnz $B241
B246  E1                 pop hl
B247  CD 90 B2           call $B290
B24A  0D                 dec c
B24B  20 F1              jr nz,$B23E
B24D  C9                 ret
B24E  E5                 push hl
B24F  6F                 ld l,a
B250  26 00              ld h,0
B252  29                 add hl,hl
B253  29                 add hl,hl
B254  29                 add hl,hl
B255  3A 96 B0           ld a,($B096)
B258  A7                 and a
B259  20 06              jr nz,$B261
B25B  CD 14 B2           call $B214
B25E  C3 00 D6           jp $D600
B261  11 00 3C           ld de,$3C00
B264  19                 add hl,de
B265  ED 5B 60 B0        ld de,($B060)
B269  06 08              ld b,8
B26B  7E                 ld a,(hl)
B26C  4F                 ld c,a
B26D  CB 3F              srl a
B26F  B1                 or c
B270  4F                 ld c,a
B271  87                 add a,a
B272  B1                 or c
B273  12                 ld (de),a
B274  23                 inc hl
B275  14                 inc d
B276  10 F3              djnz $B26B
B278  21 60 B0           ld hl,$B060
B27B  34                 inc (hl)
B27C  E1                 pop hl
B27D  C9                 ret
B27E  7C                 ld a,h
B27F  65                 ld h,l
B280  6F                 ld l,a
B281  7C                 ld a,h
B282  E6 07              and $07
B284  0F                 rrca
B285  0F                 rrca
B286  0F                 rrca
B287  B5                 or l
B288  6F                 ld l,a
B289  7C                 ld a,h
B28A  E6 18              and $18
B28C  C6 40              add a,64
B28E  67                 ld h,a
B28F  C9                 ret
B290  24                 inc h
B291  7C                 ld a,h
B292  E6 07              and $07
B294  C0                 ret nz
B295  7D                 ld a,l
B296  C6 20              add a,32
B298  6F                 ld l,a
B299  D8                 ret c
B29A  7C                 ld a,h
B29B  D6 08              sub 8
B29D  67                 ld h,a
B29E  C9                 ret
B29F  46                 ld b,(hl)
B2A0  45                 ld b,l
B2A1  46                 ld b,(hl)
B2A2  02                 ld (bc),a
B2A3  02                 ld (bc),a
B2A4  FF                 rst 38h
B2A5  4E                 ld c,(hl)
B2A6  0E 0A              ld c,10
B2A8  58                 ld e,b
B2A9  FF                 rst 38h
B2AA  FF                 rst 38h
B2AB  4E                 ld c,(hl)
B2AC  4A                 ld c,d
B2AD  59                 ld e,c
B2AE  FF                 rst 38h
B2AF  46                 ld b,(hl)
B2B0  06 02              ld b,2
B2B2  5A                 ld e,d
B2B3  FF                 rst 38h
B2B4  FF                 rst 38h
B2B5  46                 ld b,(hl)
B2B6  42                 ld b,d
B2B7  5B                 ld e,e
B2B8  FF                 rst 38h
B2B9  45                 ld b,l
B2BA  FF                 rst 38h
B2BB  06 17              ld b,23
B2BD  CD C0 B2           call $B2C0
B2C0  2A 62 B0           ld hl,($B062)
B2C3  2C                 inc l
B2C4  22 62 B0           ld ($B062),hl
B2C7  CD 7E B2           call $B27E
B2CA  22 60 B0           ld ($B060),hl
B2CD  C9                 ret
B2CE  7E                 ld a,(hl)
B2CF  23                 inc hl
B2D0  22 64 B0           ld ($B064),hl
B2D3  F5                 push af
B2D4  E6 7F              and $7F
B2D6  11 E5 B2           ld de,$B2E5
B2D9  D5                 push de
B2DA  FE 20              cp 32
B2DC  D2 4E B2           jp nc,$B24E
B2DF  11 ED B2           ld de,$B2ED
B2E2  C3 00 B3           jp $B300
B2E5  2A 64 B0           ld hl,($B064)
B2E8  F1                 pop af
B2E9  17                 rla
B2EA  30 E2              jr nc,$B2CE
B2EC  C9                 ret
B2ED  C0                 ret nz
B2EE  B2                 or d
B2EF  BD                 cp l
B2F0  B2                 or d
B2F1  F3                 di
B2F2  B2                 or d
B2F3  2A 64 B0           ld hl,($B064)
B2F6  5E                 ld e,(hl)
B2F7  23                 inc hl
B2F8  56                 ld d,(hl)
B2F9  23                 inc hl
B2FA  22 64 B0           ld ($B064),hl
B2FD  EB                 ex de,hl
B2FE  18 C4              jr $B2C4
B300  6F                 ld l,a
B301  26 00              ld h,0
B303  29                 add hl,hl
B304  19                 add hl,de
B305  7E                 ld a,(hl)
B306  23                 inc hl
B307  66                 ld h,(hl)
B308  6F                 ld l,a
B309  E9                 jp (hl)
B30A  6F                 ld l,a
B30B  26 00              ld h,0
B30D  29                 add hl,hl
B30E  19                 add hl,de
B30F  7E                 ld a,(hl)
B310  23                 inc hl
B311  66                 ld h,(hl)
B312  6F                 ld l,a
B313  C9                 ret
B314  00                 nop
B315  00                 nop
B316  3E 01              ld a,1
B318  32 95 B0           ld ($B095),a
B31B  CD FC 77           call $77FC
B31E  CD 19 76           call $7619
B321  3E 03              ld a,3
B323  32 95 B0           ld ($B095),a
B326  2A 15 D2           ld hl,($D215)
B329  3A 3D D2           ld a,($D23D)
B32C  67                 ld h,a
B32D  22 AE B0           ld ($B0AE),hl
B330  CD 47 76           call $7647
B333  2A AE B0           ld hl,($B0AE)
B336  7C                 ld a,h
B337  32 3D D2           ld ($D23D),a
B33A  7D                 ld a,l
B33B  32 15 D2           ld ($D215),a
B33E  ED 4B 72 B0        ld bc,($B072)
B342  0C                 inc c
B343  41                 ld b,c
B344  CD 29 B5           call $B529
B347  21 DE BA           ld hl,$BADE
B34A  22 E2 B4           ld ($B4E2),hl
B34D  CD 0D B4           call $B40D
B350  DD 21 6A B0        ld ix,$B06A
B354  3A 2E 78           ld a,($782E)
B357  FE 07              cp 7
B359  20 0B              jr nz,$B366
B35B  3E FF              ld a,-1
B35D  32 15 B3           ld ($B315),a
B360  DD CB 07 4E        bit 1,(ix+7)
B364  20 17              jr nz,$B37D
B366  DD CB 03 46        bit 0,(ix+3)
B36A  CA 0A B4           jp z,$B40A
B36D  DD 7E 03           ld a,(ix+3)
B370  E6 02              and $02
B372  32 15 B3           ld ($B315),a
B375  18 06              jr $B37D
B377  DD CB 04 46        bit 0,(ix+4)
B37B  28 63              jr z,$B3E0
B37D  21 88 B0           ld hl,$B088
B380  11 87 B0           ld de,$B087
B383  1A                 ld a,(de)
B384  E6 7F              and $7F
B386  FE 04              cp 4
B388  30 07              jr nc,$B391
B38A  7E                 ld a,(hl)
B38B  A7                 and a
B38C  20 03              jr nz,$B391
B38E  3E FF              ld a,-1
B390  12                 ld (de),a
B391  35                 dec (hl)
B392  7E                 ld a,(hl)
B393  E6 03              and $03
B395  20 04              jr nz,$B39B
B397  1A                 ld a,(de)
B398  EE 80              xor $80
B39A  12                 ld (de),a
B39B  1A                 ld a,(de)
B39C  E6 7F              and $7F
B39E  FE 04              cp 4
B3A0  20 19              jr nz,$B3BB
B3A2  3A 15 B3           ld a,($B315)
B3A5  A7                 and a
B3A6  28 1C              jr z,$B3C4
B3A8  3A D5 7F           ld a,($7FD5)
B3AB  E6 10              and $10
B3AD  28 07              jr z,$B3B6
B3AF  3A E1 7F           ld a,($7FE1)
B3B2  E6 10              and $10
B3B4  20 0E              jr nz,$B3C4
B3B6  3E FF              ld a,-1
B3B8  12                 ld (de),a
B3B9  18 09              jr $B3C4
B3BB  3A 15 B3           ld a,($B315)
B3BE  A7                 and a
B3BF  20 03              jr nz,$B3C4
B3C1  3E 04              ld a,4
B3C3  12                 ld (de),a
B3C4  3A 74 B0           ld a,($B074)
B3C7  A7                 and a
B3C8  28 16              jr z,$B3E0
B3CA  3A 15 D2           ld a,($D215)
B3CD  A7                 and a
B3CE  C2 4D B3           jp nz,$B34D
B3D1  3A AC B0           ld a,($B0AC)
B3D4  A7                 and a
B3D5  20 21              jr nz,$B3F8
B3D7  3A 3D D2           ld a,($D23D)
B3DA  A7                 and a
B3DB  C2 4D B3           jp nz,$B34D
B3DE  18 18              jr $B3F8
B3E0  21 72 B0           ld hl,$B072
B3E3  34                 inc (hl)
B3E4  7E                 ld a,(hl)
B3E5  FE 04              cp 4
B3E7  DA 21 B3           jp c,$B321
B3EA  06 64              ld b,100
B3EC  C5                 push bc
B3ED  CD 0D B4           call $B40D
B3F0  CD C1 7C           call $7CC1
B3F3  C1                 pop bc
B3F4  20 02              jr nz,$B3F8
B3F6  10 F4              djnz $B3EC
B3F8  FD 21 CD 7F        ld iy,$7FCD
B3FC  3E 31              ld a,49
B3FE  CD A8 74           call $74A8
B401  FD 21 D9 7F        ld iy,$7FD9
B405  3E 32              ld a,50
B407  CD A8 74           call $74A8
B40A  C3 16 B3           jp $B316
B40D  CD B8 B6           call $B6B8
B410  CD C3 CF           call $CFC3
B413  2A 10 D2           ld hl,($D210)
B416  ED 5B 12 D2        ld de,($D212)
B41A  3A AC B0           ld a,($B0AC)
B41D  A7                 and a
B41E  20 0D              jr nz,$B42D
B420  3A 15 D2           ld a,($D215)
B423  A7                 and a
B424  20 07              jr nz,$B42D
B426  2A 38 D2           ld hl,($D238)
B429  ED 5B 3A D2        ld de,($D23A)
B42D  29                 add hl,hl
B42E  29                 add hl,hl
B42F  29                 add hl,hl
B430  29                 add hl,hl
B431  CB 3A              srl d
B433  CB 1B              rr e
B435  CB 3A              srl d
B437  CB 1B              rr e
B439  6C                 ld l,h
B43A  63                 ld h,e
B43B  22 78 B0           ld ($B078),hl
B43E  CD C1 CE           call $CEC1
B441  CD 53 BD           call $BD53
B444  2A 77 B0           ld hl,($B077)
B447  26 00              ld h,0
B449  29                 add hl,hl
B44A  29                 add hl,hl
B44B  22 7E B0           ld ($B07E),hl
B44E  2A 76 B0           ld hl,($B076)
B451  26 00              ld h,0
B453  29                 add hl,hl
B454  29                 add hl,hl
B455  29                 add hl,hl
B456  29                 add hl,hl
B457  22 7C B0           ld ($B07C),hl
B45A  DD 21 F4 CF        ld ix,$CFF4
B45E  06 4A              ld b,74
B460  DD 7E 09           ld a,(ix+9)
B463  3C                 inc a
B464  CA D5 B4           jp z,$B4D5
B467  FE 0F              cp 15
B469  38 18              jr c,$B483
B46B  ED 5B 7E B0        ld de,($B07E)
B46F  DD 6E 02           ld l,(ix+2)
B472  DD 66 03           ld h,(ix+3)
B475  A7                 and a
B476  ED 52              sbc hl,de
B478  11 06 00           ld de,6
B47B  19                 add hl,de
B47C  11 B4 FF           ld de,-76
B47F  19                 add hl,de
B480  DA D5 B4           jp c,$B4D5
B483  DD 22 AA B0        ld ($B0AA),ix
B487  C5                 push bc
B488  2A AA B0           ld hl,($B0AA)
B48B  11 D8 D2           ld de,$D2D8
B48E  ED A0              ldi
B490  ED A0              ldi
B492  ED A0              ldi
B494  ED A0              ldi
B496  ED A0              ldi
B498  ED A0              ldi
B49A  ED A0              ldi
B49C  ED A0              ldi
B49E  ED A0              ldi
B4A0  ED A0              ldi
B4A2  3A 87 B0           ld a,($B087)
B4A5  E6 7F              and $7F
B4A7  FE 04              cp 4
B4A9  3A E1 D2           ld a,($D2E1)
B4AC  11 92 EA           ld de,$EA92
B4AF  C4 00 B3           call nz,$B300
B4B2  21 D8 D2           ld hl,$D2D8
B4B5  ED 5B AA B0        ld de,($B0AA)
B4B9  ED A0              ldi
B4BB  ED A0              ldi
B4BD  ED A0              ldi
B4BF  ED A0              ldi
B4C1  ED A0              ldi
B4C3  ED A0              ldi
B4C5  ED A0              ldi
B4C7  ED A0              ldi
B4C9  ED A0              ldi
B4CB  ED A0              ldi
B4CD  CD EF B6           call $B6EF
B4D0  C1                 pop bc
B4D1  DD 2A AA B0        ld ix,($B0AA)
B4D5  11 0A 00           ld de,10
B4D8  DD 19              add ix,de
B4DA  10 84              djnz $B460
B4DC  76                 halt
B4DD  CD FF B5           call $B5FF
B4E0  F3                 di
B4E1  CD 0F BB           call $BB0F
B4E4  FB                 ei
B4E5  C9                 ret
B4E6  00                 nop
B4E7  0A                 ld a,(bc)
B4E8  10 0A              djnz $B4F4
B4EA  10 0A              djnz $B4F6
B4EC  10 0A              djnz $B4F8
B4EE  10 05              djnz $B4F5
B4F0  08                 ex af,af'
B4F1  05                 dec b
B4F2  18 0F              jr $B503
B4F4  08                 ex af,af'
B4F5  0F                 rrca
B4F6  18 00              jr $B4F8
B4F8  00                 nop
B4F9  00                 nop
B4FA  1F                 rra
B4FB  13                 inc de
B4FC  00                 nop
B4FD  13                 inc de
B4FE  1F                 rra
B4FF  06 10              ld b,16
B501  06 10              ld b,16
B503  0E 08              ld c,8
B505  0E 18              ld c,24
B507  05                 dec b
B508  10 0A              djnz $B514
B50A  08                 ex af,af'
B50B  0A                 ld a,(bc)
B50C  18 0F              jr $B51D
B50E  10 B5              djnz $B4C5
B510  CD 42 B5           call $B542
B513  21 00 58           ld hl,$5800
B516  22 21 D9           ld ($D921),hl
B519  B5                 or l
B51A  CD 4C B5           call $B54C
B51D  21 00 58           ld hl,$5800
B520  22 66 B0           ld ($B066),hl
B523  C9                 ret
B524  12                 ld (de),a
B525  1B                 dec de
B526  24                 inc h
B527  2D                 dec l
B528  36 78              ld (hl),120
B52A  32 E6 B4           ld ($B4E6),a
B52D  21 24 B5           ld hl,$B524
B530  06 00              ld b,0
B532  09                 add hl,bc
B533  7E                 ld a,(hl)
B534  32 C7 B5           ld ($B5C7),a
B537  32 D0 B5           ld ($B5D0),a
B53A  32 F6 B5           ld ($B5F6),a
B53D  21 00 D3           ld hl,$D300
B540  22 66 B0           ld ($B066),hl
B543  21 B3 B5           ld hl,$B5B3
B546  CD 4C B5           call $B54C
B549  C3 39 B2           jp $B239
B54C  22 82 B5           ld ($B582),hl
B54F  22 89 B5           ld ($B589),hl
B552  22 90 B5           ld ($B590),hl
B555  22 97 B5           ld ($B597),hl
B558  2A E6 B4           ld hl,($B4E6)
B55B  26 00              ld h,0
B55D  29                 add hl,hl
B55E  29                 add hl,hl
B55F  29                 add hl,hl
B560  11 E7 B4           ld de,$B4E7
B563  19                 add hl,de
B564  11 0F B5           ld de,$B50F
B567  01 08 00           ld bc,8
B56A  ED B0              ldir
B56C  D9                 exx
B56D  06 00              ld b,0
B56F  ED 5B 66 B0        ld de,($B066)
B573  21 00 58           ld hl,$5800
B576  0E 00              ld c,0
B578  D9                 exx
B579  2E 00              ld l,0
B57B  26 00              ld h,0
B57D  ED 5B 0F B5        ld de,($B50F)
B581  CD B3 B5           call $B5B3
B584  ED 5B 11 B5        ld de,($B511)
B588  CD B3 B5           call $B5B3
B58B  ED 5B 13 B5        ld de,($B513)
B58F  CD B3 B5           call $B5B3
B592  ED 5B 15 B5        ld de,($B515)
B596  CD B3 B5           call $B5B3
B599  D9                 exx
B59A  23                 inc hl
B59B  13                 inc de
B59C  D9                 exx
B59D  24                 inc h
B59E  7C                 ld a,h
B59F  FE 20              cp 32
B5A1  C2 7D B5           jp nz,$B57D
B5A4  2C                 inc l
B5A5  7D                 ld a,l
B5A6  FE 14              cp 20
B5A8  C2 7B B5           jp nz,$B57B
B5AB  D9                 exx
B5AC  04                 inc b
B5AD  0D                 dec c
B5AE  CA 6F B5           jp z,$B56F
B5B1  D9                 exx
B5B2  C9                 ret
B5B3  7C                 ld a,h
B5B4  92                 sub d
B5B5  30 02              jr nc,$B5B9
B5B7  ED 44              neg
B5B9  4F                 ld c,a
B5BA  7D                 ld a,l
B5BB  93                 sub e
B5BC  30 02              jr nc,$B5C0
B5BE  ED 44              neg
B5C0  81                 add a,c
B5C1  D9                 exx
B5C2  90                 sub b
B5C3  C2 CA B5           jp nz,$B5CA
B5C6  36 12              ld (hl),18
B5C8  18 0D              jr $B5D7
B5CA  3D                 dec a
B5CB  C2 D7 B5           jp nz,$B5D7
B5CE  7E                 ld a,(hl)
B5CF  FE 12              cp 18
B5D1  28 04              jr z,$B5D7
B5D3  36 7F              ld (hl),127
B5D5  0E 01              ld c,1
B5D7  D9                 exx
B5D8  C9                 ret
B5D9  7C                 ld a,h
B5DA  92                 sub d
B5DB  30 02              jr nc,$B5DF
B5DD  ED 44              neg
B5DF  4F                 ld c,a
B5E0  7D                 ld a,l
B5E1  93                 sub e
B5E2  30 02              jr nc,$B5E6
B5E4  ED 44              neg
B5E6  81                 add a,c
B5E7  D9                 exx
B5E8  90                 sub b
B5E9  C2 F0 B5           jp nz,$B5F0
B5EC  1A                 ld a,(de)
B5ED  77                 ld (hl),a
B5EE  18 0D              jr $B5FD
B5F0  3D                 dec a
B5F1  C2 FD B5           jp nz,$B5FD
B5F4  7E                 ld a,(hl)
B5F5  FE 12              cp 18
B5F7  20 04              jr nz,$B5FD
B5F9  36 7F              ld (hl),127
B5FB  0E 01              ld c,1
B5FD  D9                 exx
B5FE  C9                 ret
B5FF  21 17 0A           ld hl,$0A17
B602  CD C4 B2           call $B2C4
B605  3A CA 7F           ld a,($7FCA)
B608  CD 1F B2           call $B21F
B60B  21 15 02           ld hl,533
B60E  CD C4 B2           call $B2C4
B611  21 CD 7F           ld hl,$7FCD
B614  CD 47 B6           call $B647
B617  21 17 0E           ld hl,$0E17
B61A  CD C4 B2           call $B2C4
B61D  3A 15 D2           ld a,($D215)
B620  CD 2E B2           call $B22E
B623  21 17 11           ld hl,$1117
B626  CD C4 B2           call $B2C4
B629  3A 3D D2           ld a,($D23D)
B62C  CD 2E B2           call $B22E
B62F  21 15 18           ld hl,$1815
B632  CD C4 B2           call $B2C4
B635  21 D9 7F           ld hl,$7FD9
B638  CD 47 B6           call $B647
B63B  21 17 14           ld hl,$1417
B63E  CD C4 B2           call $B2C4
B641  3A D6 7F           ld a,($7FD6)
B644  C3 1F B2           jp $B21F
B647  3A 87 B0           ld a,($B087)
B64A  FE 10              cp 16
B64C  38 0E              jr c,$B65C
B64E  06 03              ld b,3
B650  7E                 ld a,(hl)
B651  E5                 push hl
B652  C5                 push bc
B653  CD 1F B2           call $B21F
B656  C1                 pop bc
B657  E1                 pop hl
B658  23                 inc hl
B659  10 F5              djnz $B650
B65B  C9                 ret
B65C  5F                 ld e,a
B65D  87                 add a,a
B65E  83                 add a,e
B65F  87                 add a,a
B660  5F                 ld e,a
B661  16 00              ld d,0
B663  21 6A B6           ld hl,$B66A
B666  19                 add hl,de
B667  C3 CE B2           jp $B2CE
B66A  46                 ld b,(hl)
B66B  52                 ld d,d
B66C  41                 ld b,c
B66D  4E                 ld c,(hl)
B66E  43                 ld b,e
B66F  C5                 push bc
B670  53                 ld d,e
B671  57                 ld d,a
B672  49                 ld c,c
B673  53                 ld d,e
B674  53                 ld d,e
B675  A0                 and b
B676  45                 ld b,l
B677  47                 ld b,a
B678  59                 ld e,c
B679  50                 ld d,b
B67A  54                 ld d,h
B67B  A0                 and b
B67C  41                 ld b,c
B67D  46                 ld b,(hl)
B67E  52                 ld d,d
B67F  49                 ld c,c
B680  43                 ld b,e
B681  C1                 pop bc
B682  50                 ld d,b
B683  41                 ld b,c
B684  55                 ld d,l
B685  53                 ld d,e
B686  45                 ld b,l
B687  C4 EB 2A           call nz,$2AEB
B68A  AA                 xor d
B68B  B0                 or b
B68C  01 38 D2           ld bc,$D238
B68F  DD 21 CA 7F        ld ix,$7FCA
B693  A7                 and a
B694  ED 42              sbc hl,bc
B696  EB                 ex de,hl
B697  38 04              jr c,$B69D
B699  DD 21 D6 7F        ld ix,$7FD6
B69D  7D                 ld a,l
B69E  DD 86 05           add a,(ix+5)
B6A1  27                 daa
B6A2  DD 77 05           ld (ix+5),a
B6A5  7C                 ld a,h
B6A6  DD 8E 04           adc a,(ix+4)
B6A9  27                 daa
B6AA  DD 77 04           ld (ix+4),a
B6AD  D0                 ret nc
B6AE  DD 7E 03           ld a,(ix+3)
B6B1  CE 00              adc a,0
B6B3  27                 daa
B6B4  DD 77 03           ld (ix+3),a
B6B7  C9                 ret
B6B8  3A 72 B0           ld a,($B072)
B6BB  FE 01              cp 1
B6BD  C0                 ret nz
B6BE  21 20 81           ld hl,$8120
B6C1  D9                 exx
B6C2  21 20 E7           ld hl,$E720
B6C5  D9                 exx
B6C6  06 10              ld b,16
B6C8  7E                 ld a,(hl)
B6C9  4F                 ld c,a
B6CA  57                 ld d,a
B6CB  2C                 inc l
B6CC  7E                 ld a,(hl)
B6CD  5F                 ld e,a
B6CE  2D                 dec l
B6CF  87                 add a,a
B6D0  CB 11              rl c
B6D2  CE 00              adc a,0
B6D4  87                 add a,a
B6D5  CB 11              rl c
B6D7  CE 00              adc a,0
B6D9  71                 ld (hl),c
B6DA  2C                 inc l
B6DB  77                 ld (hl),a
B6DC  2C                 inc l
B6DD  AB                 xor e
B6DE  5F                 ld e,a
B6DF  79                 ld a,c
B6E0  AA                 xor d
B6E1  D9                 exx
B6E2  AE                 xor (hl)
B6E3  77                 ld (hl),a
B6E4  2C                 inc l
B6E5  D9                 exx
B6E6  7B                 ld a,e
B6E7  D9                 exx
B6E8  AE                 xor (hl)
B6E9  77                 ld (hl),a
B6EA  2C                 inc l
B6EB  D9                 exx
B6EC  10 DA              djnz $B6C8
B6EE  C9                 ret
B6EF  3A E1 D2           ld a,($D2E1)
B6F2  FE FF              cp -1
B6F4  C8                 ret z
B6F5  FE 23              cp 35
B6F7  DA FF B6           jp c,$B6FF
B6FA  FE 26              cp 38
B6FC  DA 88 B8           jp c,$B888
B6FF  3A DC D2           ld a,($D2DC)
B702  E6 7F              and $7F
B704  FE 7F              cp 127
B706  C8                 ret z
B707  2A D8 D2           ld hl,($D2D8)
B70A  ED 5B 7C B0        ld de,($B07C)
B70E  A7                 and a
B70F  ED 52              sbc hl,de
B711  7C                 ld a,h
B712  A7                 and a
B713  C0                 ret nz
B714  7D                 ld a,l
B715  FE B7              cp -73
B717  D0                 ret nc
B718  47                 ld b,a
B719  2A DA D2           ld hl,($D2DA)
B71C  ED 5B 7E B0        ld de,($B07E)
B720  A7                 and a
B721  ED 52              sbc hl,de
B723  11 06 00           ld de,6
B726  19                 add hl,de
B727  7C                 ld a,h
B728  A7                 and a
B729  C0                 ret nz
B72A  7D                 ld a,l
B72B  FE 46              cp 70
B72D  D0                 ret nc
B72E  CB 3F              srl a
B730  4F                 ld c,a
B731  ED 73 85 B8        ld ($B885),sp
B735  21 11 B8           ld hl,$B811
B738  3A DA D2           ld a,($D2DA)
B73B  E6 01              and $01
B73D  28 03              jr z,$B742
B73F  21 17 B8           ld hl,$B817
B742  22 0F B8           ld ($B80F),hl
B745  21 F2 B7           ld hl,$B7F2
B748  3A DC D2           ld a,($D2DC)
B74B  17                 rla
B74C  30 03              jr nc,$B751
B74E  21 F9 B7           ld hl,$B7F9
B751  22 F0 B7           ld ($B7F0),hl
B754  21 77 77           ld hl,$7777
B757  11 77 77           ld de,$7777
B75A  79                 ld a,c
B75B  D6 03              sub 3
B75D  D2 71 B7           jp nc,$B771
B760  26 7E              ld h,126
B762  3C                 inc a
B763  CA 83 B7           jp z,$B783
B766  2E 7E              ld l,126
B768  3C                 inc a
B769  CA 83 B7           jp z,$B783
B76C  16 7E              ld d,126
B76E  C3 83 B7           jp $B783
B771  D6 1D              sub 29
B773  DA 83 B7           jp c,$B783
B776  1E 7E              ld e,126
B778  CA 83 B7           jp z,$B783
B77B  16 7E              ld d,126
B77D  3D                 dec a
B77E  CA 83 B7           jp z,$B783
B781  2E 7E              ld l,126
B783  7C                 ld a,h
B784  32 61 B8           ld ($B861),a
B787  7D                 ld a,l
B788  32 69 B8           ld ($B869),a
B78B  7A                 ld a,d
B78C  32 73 B8           ld ($B873),a
B78F  7B                 ld a,e
B790  32 7B B8           ld ($B87B),a
B793  3A DC D2           ld a,($D2DC)
B796  E6 7F              and $7F
B798  6F                 ld l,a
B799  26 00              ld h,0
B79B  29                 add hl,hl
B79C  29                 add hl,hl
B79D  29                 add hl,hl
B79E  29                 add hl,hl
B79F  54                 ld d,h
B7A0  5D                 ld e,l
B7A1  29                 add hl,hl
B7A2  29                 add hl,hl
B7A3  29                 add hl,hl
B7A4  19                 add hl,de
B7A5  11 40 92           ld de,$9240
B7A8  19                 add hl,de
B7A9  D9                 exx
B7AA  06 00              ld b,0
B7AC  D9                 exx
B7AD  78                 ld a,b
B7AE  D6 17              sub 23
B7B0  47                 ld b,a
B7B1  D2 C6 B7           jp nc,$B7C6
B7B4  ED 44              neg
B7B6  D9                 exx
B7B7  47                 ld b,a
B7B8  D9                 exx
B7B9  87                 add a,a
B7BA  5F                 ld e,a
B7BB  87                 add a,a
B7BC  83                 add a,e
B7BD  5F                 ld e,a
B7BE  16 00              ld d,0
B7C0  19                 add hl,de
B7C1  06 00              ld b,0
B7C3  C3 CF B7           jp $B7CF
B7C6  D6 89              sub -119
B7C8  DA CF B7           jp c,$B7CF
B7CB  3C                 inc a
B7CC  D9                 exx
B7CD  47                 ld b,a
B7CE  D9                 exx
B7CF  F9                 ld sp,hl
B7D0  68                 ld l,b
B7D1  26 00              ld h,0
B7D3  29                 add hl,hl
B7D4  29                 add hl,hl
B7D5  29                 add hl,hl
B7D6  29                 add hl,hl
B7D7  29                 add hl,hl
B7D8  06 00              ld b,0
B7DA  09                 add hl,bc
B7DB  01 FD D2           ld bc,$D2FD
B7DE  09                 add hl,bc
B7DF  22 E4 B7           ld ($B7E4),hl
B7E2  DD 21 FF FF        ld ix,$FFFF
B7E6  D9                 exx
B7E7  3E 18              ld a,24
B7E9  90                 sub b
B7EA  47                 ld b,a
B7EB  11 20 00           ld de,32
B7EE  D9                 exx
B7EF  C3 F2 B7           jp $B7F2
B7F2  C1                 pop bc
B7F3  D1                 pop de
B7F4  FD E1              pop iy
B7F6  C3 0E B8           jp $B80E
B7F9  26 E8              ld h,-24
B7FB  C1                 pop bc
B7FC  68                 ld l,b
B7FD  46                 ld b,(hl)
B7FE  FD 60              ld iyh,b
B800  69                 ld l,c
B801  4E                 ld c,(hl)
B802  FD 69              ld iyl,c
B804  D1                 pop de
B805  6B                 ld l,e
B806  5E                 ld e,(hl)
B807  6A                 ld l,d
B808  56                 ld d,(hl)
B809  C1                 pop bc
B80A  69                 ld l,c
B80B  4E                 ld c,(hl)
B80C  68                 ld l,b
B80D  46                 ld b,(hl)
B80E  C3 11 B8           jp $B811
B811  21 FF 00           ld hl,255
B814  C3 5B B8           jp $B85B
B817  21 FF 00           ld hl,255
B81A  FD 7D              ld a,iyl
B81C  37                 scf
B81D  CB 19              rr c
B81F  CB 1B              rr e
B821  1F                 rra
B822  CB 1D              rr l
B824  CB 19              rr c
B826  CB 1B              rr e
B828  1F                 rra
B829  CB 1D              rr l
B82B  CB 19              rr c
B82D  CB 1B              rr e
B82F  1F                 rra
B830  CB 1D              rr l
B832  CB 19              rr c
B834  CB 1B              rr e
B836  1F                 rra
B837  CB 1D              rr l
B839  FD 6F              ld iyl,a
B83B  FD 7C              ld a,iyh
B83D  CB 38              srl b
B83F  CB 1A              rr d
B841  1F                 rra
B842  CB 1C              rr h
B844  CB 38              srl b
B846  CB 1A              rr d
B848  1F                 rra
B849  CB 1C              rr h
B84B  CB 38              srl b
B84D  CB 1A              rr d
B84F  1F                 rra
B850  CB 1C              rr h
B852  CB 38              srl b
B854  CB 1A              rr d
B856  1F                 rra
B857  CB 1C              rr h
B859  FD 67              ld iyh,a
B85B  DD 7E 00           ld a,(ix+0)
B85E  A1                 and c
B85F  B0                 or b
B860  DD 77 00           ld (ix+0),a
B863  DD 7E 01           ld a,(ix+1)
B866  A3                 and e
B867  B2                 or d
B868  DD 77 01           ld (ix+1),a
B86B  DD 7E 02           ld a,(ix+2)
B86E  FD A5              and iyl
B870  FD B4              or iyh
B872  DD 77 02           ld (ix+2),a
B875  DD 7E 03           ld a,(ix+3)
B878  A5                 and l
B879  B4                 or h
B87A  DD 77 03           ld (ix+3),a
B87D  D9                 exx
B87E  DD 19              add ix,de
B880  05                 dec b
B881  C2 EE B7           jp nz,$B7EE
B884  31 FF FF           ld sp,$FFFF
B887  C9                 ret
B888  21 D1 7C           ld hl,$7CD1
B88B  3A DC D2           ld a,($D2DC)
B88E  E6 7F              and $7F
B890  5F                 ld e,a
B891  16 00              ld d,0
B893  19                 add hl,de
B894  19                 add hl,de
B895  5E                 ld e,(hl)
B896  23                 inc hl
B897  56                 ld d,(hl)
B898  ED 53 8F B0        ld ($B08F),de
B89C  2A 8F B0           ld hl,($B08F)
B89F  5E                 ld e,(hl)
B8A0  23                 inc hl
B8A1  56                 ld d,(hl)
B8A2  23                 inc hl
B8A3  7A                 ld a,d
B8A4  A7                 and a
B8A5  C8                 ret z
B8A6  ED 53 89 B0        ld ($B089),de
B8AA  5E                 ld e,(hl)
B8AB  23                 inc hl
B8AC  56                 ld d,(hl)
B8AD  23                 inc hl
B8AE  ED 53 8B B0        ld ($B08B),de
B8B2  ED 53 8D B0        ld ($B08D),de
B8B6  5E                 ld e,(hl)
B8B7  16 00              ld d,0
B8B9  23                 inc hl
B8BA  4E                 ld c,(hl)
B8BB  23                 inc hl
B8BC  06 00              ld b,0
B8BE  22 8F B0           ld ($B08F),hl
B8C1  2A D8 D2           ld hl,($D2D8)
B8C4  A7                 and a
B8C5  ED 52              sbc hl,de
B8C7  22 93 B0           ld ($B093),hl
B8CA  2A DA D2           ld hl,($D2DA)
B8CD  09                 add hl,bc
B8CE  22 91 B0           ld ($B091),hl
B8D1  CD D6 B8           call $B8D6
B8D4  18 C6              jr $B89C
B8D6  ED 5B 7E B0        ld de,($B07E)
B8DA  2A 91 B0           ld hl,($B091)
B8DD  A7                 and a
B8DE  ED 52              sbc hl,de
B8E0  CB 2C              sra h
B8E2  CB 1D              rr l
B8E4  ED 5B 8C B0        ld de,($B08C)
B8E8  16 00              ld d,0
B8EA  1B                 dec de
B8EB  19                 add hl,de
B8EC  7C                 ld a,h
B8ED  A7                 and a
B8EE  C0                 ret nz
B8EF  ED 52              sbc hl,de
B8F1  11 20 00           ld de,32
B8F4  A7                 and a
B8F5  ED 52              sbc hl,de
B8F7  F0                 ret p
B8F8  19                 add hl,de
B8F9  4D                 ld c,l
B8FA  ED 5B 7C B0        ld de,($B07C)
B8FE  2A 93 B0           ld hl,($B093)
B901  A7                 and a
B902  ED 52              sbc hl,de
B904  7C                 ld a,h
B905  A7                 and a
B906  C0                 ret nz
B907  3A 8B B0           ld a,($B08B)
B90A  C6 9E              add a,-98
B90C  BD                 cp l
B90D  D8                 ret c
B90E  45                 ld b,l
B90F  79                 ld a,c
B910  A7                 and a
B911  F2 2B B9           jp p,$B92B
B914  3A 8C B0           ld a,($B08C)
B917  81                 add a,c
B918  32 8C B0           ld ($B08C),a
B91B  2A 89 B0           ld hl,($B089)
B91E  59                 ld e,c
B91F  16 FF              ld d,-1
B921  A7                 and a
B922  ED 52              sbc hl,de
B924  22 89 B0           ld ($B089),hl
B927  0E 00              ld c,0
B929  18 0D              jr $B938
B92B  3E 20              ld a,32
B92D  91                 sub c
B92E  ED 5B 8C B0        ld de,($B08C)
B932  BB                 cp e
B933  30 03              jr nc,$B938
B935  32 8C B0           ld ($B08C),a
B938  3A 8B B0           ld a,($B08B)
B93B  3D                 dec a
B93C  5F                 ld e,a
B93D  78                 ld a,b
B93E  93                 sub e
B93F  47                 ld b,a
B940  30 24              jr nc,$B966
B942  ED 44              neg
B944  6F                 ld l,a
B945  26 00              ld h,0
B947  3A 8E B0           ld a,($B08E)
B94A  CB 3F              srl a
B94C  28 05              jr z,$B953
B94E  29                 add hl,hl
B94F  CB 3F              srl a
B951  20 FB              jr nz,$B94E
B953  ED 5B 89 B0        ld de,($B089)
B957  19                 add hl,de
B958  22 89 B0           ld ($B089),hl
B95B  3A 8B B0           ld a,($B08B)
B95E  80                 add a,b
B95F  32 8B B0           ld ($B08B),a
B962  06 00              ld b,0
B964  18 0D              jr $B973
B966  3E A0              ld a,-96
B968  90                 sub b
B969  ED 5B 8B B0        ld de,($B08B)
B96D  BB                 cp e
B96E  30 03              jr nc,$B973
B970  32 8B B0           ld ($B08B),a
B973  68                 ld l,b
B974  26 00              ld h,0
B976  29                 add hl,hl
B977  29                 add hl,hl
B978  29                 add hl,hl
B979  29                 add hl,hl
B97A  29                 add hl,hl
B97B  06 00              ld b,0
B97D  09                 add hl,bc
B97E  01 00 D3           ld bc,$D300
B981  09                 add hl,bc
B982  ED 5B 8C B0        ld de,($B08C)
B986  3A 8E B0           ld a,($B08E)
B989  93                 sub e
B98A  32 C4 B9           ld ($B9C4),a
B98D  3A 8C B0           ld a,($B08C)
B990  ED 44              neg
B992  C6 20              add a,32
B994  32 BD B9           ld ($B9BD),a
B997  3A 8C B0           ld a,($B08C)
B99A  3D                 dec a
B99B  EE 03              xor $03
B99D  87                 add a,a
B99E  87                 add a,a
B99F  32 AB B9           ld ($B9AB),a
B9A2  ED 5B 89 B0        ld de,($B089)
B9A6  ED 4B 8A B0        ld bc,($B08A)
B9AA  18 00              jr $B9AC
B9AC  1A                 ld a,(de)
B9AD  77                 ld (hl),a
B9AE  13                 inc de
B9AF  23                 inc hl
B9B0  1A                 ld a,(de)
B9B1  77                 ld (hl),a
B9B2  13                 inc de
B9B3  23                 inc hl
B9B4  1A                 ld a,(de)
B9B5  77                 ld (hl),a
B9B6  13                 inc de
B9B7  23                 inc hl
B9B8  1A                 ld a,(de)
B9B9  77                 ld (hl),a
B9BA  13                 inc de
B9BB  23                 inc hl
B9BC  3E 20              ld a,32
B9BE  85                 add a,l
B9BF  6F                 ld l,a
B9C0  30 01              jr nc,$B9C3
B9C2  24                 inc h
B9C3  3E FF              ld a,-1
B9C5  83                 add a,e
B9C6  5F                 ld e,a
B9C7  30 01              jr nc,$B9CA
B9C9  14                 inc d
B9CA  10 DE              djnz $B9AA
B9CC  C9                 ret
B9CD  3A 75 B0           ld a,($B075)
B9D0  A7                 and a
B9D1  C0                 ret nz
B9D2  2A 75 B0           ld hl,($B075)
B9D5  2E 00              ld l,0
B9D7  CB 3C              srl h
B9D9  CB 1D              rr l
B9DB  11 00 EB           ld de,$EB00
B9DE  19                 add hl,de
B9DF  ED 4B 77 B0        ld bc,($B077)
B9E3  06 00              ld b,0
B9E5  09                 add hl,bc
B9E6  ED 5B 66 B0        ld de,($B066)
B9EA  D9                 exx
B9EB  06 0A              ld b,10
B9ED  D9                 exx
B9EE  E5                 push hl
B9EF  06 EA              ld b,-22
B9F1  4E                 ld c,(hl)
B9F2  0A                 ld a,(bc)
B9F3  12                 ld (de),a
B9F4  1C                 inc e
B9F5  12                 ld (de),a
B9F6  1C                 inc e
B9F7  2C                 inc l
B9F8  4E                 ld c,(hl)
B9F9  0A                 ld a,(bc)
B9FA  12                 ld (de),a
B9FB  1C                 inc e
B9FC  12                 ld (de),a
B9FD  1C                 inc e
B9FE  2C                 inc l
B9FF  4E                 ld c,(hl)
BA00  0A                 ld a,(bc)
BA01  12                 ld (de),a
BA02  1C                 inc e
BA03  12                 ld (de),a
BA04  1C                 inc e
BA05  2C                 inc l
BA06  4E                 ld c,(hl)
BA07  0A                 ld a,(bc)
BA08  12                 ld (de),a
BA09  1C                 inc e
BA0A  12                 ld (de),a
BA0B  1C                 inc e
BA0C  2C                 inc l
BA0D  4E                 ld c,(hl)
BA0E  0A                 ld a,(bc)
BA0F  12                 ld (de),a
BA10  1C                 inc e
BA11  12                 ld (de),a
BA12  1C                 inc e
BA13  2C                 inc l
BA14  4E                 ld c,(hl)
BA15  0A                 ld a,(bc)
BA16  12                 ld (de),a
BA17  1C                 inc e
BA18  12                 ld (de),a
BA19  1C                 inc e
BA1A  2C                 inc l
BA1B  4E                 ld c,(hl)
BA1C  0A                 ld a,(bc)
BA1D  12                 ld (de),a
BA1E  1C                 inc e
BA1F  12                 ld (de),a
BA20  1C                 inc e
BA21  2C                 inc l
BA22  4E                 ld c,(hl)
BA23  0A                 ld a,(bc)
BA24  12                 ld (de),a
BA25  1C                 inc e
BA26  12                 ld (de),a
BA27  1C                 inc e
BA28  2C                 inc l
BA29  4E                 ld c,(hl)
BA2A  0A                 ld a,(bc)
BA2B  12                 ld (de),a
BA2C  1C                 inc e
BA2D  12                 ld (de),a
BA2E  1C                 inc e
BA2F  2C                 inc l
BA30  4E                 ld c,(hl)
BA31  0A                 ld a,(bc)
BA32  12                 ld (de),a
BA33  1C                 inc e
BA34  12                 ld (de),a
BA35  1C                 inc e
BA36  2C                 inc l
BA37  4E                 ld c,(hl)
BA38  0A                 ld a,(bc)
BA39  12                 ld (de),a
BA3A  1C                 inc e
BA3B  12                 ld (de),a
BA3C  1C                 inc e
BA3D  2C                 inc l
BA3E  4E                 ld c,(hl)
BA3F  0A                 ld a,(bc)
BA40  12                 ld (de),a
BA41  1C                 inc e
BA42  12                 ld (de),a
BA43  1C                 inc e
BA44  2C                 inc l
BA45  4E                 ld c,(hl)
BA46  0A                 ld a,(bc)
BA47  12                 ld (de),a
BA48  1C                 inc e
BA49  12                 ld (de),a
BA4A  1C                 inc e
BA4B  2C                 inc l
BA4C  4E                 ld c,(hl)
BA4D  0A                 ld a,(bc)
BA4E  12                 ld (de),a
BA4F  1C                 inc e
BA50  12                 ld (de),a
BA51  1C                 inc e
BA52  2C                 inc l
BA53  4E                 ld c,(hl)
BA54  0A                 ld a,(bc)
BA55  12                 ld (de),a
BA56  1C                 inc e
BA57  12                 ld (de),a
BA58  1C                 inc e
BA59  2C                 inc l
BA5A  4E                 ld c,(hl)
BA5B  0A                 ld a,(bc)
BA5C  12                 ld (de),a
BA5D  1C                 inc e
BA5E  12                 ld (de),a
BA5F  1C                 inc e
BA60  2C                 inc l
BA61  E1                 pop hl
BA62  E5                 push hl
BA63  4E                 ld c,(hl)
BA64  0A                 ld a,(bc)
BA65  12                 ld (de),a
BA66  1C                 inc e
BA67  12                 ld (de),a
BA68  1C                 inc e
BA69  2C                 inc l
BA6A  4E                 ld c,(hl)
BA6B  0A                 ld a,(bc)
BA6C  12                 ld (de),a
BA6D  1C                 inc e
BA6E  12                 ld (de),a
BA6F  1C                 inc e
BA70  2C                 inc l
BA71  4E                 ld c,(hl)
BA72  0A                 ld a,(bc)
BA73  12                 ld (de),a
BA74  1C                 inc e
BA75  12                 ld (de),a
BA76  1C                 inc e
BA77  2C                 inc l
BA78  4E                 ld c,(hl)
BA79  0A                 ld a,(bc)
BA7A  12                 ld (de),a
BA7B  1C                 inc e
BA7C  12                 ld (de),a
BA7D  1C                 inc e
BA7E  2C                 inc l
BA7F  4E                 ld c,(hl)
BA80  0A                 ld a,(bc)
BA81  12                 ld (de),a
BA82  1C                 inc e
BA83  12                 ld (de),a
BA84  1C                 inc e
BA85  2C                 inc l
BA86  4E                 ld c,(hl)
BA87  0A                 ld a,(bc)
BA88  12                 ld (de),a
BA89  1C                 inc e
BA8A  12                 ld (de),a
BA8B  1C                 inc e
BA8C  2C                 inc l
BA8D  4E                 ld c,(hl)
BA8E  0A                 ld a,(bc)
BA8F  12                 ld (de),a
BA90  1C                 inc e
BA91  12                 ld (de),a
BA92  1C                 inc e
BA93  2C                 inc l
BA94  4E                 ld c,(hl)
BA95  0A                 ld a,(bc)
BA96  12                 ld (de),a
BA97  1C                 inc e
BA98  12                 ld (de),a
BA99  1C                 inc e
BA9A  2C                 inc l
BA9B  4E                 ld c,(hl)
BA9C  0A                 ld a,(bc)
BA9D  12                 ld (de),a
BA9E  1C                 inc e
BA9F  12                 ld (de),a
BAA0  1C                 inc e
BAA1  2C                 inc l
BAA2  4E                 ld c,(hl)
BAA3  0A                 ld a,(bc)
BAA4  12                 ld (de),a
BAA5  1C                 inc e
BAA6  12                 ld (de),a
BAA7  1C                 inc e
BAA8  2C                 inc l
BAA9  4E                 ld c,(hl)
BAAA  0A                 ld a,(bc)
BAAB  12                 ld (de),a
BAAC  1C                 inc e
BAAD  12                 ld (de),a
BAAE  1C                 inc e
BAAF  2C                 inc l
BAB0  4E                 ld c,(hl)
BAB1  0A                 ld a,(bc)
BAB2  12                 ld (de),a
BAB3  1C                 inc e
BAB4  12                 ld (de),a
BAB5  1C                 inc e
BAB6  2C                 inc l
BAB7  4E                 ld c,(hl)
BAB8  0A                 ld a,(bc)
BAB9  12                 ld (de),a
BABA  1C                 inc e
BABB  12                 ld (de),a
BABC  1C                 inc e
BABD  2C                 inc l
BABE  4E                 ld c,(hl)
BABF  0A                 ld a,(bc)
BAC0  12                 ld (de),a
BAC1  1C                 inc e
BAC2  12                 ld (de),a
BAC3  1C                 inc e
BAC4  2C                 inc l
BAC5  4E                 ld c,(hl)
BAC6  0A                 ld a,(bc)
BAC7  12                 ld (de),a
BAC8  1C                 inc e
BAC9  12                 ld (de),a
BACA  1C                 inc e
BACB  2C                 inc l
BACC  4E                 ld c,(hl)
BACD  0A                 ld a,(bc)
BACE  12                 ld (de),a
BACF  1C                 inc e
BAD0  12                 ld (de),a
BAD1  13                 inc de
BAD2  2C                 inc l
BAD3  E1                 pop hl
BAD4  01 80 00           ld bc,128
BAD7  09                 add hl,bc
BAD8  D9                 exx
BAD9  05                 dec b
BADA  C2 ED B9           jp nz,$B9ED
BADD  C9                 ret
BADE  21 0F 40           ld hl,$400F
BAE1  D9                 exx
BAE2  21 00 D3           ld hl,$D300
BAE5  3E 08              ld a,8
BAE7  CD 48 BB           call $BB48
BAEA  D9                 exx
BAEB  26 48              ld h,72
BAED  D9                 exx
BAEE  3E 08              ld a,8
BAF0  21 00 DB           ld hl,$DB00
BAF3  CD 48 BB           call $BB48
BAF6  D9                 exx
BAF7  26 50              ld h,80
BAF9  D9                 exx
BAFA  3E 04              ld a,4
BAFC  21 00 E3           ld hl,$E300
BAFF  CD 48 BB           call $BB48
BB02  CD CD B9           call $B9CD
BB05  21 0F BB           ld hl,$BB0F
BB08  22 E2 B4           ld ($B4E2),hl
BB0B  FB                 ei
BB0C  C3 17 B5           jp $B517
BB0F  21 0F 40           ld hl,$400F
BB12  D9                 exx
BB13  21 00 D3           ld hl,$D300
BB16  3E 08              ld a,8
BB18  CD 48 BB           call $BB48
BB1B  D9                 exx
BB1C  CD CD B9           call $B9CD
BB1F  3A 96 B0           ld a,($B096)
BB22  A7                 and a
BB23  20 0D              jr nz,$BB32
BB25  3E 14              ld a,20
BB27  CD 19 B2           call $B219
BB2A  CD 0E D5           call $D50E
BB2D  3E 10              ld a,16
BB2F  CD 19 B2           call $B219
BB32  21 0F 48           ld hl,$480F
BB35  D9                 exx
BB36  3E 08              ld a,8
BB38  21 00 DB           ld hl,$DB00
BB3B  CD 48 BB           call $BB48
BB3E  D9                 exx
BB3F  21 0F 50           ld hl,$500F
BB42  D9                 exx
BB43  3E 04              ld a,4
BB45  21 00 E3           ld hl,$E300
BB48  ED 73 50 BD        ld ($BD50),sp
BB4C  32 4A BD           ld ($BD4A),a
BB4F  F9                 ld sp,hl
BB50  2E 10              ld l,16
BB52  F1                 pop af
BB53  C1                 pop bc
BB54  D1                 pop de
BB55  DD E1              pop ix
BB57  08                 ex af,af'
BB58  D9                 exx
BB59  F1                 pop af
BB5A  C1                 pop bc
BB5B  D1                 pop de
BB5C  FD E1              pop iy
BB5E  F9                 ld sp,hl
BB5F  33                 inc sp
BB60  CB E5              set 4,l
BB62  FD E5              push iy
BB64  D5                 push de
BB65  C5                 push bc
BB66  F5                 push af
BB67  08                 ex af,af'
BB68  D9                 exx
BB69  DD E5              push ix
BB6B  D5                 push de
BB6C  C5                 push bc
BB6D  F5                 push af
BB6E  F9                 ld sp,hl
BB6F  2E 20              ld l,32
BB71  F1                 pop af
BB72  C1                 pop bc
BB73  D1                 pop de
BB74  DD E1              pop ix
BB76  08                 ex af,af'
BB77  D9                 exx
BB78  F1                 pop af
BB79  C1                 pop bc
BB7A  D1                 pop de
BB7B  FD E1              pop iy
BB7D  F9                 ld sp,hl
BB7E  33                 inc sp
BB7F  FD E5              push iy
BB81  D5                 push de
BB82  C5                 push bc
BB83  F5                 push af
BB84  CB A5              res 4,l
BB86  24                 inc h
BB87  08                 ex af,af'
BB88  D9                 exx
BB89  DD E5              push ix
BB8B  D5                 push de
BB8C  C5                 push bc
BB8D  F5                 push af
BB8E  F9                 ld sp,hl
BB8F  2E 30              ld l,48
BB91  F1                 pop af
BB92  C1                 pop bc
BB93  D1                 pop de
BB94  DD E1              pop ix
BB96  08                 ex af,af'
BB97  D9                 exx
BB98  F1                 pop af
BB99  C1                 pop bc
BB9A  D1                 pop de
BB9B  FD E1              pop iy
BB9D  F9                 ld sp,hl
BB9E  33                 inc sp
BB9F  CB E5              set 4,l
BBA1  FD E5              push iy
BBA3  D5                 push de
BBA4  C5                 push bc
BBA5  F5                 push af
BBA6  08                 ex af,af'
BBA7  D9                 exx
BBA8  DD E5              push ix
BBAA  D5                 push de
BBAB  C5                 push bc
BBAC  F5                 push af
BBAD  F9                 ld sp,hl
BBAE  2E 40              ld l,64
BBB0  F1                 pop af
BBB1  C1                 pop bc
BBB2  D1                 pop de
BBB3  DD E1              pop ix
BBB5  08                 ex af,af'
BBB6  D9                 exx
BBB7  F1                 pop af
BBB8  C1                 pop bc
BBB9  D1                 pop de
BBBA  FD E1              pop iy
BBBC  F9                 ld sp,hl
BBBD  33                 inc sp
BBBE  FD E5              push iy
BBC0  D5                 push de
BBC1  C5                 push bc
BBC2  F5                 push af
BBC3  CB A5              res 4,l
BBC5  24                 inc h
BBC6  08                 ex af,af'
BBC7  D9                 exx
BBC8  DD E5              push ix
BBCA  D5                 push de
BBCB  C5                 push bc
BBCC  F5                 push af
BBCD  F9                 ld sp,hl
BBCE  2E 50              ld l,80
BBD0  F1                 pop af
BBD1  C1                 pop bc
BBD2  D1                 pop de
BBD3  DD E1              pop ix
BBD5  08                 ex af,af'
BBD6  D9                 exx
BBD7  F1                 pop af
BBD8  C1                 pop bc
BBD9  D1                 pop de
BBDA  FD E1              pop iy
BBDC  F9                 ld sp,hl
BBDD  33                 inc sp
BBDE  CB E5              set 4,l
BBE0  FD E5              push iy
BBE2  D5                 push de
BBE3  C5                 push bc
BBE4  F5                 push af
BBE5  08                 ex af,af'
BBE6  D9                 exx
BBE7  DD E5              push ix
BBE9  D5                 push de
BBEA  C5                 push bc
BBEB  F5                 push af
BBEC  F9                 ld sp,hl
BBED  2E 60              ld l,96
BBEF  F1                 pop af
BBF0  C1                 pop bc
BBF1  D1                 pop de
BBF2  DD E1              pop ix
BBF4  08                 ex af,af'
BBF5  D9                 exx
BBF6  F1                 pop af
BBF7  C1                 pop bc
BBF8  D1                 pop de
BBF9  FD E1              pop iy
BBFB  F9                 ld sp,hl
BBFC  33                 inc sp
BBFD  FD E5              push iy
BBFF  D5                 push de
BC00  C5                 push bc
BC01  F5                 push af
BC02  CB A5              res 4,l
BC04  24                 inc h
BC05  08                 ex af,af'
BC06  D9                 exx
BC07  DD E5              push ix
BC09  D5                 push de
BC0A  C5                 push bc
BC0B  F5                 push af
BC0C  F9                 ld sp,hl
BC0D  2E 70              ld l,112
BC0F  F1                 pop af
BC10  C1                 pop bc
BC11  D1                 pop de
BC12  DD E1              pop ix
BC14  08                 ex af,af'
BC15  D9                 exx
BC16  F1                 pop af
BC17  C1                 pop bc
BC18  D1                 pop de
BC19  FD E1              pop iy
BC1B  F9                 ld sp,hl
BC1C  33                 inc sp
BC1D  CB E5              set 4,l
BC1F  FD E5              push iy
BC21  D5                 push de
BC22  C5                 push bc
BC23  F5                 push af
BC24  08                 ex af,af'
BC25  D9                 exx
BC26  DD E5              push ix
BC28  D5                 push de
BC29  C5                 push bc
BC2A  F5                 push af
BC2B  F9                 ld sp,hl
BC2C  2E 80              ld l,-128
BC2E  F1                 pop af
BC2F  C1                 pop bc
BC30  D1                 pop de
BC31  DD E1              pop ix
BC33  08                 ex af,af'
BC34  D9                 exx
BC35  F1                 pop af
BC36  C1                 pop bc
BC37  D1                 pop de
BC38  FD E1              pop iy
BC3A  F9                 ld sp,hl
BC3B  33                 inc sp
BC3C  FD E5              push iy
BC3E  D5                 push de
BC3F  C5                 push bc
BC40  F5                 push af
BC41  CB A5              res 4,l
BC43  24                 inc h
BC44  08                 ex af,af'
BC45  D9                 exx
BC46  DD E5              push ix
BC48  D5                 push de
BC49  C5                 push bc
BC4A  F5                 push af
BC4B  F9                 ld sp,hl
BC4C  2E 90              ld l,-112
BC4E  F1                 pop af
BC4F  C1                 pop bc
BC50  D1                 pop de
BC51  DD E1              pop ix
BC53  08                 ex af,af'
BC54  D9                 exx
BC55  F1                 pop af
BC56  C1                 pop bc
BC57  D1                 pop de
BC58  FD E1              pop iy
BC5A  F9                 ld sp,hl
BC5B  33                 inc sp
BC5C  CB E5              set 4,l
BC5E  FD E5              push iy
BC60  D5                 push de
BC61  C5                 push bc
BC62  F5                 push af
BC63  08                 ex af,af'
BC64  D9                 exx
BC65  DD E5              push ix
BC67  D5                 push de
BC68  C5                 push bc
BC69  F5                 push af
BC6A  F9                 ld sp,hl
BC6B  2E A0              ld l,-96
BC6D  F1                 pop af
BC6E  C1                 pop bc
BC6F  D1                 pop de
BC70  DD E1              pop ix
BC72  08                 ex af,af'
BC73  D9                 exx
BC74  F1                 pop af
BC75  C1                 pop bc
BC76  D1                 pop de
BC77  FD E1              pop iy
BC79  F9                 ld sp,hl
BC7A  33                 inc sp
BC7B  FD E5              push iy
BC7D  D5                 push de
BC7E  C5                 push bc
BC7F  F5                 push af
BC80  CB A5              res 4,l
BC82  24                 inc h
BC83  08                 ex af,af'
BC84  D9                 exx
BC85  DD E5              push ix
BC87  D5                 push de
BC88  C5                 push bc
BC89  F5                 push af
BC8A  F9                 ld sp,hl
BC8B  2E B0              ld l,-80
BC8D  F1                 pop af
BC8E  C1                 pop bc
BC8F  D1                 pop de
BC90  DD E1              pop ix
BC92  08                 ex af,af'
BC93  D9                 exx
BC94  F1                 pop af
BC95  C1                 pop bc
BC96  D1                 pop de
BC97  FD E1              pop iy
BC99  F9                 ld sp,hl
BC9A  33                 inc sp
BC9B  CB E5              set 4,l
BC9D  FD E5              push iy
BC9F  D5                 push de
BCA0  C5                 push bc
BCA1  F5                 push af
BCA2  08                 ex af,af'
BCA3  D9                 exx
BCA4  DD E5              push ix
BCA6  D5                 push de
BCA7  C5                 push bc
BCA8  F5                 push af
BCA9  F9                 ld sp,hl
BCAA  2E C0              ld l,-64
BCAC  F1                 pop af
BCAD  C1                 pop bc
BCAE  D1                 pop de
BCAF  DD E1              pop ix
BCB1  08                 ex af,af'
BCB2  D9                 exx
BCB3  F1                 pop af
BCB4  C1                 pop bc
BCB5  D1                 pop de
BCB6  FD E1              pop iy
BCB8  F9                 ld sp,hl
BCB9  33                 inc sp
BCBA  FD E5              push iy
BCBC  D5                 push de
BCBD  C5                 push bc
BCBE  F5                 push af
BCBF  CB A5              res 4,l
BCC1  24                 inc h
BCC2  08                 ex af,af'
BCC3  D9                 exx
BCC4  DD E5              push ix
BCC6  D5                 push de
BCC7  C5                 push bc
BCC8  F5                 push af
BCC9  F9                 ld sp,hl
BCCA  2E D0              ld l,-48
BCCC  F1                 pop af
BCCD  C1                 pop bc
BCCE  D1                 pop de
BCCF  DD E1              pop ix
BCD1  08                 ex af,af'
BCD2  D9                 exx
BCD3  F1                 pop af
BCD4  C1                 pop bc
BCD5  D1                 pop de
BCD6  FD E1              pop iy
BCD8  F9                 ld sp,hl
BCD9  33                 inc sp
BCDA  CB E5              set 4,l
BCDC  FD E5              push iy
BCDE  D5                 push de
BCDF  C5                 push bc
BCE0  F5                 push af
BCE1  08                 ex af,af'
BCE2  D9                 exx
BCE3  DD E5              push ix
BCE5  D5                 push de
BCE6  C5                 push bc
BCE7  F5                 push af
BCE8  F9                 ld sp,hl
BCE9  2E E0              ld l,-32
BCEB  F1                 pop af
BCEC  C1                 pop bc
BCED  D1                 pop de
BCEE  DD E1              pop ix
BCF0  08                 ex af,af'
BCF1  D9                 exx
BCF2  F1                 pop af
BCF3  C1                 pop bc
BCF4  D1                 pop de
BCF5  FD E1              pop iy
BCF7  F9                 ld sp,hl
BCF8  33                 inc sp
BCF9  FD E5              push iy
BCFB  D5                 push de
BCFC  C5                 push bc
BCFD  F5                 push af
BCFE  CB A5              res 4,l
BD00  24                 inc h
BD01  08                 ex af,af'
BD02  D9                 exx
BD03  DD E5              push ix
BD05  D5                 push de
BD06  C5                 push bc
BD07  F5                 push af
BD08  F9                 ld sp,hl
BD09  2E F0              ld l,-16
BD0B  F1                 pop af
BD0C  C1                 pop bc
BD0D  D1                 pop de
BD0E  DD E1              pop ix
BD10  08                 ex af,af'
BD11  D9                 exx
BD12  F1                 pop af
BD13  C1                 pop bc
BD14  D1                 pop de
BD15  FD E1              pop iy
BD17  F9                 ld sp,hl
BD18  33                 inc sp
BD19  CB E5              set 4,l
BD1B  FD E5              push iy
BD1D  D5                 push de
BD1E  C5                 push bc
BD1F  F5                 push af
BD20  08                 ex af,af'
BD21  D9                 exx
BD22  DD E5              push ix
BD24  D5                 push de
BD25  C5                 push bc
BD26  F5                 push af
BD27  F9                 ld sp,hl
BD28  2E 00              ld l,0
BD2A  F1                 pop af
BD2B  C1                 pop bc
BD2C  D1                 pop de
BD2D  DD E1              pop ix
BD2F  08                 ex af,af'
BD30  D9                 exx
BD31  F1                 pop af
BD32  C1                 pop bc
BD33  D1                 pop de
BD34  FD E1              pop iy
BD36  F9                 ld sp,hl
BD37  33                 inc sp
BD38  FD E5              push iy
BD3A  D5                 push de
BD3B  C5                 push bc
BD3C  F5                 push af
BD3D  11 10 F9           ld de,-1776
BD40  19                 add hl,de
BD41  08                 ex af,af'
BD42  D9                 exx
BD43  DD E5              push ix
BD45  D5                 push de
BD46  C5                 push bc
BD47  F5                 push af
BD48  24                 inc h
BD49  3E 00              ld a,0
BD4B  3D                 dec a
BD4C  C2 4C BB           jp nz,$BB4C
BD4F  31 00 00           ld sp,$0000
BD52  C9                 ret
BD53  ED 73 D7 C3        ld ($C3D7),sp
BD57  21 00 D3           ld hl,$D300
BD5A  11 1F 00           ld de,31
BD5D  D9                 exx
BD5E  2A 75 B0           ld hl,($B075)
BD61  2E 00              ld l,0
BD63  CB 3C              srl h
BD65  CB 1D              rr l
BD67  11 00 EB           ld de,$EB00
BD6A  19                 add hl,de
BD6B  ED 4B 77 B0        ld bc,($B077)
BD6F  06 00              ld b,0
BD71  09                 add hl,bc
BD72  EB                 ex de,hl
BD73  01 80 0A           ld bc,$0A80
BD76  D9                 exx
BD77  D9                 exx
BD78  1A                 ld a,(de)
BD79  1C                 inc e
BD7A  2E 02              ld l,2
BD7C  CB 3F              srl a
BD7E  CB 1D              rr l
BD80  1F                 rra
BD81  CB 1D              rr l
BD83  1F                 rra
BD84  CB 1D              rr l
BD86  67                 ld h,a
BD87  F9                 ld sp,hl
BD88  D9                 exx
BD89  C1                 pop bc
BD8A  71                 ld (hl),c
BD8B  2C                 inc l
BD8C  70                 ld (hl),b
BD8D  19                 add hl,de
BD8E  C1                 pop bc
BD8F  71                 ld (hl),c
BD90  2C                 inc l
BD91  70                 ld (hl),b
BD92  19                 add hl,de
BD93  C1                 pop bc
BD94  71                 ld (hl),c
BD95  2C                 inc l
BD96  70                 ld (hl),b
BD97  19                 add hl,de
BD98  C1                 pop bc
BD99  71                 ld (hl),c
BD9A  2C                 inc l
BD9B  70                 ld (hl),b
BD9C  19                 add hl,de
BD9D  C1                 pop bc
BD9E  71                 ld (hl),c
BD9F  2C                 inc l
BDA0  70                 ld (hl),b
BDA1  19                 add hl,de
BDA2  C1                 pop bc
BDA3  71                 ld (hl),c
BDA4  2C                 inc l
BDA5  70                 ld (hl),b
BDA6  19                 add hl,de
BDA7  C1                 pop bc
BDA8  71                 ld (hl),c
BDA9  2C                 inc l
BDAA  70                 ld (hl),b
BDAB  19                 add hl,de
BDAC  C1                 pop bc
BDAD  71                 ld (hl),c
BDAE  2C                 inc l
BDAF  70                 ld (hl),b
BDB0  19                 add hl,de
BDB1  C1                 pop bc
BDB2  71                 ld (hl),c
BDB3  2C                 inc l
BDB4  70                 ld (hl),b
BDB5  19                 add hl,de
BDB6  C1                 pop bc
BDB7  71                 ld (hl),c
BDB8  2C                 inc l
BDB9  70                 ld (hl),b
BDBA  19                 add hl,de
BDBB  C1                 pop bc
BDBC  71                 ld (hl),c
BDBD  2C                 inc l
BDBE  70                 ld (hl),b
BDBF  19                 add hl,de
BDC0  C1                 pop bc
BDC1  71                 ld (hl),c
BDC2  2C                 inc l
BDC3  70                 ld (hl),b
BDC4  19                 add hl,de
BDC5  C1                 pop bc
BDC6  71                 ld (hl),c
BDC7  2C                 inc l
BDC8  70                 ld (hl),b
BDC9  19                 add hl,de
BDCA  C1                 pop bc
BDCB  71                 ld (hl),c
BDCC  2C                 inc l
BDCD  70                 ld (hl),b
BDCE  19                 add hl,de
BDCF  C1                 pop bc
BDD0  71                 ld (hl),c
BDD1  2C                 inc l
BDD2  70                 ld (hl),b
BDD3  19                 add hl,de
BDD4  C1                 pop bc
BDD5  71                 ld (hl),c
BDD6  2C                 inc l
BDD7  70                 ld (hl),b
BDD8  01 21 FE           ld bc,-479
BDDB  09                 add hl,bc
BDDC  D9                 exx
BDDD  1A                 ld a,(de)
BDDE  1C                 inc e
BDDF  2E 02              ld l,2
BDE1  CB 3F              srl a
BDE3  CB 1D              rr l
BDE5  1F                 rra
BDE6  CB 1D              rr l
BDE8  1F                 rra
BDE9  CB 1D              rr l
BDEB  67                 ld h,a
BDEC  F9                 ld sp,hl
BDED  D9                 exx
BDEE  C1                 pop bc
BDEF  71                 ld (hl),c
BDF0  2C                 inc l
BDF1  70                 ld (hl),b
BDF2  19                 add hl,de
BDF3  C1                 pop bc
BDF4  71                 ld (hl),c
BDF5  2C                 inc l
BDF6  70                 ld (hl),b
BDF7  19                 add hl,de
BDF8  C1                 pop bc
BDF9  71                 ld (hl),c
BDFA  2C                 inc l
BDFB  70                 ld (hl),b
BDFC  19                 add hl,de
BDFD  C1                 pop bc
BDFE  71                 ld (hl),c
BDFF  2C                 inc l
BE00  70                 ld (hl),b
BE01  19                 add hl,de
BE02  C1                 pop bc
BE03  71                 ld (hl),c
BE04  2C                 inc l
BE05  70                 ld (hl),b
BE06  19                 add hl,de
BE07  C1                 pop bc
BE08  71                 ld (hl),c
BE09  2C                 inc l
BE0A  70                 ld (hl),b
BE0B  19                 add hl,de
BE0C  C1                 pop bc
BE0D  71                 ld (hl),c
BE0E  2C                 inc l
BE0F  70                 ld (hl),b
BE10  19                 add hl,de
BE11  C1                 pop bc
BE12  71                 ld (hl),c
BE13  2C                 inc l
BE14  70                 ld (hl),b
BE15  19                 add hl,de
BE16  C1                 pop bc
BE17  71                 ld (hl),c
BE18  2C                 inc l
BE19  70                 ld (hl),b
BE1A  19                 add hl,de
BE1B  C1                 pop bc
BE1C  71                 ld (hl),c
BE1D  2C                 inc l
BE1E  70                 ld (hl),b
BE1F  19                 add hl,de
BE20  C1                 pop bc
BE21  71                 ld (hl),c
BE22  2C                 inc l
BE23  70                 ld (hl),b
BE24  19                 add hl,de
BE25  C1                 pop bc
BE26  71                 ld (hl),c
BE27  2C                 inc l
BE28  70                 ld (hl),b
BE29  19                 add hl,de
BE2A  C1                 pop bc
BE2B  71                 ld (hl),c
BE2C  2C                 inc l
BE2D  70                 ld (hl),b
BE2E  19                 add hl,de
BE2F  C1                 pop bc
BE30  71                 ld (hl),c
BE31  2C                 inc l
BE32  70                 ld (hl),b
BE33  19                 add hl,de
BE34  C1                 pop bc
BE35  71                 ld (hl),c
BE36  2C                 inc l
BE37  70                 ld (hl),b
BE38  19                 add hl,de
BE39  C1                 pop bc
BE3A  71                 ld (hl),c
BE3B  2C                 inc l
BE3C  70                 ld (hl),b
BE3D  01 21 FE           ld bc,-479
BE40  09                 add hl,bc
BE41  D9                 exx
BE42  1A                 ld a,(de)
BE43  1C                 inc e
BE44  2E 02              ld l,2
BE46  CB 3F              srl a
BE48  CB 1D              rr l
BE4A  1F                 rra
BE4B  CB 1D              rr l
BE4D  1F                 rra
BE4E  CB 1D              rr l
BE50  67                 ld h,a
BE51  F9                 ld sp,hl
BE52  D9                 exx
BE53  C1                 pop bc
BE54  71                 ld (hl),c
BE55  2C                 inc l
BE56  70                 ld (hl),b
BE57  19                 add hl,de
BE58  C1                 pop bc
BE59  71                 ld (hl),c
BE5A  2C                 inc l
BE5B  70                 ld (hl),b
BE5C  19                 add hl,de
BE5D  C1                 pop bc
BE5E  71                 ld (hl),c
BE5F  2C                 inc l
BE60  70                 ld (hl),b
BE61  19                 add hl,de
BE62  C1                 pop bc
BE63  71                 ld (hl),c
BE64  2C                 inc l
BE65  70                 ld (hl),b
BE66  19                 add hl,de
BE67  C1                 pop bc
BE68  71                 ld (hl),c
BE69  2C                 inc l
BE6A  70                 ld (hl),b
BE6B  19                 add hl,de
BE6C  C1                 pop bc
BE6D  71                 ld (hl),c
BE6E  2C                 inc l
BE6F  70                 ld (hl),b
BE70  19                 add hl,de
BE71  C1                 pop bc
BE72  71                 ld (hl),c
BE73  2C                 inc l
BE74  70                 ld (hl),b
BE75  19                 add hl,de
BE76  C1                 pop bc
BE77  71                 ld (hl),c
BE78  2C                 inc l
BE79  70                 ld (hl),b
BE7A  19                 add hl,de
BE7B  C1                 pop bc
BE7C  71                 ld (hl),c
BE7D  2C                 inc l
BE7E  70                 ld (hl),b
BE7F  19                 add hl,de
BE80  C1                 pop bc
BE81  71                 ld (hl),c
BE82  2C                 inc l
BE83  70                 ld (hl),b
BE84  19                 add hl,de
BE85  C1                 pop bc
BE86  71                 ld (hl),c
BE87  2C                 inc l
BE88  70                 ld (hl),b
BE89  19                 add hl,de
BE8A  C1                 pop bc
BE8B  71                 ld (hl),c
BE8C  2C                 inc l
BE8D  70                 ld (hl),b
BE8E  19                 add hl,de
BE8F  C1                 pop bc
BE90  71                 ld (hl),c
BE91  2C                 inc l
BE92  70                 ld (hl),b
BE93  19                 add hl,de
BE94  C1                 pop bc
BE95  71                 ld (hl),c
BE96  2C                 inc l
BE97  70                 ld (hl),b
BE98  19                 add hl,de
BE99  C1                 pop bc
BE9A  71                 ld (hl),c
BE9B  2C                 inc l
BE9C  70                 ld (hl),b
BE9D  19                 add hl,de
BE9E  C1                 pop bc
BE9F  71                 ld (hl),c
BEA0  2C                 inc l
BEA1  70                 ld (hl),b
BEA2  01 21 FE           ld bc,-479
BEA5  09                 add hl,bc
BEA6  D9                 exx
BEA7  1A                 ld a,(de)
BEA8  1C                 inc e
BEA9  2E 02              ld l,2
BEAB  CB 3F              srl a
BEAD  CB 1D              rr l
BEAF  1F                 rra
BEB0  CB 1D              rr l
BEB2  1F                 rra
BEB3  CB 1D              rr l
BEB5  67                 ld h,a
BEB6  F9                 ld sp,hl
BEB7  D9                 exx
BEB8  C1                 pop bc
BEB9  71                 ld (hl),c
BEBA  2C                 inc l
BEBB  70                 ld (hl),b
BEBC  19                 add hl,de
BEBD  C1                 pop bc
BEBE  71                 ld (hl),c
BEBF  2C                 inc l
BEC0  70                 ld (hl),b
BEC1  19                 add hl,de
BEC2  C1                 pop bc
BEC3  71                 ld (hl),c
BEC4  2C                 inc l
BEC5  70                 ld (hl),b
BEC6  19                 add hl,de
BEC7  C1                 pop bc
BEC8  71                 ld (hl),c
BEC9  2C                 inc l
BECA  70                 ld (hl),b
BECB  19                 add hl,de
BECC  C1                 pop bc
BECD  71                 ld (hl),c
BECE  2C                 inc l
BECF  70                 ld (hl),b
BED0  19                 add hl,de
BED1  C1                 pop bc
BED2  71                 ld (hl),c
BED3  2C                 inc l
BED4  70                 ld (hl),b
BED5  19                 add hl,de
BED6  C1                 pop bc
BED7  71                 ld (hl),c
BED8  2C                 inc l
BED9  70                 ld (hl),b
BEDA  19                 add hl,de
BEDB  C1                 pop bc
BEDC  71                 ld (hl),c
BEDD  2C                 inc l
BEDE  70                 ld (hl),b
BEDF  19                 add hl,de
BEE0  C1                 pop bc
BEE1  71                 ld (hl),c
BEE2  2C                 inc l
BEE3  70                 ld (hl),b
BEE4  19                 add hl,de
BEE5  C1                 pop bc
BEE6  71                 ld (hl),c
BEE7  2C                 inc l
BEE8  70                 ld (hl),b
BEE9  19                 add hl,de
BEEA  C1                 pop bc
BEEB  71                 ld (hl),c
BEEC  2C                 inc l
BEED  70                 ld (hl),b
BEEE  19                 add hl,de
BEEF  C1                 pop bc
BEF0  71                 ld (hl),c
BEF1  2C                 inc l
BEF2  70                 ld (hl),b
BEF3  19                 add hl,de
BEF4  C1                 pop bc
BEF5  71                 ld (hl),c
BEF6  2C                 inc l
BEF7  70                 ld (hl),b
BEF8  19                 add hl,de
BEF9  C1                 pop bc
BEFA  71                 ld (hl),c
BEFB  2C                 inc l
BEFC  70                 ld (hl),b
BEFD  19                 add hl,de
BEFE  C1                 pop bc
BEFF  71                 ld (hl),c
BF00  2C                 inc l
BF01  70                 ld (hl),b
BF02  19                 add hl,de
BF03  C1                 pop bc
BF04  71                 ld (hl),c
BF05  2C                 inc l
BF06  70                 ld (hl),b
BF07  01 21 FE           ld bc,-479
BF0A  09                 add hl,bc
BF0B  D9                 exx
BF0C  1A                 ld a,(de)
BF0D  1C                 inc e
BF0E  2E 02              ld l,2
BF10  CB 3F              srl a
BF12  CB 1D              rr l
BF14  1F                 rra
BF15  CB 1D              rr l
BF17  1F                 rra
BF18  CB 1D              rr l
BF1A  67                 ld h,a
BF1B  F9                 ld sp,hl
BF1C  D9                 exx
BF1D  C1                 pop bc
BF1E  71                 ld (hl),c
BF1F  2C                 inc l
BF20  70                 ld (hl),b
BF21  19                 add hl,de
BF22  C1                 pop bc
BF23  71                 ld (hl),c
BF24  2C                 inc l
BF25  70                 ld (hl),b
BF26  19                 add hl,de
BF27  C1                 pop bc
BF28  71                 ld (hl),c
BF29  2C                 inc l
BF2A  70                 ld (hl),b
BF2B  19                 add hl,de
BF2C  C1                 pop bc
BF2D  71                 ld (hl),c
BF2E  2C                 inc l
BF2F  70                 ld (hl),b
BF30  19                 add hl,de
BF31  C1                 pop bc
BF32  71                 ld (hl),c
BF33  2C                 inc l
BF34  70                 ld (hl),b
BF35  19                 add hl,de
BF36  C1                 pop bc
BF37  71                 ld (hl),c
BF38  2C                 inc l
BF39  70                 ld (hl),b
BF3A  19                 add hl,de
BF3B  C1                 pop bc
BF3C  71                 ld (hl),c
BF3D  2C                 inc l
BF3E  70                 ld (hl),b
BF3F  19                 add hl,de
BF40  C1                 pop bc
BF41  71                 ld (hl),c
BF42  2C                 inc l
BF43  70                 ld (hl),b
BF44  19                 add hl,de
BF45  C1                 pop bc
BF46  71                 ld (hl),c
BF47  2C                 inc l
BF48  70                 ld (hl),b
BF49  19                 add hl,de
BF4A  C1                 pop bc
BF4B  71                 ld (hl),c
BF4C  2C                 inc l
BF4D  70                 ld (hl),b
BF4E  19                 add hl,de
BF4F  C1                 pop bc
BF50  71                 ld (hl),c
BF51  2C                 inc l
BF52  70                 ld (hl),b
BF53  19                 add hl,de
BF54  C1                 pop bc
BF55  71                 ld (hl),c
BF56  2C                 inc l
BF57  70                 ld (hl),b
BF58  19                 add hl,de
BF59  C1                 pop bc
BF5A  71                 ld (hl),c
BF5B  2C                 inc l
BF5C  70                 ld (hl),b
BF5D  19                 add hl,de
BF5E  C1                 pop bc
BF5F  71                 ld (hl),c
BF60  2C                 inc l
BF61  70                 ld (hl),b
BF62  19                 add hl,de
BF63  C1                 pop bc
BF64  71                 ld (hl),c
BF65  2C                 inc l
BF66  70                 ld (hl),b
BF67  19                 add hl,de
BF68  C1                 pop bc
BF69  71                 ld (hl),c
BF6A  2C                 inc l
BF6B  70                 ld (hl),b
BF6C  01 21 FE           ld bc,-479
BF6F  09                 add hl,bc
BF70  D9                 exx
BF71  1A                 ld a,(de)
BF72  1C                 inc e
BF73  2E 02              ld l,2
BF75  CB 3F              srl a
BF77  CB 1D              rr l
BF79  1F                 rra
BF7A  CB 1D              rr l
BF7C  1F                 rra
BF7D  CB 1D              rr l
BF7F  67                 ld h,a
BF80  F9                 ld sp,hl
BF81  D9                 exx
BF82  C1                 pop bc
BF83  71                 ld (hl),c
BF84  2C                 inc l
BF85  70                 ld (hl),b
BF86  19                 add hl,de
BF87  C1                 pop bc
BF88  71                 ld (hl),c
BF89  2C                 inc l
BF8A  70                 ld (hl),b
BF8B  19                 add hl,de
BF8C  C1                 pop bc
BF8D  71                 ld (hl),c
BF8E  2C                 inc l
BF8F  70                 ld (hl),b
BF90  19                 add hl,de
BF91  C1                 pop bc
BF92  71                 ld (hl),c
BF93  2C                 inc l
BF94  70                 ld (hl),b
BF95  19                 add hl,de
BF96  C1                 pop bc
BF97  71                 ld (hl),c
BF98  2C                 inc l
BF99  70                 ld (hl),b
BF9A  19                 add hl,de
BF9B  C1                 pop bc
BF9C  71                 ld (hl),c
BF9D  2C                 inc l
BF9E  70                 ld (hl),b
BF9F  19                 add hl,de
BFA0  C1                 pop bc
BFA1  71                 ld (hl),c
BFA2  2C                 inc l
BFA3  70                 ld (hl),b
BFA4  19                 add hl,de
BFA5  C1                 pop bc
BFA6  71                 ld (hl),c
BFA7  2C                 inc l
BFA8  70                 ld (hl),b
BFA9  19                 add hl,de
BFAA  C1                 pop bc
BFAB  71                 ld (hl),c
BFAC  2C                 inc l
BFAD  70                 ld (hl),b
BFAE  19                 add hl,de
BFAF  C1                 pop bc
BFB0  71                 ld (hl),c
BFB1  2C                 inc l
BFB2  70                 ld (hl),b
BFB3  19                 add hl,de
BFB4  C1                 pop bc
BFB5  71                 ld (hl),c
BFB6  2C                 inc l
BFB7  70                 ld (hl),b
BFB8  19                 add hl,de
BFB9  C1                 pop bc
BFBA  71                 ld (hl),c
BFBB  2C                 inc l
BFBC  70                 ld (hl),b
BFBD  19                 add hl,de
BFBE  C1                 pop bc
BFBF  71                 ld (hl),c
BFC0  2C                 inc l
BFC1  70                 ld (hl),b
BFC2  19                 add hl,de
BFC3  C1                 pop bc
BFC4  71                 ld (hl),c
BFC5  2C                 inc l
BFC6  70                 ld (hl),b
BFC7  19                 add hl,de
BFC8  C1                 pop bc
BFC9  71                 ld (hl),c
BFCA  2C                 inc l
BFCB  70                 ld (hl),b
BFCC  19                 add hl,de
BFCD  C1                 pop bc
BFCE  71                 ld (hl),c
BFCF  2C                 inc l
BFD0  70                 ld (hl),b
BFD1  01 21 FE           ld bc,-479
BFD4  09                 add hl,bc
BFD5  D9                 exx
BFD6  1A                 ld a,(de)
BFD7  1C                 inc e
BFD8  2E 02              ld l,2
BFDA  CB 3F              srl a
BFDC  CB 1D              rr l
BFDE  1F                 rra
BFDF  CB 1D              rr l
BFE1  1F                 rra
BFE2  CB 1D              rr l
BFE4  67                 ld h,a
BFE5  F9                 ld sp,hl
BFE6  D9                 exx
BFE7  C1                 pop bc
BFE8  71                 ld (hl),c
BFE9  2C                 inc l
BFEA  70                 ld (hl),b
BFEB  19                 add hl,de
BFEC  C1                 pop bc
BFED  71                 ld (hl),c
BFEE  2C                 inc l
BFEF  70                 ld (hl),b
BFF0  19                 add hl,de
BFF1  C1                 pop bc
BFF2  71                 ld (hl),c
BFF3  2C                 inc l
BFF4  70                 ld (hl),b
BFF5  19                 add hl,de
BFF6  C1                 pop bc
BFF7  71                 ld (hl),c
BFF8  2C                 inc l
BFF9  70                 ld (hl),b
BFFA  19                 add hl,de
BFFB  C1                 pop bc
BFFC  71                 ld (hl),c
BFFD  2C                 inc l
BFFE  70                 ld (hl),b
BFFF  19                 add hl,de
C000  C1                 pop bc
C001  71                 ld (hl),c
C002  2C                 inc l
C003  70                 ld (hl),b
C004  19                 add hl,de
C005  C1                 pop bc
C006  71                 ld (hl),c
C007  2C                 inc l
C008  70                 ld (hl),b
C009  19                 add hl,de
C00A  C1                 pop bc
C00B  71                 ld (hl),c
C00C  2C                 inc l
C00D  70                 ld (hl),b
C00E  19                 add hl,de
C00F  C1                 pop bc
C010  71                 ld (hl),c
C011  2C                 inc l
C012  70                 ld (hl),b
C013  19                 add hl,de
C014  C1                 pop bc
C015  71                 ld (hl),c
C016  2C                 inc l
C017  70                 ld (hl),b
C018  19                 add hl,de
C019  C1                 pop bc
C01A  71                 ld (hl),c
C01B  2C                 inc l
C01C  70                 ld (hl),b
C01D  19                 add hl,de
C01E  C1                 pop bc
C01F  71                 ld (hl),c
C020  2C                 inc l
C021  70                 ld (hl),b
C022  19                 add hl,de
C023  C1                 pop bc
C024  71                 ld (hl),c
C025  2C                 inc l
C026  70                 ld (hl),b
C027  19                 add hl,de
C028  C1                 pop bc
C029  71                 ld (hl),c
C02A  2C                 inc l
C02B  70                 ld (hl),b
C02C  19                 add hl,de
C02D  C1                 pop bc
C02E  71                 ld (hl),c
C02F  2C                 inc l
C030  70                 ld (hl),b
C031  19                 add hl,de
C032  C1                 pop bc
C033  71                 ld (hl),c
C034  2C                 inc l
C035  70                 ld (hl),b
C036  01 21 FE           ld bc,-479
C039  09                 add hl,bc
C03A  D9                 exx
C03B  1A                 ld a,(de)
C03C  1C                 inc e
C03D  2E 02              ld l,2
C03F  CB 3F              srl a
C041  CB 1D              rr l
C043  1F                 rra
C044  CB 1D              rr l
C046  1F                 rra
C047  CB 1D              rr l
C049  67                 ld h,a
C04A  F9                 ld sp,hl
C04B  D9                 exx
C04C  C1                 pop bc
C04D  71                 ld (hl),c
C04E  2C                 inc l
C04F  70                 ld (hl),b
C050  19                 add hl,de
C051  C1                 pop bc
C052  71                 ld (hl),c
C053  2C                 inc l
C054  70                 ld (hl),b
C055  19                 add hl,de
C056  C1                 pop bc
C057  71                 ld (hl),c
C058  2C                 inc l
C059  70                 ld (hl),b
C05A  19                 add hl,de
C05B  C1                 pop bc
C05C  71                 ld (hl),c
C05D  2C                 inc l
C05E  70                 ld (hl),b
C05F  19                 add hl,de
C060  C1                 pop bc
C061  71                 ld (hl),c
C062  2C                 inc l
C063  70                 ld (hl),b
C064  19                 add hl,de
C065  C1                 pop bc
C066  71                 ld (hl),c
C067  2C                 inc l
C068  70                 ld (hl),b
C069  19                 add hl,de
C06A  C1                 pop bc
C06B  71                 ld (hl),c
C06C  2C                 inc l
C06D  70                 ld (hl),b
C06E  19                 add hl,de
C06F  C1                 pop bc
C070  71                 ld (hl),c
C071  2C                 inc l
C072  70                 ld (hl),b
C073  19                 add hl,de
C074  C1                 pop bc
C075  71                 ld (hl),c
C076  2C                 inc l
C077  70                 ld (hl),b
C078  19                 add hl,de
C079  C1                 pop bc
C07A  71                 ld (hl),c
C07B  2C                 inc l
C07C  70                 ld (hl),b
C07D  19                 add hl,de
C07E  C1                 pop bc
C07F  71                 ld (hl),c
C080  2C                 inc l
C081  70                 ld (hl),b
C082  19                 add hl,de
C083  C1                 pop bc
C084  71                 ld (hl),c
C085  2C                 inc l
C086  70                 ld (hl),b
C087  19                 add hl,de
C088  C1                 pop bc
C089  71                 ld (hl),c
C08A  2C                 inc l
C08B  70                 ld (hl),b
C08C  19                 add hl,de
C08D  C1                 pop bc
C08E  71                 ld (hl),c
C08F  2C                 inc l
C090  70                 ld (hl),b
C091  19                 add hl,de
C092  C1                 pop bc
C093  71                 ld (hl),c
C094  2C                 inc l
C095  70                 ld (hl),b
C096  19                 add hl,de
C097  C1                 pop bc
C098  71                 ld (hl),c
C099  2C                 inc l
C09A  70                 ld (hl),b
C09B  01 21 FE           ld bc,-479
C09E  09                 add hl,bc
C09F  D9                 exx
C0A0  1A                 ld a,(de)
C0A1  1C                 inc e
C0A2  2E 02              ld l,2
C0A4  CB 3F              srl a
C0A6  CB 1D              rr l
C0A8  1F                 rra
C0A9  CB 1D              rr l
C0AB  1F                 rra
C0AC  CB 1D              rr l
C0AE  67                 ld h,a
C0AF  F9                 ld sp,hl
C0B0  D9                 exx
C0B1  C1                 pop bc
C0B2  71                 ld (hl),c
C0B3  2C                 inc l
C0B4  70                 ld (hl),b
C0B5  19                 add hl,de
C0B6  C1                 pop bc
C0B7  71                 ld (hl),c
C0B8  2C                 inc l
C0B9  70                 ld (hl),b
C0BA  19                 add hl,de
C0BB  C1                 pop bc
C0BC  71                 ld (hl),c
C0BD  2C                 inc l
C0BE  70                 ld (hl),b
C0BF  19                 add hl,de
C0C0  C1                 pop bc
C0C1  71                 ld (hl),c
C0C2  2C                 inc l
C0C3  70                 ld (hl),b
C0C4  19                 add hl,de
C0C5  C1                 pop bc
C0C6  71                 ld (hl),c
C0C7  2C                 inc l
C0C8  70                 ld (hl),b
C0C9  19                 add hl,de
C0CA  C1                 pop bc
C0CB  71                 ld (hl),c
C0CC  2C                 inc l
C0CD  70                 ld (hl),b
C0CE  19                 add hl,de
C0CF  C1                 pop bc
C0D0  71                 ld (hl),c
C0D1  2C                 inc l
C0D2  70                 ld (hl),b
C0D3  19                 add hl,de
C0D4  C1                 pop bc
C0D5  71                 ld (hl),c
C0D6  2C                 inc l
C0D7  70                 ld (hl),b
C0D8  19                 add hl,de
C0D9  C1                 pop bc
C0DA  71                 ld (hl),c
C0DB  2C                 inc l
C0DC  70                 ld (hl),b
C0DD  19                 add hl,de
C0DE  C1                 pop bc
C0DF  71                 ld (hl),c
C0E0  2C                 inc l
C0E1  70                 ld (hl),b
C0E2  19                 add hl,de
C0E3  C1                 pop bc
C0E4  71                 ld (hl),c
C0E5  2C                 inc l
C0E6  70                 ld (hl),b
C0E7  19                 add hl,de
C0E8  C1                 pop bc
C0E9  71                 ld (hl),c
C0EA  2C                 inc l
C0EB  70                 ld (hl),b
C0EC  19                 add hl,de
C0ED  C1                 pop bc
C0EE  71                 ld (hl),c
C0EF  2C                 inc l
C0F0  70                 ld (hl),b
C0F1  19                 add hl,de
C0F2  C1                 pop bc
C0F3  71                 ld (hl),c
C0F4  2C                 inc l
C0F5  70                 ld (hl),b
C0F6  19                 add hl,de
C0F7  C1                 pop bc
C0F8  71                 ld (hl),c
C0F9  2C                 inc l
C0FA  70                 ld (hl),b
C0FB  19                 add hl,de
C0FC  C1                 pop bc
C0FD  71                 ld (hl),c
C0FE  2C                 inc l
C0FF  70                 ld (hl),b
C100  01 21 FE           ld bc,-479
C103  09                 add hl,bc
C104  D9                 exx
C105  1A                 ld a,(de)
C106  1C                 inc e
C107  2E 02              ld l,2
C109  CB 3F              srl a
C10B  CB 1D              rr l
C10D  1F                 rra
C10E  CB 1D              rr l
C110  1F                 rra
C111  CB 1D              rr l
C113  67                 ld h,a
C114  F9                 ld sp,hl
C115  D9                 exx
C116  C1                 pop bc
C117  71                 ld (hl),c
C118  2C                 inc l
C119  70                 ld (hl),b
C11A  19                 add hl,de
C11B  C1                 pop bc
C11C  71                 ld (hl),c
C11D  2C                 inc l
C11E  70                 ld (hl),b
C11F  19                 add hl,de
C120  C1                 pop bc
C121  71                 ld (hl),c
C122  2C                 inc l
C123  70                 ld (hl),b
C124  19                 add hl,de
C125  C1                 pop bc
C126  71                 ld (hl),c
C127  2C                 inc l
C128  70                 ld (hl),b
C129  19                 add hl,de
C12A  C1                 pop bc
C12B  71                 ld (hl),c
C12C  2C                 inc l
C12D  70                 ld (hl),b
C12E  19                 add hl,de
C12F  C1                 pop bc
C130  71                 ld (hl),c
C131  2C                 inc l
C132  70                 ld (hl),b
C133  19                 add hl,de
C134  C1                 pop bc
C135  71                 ld (hl),c
C136  2C                 inc l
C137  70                 ld (hl),b
C138  19                 add hl,de
C139  C1                 pop bc
C13A  71                 ld (hl),c
C13B  2C                 inc l
C13C  70                 ld (hl),b
C13D  19                 add hl,de
C13E  C1                 pop bc
C13F  71                 ld (hl),c
C140  2C                 inc l
C141  70                 ld (hl),b
C142  19                 add hl,de
C143  C1                 pop bc
C144  71                 ld (hl),c
C145  2C                 inc l
C146  70                 ld (hl),b
C147  19                 add hl,de
C148  C1                 pop bc
C149  71                 ld (hl),c
C14A  2C                 inc l
C14B  70                 ld (hl),b
C14C  19                 add hl,de
C14D  C1                 pop bc
C14E  71                 ld (hl),c
C14F  2C                 inc l
C150  70                 ld (hl),b
C151  19                 add hl,de
C152  C1                 pop bc
C153  71                 ld (hl),c
C154  2C                 inc l
C155  70                 ld (hl),b
C156  19                 add hl,de
C157  C1                 pop bc
C158  71                 ld (hl),c
C159  2C                 inc l
C15A  70                 ld (hl),b
C15B  19                 add hl,de
C15C  C1                 pop bc
C15D  71                 ld (hl),c
C15E  2C                 inc l
C15F  70                 ld (hl),b
C160  19                 add hl,de
C161  C1                 pop bc
C162  71                 ld (hl),c
C163  2C                 inc l
C164  70                 ld (hl),b
C165  01 21 FE           ld bc,-479
C168  09                 add hl,bc
C169  D9                 exx
C16A  1A                 ld a,(de)
C16B  1C                 inc e
C16C  2E 02              ld l,2
C16E  CB 3F              srl a
C170  CB 1D              rr l
C172  1F                 rra
C173  CB 1D              rr l
C175  1F                 rra
C176  CB 1D              rr l
C178  67                 ld h,a
C179  F9                 ld sp,hl
C17A  D9                 exx
C17B  C1                 pop bc
C17C  71                 ld (hl),c
C17D  2C                 inc l
C17E  70                 ld (hl),b
C17F  19                 add hl,de
C180  C1                 pop bc
C181  71                 ld (hl),c
C182  2C                 inc l
C183  70                 ld (hl),b
C184  19                 add hl,de
C185  C1                 pop bc
C186  71                 ld (hl),c
C187  2C                 inc l
C188  70                 ld (hl),b
C189  19                 add hl,de
C18A  C1                 pop bc
C18B  71                 ld (hl),c
C18C  2C                 inc l
C18D  70                 ld (hl),b
C18E  19                 add hl,de
C18F  C1                 pop bc
C190  71                 ld (hl),c
C191  2C                 inc l
C192  70                 ld (hl),b
C193  19                 add hl,de
C194  C1                 pop bc
C195  71                 ld (hl),c
C196  2C                 inc l
C197  70                 ld (hl),b
C198  19                 add hl,de
C199  C1                 pop bc
C19A  71                 ld (hl),c
C19B  2C                 inc l
C19C  70                 ld (hl),b
C19D  19                 add hl,de
C19E  C1                 pop bc
C19F  71                 ld (hl),c
C1A0  2C                 inc l
C1A1  70                 ld (hl),b
C1A2  19                 add hl,de
C1A3  C1                 pop bc
C1A4  71                 ld (hl),c
C1A5  2C                 inc l
C1A6  70                 ld (hl),b
C1A7  19                 add hl,de
C1A8  C1                 pop bc
C1A9  71                 ld (hl),c
C1AA  2C                 inc l
C1AB  70                 ld (hl),b
C1AC  19                 add hl,de
C1AD  C1                 pop bc
C1AE  71                 ld (hl),c
C1AF  2C                 inc l
C1B0  70                 ld (hl),b
C1B1  19                 add hl,de
C1B2  C1                 pop bc
C1B3  71                 ld (hl),c
C1B4  2C                 inc l
C1B5  70                 ld (hl),b
C1B6  19                 add hl,de
C1B7  C1                 pop bc
C1B8  71                 ld (hl),c
C1B9  2C                 inc l
C1BA  70                 ld (hl),b
C1BB  19                 add hl,de
C1BC  C1                 pop bc
C1BD  71                 ld (hl),c
C1BE  2C                 inc l
C1BF  70                 ld (hl),b
C1C0  19                 add hl,de
C1C1  C1                 pop bc
C1C2  71                 ld (hl),c
C1C3  2C                 inc l
C1C4  70                 ld (hl),b
C1C5  19                 add hl,de
C1C6  C1                 pop bc
C1C7  71                 ld (hl),c
C1C8  2C                 inc l
C1C9  70                 ld (hl),b
C1CA  01 21 FE           ld bc,-479
C1CD  09                 add hl,bc
C1CE  D9                 exx
C1CF  1A                 ld a,(de)
C1D0  1C                 inc e
C1D1  2E 02              ld l,2
C1D3  CB 3F              srl a
C1D5  CB 1D              rr l
C1D7  1F                 rra
C1D8  CB 1D              rr l
C1DA  1F                 rra
C1DB  CB 1D              rr l
C1DD  67                 ld h,a
C1DE  F9                 ld sp,hl
C1DF  D9                 exx
C1E0  C1                 pop bc
C1E1  71                 ld (hl),c
C1E2  2C                 inc l
C1E3  70                 ld (hl),b
C1E4  19                 add hl,de
C1E5  C1                 pop bc
C1E6  71                 ld (hl),c
C1E7  2C                 inc l
C1E8  70                 ld (hl),b
C1E9  19                 add hl,de
C1EA  C1                 pop bc
C1EB  71                 ld (hl),c
C1EC  2C                 inc l
C1ED  70                 ld (hl),b
C1EE  19                 add hl,de
C1EF  C1                 pop bc
C1F0  71                 ld (hl),c
C1F1  2C                 inc l
C1F2  70                 ld (hl),b
C1F3  19                 add hl,de
C1F4  C1                 pop bc
C1F5  71                 ld (hl),c
C1F6  2C                 inc l
C1F7  70                 ld (hl),b
C1F8  19                 add hl,de
C1F9  C1                 pop bc
C1FA  71                 ld (hl),c
C1FB  2C                 inc l
C1FC  70                 ld (hl),b
C1FD  19                 add hl,de
C1FE  C1                 pop bc
C1FF  71                 ld (hl),c
C200  2C                 inc l
C201  70                 ld (hl),b
C202  19                 add hl,de
C203  C1                 pop bc
C204  71                 ld (hl),c
C205  2C                 inc l
C206  70                 ld (hl),b
C207  19                 add hl,de
C208  C1                 pop bc
C209  71                 ld (hl),c
C20A  2C                 inc l
C20B  70                 ld (hl),b
C20C  19                 add hl,de
C20D  C1                 pop bc
C20E  71                 ld (hl),c
C20F  2C                 inc l
C210  70                 ld (hl),b
C211  19                 add hl,de
C212  C1                 pop bc
C213  71                 ld (hl),c
C214  2C                 inc l
C215  70                 ld (hl),b
C216  19                 add hl,de
C217  C1                 pop bc
C218  71                 ld (hl),c
C219  2C                 inc l
C21A  70                 ld (hl),b
C21B  19                 add hl,de
C21C  C1                 pop bc
C21D  71                 ld (hl),c
C21E  2C                 inc l
C21F  70                 ld (hl),b
C220  19                 add hl,de
C221  C1                 pop bc
C222  71                 ld (hl),c
C223  2C                 inc l
C224  70                 ld (hl),b
C225  19                 add hl,de
C226  C1                 pop bc
C227  71                 ld (hl),c
C228  2C                 inc l
C229  70                 ld (hl),b
C22A  19                 add hl,de
C22B  C1                 pop bc
C22C  71                 ld (hl),c
C22D  2C                 inc l
C22E  70                 ld (hl),b
C22F  01 21 FE           ld bc,-479
C232  09                 add hl,bc
C233  D9                 exx
C234  1A                 ld a,(de)
C235  1C                 inc e
C236  2E 02              ld l,2
C238  CB 3F              srl a
C23A  CB 1D              rr l
C23C  1F                 rra
C23D  CB 1D              rr l
C23F  1F                 rra
C240  CB 1D              rr l
C242  67                 ld h,a
C243  F9                 ld sp,hl
C244  D9                 exx
C245  C1                 pop bc
C246  71                 ld (hl),c
C247  2C                 inc l
C248  70                 ld (hl),b
C249  19                 add hl,de
C24A  C1                 pop bc
C24B  71                 ld (hl),c
C24C  2C                 inc l
C24D  70                 ld (hl),b
C24E  19                 add hl,de
C24F  C1                 pop bc
C250  71                 ld (hl),c
C251  2C                 inc l
C252  70                 ld (hl),b
C253  19                 add hl,de
C254  C1                 pop bc
C255  71                 ld (hl),c
C256  2C                 inc l
C257  70                 ld (hl),b
C258  19                 add hl,de
C259  C1                 pop bc
C25A  71                 ld (hl),c
C25B  2C                 inc l
C25C  70                 ld (hl),b
C25D  19                 add hl,de
C25E  C1                 pop bc
C25F  71                 ld (hl),c
C260  2C                 inc l
C261  70                 ld (hl),b
C262  19                 add hl,de
C263  C1                 pop bc
C264  71                 ld (hl),c
C265  2C                 inc l
C266  70                 ld (hl),b
C267  19                 add hl,de
C268  C1                 pop bc
C269  71                 ld (hl),c
C26A  2C                 inc l
C26B  70                 ld (hl),b
C26C  19                 add hl,de
C26D  C1                 pop bc
C26E  71                 ld (hl),c
C26F  2C                 inc l
C270  70                 ld (hl),b
C271  19                 add hl,de
C272  C1                 pop bc
C273  71                 ld (hl),c
C274  2C                 inc l
C275  70                 ld (hl),b
C276  19                 add hl,de
C277  C1                 pop bc
C278  71                 ld (hl),c
C279  2C                 inc l
C27A  70                 ld (hl),b
C27B  19                 add hl,de
C27C  C1                 pop bc
C27D  71                 ld (hl),c
C27E  2C                 inc l
C27F  70                 ld (hl),b
C280  19                 add hl,de
C281  C1                 pop bc
C282  71                 ld (hl),c
C283  2C                 inc l
C284  70                 ld (hl),b
C285  19                 add hl,de
C286  C1                 pop bc
C287  71                 ld (hl),c
C288  2C                 inc l
C289  70                 ld (hl),b
C28A  19                 add hl,de
C28B  C1                 pop bc
C28C  71                 ld (hl),c
C28D  2C                 inc l
C28E  70                 ld (hl),b
C28F  19                 add hl,de
C290  C1                 pop bc
C291  71                 ld (hl),c
C292  2C                 inc l
C293  70                 ld (hl),b
C294  01 21 FE           ld bc,-479
C297  09                 add hl,bc
C298  D9                 exx
C299  1A                 ld a,(de)
C29A  1C                 inc e
C29B  2E 02              ld l,2
C29D  CB 3F              srl a
C29F  CB 1D              rr l
C2A1  1F                 rra
C2A2  CB 1D              rr l
C2A4  1F                 rra
C2A5  CB 1D              rr l
C2A7  67                 ld h,a
C2A8  F9                 ld sp,hl
C2A9  D9                 exx
C2AA  C1                 pop bc
C2AB  71                 ld (hl),c
C2AC  2C                 inc l
C2AD  70                 ld (hl),b
C2AE  19                 add hl,de
C2AF  C1                 pop bc
C2B0  71                 ld (hl),c
C2B1  2C                 inc l
C2B2  70                 ld (hl),b
C2B3  19                 add hl,de
C2B4  C1                 pop bc
C2B5  71                 ld (hl),c
C2B6  2C                 inc l
C2B7  70                 ld (hl),b
C2B8  19                 add hl,de
C2B9  C1                 pop bc
C2BA  71                 ld (hl),c
C2BB  2C                 inc l
C2BC  70                 ld (hl),b
C2BD  19                 add hl,de
C2BE  C1                 pop bc
C2BF  71                 ld (hl),c
C2C0  2C                 inc l
C2C1  70                 ld (hl),b
C2C2  19                 add hl,de
C2C3  C1                 pop bc
C2C4  71                 ld (hl),c
C2C5  2C                 inc l
C2C6  70                 ld (hl),b
C2C7  19                 add hl,de
C2C8  C1                 pop bc
C2C9  71                 ld (hl),c
C2CA  2C                 inc l
C2CB  70                 ld (hl),b
C2CC  19                 add hl,de
C2CD  C1                 pop bc
C2CE  71                 ld (hl),c
C2CF  2C                 inc l
C2D0  70                 ld (hl),b
C2D1  19                 add hl,de
C2D2  C1                 pop bc
C2D3  71                 ld (hl),c
C2D4  2C                 inc l
C2D5  70                 ld (hl),b
C2D6  19                 add hl,de
C2D7  C1                 pop bc
C2D8  71                 ld (hl),c
C2D9  2C                 inc l
C2DA  70                 ld (hl),b
C2DB  19                 add hl,de
C2DC  C1                 pop bc
C2DD  71                 ld (hl),c
C2DE  2C                 inc l
C2DF  70                 ld (hl),b
C2E0  19                 add hl,de
C2E1  C1                 pop bc
C2E2  71                 ld (hl),c
C2E3  2C                 inc l
C2E4  70                 ld (hl),b
C2E5  19                 add hl,de
C2E6  C1                 pop bc
C2E7  71                 ld (hl),c
C2E8  2C                 inc l
C2E9  70                 ld (hl),b
C2EA  19                 add hl,de
C2EB  C1                 pop bc
C2EC  71                 ld (hl),c
C2ED  2C                 inc l
C2EE  70                 ld (hl),b
C2EF  19                 add hl,de
C2F0  C1                 pop bc
C2F1  71                 ld (hl),c
C2F2  2C                 inc l
C2F3  70                 ld (hl),b
C2F4  19                 add hl,de
C2F5  C1                 pop bc
C2F6  71                 ld (hl),c
C2F7  2C                 inc l
C2F8  70                 ld (hl),b
C2F9  01 21 FE           ld bc,-479
C2FC  09                 add hl,bc
C2FD  D9                 exx
C2FE  1A                 ld a,(de)
C2FF  1C                 inc e
C300  2E 02              ld l,2
C302  CB 3F              srl a
C304  CB 1D              rr l
C306  1F                 rra
C307  CB 1D              rr l
C309  1F                 rra
C30A  CB 1D              rr l
C30C  67                 ld h,a
C30D  F9                 ld sp,hl
C30E  D9                 exx
C30F  C1                 pop bc
C310  71                 ld (hl),c
C311  2C                 inc l
C312  70                 ld (hl),b
C313  19                 add hl,de
C314  C1                 pop bc
C315  71                 ld (hl),c
C316  2C                 inc l
C317  70                 ld (hl),b
C318  19                 add hl,de
C319  C1                 pop bc
C31A  71                 ld (hl),c
C31B  2C                 inc l
C31C  70                 ld (hl),b
C31D  19                 add hl,de
C31E  C1                 pop bc
C31F  71                 ld (hl),c
C320  2C                 inc l
C321  70                 ld (hl),b
C322  19                 add hl,de
C323  C1                 pop bc
C324  71                 ld (hl),c
C325  2C                 inc l
C326  70                 ld (hl),b
C327  19                 add hl,de
C328  C1                 pop bc
C329  71                 ld (hl),c
C32A  2C                 inc l
C32B  70                 ld (hl),b
C32C  19                 add hl,de
C32D  C1                 pop bc
C32E  71                 ld (hl),c
C32F  2C                 inc l
C330  70                 ld (hl),b
C331  19                 add hl,de
C332  C1                 pop bc
C333  71                 ld (hl),c
C334  2C                 inc l
C335  70                 ld (hl),b
C336  19                 add hl,de
C337  C1                 pop bc
C338  71                 ld (hl),c
C339  2C                 inc l
C33A  70                 ld (hl),b
C33B  19                 add hl,de
C33C  C1                 pop bc
C33D  71                 ld (hl),c
C33E  2C                 inc l
C33F  70                 ld (hl),b
C340  19                 add hl,de
C341  C1                 pop bc
C342  71                 ld (hl),c
C343  2C                 inc l
C344  70                 ld (hl),b
C345  19                 add hl,de
C346  C1                 pop bc
C347  71                 ld (hl),c
C348  2C                 inc l
C349  70                 ld (hl),b
C34A  19                 add hl,de
C34B  C1                 pop bc
C34C  71                 ld (hl),c
C34D  2C                 inc l
C34E  70                 ld (hl),b
C34F  19                 add hl,de
C350  C1                 pop bc
C351  71                 ld (hl),c
C352  2C                 inc l
C353  70                 ld (hl),b
C354  19                 add hl,de
C355  C1                 pop bc
C356  71                 ld (hl),c
C357  2C                 inc l
C358  70                 ld (hl),b
C359  19                 add hl,de
C35A  C1                 pop bc
C35B  71                 ld (hl),c
C35C  2C                 inc l
C35D  70                 ld (hl),b
C35E  01 21 FE           ld bc,-479
C361  09                 add hl,bc
C362  D9                 exx
C363  1A                 ld a,(de)
C364  1C                 inc e
C365  2E 02              ld l,2
C367  CB 3F              srl a
C369  CB 1D              rr l
C36B  1F                 rra
C36C  CB 1D              rr l
C36E  1F                 rra
C36F  CB 1D              rr l
C371  67                 ld h,a
C372  F9                 ld sp,hl
C373  D9                 exx
C374  C1                 pop bc
C375  71                 ld (hl),c
C376  2C                 inc l
C377  70                 ld (hl),b
C378  19                 add hl,de
C379  C1                 pop bc
C37A  71                 ld (hl),c
C37B  2C                 inc l
C37C  70                 ld (hl),b
C37D  19                 add hl,de
C37E  C1                 pop bc
C37F  71                 ld (hl),c
C380  2C                 inc l
C381  70                 ld (hl),b
C382  19                 add hl,de
C383  C1                 pop bc
C384  71                 ld (hl),c
C385  2C                 inc l
C386  70                 ld (hl),b
C387  19                 add hl,de
C388  C1                 pop bc
C389  71                 ld (hl),c
C38A  2C                 inc l
C38B  70                 ld (hl),b
C38C  19                 add hl,de
C38D  C1                 pop bc
C38E  71                 ld (hl),c
C38F  2C                 inc l
C390  70                 ld (hl),b
C391  19                 add hl,de
C392  C1                 pop bc
C393  71                 ld (hl),c
C394  2C                 inc l
C395  70                 ld (hl),b
C396  19                 add hl,de
C397  C1                 pop bc
C398  71                 ld (hl),c
C399  2C                 inc l
C39A  70                 ld (hl),b
C39B  19                 add hl,de
C39C  C1                 pop bc
C39D  71                 ld (hl),c
C39E  2C                 inc l
C39F  70                 ld (hl),b
C3A0  19                 add hl,de
C3A1  C1                 pop bc
C3A2  71                 ld (hl),c
C3A3  2C                 inc l
C3A4  70                 ld (hl),b
C3A5  19                 add hl,de
C3A6  C1                 pop bc
C3A7  71                 ld (hl),c
C3A8  2C                 inc l
C3A9  70                 ld (hl),b
C3AA  19                 add hl,de
C3AB  C1                 pop bc
C3AC  71                 ld (hl),c
C3AD  2C                 inc l
C3AE  70                 ld (hl),b
C3AF  19                 add hl,de
C3B0  C1                 pop bc
C3B1  71                 ld (hl),c
C3B2  2C                 inc l
C3B3  70                 ld (hl),b
C3B4  19                 add hl,de
C3B5  C1                 pop bc
C3B6  71                 ld (hl),c
C3B7  2C                 inc l
C3B8  70                 ld (hl),b
C3B9  19                 add hl,de
C3BA  C1                 pop bc
C3BB  71                 ld (hl),c
C3BC  2C                 inc l
C3BD  70                 ld (hl),b
C3BE  19                 add hl,de
C3BF  C1                 pop bc
C3C0  71                 ld (hl),c
C3C1  2C                 inc l
C3C2  70                 ld (hl),b
C3C3  01 01 00           ld bc,1
C3C6  09                 add hl,bc
C3C7  D9                 exx
C3C8  7B                 ld a,e
C3C9  D6 10              sub 16
C3CB  81                 add a,c
C3CC  5F                 ld e,a
C3CD  D2 D1 C3           jp nc,$C3D1
C3D0  14                 inc d
C3D1  05                 dec b
C3D2  C2 76 BD           jp nz,$BD76
C3D5  D9                 exx
C3D6  31 FF FF           ld sp,$FFFF
C3D9  C9                 ret
C3DA  21 E0 D2           ld hl,$D2E0
C3DD  7E                 ld a,(hl)
C3DE  35                 dec (hl)
C3DF  F2 E4 C3           jp p,$C3E4
C3E2  36 30              ld (hl),48
C3E4  FE 1C              cp 28
C3E6  28 08              jr z,$C3F0
C3E8  30 14              jr nc,$C3FE
C3EA  3E 01              ld a,1
C3EC  32 DC D2           ld ($D2DC),a
C3EF  C9                 ret
C3F0  3E 03              ld a,3
C3F2  32 DC D2           ld ($D2DC),a
C3F5  C3 B7 C9           jp $C9B7
C3F8  3E 1C              ld a,28
C3FA  32 E0 D2           ld ($D2E0),a
C3FD  C9                 ret
C3FE  2A D8 D2           ld hl,($D2D8)
C401  11 F0 FF           ld de,-16
C404  19                 add hl,de
C405  ED 5B 10 D2        ld de,($D210)
C409  A7                 and a
C40A  ED 52              sbc hl,de
C40C  38 18              jr c,$C426
C40E  11 FC FF           ld de,-4
C411  7C                 ld a,h
C412  A7                 and a
C413  20 05              jr nz,$C41A
C415  7D                 ld a,l
C416  FE 04              cp 4
C418  38 DE              jr c,$C3F8
C41A  2A D8 D2           ld hl,($D2D8)
C41D  01 A8 00           ld bc,168
C420  A7                 and a
C421  ED 42              sbc hl,bc
C423  D8                 ret c
C424  18 18              jr $C43E
C426  11 04 00           ld de,4
C429  7C                 ld a,h
C42A  3C                 inc a
C42B  20 05              jr nz,$C432
C42D  7D                 ld a,l
C42E  FE F8              cp -8
C430  30 C6              jr nc,$C3F8
C432  2A D8 D2           ld hl,($D2D8)
C435  01 20 01           ld bc,288
C438  A7                 and a
C439  ED 42              sbc hl,bc
C43B  D0                 ret nc
C43C  18 00              jr $C43E
C43E  2A D8 D2           ld hl,($D2D8)
C441  A7                 and a
C442  19                 add hl,de
C443  22 D8 D2           ld ($D2D8),hl
C446  7D                 ld a,l
C447  CB 3F              srl a
C449  CB 3F              srl a
C44B  CB 3F              srl a
C44D  E6 03              and $03
C44F  FE 03              cp 3
C451  20 02              jr nz,$C455
C453  3E 01              ld a,1
C455  32 DC D2           ld ($D2DC),a
C458  C9                 ret
C459  3A 84 B0           ld a,($B084)
C45C  A7                 and a
C45D  C0                 ret nz
C45E  21 E0 D2           ld hl,$D2E0
C461  7E                 ld a,(hl)
C462  FE 64              cp 100
C464  28 31              jr z,$C497
C466  CD DE E9           call $E9DE
C469  E6 80              and $80
C46B  F6 7F              or $7F
C46D  32 DC D2           ld ($D2DC),a
C470  CD DE E9           call $E9DE
C473  E6 01              and $01
C475  C0                 ret nz
C476  CD 3D CA           call $CA3D
C479  C0                 ret nz
C47A  21 E0 D2           ld hl,$D2E0
C47D  34                 inc (hl)
C47E  CD DE E9           call $E9DE
C481  E6 03              and $03
C483  F6 F8              or $F8
C485  FD 77 06           ld (iy+6),a
C488  CD DE E9           call $E9DE
C48B  E6 0F              and $0F
C48D  D6 08              sub 8
C48F  FD 77 07           ld (iy+7),a
C492  AF                 xor a
C493  FD 77 08           ld (iy+8),a
C496  C9                 ret
C497  21 69 D2           ld hl,$D269
C49A  06 0C              ld b,12
C49C  11 0A 00           ld de,10
C49F  7E                 ld a,(hl)
C4A0  FE FF              cp -1
C4A2  C0                 ret nz
C4A3  19                 add hl,de
C4A4  10 F9              djnz $C49F
C4A6  CD D2 7B           call $7BD2
C4A9  3E 1C              ld a,28
C4AB  32 19 D2           ld ($D219),a
C4AE  3A AC B0           ld a,($B0AC)
C4B1  A7                 and a
C4B2  C0                 ret nz
C4B3  32 41 D2           ld ($D241),a
C4B6  C9                 ret
C4B7  2A DA D2           ld hl,($D2DA)
C4BA  FD 5E 02           ld e,(iy+2)
C4BD  FD 56 03           ld d,(iy+3)
C4C0  A7                 and a
C4C1  ED 52              sbc hl,de
C4C3  11 03 00           ld de,3
C4C6  19                 add hl,de
C4C7  7C                 ld a,h
C4C8  A7                 and a
C4C9  C0                 ret nz
C4CA  7D                 ld a,l
C4CB  FE 06              cp 6
C4CD  D0                 ret nc
C4CE  2A D8 D2           ld hl,($D2D8)
C4D1  FD 5E 00           ld e,(iy+0)
C4D4  FD 56 01           ld d,(iy+1)
C4D7  A7                 and a
C4D8  ED 52              sbc hl,de
C4DA  11 0C 00           ld de,12
C4DD  19                 add hl,de
C4DE  7C                 ld a,h
C4DF  A7                 and a
C4E0  C0                 ret nz
C4E1  7D                 ld a,l
C4E2  FE 18              cp 24
C4E4  C9                 ret
C4E5  2A DA D2           ld hl,($D2DA)
C4E8  FD 5E 02           ld e,(iy+2)
C4EB  FD 56 03           ld d,(iy+3)
C4EE  A7                 and a
C4EF  ED 52              sbc hl,de
C4F1  11 04 00           ld de,4
C4F4  19                 add hl,de
C4F5  7C                 ld a,h
C4F6  A7                 and a
C4F7  C0                 ret nz
C4F8  7D                 ld a,l
C4F9  FE 0E              cp 14
C4FB  D0                 ret nc
C4FC  2A D8 D2           ld hl,($D2D8)
C4FF  FD 5E 00           ld e,(iy+0)
C502  FD 56 01           ld d,(iy+1)
C505  A7                 and a
C506  ED 52              sbc hl,de
C508  11 24 00           ld de,36
C50B  19                 add hl,de
C50C  7C                 ld a,h
C50D  A7                 and a
C50E  C0                 ret nz
C50F  7D                 ld a,l
C510  FE 3C              cp 60
C512  C9                 ret
C513  FD 21 F4 CF        ld iy,$CFF4
C517  11 0A 00           ld de,10
C51A  06 4A              ld b,74
C51C  FD 7E 09           ld a,(iy+9)
C51F  FE 1C              cp 28
C521  D2 30 C5           jp nc,$C530
C524  FE 08              cp 8
C526  DA 30 C5           jp c,$C530
C529  CD B7 C4           call $C4B7
C52C  38 19              jr c,$C547
C52E  18 0F              jr $C53F
C530  FE 26              cp 38
C532  D2 42 C5           jp nc,$C542
C535  FE 23              cp 35
C537  DA 42 C5           jp c,$C542
C53A  CD E5 C4           call $C4E5
C53D  38 08              jr c,$C547
C53F  11 0A 00           ld de,10
C542  FD 19              add iy,de
C544  10 D6              djnz $C51C
C546  C9                 ret
C547  21 E1 D2           ld hl,$D2E1
C54A  7E                 ld a,(hl)
C54B  06 FF              ld b,-1
C54D  FE 00              cp 0
C54F  28 0A              jr z,$C55B
C551  36 FF              ld (hl),-1
C553  06 01              ld b,1
C555  FE 04              cp 4
C557  28 02              jr z,$C55B
C559  06 05              ld b,5
C55B  FD 7E 05           ld a,(iy+5)
C55E  90                 sub b
C55F  FD 77 05           ld (iy+5),a
C562  D0                 ret nc
C563  FD 7E 09           ld a,(iy+9)
C566  FE 23              cp 35
C568  38 19              jr c,$C583
C56A  FE 26              cp 38
C56C  30 15              jr nc,$C583
C56E  3E 20              ld a,32
C570  FD 77 09           ld (iy+9),a
C573  3E F5              ld a,-11
C575  FD 77 06           ld (iy+6),a
C578  3E 1C              ld a,28
C57A  FD 77 04           ld (iy+4),a
C57D  21 00 10           ld hl,$1000
C580  C3 88 B6           jp $B688
C583  3E 02              ld a,2
C585  32 FF E9           ld ($E9FF),a
C588  3A E1 D2           ld a,($D2E1)
C58B  FD 77 05           ld (iy+5),a
C58E  FE 00              cp 0
C590  21 50 00           ld hl,80
C593  FD E5              push iy
C595  C4 88 B6           call nz,$B688
C598  FD E1              pop iy
C59A  C3 91 CA           jp $CA91
C59D  21 4D FF           ld hl,-179
C5A0  ED 4B E1 D2        ld bc,($D2E1)
C5A4  06 00              ld b,0
C5A6  09                 add hl,bc
C5A7  7E                 ld a,(hl)
C5A8  18 64              jr $C60E
C5AA  21 92 E9           ld hl,$E992
C5AD  ED 4B E1 D2        ld bc,($D2E1)
C5B1  06 00              ld b,0
C5B3  09                 add hl,bc
C5B4  7E                 ld a,(hl)
C5B5  18 57              jr $C60E
C5B7  21 73 FF           ld hl,-141
C5BA  ED 4B E1 D2        ld bc,($D2E1)
C5BE  06 00              ld b,0
C5C0  09                 add hl,bc
C5C1  7E                 ld a,(hl)
C5C2  18 4A              jr $C60E
C5C4  21 B8 E9           ld hl,$E9B8
C5C7  ED 4B E1 D2        ld bc,($D2E1)
C5CB  06 00              ld b,0
C5CD  09                 add hl,bc
C5CE  7E                 ld a,(hl)
C5CF  A7                 and a
C5D0  28 D8              jr z,$C5AA
C5D2  FE 07              cp 7
C5D4  28 21              jr z,$C5F7
C5D6  FE 02              cp 2
C5D8  28 0F              jr z,$C5E9
C5DA  21 E0 D2           ld hl,$D2E0
C5DD  34                 inc (hl)
C5DE  7E                 ld a,(hl)
C5DF  E6 03              and $03
C5E1  FE 03              cp 3
C5E3  20 02              jr nz,$C5E7
C5E5  3E 01              ld a,1
C5E7  18 1A              jr $C603
C5E9  01 49 15           ld bc,$1549
C5EC  09                 add hl,bc
C5ED  3A DC D2           ld a,($D2DC)
C5F0  E6 7F              and $7F
C5F2  96                 sub (hl)
C5F3  EE 01              xor $01
C5F5  18 0C              jr $C603
C5F7  21 DC D2           ld hl,$D2DC
C5FA  7E                 ld a,(hl)
C5FB  3C                 inc a
C5FC  E6 07              and $07
C5FE  FE 07              cp 7
C600  38 01              jr c,$C603
C602  AF                 xor a
C603  21 01 FF           ld hl,-255
C606  ED 4B E1 D2        ld bc,($D2E1)
C60A  06 00              ld b,0
C60C  09                 add hl,bc
C60D  86                 add a,(hl)
C60E  21 DC D2           ld hl,$D2DC
C611  AE                 xor (hl)
C612  E6 7F              and $7F
C614  AE                 xor (hl)
C615  77                 ld (hl),a
C616  C9                 ret
C617  21 DC D2           ld hl,$D2DC
C61A  AE                 xor (hl)
C61B  E6 80              and $80
C61D  AE                 xor (hl)
C61E  77                 ld (hl),a
C61F  C9                 ret
C620  2A DA D2           ld hl,($D2DA)
C623  ED 5B 12 D2        ld de,($D212)
C627  A7                 and a
C628  ED 52              sbc hl,de
C62A  30 05              jr nc,$C631
C62C  3E 80              ld a,-128
C62E  C3 17 C6           jp $C617
C631  AF                 xor a
C632  C3 17 C6           jp $C617
C635  3A DC D2           ld a,($D2DC)
C638  21 DF D2           ld hl,$D2DF
C63B  06 01              ld b,1
C63D  17                 rla
C63E  30 02              jr nc,$C642
C640  06 02              ld b,2
C642  7E                 ld a,(hl)
C643  E6 FC              and $FC
C645  B0                 or b
C646  77                 ld (hl),a
C647  C9                 ret
C648  21 DF D2           ld hl,$D2DF
C64B  2F                 cpl
C64C  A6                 and (hl)
C64D  77                 ld (hl),a
C64E  C9                 ret
C64F  CD 9D C5           call $C59D
C652  21 DE D2           ld hl,$D2DE
C655  34                 inc (hl)
C656  3A E0 D2           ld a,($D2E0)
C659  2A DA D2           ld hl,($D2DA)
C65C  87                 add a,a
C65D  ED 6A              adc hl,hl
C65F  87                 add a,a
C660  ED 6A              adc hl,hl
C662  ED 4B DF D2        ld bc,($D2DF)
C666  CD C8 C8           call $C8C8
C669  CB 3C              srl h
C66B  CB 1D              rr l
C66D  1F                 rra
C66E  CB 3C              srl h
C670  CB 1D              rr l
C672  1F                 rra
C673  22 DA D2           ld ($D2DA),hl
C676  32 E0 D2           ld ($D2E0),a
C679  2A D8 D2           ld hl,($D2D8)
C67C  ED 4B DE D2        ld bc,($D2DE)
C680  CD C8 C8           call $C8C8
C683  22 D8 D2           ld ($D2D8),hl
C686  11 7F 02           ld de,639
C689  A7                 and a
C68A  ED 52              sbc hl,de
C68C  D8                 ret c
C68D  C3 C8 CA           jp $CAC8
C690  21 E0 D2           ld hl,$D2E0
C693  7E                 ld a,(hl)
C694  35                 dec (hl)
C695  A7                 and a
C696  28 20              jr z,$C6B8
C698  FE 18              cp 24
C69A  38 0C              jr c,$C6A8
C69C  FE 20              cp 32
C69E  D0                 ret nc
C69F  3A DC D2           ld a,($D2DC)
C6A2  EE 80              xor $80
C6A4  32 DC D2           ld ($D2DC),a
C6A7  C9                 ret
C6A8  21 DE D2           ld hl,$D2DE
C6AB  34                 inc (hl)
C6AC  34                 inc (hl)
C6AD  5E                 ld e,(hl)
C6AE  16 00              ld d,0
C6B0  2A D8 D2           ld hl,($D2D8)
C6B3  19                 add hl,de
C6B4  22 D8 D2           ld ($D2D8),hl
C6B7  C9                 ret
C6B8  CD DE E9           call $E9DE
C6BB  E6 0F              and $0F
C6BD  C6 30              add a,48
C6BF  77                 ld (hl),a
C6C0  CD DE E9           call $E9DE
C6C3  E6 1F              and $1F
C6C5  5F                 ld e,a
C6C6  16 00              ld d,0
C6C8  21 BC 01           ld hl,444
C6CB  19                 add hl,de
C6CC  22 DA D2           ld ($D2DA),hl
C6CF  21 74 01           ld hl,372
C6D2  22 D8 D2           ld ($D2D8),hl
C6D5  AF                 xor a
C6D6  32 DE D2           ld ($D2DE),a
C6D9  C3 AA C5           jp $C5AA
C6DC  3A 16 D2           ld a,($D216)
C6DF  A7                 and a
C6E0  20 1F              jr nz,$C701
C6E2  2A 10 D2           ld hl,($D210)
C6E5  11 C8 FF           ld de,-56
C6E8  19                 add hl,de
C6E9  ED 5B D8 D2        ld de,($D2D8)
C6ED  A7                 and a
C6EE  ED 52              sbc hl,de
C6F0  28 0F              jr z,$C701
C6F2  11 FC FF           ld de,-4
C6F5  38 03              jr c,$C6FA
C6F7  11 04 00           ld de,4
C6FA  2A D8 D2           ld hl,($D2D8)
C6FD  19                 add hl,de
C6FE  22 D8 D2           ld ($D2D8),hl
C701  CD 20 C6           call $C620
C704  21 E0 D2           ld hl,$D2E0
C707  35                 dec (hl)
C708  7E                 ld a,(hl)
C709  E6 10              and $10
C70B  11 10 00           ld de,16
C70E  28 03              jr z,$C713
C710  11 F0 FF           ld de,-16
C713  2A 12 D2           ld hl,($D212)
C716  19                 add hl,de
C717  ED 5B DA D2        ld de,($D2DA)
C71B  A7                 and a
C71C  ED 52              sbc hl,de
C71E  7D                 ld a,l
C71F  A7                 and a
C720  F2 2B C7           jp p,$C72B
C723  FE F0              cp -16
C725  30 0A              jr nc,$C731
C727  3E F0              ld a,-16
C729  18 06              jr $C731
C72B  FE 10              cp 16
C72D  38 02              jr c,$C731
C72F  3E 0F              ld a,15
C731  C6 10              add a,16
C733  5F                 ld e,a
C734  16 00              ld d,0
C736  21 DE EA           ld hl,$EADE
C739  19                 add hl,de
C73A  5E                 ld e,(hl)
C73B  7B                 ld a,e
C73C  07                 rlca
C73D  30 01              jr nc,$C740
C73F  15                 dec d
C740  2A DA D2           ld hl,($D2DA)
C743  19                 add hl,de
C744  22 DA D2           ld ($D2DA),hl
C747  A7                 and a
C748  C0                 ret nz
C749  CD DE E9           call $E9DE
C74C  E6 07              and $07
C74E  C0                 ret nz
C74F  C3 3D CA           jp $CA3D
C752  CD DA CC           call $CCDA
C755  3A 9C B0           ld a,($B09C)
C758  A7                 and a
C759  28 04              jr z,$C75F
C75B  CD 35 C6           call $C635
C75E  C9                 ret
C75F  3A E0 D2           ld a,($D2E0)
C762  FE 0F              cp 15
C764  DC 20 C6           call c,$C620
C767  21 E0 D2           ld hl,$D2E0
C76A  35                 dec (hl)
C76B  F0                 ret p
C76C  CD DE E9           call $E9DE
C76F  E6 0F              and $0F
C771  C6 0A              add a,10
C773  32 E0 D2           ld ($D2E0),a
C776  3E 08              ld a,8
C778  CD 48 C6           call $C648
C77B  C9                 ret
C77C  CD 35 C6           call $C635
C77F  CD DE E9           call $E9DE
C782  E6 0F              and $0F
C784  20 05              jr nz,$C78B
C786  3E 08              ld a,8
C788  CD 48 C6           call $C648
C78B  CD DA CC           call $CCDA
C78E  3A 9C B0           ld a,($B09C)
C791  A7                 and a
C792  C0                 ret nz
C793  CD DE E9           call $E9DE
C796  E6 07              and $07
C798  C0                 ret nz
C799  3A DC D2           ld a,($D2DC)
C79C  EE 80              xor $80
C79E  32 DC D2           ld ($D2DC),a
C7A1  C9                 ret
C7A2  CD DE E9           call $E9DE
C7A5  E6 3F              and $3F
C7A7  20 05              jr nz,$C7AE
C7A9  3E 10              ld a,16
C7AB  CD 48 C6           call $C648
C7AE  CD 35 C6           call $C635
C7B1  C3 D3 CC           jp $CCD3
C7B4  3A DC D2           ld a,($D2DC)
C7B7  EE 80              xor $80
C7B9  32 DC D2           ld ($D2DC),a
C7BC  CD DE E9           call $E9DE
C7BF  E6 0F              and $0F
C7C1  C0                 ret nz
C7C2  C3 3D CA           jp $CA3D
C7C5  CD AA C5           call $C5AA
C7C8  2A AA B0           ld hl,($B0AA)
C7CB  11 F6 FF           ld de,-10
C7CE  19                 add hl,de
C7CF  5E                 ld e,(hl)
C7D0  23                 inc hl
C7D1  56                 ld d,(hl)
C7D2  23                 inc hl
C7D3  4E                 ld c,(hl)
C7D4  23                 inc hl
C7D5  46                 ld b,(hl)
C7D6  ED 43 DA D2        ld ($D2DA),bc
C7DA  21 E8 FF           ld hl,-24
C7DD  19                 add hl,de
C7DE  22 D8 D2           ld ($D2D8),hl
C7E1  C9                 ret
C7E2  CD DA CC           call $CCDA
C7E5  FD 21 10 D2        ld iy,$D210
C7E9  CD B7 C4           call $C4B7
C7EC  38 08              jr c,$C7F6
C7EE  FD 21 38 D2        ld iy,$D238
C7F2  CD B7 C4           call $C4B7
C7F5  D0                 ret nc
C7F6  AF                 xor a
C7F7  32 74 B0           ld ($B074),a
C7FA  C9                 ret
C7FB  DD 21 CA 7F        ld ix,$7FCA
C7FF  FD 21 10 D2        ld iy,$D210
C803  CD 2F C9           call $C92F
C806  DD 21 D6 7F        ld ix,$7FD6
C80A  FD 21 38 D2        ld iy,$D238
C80E  CD 2F C9           call $C92F
C811  CD 75 CE           call $CE75
C814  3A D0 7F           ld a,($7FD0)
C817  A7                 and a
C818  28 0E              jr z,$C828
C81A  3A DC 7F           ld a,($7FDC)
C81D  A7                 and a
C81E  28 08              jr z,$C828
C820  7E                 ld a,(hl)
C821  FE 81              cp -127
C823  C8                 ret z
C824  35                 dec (hl)
C825  23                 inc hl
C826  35                 dec (hl)
C827  C9                 ret
C828  7E                 ld a,(hl)
C829  FE 84              cp -124
C82B  C8                 ret z
C82C  34                 inc (hl)
C82D  23                 inc hl
C82E  34                 inc (hl)
C82F  C9                 ret
C830  DD 21 CA 7F        ld ix,$7FCA
C834  FD 21 10 D2        ld iy,$D210
C838  CD 2F C9           call $C92F
C83B  DD 21 D6 7F        ld ix,$7FD6
C83F  FD 21 38 D2        ld iy,$D238
C843  CD 2F C9           call $C92F
C846  2A D8 D2           ld hl,($D2D8)
C849  E5                 push hl
C84A  11 F0 FF           ld de,-16
C84D  19                 add hl,de
C84E  22 D8 D2           ld ($D2D8),hl
C851  CD 75 CE           call $CE75
C854  E1                 pop hl
C855  22 D8 D2           ld ($D2D8),hl
C858  3A E1 D2           ld a,($D2E1)
C85B  FE 05              cp 5
C85D  CC D1 C8           call z,$C8D1
C860  3A E1 D2           ld a,($D2E1)
C863  FE 05              cp 5
C865  C4 FD C8           call nz,$C8FD
C868  2A D8 D2           ld hl,($D2D8)
C86B  ED 4B DE D2        ld bc,($D2DE)
C86F  3A E1 D2           ld a,($D2E1)
C872  FE 07              cp 7
C874  C4 C8 C8           call nz,$C8C8
C877  22 D8 D2           ld ($D2D8),hl
C87A  2A DA D2           ld hl,($D2DA)
C87D  ED 4B E0 D2        ld bc,($D2E0)
C881  CD C8 C8           call $C8C8
C884  22 DA D2           ld ($D2DA),hl
C887  DD 21 CA 7F        ld ix,$7FCA
C88B  FD 21 10 D2        ld iy,$D210
C88F  CD 9A C8           call $C89A
C892  DD 21 D6 7F        ld ix,$7FD6
C896  FD 21 38 D2        ld iy,$D238
C89A  DD 7E 06           ld a,(ix+6)
C89D  A7                 and a
C89E  C0                 ret nz
C89F  FD CB 08 FE        set 7,(iy+8)
C8A3  2A D8 D2           ld hl,($D2D8)
C8A6  11 F0 FF           ld de,-16
C8A9  19                 add hl,de
C8AA  FD 75 00           ld (iy+0),l
C8AD  FD 74 01           ld (iy+1),h
C8B0  06 FF              ld b,-1
C8B2  3A E1 D2           ld a,($D2E1)
C8B5  FE 05              cp 5
C8B7  28 0A              jr z,$C8C3
C8B9  06 FD              ld b,-3
C8BB  3A E0 D2           ld a,($D2E0)
C8BE  17                 rla
C8BF  38 02              jr c,$C8C3
C8C1  06 FE              ld b,-2
C8C3  78                 ld a,b
C8C4  DD 77 01           ld (ix+1),a
C8C7  C9                 ret
C8C8  06 00              ld b,0
C8CA  79                 ld a,c
C8CB  17                 rla
C8CC  30 01              jr nc,$C8CF
C8CE  05                 dec b
C8CF  09                 add hl,bc
C8D0  C9                 ret
C8D1  21 DE D2           ld hl,$D2DE
C8D4  7E                 ld a,(hl)
C8D5  A7                 and a
C8D6  20 02              jr nz,$C8DA
C8D8  36 04              ld (hl),4
C8DA  3A 99 B0           ld a,($B099)
C8DD  E6 90              and $90
C8DF  28 04              jr z,$C8E5
C8E1  7E                 ld a,(hl)
C8E2  ED 44              neg
C8E4  77                 ld (hl),a
C8E5  2A D8 D2           ld hl,($D2D8)
C8E8  7C                 ld a,h
C8E9  A7                 and a
C8EA  C0                 ret nz
C8EB  7D                 ld a,l
C8EC  FE 50              cp 80
C8EE  D0                 ret nc
C8EF  3E 04              ld a,4
C8F1  32 DE D2           ld ($D2DE),a
C8F4  C9                 ret
C8F5  00                 nop
C8F6  00                 nop
C8F7  07                 rlca
C8F8  06 08              ld b,8
C8FA  08                 ex af,af'
C8FB  08                 ex af,af'
C8FC  08                 ex af,af'
C8FD  21 E0 D2           ld hl,$D2E0
C900  7E                 ld a,(hl)
C901  A7                 and a
C902  20 07              jr nz,$C90B
C904  36 FF              ld (hl),-1
C906  3E 02              ld a,2
C908  32 DE D2           ld ($D2DE),a
C90B  3A 98 B0           ld a,($B098)
C90E  CD 2A C9           call $C92A
C911  28 0F              jr z,$C922
C913  3A 9A B0           ld a,($B09A)
C916  CD 2A C9           call $C92A
C919  C0                 ret nz
C91A  36 FF              ld (hl),-1
C91C  3E 02              ld a,2
C91E  32 DE D2           ld ($D2DE),a
C921  C9                 ret
C922  36 01              ld (hl),1
C924  3E FE              ld a,-2
C926  32 DE D2           ld ($D2DE),a
C929  C9                 ret
C92A  E6 07              and $07
C92C  FE 06              cp 6
C92E  C9                 ret
C92F  3E FF              ld a,-1
C931  DD 77 06           ld (ix+6),a
C934  FD 7E 06           ld a,(iy+6)
C937  A7                 and a
C938  F8                 ret m
C939  2A DA D2           ld hl,($D2DA)
C93C  FD 5E 02           ld e,(iy+2)
C93F  FD 56 03           ld d,(iy+3)
C942  A7                 and a
C943  ED 52              sbc hl,de
C945  11 03 00           ld de,3
C948  19                 add hl,de
C949  7C                 ld a,h
C94A  A7                 and a
C94B  C0                 ret nz
C94C  7D                 ld a,l
C94D  FE 09              cp 9
C94F  D0                 ret nc
C950  FD 6E 00           ld l,(iy+0)
C953  FD 66 01           ld h,(iy+1)
C956  FD 5E 06           ld e,(iy+6)
C959  16 00              ld d,0
C95B  19                 add hl,de
C95C  EB                 ex de,hl
C95D  2A D8 D2           ld hl,($D2D8)
C960  A7                 and a
C961  ED 52              sbc hl,de
C963  7C                 ld a,h
C964  A7                 and a
C965  C0                 ret nz
C966  7D                 ld a,l
C967  FE 11              cp 17
C969  D0                 ret nc
C96A  AF                 xor a
C96B  DD 77 06           ld (ix+6),a
C96E  C9                 ret
C96F  FD 21 60 D2        ld iy,$D260
C973  06 0C              ld b,12
C975  11 0A 00           ld de,10
C978  FD 7E 09           ld a,(iy+9)
C97B  3C                 inc a
C97C  C8                 ret z
C97D  FD 19              add iy,de
C97F  10 F7              djnz $C978
C981  A7                 and a
C982  C9                 ret
C983  FD 2A AA B0        ld iy,($B0AA)
C987  11 14 00           ld de,20
C98A  FD 19              add iy,de
C98C  FD 7E 09           ld a,(iy+9)
C98F  3C                 inc a
C990  C0                 ret nz
C991  DD 7E 00           ld a,(ix+0)
C994  A7                 and a
C995  C8                 ret z
C996  D6 01              sub 1
C998  27                 daa
C999  DD 77 00           ld (ix+0),a
C99C  FD 36 09 03        ld (iy+9),3
C9A0  C3 59 CA           jp $CA59
C9A3  CD 35 C6           call $C635
C9A6  CD DA CC           call $CCDA
C9A9  CD 13 C5           call $C513
C9AC  21 E0 D2           ld hl,$D2E0
C9AF  35                 dec (hl)
C9B0  C0                 ret nz
C9B1  3E FF              ld a,-1
C9B3  32 E1 D2           ld ($D2E1),a
C9B6  C9                 ret
C9B7  CD 6F C9           call $C96F
C9BA  C0                 ret nz
C9BB  FD 36 08 1E        ld (iy+8),30
C9BF  FD 36 09 0C        ld (iy+9),12
C9C3  2A DA D2           ld hl,($D2DA)
C9C6  FD 75 02           ld (iy+2),l
C9C9  FD 74 03           ld (iy+3),h
C9CC  2A D8 D2           ld hl,($D2D8)
C9CF  11 F4 FF           ld de,-12
C9D2  19                 add hl,de
C9D3  FD 75 00           ld (iy+0),l
C9D6  FD 74 01           ld (iy+1),h
C9D9  3A DC D2           ld a,($D2DC)
C9DC  E6 80              and $80
C9DE  FD 77 04           ld (iy+4),a
C9E1  C9                 ret
C9E2  CD AA C5           call $C5AA
C9E5  01 01 00           ld bc,1
C9E8  3A DC D2           ld a,($D2DC)
C9EB  17                 rla
C9EC  38 02              jr c,$C9F0
C9EE  0B                 dec bc
C9EF  0B                 dec bc
C9F0  2A DA D2           ld hl,($D2DA)
C9F3  09                 add hl,bc
C9F4  22 DA D2           ld ($D2DA),hl
C9F7  2A D8 D2           ld hl,($D2D8)
C9FA  01 04 00           ld bc,4
C9FD  09                 add hl,bc
C9FE  22 D8 D2           ld ($D2D8),hl
CA01  21 E0 D2           ld hl,$D2E0
CA04  35                 dec (hl)
CA05  C0                 ret nz
CA06  C3 C8 CA           jp $CAC8
CA09  CD 13 C5           call $C513
CA0C  3A DC D2           ld a,($D2DC)
CA0F  F5                 push af
CA10  06 03              ld b,3
CA12  C5                 push bc
CA13  CD 35 C6           call $C635
CA16  CD D3 CC           call $CCD3
CA19  C1                 pop bc
CA1A  10 F6              djnz $CA12
CA1C  C1                 pop bc
CA1D  3A DC D2           ld a,($D2DC)
CA20  A8                 xor b
CA21  17                 rla
CA22  38 05              jr c,$CA29
CA24  21 E0 D2           ld hl,$D2E0
CA27  35                 dec (hl)
CA28  C0                 ret nz
CA29  3E FF              ld a,-1
CA2B  32 E1 D2           ld ($D2E1),a
CA2E  C9                 ret
CA2F  CD 35 C6           call $C635
CA32  CD DA CC           call $CCDA
CA35  21 E0 D2           ld hl,$D2E0
CA38  35                 dec (hl)
CA39  C0                 ret nz
CA3A  C3 C8 CA           jp $CAC8
CA3D  CD 6F C9           call $C96F
CA40  C0                 ret nz
CA41  18 08              jr $CA4B
CA43  3E 9F              ld a,-97
CA45  32 DF D2           ld ($D2DF),a
CA48  CD B7 C5           call $C5B7
CA4B  21 99 FF           ld hl,-103
CA4E  ED 4B E1 D2        ld bc,($D2E1)
CA52  06 00              ld b,0
CA54  09                 add hl,bc
CA55  7E                 ld a,(hl)
CA56  FD 77 09           ld (iy+9),a
CA59  FD 36 08 1E        ld (iy+8),30
CA5D  FD 7E 09           ld a,(iy+9)
CA60  FE 04              cp 4
CA62  20 04              jr nz,$CA68
CA64  FD 36 08 19        ld (iy+8),25
CA68  FD 36 05 01        ld (iy+5),1
CA6C  2A DA D2           ld hl,($D2DA)
CA6F  FD 75 02           ld (iy+2),l
CA72  FD 74 03           ld (iy+3),h
CA75  2A D8 D2           ld hl,($D2D8)
CA78  FD 75 00           ld (iy+0),l
CA7B  FD 74 01           ld (iy+1),h
CA7E  FD 36 06 F3        ld (iy+6),-13
CA82  3A DC D2           ld a,($D2DC)
CA85  E6 80              and $80
CA87  FD 77 04           ld (iy+4),a
CA8A  3E 01              ld a,1
CA8C  32 FF E9           ld ($E9FF),a
CA8F  AF                 xor a
CA90  C9                 ret
CA91  FD 36 09 1E        ld (iy+9),30
CA95  FD 36 08 05        ld (iy+8),5
CA99  FD 36 04 1B        ld (iy+4),27
CA9D  C9                 ret
CA9E  21 E0 D2           ld hl,$D2E0
CAA1  35                 dec (hl)
CAA2  C0                 ret nz
CAA3  3A DD D2           ld a,($D2DD)
CAA6  FE 00              cp 0
CAA8  28 1E              jr z,$CAC8
CAAA  3E 1F              ld a,31
CAAC  32 E1 D2           ld ($D2E1),a
CAAF  CD DE E9           call $E9DE
CAB2  E6 07              and $07
CAB4  32 DE D2           ld ($D2DE),a
CAB7  FE 04              cp 4
CAB9  30 0D              jr nc,$CAC8
CABB  CD 03 C6           call $C603
CABE  AF                 xor a
CABF  CD 17 C6           call $C617
CAC2  3E 30              ld a,48
CAC4  32 E0 D2           ld ($D2E0),a
CAC7  C9                 ret
CAC8  3E FF              ld a,-1
CACA  32 E1 D2           ld ($D2E1),a
CACD  C9                 ret
CACE  3A DE D2           ld a,($D2DE)
CAD1  CD 03 C6           call $C603
CAD4  21 E0 D2           ld hl,$D2E0
CAD7  35                 dec (hl)
CAD8  28 EE              jr z,$CAC8
CADA  7E                 ld a,(hl)
CADB  FE 18              cp 24
CADD  30 0A              jr nc,$CAE9
CADF  E6 01              and $01
CAE1  20 06              jr nz,$CAE9
CAE3  3D                 dec a
CAE4  E6 7F              and $7F
CAE6  32 DC D2           ld ($D2DC),a
CAE9  FD 21 10 D2        ld iy,$D210
CAED  DD 21 CA 7F        ld ix,$7FCA
CAF1  CD B7 C4           call $C4B7
CAF4  38 0C              jr c,$CB02
CAF6  FD 21 38 D2        ld iy,$D238
CAFA  DD 21 D6 7F        ld ix,$7FD6
CAFE  CD B7 C4           call $C4B7
CB01  D0                 ret nc
CB02  3E 02              ld a,2
CB04  32 FF E9           ld ($E9FF),a
CB07  CD C8 CA           call $CAC8
CB0A  3A DE D2           ld a,($D2DE)
CB0D  E6 03              and $03
CB0F  11 15 CB           ld de,$CB15
CB12  C3 00 B3           jp $B300
CB15  2B                 dec hl
CB16  CB 36              sll (hl)
CB18  CB 36              sll (hl)
CB1A  CB 1D              rr l
CB1C  CB DD              set 3,l
CB1E  7E                 ld a,(hl)
CB1F  00                 nop
CB20  C6 10              add a,16
CB22  27                 daa
CB23  30 02              jr nc,$CB27
CB25  3E 99              ld a,-103
CB27  DD 77 00           ld (ix+0),a
CB2A  C9                 ret
CB2B  FD 7E 08           ld a,(iy+8)
CB2E  E6 80              and $80
CB30  F6 40              or $40
CB32  FD 77 08           ld (iy+8),a
CB35  C9                 ret
CB36  21 60 00           ld hl,96
CB39  C3 9D B6           jp $B69D
CB3C  2A DA D2           ld hl,($D2DA)
CB3F  11 D0 01           ld de,464
CB42  A7                 and a
CB43  ED 52              sbc hl,de
CB45  3E FE              ld a,-2
CB47  20 02              jr nz,$CB4B
CB49  3E F6              ld a,-10
CB4B  CD 7A CB           call $CB7A
CB4E  2A DA D2           ld hl,($D2DA)
CB51  11 F0 01           ld de,496
CB54  A7                 and a
CB55  ED 52              sbc hl,de
CB57  D8                 ret c
CB58  3E 1D              ld a,29
CB5A  32 E1 D2           ld ($D2E1),a
CB5D  AF                 xor a
CB5E  32 74 B0           ld ($B074),a
CB61  C9                 ret
CB62  2A AA B0           ld hl,($B0AA)
CB65  11 10 D2           ld de,$D210
CB68  A7                 and a
CB69  ED 52              sbc hl,de
CB6B  20 06              jr nz,$CB73
CB6D  DD 21 CA 7F        ld ix,$7FCA
CB71  18 04              jr $CB77
CB73  DD 21 D6 7F        ld ix,$7FD6
CB77  DD 7E 0B           ld a,(ix+11)
CB7A  DD 4E 02           ld c,(ix+2)
CB7D  DD 77 02           ld (ix+2),a
CB80  CB 51              bit 2,c
CB82  20 02              jr nz,$CB86
CB84  CB D7              set 2,a
CB86  21 DF D2           ld hl,$D2DF
CB89  AE                 xor (hl)
CB8A  E6 1F              and $1F
CB8C  AE                 xor (hl)
CB8D  77                 ld (hl),a
CB8E  CB 57              bit 2,a
CB90  CC 83 C9           call z,$C983
CB93  21 E0 D2           ld hl,$D2E0
CB96  CB 7E              bit 7,(hl)
CB98  28 24              jr z,$CBBE
CB9A  CB BE              res 7,(hl)
CB9C  AF                 xor a
CB9D  32 9C B0           ld ($B09C),a
CBA0  FD 2A AA B0        ld iy,($B0AA)
CBA4  FD 36 13 FF        ld (iy+19),-1
CBA8  CD D3 CC           call $CCD3
CBAB  DD 7E 01           ld a,(ix+1)
CBAE  32 DF D2           ld ($D2DF),a
CBB1  3A DC D2           ld a,($D2DC)
CBB4  F5                 push af
CBB5  CD D3 CC           call $CCD3
CBB8  F1                 pop af
CBB9  32 DC D2           ld ($D2DC),a
CBBC  18 03              jr $CBC1
CBBE  CD DA CC           call $CCDA
CBC1  FD 2A AA B0        ld iy,($B0AA)
CBC5  FD 7E 13           ld a,(iy+19)
CBC8  3C                 inc a
CBC9  28 0A              jr z,$CBD5
CBCB  3E 02              ld a,2
CBCD  32 DE D2           ld ($D2DE),a
CBD0  3E 09              ld a,9
CBD2  CD 0E C6           call $C60E
CBD5  21 E0 D2           ld hl,$D2E0
CBD8  7E                 ld a,(hl)
CBD9  E6 80              and $80
CBDB  47                 ld b,a
CBDC  7E                 ld a,(hl)
CBDD  E6 7F              and $7F
CBDF  28 0C              jr z,$CBED
CBE1  3D                 dec a
CBE2  B0                 or b
CBE3  77                 ld (hl),a
CBE4  E6 01              and $01
CBE6  20 05              jr nz,$CBED
CBE8  3E 7F              ld a,127
CBEA  CD 0E C6           call $C60E
CBED  3A DE D2           ld a,($D2DE)
CBF0  A7                 and a
CBF1  FA 00 CC           jp m,$CC00
CBF4  FE 0C              cp 12
CBF6  38 08              jr c,$CC00
CBF8  FD 2A AA B0        ld iy,($B0AA)
CBFC  FD 36 13 01        ld (iy+19),1
CC00  3A 9C B0           ld a,($B09C)
CC03  A7                 and a
CC04  20 10              jr nz,$CC16
CC06  FD 2A AA B0        ld iy,($B0AA)
CC0A  FD 36 13 FF        ld (iy+19),-1
CC0E  3A 99 B0           ld a,($B099)
CC11  E6 80              and $80
CC13  C2 7E CC           jp nz,$CC7E
CC16  FD 21 F4 CF        ld iy,$CFF4
CC1A  11 0A 00           ld de,10
CC1D  06 4A              ld b,74
CC1F  FD 7E 09           ld a,(iy+9)
CC22  FE 1C              cp 28
CC24  D2 33 CC           jp nc,$CC33
CC27  FE 08              cp 8
CC29  DA 33 CC           jp c,$CC33
CC2C  CD B7 C4           call $C4B7
CC2F  38 45              jr c,$CC76
CC31  18 0F              jr $CC42
CC33  FE 26              cp 38
CC35  D2 45 CC           jp nc,$CC45
CC38  FE 23              cp 35
CC3A  DA 45 CC           jp c,$CC45
CC3D  CD E5 C4           call $C4E5
CC40  38 3C              jr c,$CC7E
CC42  11 0A 00           ld de,10
CC45  FD 19              add iy,de
CC47  10 D6              djnz $CC1F
CC49  3A 9C B0           ld a,($B09C)
CC4C  A7                 and a
CC4D  C0                 ret nz
CC4E  CD 75 CE           call $CE75
CC51  3A 99 B0           ld a,($B099)
CC54  E6 90              and $90
CC56  EE 10              xor $10
CC58  C0                 ret nz
CC59  2A AA B0           ld hl,($B0AA)
CC5C  11 10 D2           ld de,$D210
CC5F  A7                 and a
CC60  ED 52              sbc hl,de
CC62  11 D1 7F           ld de,$7FD1
CC65  28 03              jr z,$CC6A
CC67  11 DD 7F           ld de,$7FDD
CC6A  21 D8 D2           ld hl,$D2D8
CC6D  ED A0              ldi
CC6F  ED A0              ldi
CC71  ED A0              ldi
CC73  ED A0              ldi
CC75  C9                 ret
CC76  3A E0 D2           ld a,($D2E0)
CC79  E6 7F              and $7F
CC7B  C2 47 C5           jp nz,$C547
CC7E  21 E0 D2           ld hl,$D2E0
CC81  7E                 ld a,(hl)
CC82  E6 7F              and $7F
CC84  C0                 ret nz
CC85  FD 2A AA B0        ld iy,($B0AA)
CC89  FD 36 13 FF        ld (iy+19),-1
CC8D  3E 02              ld a,2
CC8F  32 E1 D2           ld ($D2E1),a
CC92  3E 1B              ld a,27
CC94  32 DC D2           ld ($D2DC),a
CC97  3E F5              ld a,-11
CC99  32 DE D2           ld ($D2DE),a
CC9C  C9                 ret
CC9D  CD DA CC           call $CCDA
CCA0  3A 9C B0           ld a,($B09C)
CCA3  A7                 and a
CCA4  C0                 ret nz
CCA5  2A AA B0           ld hl,($B0AA)
CCA8  11 10 D2           ld de,$D210
CCAB  A7                 and a
CCAC  ED 52              sbc hl,de
CCAE  21 D1 7F           ld hl,$7FD1
CCB1  28 03              jr z,$CCB6
CCB3  21 DD 7F           ld hl,$7FDD
CCB6  11 D8 D2           ld de,$D2D8
CCB9  ED A0              ldi
CCBB  ED A0              ldi
CCBD  ED A0              ldi
CCBF  ED A0              ldi
CCC1  3E 00              ld a,0
CCC3  32 E1 D2           ld ($D2E1),a
CCC6  21 DD D2           ld hl,$D2DD
CCC9  35                 dec (hl)
CCCA  CA C8 CA           jp z,$CAC8
CCCD  3E 40              ld a,64
CCCF  32 E0 D2           ld ($D2E0),a
CCD2  C9                 ret
CCD3  AF                 xor a
CCD4  32 9C B0           ld ($B09C),a
CCD7  C3 96 CD           jp $CD96
CCDA  3E FF              ld a,-1
CCDC  32 9C B0           ld ($B09C),a
CCDF  21 DE D2           ld hl,$D2DE
CCE2  7E                 ld a,(hl)
CCE3  3C                 inc a
CCE4  3C                 inc a
CCE5  FE 0F              cp 15
CCE7  20 02              jr nz,$CCEB
CCE9  3D                 dec a
CCEA  3D                 dec a
CCEB  77                 ld (hl),a
CCEC  A7                 and a
CCED  FA 88 CD           jp m,$CD88
CCF0  CD 75 CE           call $CE75
CCF3  3A 99 B0           ld a,($B099)
CCF6  E6 10              and $10
CCF8  20 0A              jr nz,$CD04
CCFA  3A 99 B0           ld a,($B099)
CCFD  E6 20              and $20
CCFF  20 0E              jr nz,$CD0F
CD01  C3 85 CD           jp $CD85
CD04  3A D8 D2           ld a,($D2D8)
CD07  E6 0F              and $0F
CD09  EE 0F              xor $0F
CD0B  47                 ld b,a
CD0C  C3 56 CD           jp $CD56
CD0F  3A 9B B0           ld a,($B09B)
CD12  0E 00              ld c,0
CD14  FE 0C              cp 12
CD16  28 06              jr z,$CD1E
CD18  FE 0F              cp 15
CD1A  28 02              jr z,$CD1E
CD1C  0E 08              ld c,8
CD1E  FE 0E              cp 14
CD20  30 04              jr nc,$CD26
CD22  79                 ld a,c
CD23  C6 10              add a,16
CD25  4F                 ld c,a
CD26  3A DA D2           ld a,($D2DA)
CD29  3C                 inc a
CD2A  E6 03              and $03
CD2C  87                 add a,a
CD2D  81                 add a,c
CD2E  4F                 ld c,a
CD2F  2A D8 D2           ld hl,($D2D8)
CD32  7D                 ld a,l
CD33  E6 F0              and $F0
CD35  6F                 ld l,a
CD36  06 00              ld b,0
CD38  09                 add hl,bc
CD39  01 FE FF           ld bc,-2
CD3C  09                 add hl,bc
CD3D  EB                 ex de,hl
CD3E  2A D8 D2           ld hl,($D2D8)
CD41  ED 4B DE D2        ld bc,($D2DE)
CD45  06 00              ld b,0
CD47  09                 add hl,bc
CD48  A7                 and a
CD49  ED 52              sbc hl,de
CD4B  19                 add hl,de
CD4C  38 37              jr c,$CD85
CD4E  ED 53 D8 D2        ld ($D2D8),de
CD52  06 00              ld b,0
CD54  18 06              jr $CD5C
CD56  3A DE D2           ld a,($D2DE)
CD59  B8                 cp b
CD5A  38 2C              jr c,$CD88
CD5C  AF                 xor a
CD5D  32 9C B0           ld ($B09C),a
CD60  ED 5B E1 D2        ld de,($D2E1)
CD64  16 00              ld d,0
CD66  21 27 FF           ld hl,-217
CD69  19                 add hl,de
CD6A  7E                 ld a,(hl)
CD6B  E6 04              and $04
CD6D  28 12              jr z,$CD81
CD6F  3A DE D2           ld a,($D2DE)
CD72  57                 ld d,a
CD73  CB 3A              srl d
CD75  CB 3A              srl d
CD77  92                 sub d
CD78  3C                 inc a
CD79  ED 44              neg
CD7B  32 DE D2           ld ($D2DE),a
CD7E  78                 ld a,b
CD7F  18 07              jr $CD88
CD81  78                 ld a,b
CD82  32 DE D2           ld ($D2DE),a
CD85  3A DE D2           ld a,($D2DE)
CD88  4F                 ld c,a
CD89  06 00              ld b,0
CD8B  17                 rla
CD8C  30 01              jr nc,$CD8F
CD8E  05                 dec b
CD8F  2A D8 D2           ld hl,($D2D8)
CD92  09                 add hl,bc
CD93  22 D8 D2           ld ($D2D8),hl
CD96  3A DF D2           ld a,($D2DF)
CD99  E6 10              and $10
CD9B  20 26              jr nz,$CDC3
CD9D  FD 21 60 D2        ld iy,$D260
CDA1  06 0C              ld b,12
CDA3  11 0A 00           ld de,10
CDA6  3A E1 D2           ld a,($D2E1)
CDA9  FE 00              cp 0
CDAB  20 0B              jr nz,$CDB8
CDAD  FD 2A AA B0        ld iy,($B0AA)
CDB1  11 1E 00           ld de,30
CDB4  FD 19              add iy,de
CDB6  06 01              ld b,1
CDB8  FD 7E 09           ld a,(iy+9)
CDBB  3C                 inc a
CDBC  CA 43 CA           jp z,$CA43
CDBF  FD 19              add iy,de
CDC1  10 F5              djnz $CDB8
CDC3  21 DF D2           ld hl,$D2DF
CDC6  7E                 ld a,(hl)
CDC7  E6 E0              and $E0
CDC9  FE E0              cp -32
CDCB  28 08              jr z,$CDD5
CDCD  C6 20              add a,32
CDCF  F6 1F              or $1F
CDD1  77                 ld (hl),a
CDD2  C3 B7 C5           jp $C5B7
CDD5  CD 75 CE           call $CE75
CDD8  ED 5B E1 D2        ld de,($D2E1)
CDDC  16 00              ld d,0
CDDE  21 27 FF           ld hl,-217
CDE1  19                 add hl,de
CDE2  7E                 ld a,(hl)
CDE3  11 00 08           ld de,2048
CDE6  E6 01              and $01
CDE8  28 03              jr z,$CDED
CDEA  11 10 98           ld de,$9810
CDED  7E                 ld a,(hl)
CDEE  E6 02              and $02
CDF0  47                 ld b,a
CDF1  3A DF D2           ld a,($D2DF)
CDF4  E6 08              and $08
CDF6  20 11              jr nz,$CE09
CDF8  21 9C B0           ld hl,$B09C
CDFB  7E                 ld a,(hl)
CDFC  A7                 and a
CDFD  20 0A              jr nz,$CE09
CDFF  3E F3              ld a,-13
CE01  32 DE D2           ld ($D2DE),a
CE04  3E 00              ld a,0
CE06  32 FF E9           ld ($E9FF),a
CE09  3A DF D2           ld a,($D2DF)
CE0C  E6 02              and $02
CE0E  20 1D              jr nz,$CE2D
CE10  3A 98 B0           ld a,($B098)
CE13  A2                 and d
CE14  AB                 xor e
CE15  20 0D              jr nz,$CE24
CE17  2A DA D2           ld hl,($D2DA)
CE1A  2B                 dec hl
CE1B  22 DA D2           ld ($D2DA),hl
CE1E  AF                 xor a
CE1F  CD 17 C6           call $C617
CE22  18 09              jr $CE2D
CE24  78                 ld a,b
CE25  A7                 and a
CE26  28 05              jr z,$CE2D
CE28  3E 80              ld a,-128
CE2A  CD 17 C6           call $C617
CE2D  3A DF D2           ld a,($D2DF)
CE30  E6 01              and $01
CE32  20 1D              jr nz,$CE51
CE34  3A 9A B0           ld a,($B09A)
CE37  A2                 and d
CE38  AB                 xor e
CE39  20 0E              jr nz,$CE49
CE3B  2A DA D2           ld hl,($D2DA)
CE3E  23                 inc hl
CE3F  22 DA D2           ld ($D2DA),hl
CE42  3E 80              ld a,-128
CE44  CD 17 C6           call $C617
CE47  18 08              jr $CE51
CE49  78                 ld a,b
CE4A  A7                 and a
CE4B  28 04              jr z,$CE51
CE4D  AF                 xor a
CE4E  CD 17 C6           call $C617
CE51  3A 9C B0           ld a,($B09C)
CE54  A7                 and a
CE55  20 13              jr nz,$CE6A
CE57  3A DF D2           ld a,($D2DF)
CE5A  E6 03              and $03
CE5C  EE 03              xor $03
CE5E  28 05              jr z,$CE65
CE60  CD C4 C5           call $C5C4
CE63  18 08              jr $CE6D
CE65  C3 AA C5           jp $C5AA
CE68  18 03              jr $CE6D
CE6A  CD 9D C5           call $C59D
CE6D  21 DF D2           ld hl,$D2DF
CE70  7E                 ld a,(hl)
CE71  F6 1F              or $1F
CE73  77                 ld (hl),a
CE74  C9                 ret
CE75  2A D8 D2           ld hl,($D2D8)
CE78  7D                 ld a,l
CE79  E6 F0              and $F0
CE7B  6F                 ld l,a
CE7C  29                 add hl,hl
CE7D  29                 add hl,hl
CE7E  29                 add hl,hl
CE7F  11 80 EB           ld de,$EB80
CE82  19                 add hl,de
CE83  ED 4B DA D2        ld bc,($D2DA)
CE87  79                 ld a,c
CE88  E6 03              and $03
CE8A  5F                 ld e,a
CE8B  CB 38              srl b
CE8D  CB 19              rr c
CE8F  CB 38              srl b
CE91  CB 19              rr c
CE93  09                 add hl,bc
CE94  01 00 E9           ld bc,$E900
CE97  4E                 ld c,(hl)
CE98  0A                 ld a,(bc)
CE99  32 98 B0           ld ($B098),a
CE9C  7B                 ld a,e
CE9D  FE 03              cp 3
CE9F  20 01              jr nz,$CEA2
CEA1  23                 inc hl
CEA2  4E                 ld c,(hl)
CEA3  0A                 ld a,(bc)
CEA4  57                 ld d,a
CEA5  79                 ld a,c
CEA6  32 9B B0           ld ($B09B),a
CEA9  7B                 ld a,e
CEAA  3D                 dec a
CEAB  FE 02              cp 2
CEAD  30 01              jr nc,$CEB0
CEAF  23                 inc hl
CEB0  4E                 ld c,(hl)
CEB1  0A                 ld a,(bc)
CEB2  B2                 or d
CEB3  32 99 B0           ld ($B099),a
CEB6  7B                 ld a,e
CEB7  A7                 and a
CEB8  20 01              jr nz,$CEBB
CEBA  23                 inc hl
CEBB  4E                 ld c,(hl)
CEBC  0A                 ld a,(bc)
CEBD  32 9A B0           ld ($B09A),a
CEC0  C9                 ret
CEC1  ED 4B 80 B0        ld bc,($B080)
CEC5  78                 ld a,b
CEC6  D6 10              sub 16
CEC8  47                 ld b,a
CEC9  79                 ld a,c
CECA  D6 0A              sub 10
CECC  4F                 ld c,a
CECD  2A 78 B0           ld hl,($B078)
CED0  3A 76 B0           ld a,($B076)
CED3  95                 sub l
CED4  ED 44              neg
CED6  D6 02              sub 2
CED8  F2 E7 CE           jp p,$CEE7
CEDB  7D                 ld a,l
CEDC  D6 06              sub 6
CEDE  F2 E2 CE           jp p,$CEE2
CEE1  AF                 xor a
CEE2  32 7A B0           ld ($B07A),a
CEE5  18 0F              jr $CEF6
CEE7  D6 06              sub 6
CEE9  FA F6 CE           jp m,$CEF6
CEEC  7D                 ld a,l
CEED  D6 03              sub 3
CEEF  B9                 cp c
CEF0  38 01              jr c,$CEF3
CEF2  79                 ld a,c
CEF3  32 7A B0           ld ($B07A),a
CEF6  3A 77 B0           ld a,($B077)
CEF9  94                 sub h
CEFA  ED 44              neg
CEFC  D6 03              sub 3
CEFE  F2 0D CF           jp p,$CF0D
CF01  7C                 ld a,h
CF02  D6 0A              sub 10
CF04  F2 08 CF           jp p,$CF08
CF07  AF                 xor a
CF08  32 7B B0           ld ($B07B),a
CF0B  18 0F              jr $CF1C
CF0D  D6 09              sub 9
CF0F  FA 1C CF           jp m,$CF1C
CF12  7C                 ld a,h
CF13  D6 05              sub 5
CF15  B8                 cp b
CF16  38 01              jr c,$CF19
CF18  78                 ld a,b
CF19  32 7B B0           ld ($B07B),a
CF1C  ED 5B E2 D2        ld de,($D2E2)
CF20  A7                 and a
CF21  ED 52              sbc hl,de
CF23  20 0E              jr nz,$CF33
CF25  19                 add hl,de
CF26  3A AD B0           ld a,($B0AD)
CF29  3C                 inc a
CF2A  32 AD B0           ld ($B0AD),a
CF2D  FE 12              cp 18
CF2F  30 0D              jr nc,$CF3E
CF31  18 08              jr $CF3B
CF33  19                 add hl,de
CF34  22 E2 D2           ld ($D2E2),hl
CF37  AF                 xor a
CF38  32 AD B0           ld ($B0AD),a
CF3B  C3 58 CF           jp $CF58
CF3E  7C                 ld a,h
CF3F  D6 07              sub 7
CF41  30 01              jr nc,$CF44
CF43  AF                 xor a
CF44  B8                 cp b
CF45  38 01              jr c,$CF48
CF47  78                 ld a,b
CF48  32 7B B0           ld ($B07B),a
CF4B  7D                 ld a,l
CF4C  D6 05              sub 5
CF4E  30 01              jr nc,$CF51
CF50  AF                 xor a
CF51  B9                 cp c
CF52  38 01              jr c,$CF55
CF54  79                 ld a,c
CF55  32 7A B0           ld ($B07A),a
CF58  ED 5B 82 B0        ld de,($B082)
CF5C  3A 84 B0           ld a,($B084)
CF5F  A7                 and a
CF60  28 18              jr z,$CF7A
CF62  7D                 ld a,l
CF63  BB                 cp e
CF64  38 34              jr c,$CF9A
CF66  7C                 ld a,h
CF67  92                 sub d
CF68  38 30              jr c,$CF9A
CF6A  D6 03              sub 3
CF6C  38 2C              jr c,$CF9A
CF6E  AF                 xor a
CF6F  32 84 B0           ld ($B084),a
CF72  3E 02              ld a,2
CF74  32 95 B0           ld ($B095),a
CF77  CD A7 7B           call $7BA7
CF7A  3A 19 D2           ld a,($D219)
CF7D  FE 1C              cp 28
CF7F  28 19              jr z,$CF9A
CF81  FE 1D              cp 29
CF83  28 15              jr z,$CF9A
CF85  3A 72 B0           ld a,($B072)
CF88  A7                 and a
CF89  28 04              jr z,$CF8F
CF8B  7B                 ld a,e
CF8C  32 7A B0           ld ($B07A),a
CF8F  3A 72 B0           ld a,($B072)
CF92  FE 02              cp 2
CF94  28 04              jr z,$CF9A
CF96  7A                 ld a,d
CF97  32 7B B0           ld ($B07B),a
CF9A  06 FF              ld b,-1
CF9C  2A 76 B0           ld hl,($B076)
CF9F  ED 5B 7A B0        ld de,($B07A)
CFA3  7B                 ld a,e
CFA4  95                 sub l
CFA5  28 08              jr z,$CFAF
CFA7  06 00              ld b,0
CFA9  F2 AE CF           jp p,$CFAE
CFAC  2D                 dec l
CFAD  2D                 dec l
CFAE  2C                 inc l
CFAF  7A                 ld a,d
CFB0  94                 sub h
CFB1  28 08              jr z,$CFBB
CFB3  06 00              ld b,0
CFB5  F2 BA CF           jp p,$CFBA
CFB8  25                 dec h
CFB9  25                 dec h
CFBA  24                 inc h
CFBB  22 76 B0           ld ($B076),hl
CFBE  78                 ld a,b
CFBF  32 75 B0           ld ($B075),a
CFC2  C9                 ret
CFC3  21 FF E9           ld hl,$E9FF
CFC6  5E                 ld e,(hl)
CFC7  1C                 inc e
CFC8  C8                 ret z
CFC9  36 FF              ld (hl),-1
CFCB  21 E2 FF           ld hl,-30
CFCE  16 00              ld d,0
CFD0  19                 add hl,de
CFD1  19                 add hl,de
CFD2  19                 add hl,de
CFD3  46                 ld b,(hl)
CFD4  23                 inc hl
CFD5  4E                 ld c,(hl)
CFD6  23                 inc hl
CFD7  56                 ld d,(hl)
CFD8  1E 00              ld e,0
CFDA  79                 ld a,c
CFDB  3D                 dec a
CFDC  20 FD              jr nz,$CFDB
CFDE  7B                 ld a,e
CFDF  EE 10              xor $10
CFE1  5F                 ld e,a
CFE2  D3 FE              out ($FE),a
CFE4  79                 ld a,c
CFE5  82                 add a,d
CFE6  4F                 ld c,a
CFE7  10 F1              djnz $CFDA
CFE9  CD DE E9           call $E9DE
CFEC  E6 1F              and $1F
CFEE  C6 50              add a,80
CFF0  32 EC FF           ld ($FFEC),a
CFF3  C9                 ret
CFF4  AA                 xor d
CFF5  AA                 xor d
CFF6  AA                 xor d
CFF7  AA                 xor d
CFF8  AA                 xor d
CFF9  AA                 xor d
CFFA  AA                 xor d
CFFB  AA                 xor d
CFFC  AA                 xor d
CFFD  AA                 xor d
CFFE  AA                 xor d
CFFF  AA                 xor d
D000  AA                 xor d
D001  AA                 xor d
D002  AA                 xor d
D003  AA                 xor d
D004  AA                 xor d
D005  AA                 xor d
D006  AA                 xor d
D007  AA                 xor d
D008  AA                 xor d
D009  AA                 xor d
D00A  AA                 xor d
D00B  AA                 xor d
D00C  AA                 xor d
D00D  AA                 xor d
D00E  AA                 xor d
D00F  AA                 xor d
D010  AA                 xor d
D011  AA                 xor d
D012  AA                 xor d
D013  AA                 xor d
D014  AA                 xor d
D015  AA                 xor d
D016  AA                 xor d
D017  AA                 xor d
D018  AA                 xor d
D019  AA                 xor d
D01A  AA                 xor d
D01B  AA                 xor d
D01C  AA                 xor d
D01D  AA                 xor d
D01E  AA                 xor d
D01F  AA                 xor d
D020  AA                 xor d
D021  AA                 xor d
D022  AA                 xor d
D023  AA                 xor d
D024  AA                 xor d
D025  AA                 xor d
D026  AA                 xor d
D027  AA                 xor d
D028  AA                 xor d
D029  AA                 xor d
D02A  AA                 xor d
D02B  AA                 xor d
D02C  AA                 xor d
D02D  AA                 xor d
D02E  AA                 xor d
D02F  AA                 xor d
D030  AA                 xor d
D031  AA                 xor d
D032  AA                 xor d
D033  AA                 xor d
D034  AA                 xor d
D035  AA                 xor d
D036  AA                 xor d
D037  AA                 xor d
D038  AA                 xor d
D039  AA                 xor d
D03A  AA                 xor d
D03B  AA                 xor d
D03C  AA                 xor d
D03D  AA                 xor d
D03E  AA                 xor d
D03F  AA                 xor d
D040  AA                 xor d
D041  AA                 xor d
D042  AA                 xor d
D043  AA                 xor d
D044  AA                 xor d
D045  AA                 xor d
D046  AA                 xor d
D047  AA                 xor d
D048  AA                 xor d
D049  AA                 xor d
D04A  AA                 xor d
D04B  AA                 xor d
D04C  AA                 xor d
D04D  AA                 xor d
D04E  AA                 xor d
D04F  AA                 xor d
D050  AA                 xor d
D051  AA                 xor d
D052  AA                 xor d
D053  AA                 xor d
D054  AA                 xor d
D055  AA                 xor d
D056  AA                 xor d
D057  AA                 xor d
D058  AA                 xor d
D059  AA                 xor d
D05A  AA                 xor d
D05B  AA                 xor d
D05C  AA                 xor d
D05D  AA                 xor d
D05E  AA                 xor d
D05F  AA                 xor d
D060  AA                 xor d
D061  AA                 xor d
D062  AA                 xor d
D063  AA                 xor d
D064  AA                 xor d
D065  AA                 xor d
D066  AA                 xor d
D067  AA                 xor d
D068  AA                 xor d
D069  AA                 xor d
D06A  AA                 xor d
D06B  AA                 xor d
D06C  AA                 xor d
D06D  AA                 xor d
D06E  AA                 xor d
D06F  AA                 xor d
D070  AA                 xor d
D071  AA                 xor d
D072  AA                 xor d
D073  AA                 xor d
D074  AA                 xor d
D075  AA                 xor d
D076  AA                 xor d
D077  AA                 xor d
D078  AA                 xor d
D079  AA                 xor d
D07A  AA                 xor d
D07B  AA                 xor d
D07C  AA                 xor d
D07D  AA                 xor d
D07E  AA                 xor d
D07F  AA                 xor d
D080  AA                 xor d
D081  AA                 xor d
D082  AA                 xor d
D083  AA                 xor d
D084  AA                 xor d
D085  AA                 xor d
D086  AA                 xor d
D087  AA                 xor d
D088  AA                 xor d
D089  AA                 xor d
D08A  AA                 xor d
D08B  AA                 xor d
D08C  AA                 xor d
D08D  AA                 xor d
D08E  AA                 xor d
D08F  AA                 xor d
D090  AA                 xor d
D091  AA                 xor d
D092  AA                 xor d
D093  AA                 xor d
D094  AA                 xor d
D095  AA                 xor d
D096  AA                 xor d
D097  AA                 xor d
D098  AA                 xor d
D099  AA                 xor d
D09A  AA                 xor d
D09B  AA                 xor d
D09C  AA                 xor d
D09D  AA                 xor d
D09E  AA                 xor d
D09F  AA                 xor d
D0A0  AA                 xor d
D0A1  AA                 xor d
D0A2  AA                 xor d
D0A3  AA                 xor d
D0A4  AA                 xor d
D0A5  AA                 xor d
D0A6  AA                 xor d
D0A7  AA                 xor d
D0A8  AA                 xor d
D0A9  AA                 xor d
D0AA  AA                 xor d
D0AB  AA                 xor d
D0AC  AA                 xor d
D0AD  AA                 xor d
D0AE  AA                 xor d
D0AF  AA                 xor d
D0B0  AA                 xor d
D0B1  AA                 xor d
D0B2  AA                 xor d
D0B3  AA                 xor d
D0B4  AA                 xor d
D0B5  AA                 xor d
D0B6  AA                 xor d
D0B7  AA                 xor d
D0B8  AA                 xor d
D0B9  AA                 xor d
D0BA  AA                 xor d
D0BB  AA                 xor d
D0BC  AA                 xor d
D0BD  AA                 xor d
D0BE  AA                 xor d
D0BF  AA                 xor d
D0C0  AA                 xor d
D0C1  AA                 xor d
D0C2  AA                 xor d
D0C3  AA                 xor d
D0C4  AA                 xor d
D0C5  AA                 xor d
D0C6  AA                 xor d
D0C7  AA                 xor d
D0C8  AA                 xor d
D0C9  AA                 xor d
D0CA  AA                 xor d
D0CB  AA                 xor d
D0CC  AA                 xor d
D0CD  AA                 xor d
D0CE  AA                 xor d
D0CF  AA                 xor d
D0D0  AA                 xor d
D0D1  AA                 xor d
D0D2  AA                 xor d
D0D3  AA                 xor d
D0D4  AA                 xor d
D0D5  AA                 xor d
D0D6  AA                 xor d
D0D7  AA                 xor d
D0D8  AA                 xor d
D0D9  AA                 xor d
D0DA  AA                 xor d
D0DB  AA                 xor d
D0DC  AA                 xor d
D0DD  AA                 xor d
D0DE  AA                 xor d
D0DF  AA                 xor d
D0E0  AA                 xor d
D0E1  AA                 xor d
D0E2  AA                 xor d
D0E3  AA                 xor d
D0E4  AA                 xor d
D0E5  AA                 xor d
D0E6  AA                 xor d
D0E7  AA                 xor d
D0E8  AA                 xor d
D0E9  AA                 xor d
D0EA  AA                 xor d
D0EB  AA                 xor d
D0EC  AA                 xor d
D0ED  AA                 xor d
D0EE  AA                 xor d
D0EF  AA                 xor d
D0F0  AA                 xor d
D0F1  AA                 xor d
D0F2  AA                 xor d
D0F3  AA                 xor d
D0F4  AA                 xor d
D0F5  AA                 xor d
D0F6  AA                 xor d
D0F7  AA                 xor d
D0F8  AA                 xor d
D0F9  AA                 xor d
D0FA  AA                 xor d
D0FB  AA                 xor d
D0FC  AA                 xor d
D0FD  AA                 xor d
D0FE  AA                 xor d
D0FF  AA                 xor d
D100  AA                 xor d
D101  AA                 xor d
D102  AA                 xor d
D103  AA                 xor d
D104  AA                 xor d
D105  AA                 xor d
D106  AA                 xor d
D107  AA                 xor d
D108  AA                 xor d
D109  AA                 xor d
D10A  AA                 xor d
D10B  AA                 xor d
D10C  AA                 xor d
D10D  AA                 xor d
D10E  AA                 xor d
D10F  AA                 xor d
D110  AA                 xor d
D111  AA                 xor d
D112  AA                 xor d
D113  AA                 xor d
D114  AA                 xor d
D115  AA                 xor d
D116  AA                 xor d
D117  AA                 xor d
D118  AA                 xor d
D119  AA                 xor d
D11A  AA                 xor d
D11B  AA                 xor d
D11C  AA                 xor d
D11D  AA                 xor d
D11E  AA                 xor d
D11F  AA                 xor d
D120  AA                 xor d
D121  AA                 xor d
D122  AA                 xor d
D123  AA                 xor d
D124  AA                 xor d
D125  AA                 xor d
D126  AA                 xor d
D127  AA                 xor d
D128  AA                 xor d
D129  AA                 xor d
D12A  AA                 xor d
D12B  AA                 xor d
D12C  AA                 xor d
D12D  AA                 xor d
D12E  AA                 xor d
D12F  AA                 xor d
D130  AA                 xor d
D131  AA                 xor d
D132  AA                 xor d
D133  AA                 xor d
D134  AA                 xor d
D135  AA                 xor d
D136  AA                 xor d
D137  AA                 xor d
D138  AA                 xor d
D139  AA                 xor d
D13A  AA                 xor d
D13B  AA                 xor d
D13C  AA                 xor d
D13D  AA                 xor d
D13E  AA                 xor d
D13F  AA                 xor d
D140  AA                 xor d
D141  AA                 xor d
D142  AA                 xor d
D143  AA                 xor d
D144  AA                 xor d
D145  AA                 xor d
D146  AA                 xor d
D147  AA                 xor d
D148  AA                 xor d
D149  AA                 xor d
D14A  AA                 xor d
D14B  AA                 xor d
D14C  AA                 xor d
D14D  AA                 xor d
D14E  AA                 xor d
D14F  AA                 xor d
D150  AA                 xor d
D151  AA                 xor d
D152  AA                 xor d
D153  AA                 xor d
D154  AA                 xor d
D155  AA                 xor d
D156  AA                 xor d
D157  AA                 xor d
D158  AA                 xor d
D159  AA                 xor d
D15A  AA                 xor d
D15B  AA                 xor d
D15C  AA                 xor d
D15D  AA                 xor d
D15E  AA                 xor d
D15F  AA                 xor d
D160  AA                 xor d
D161  AA                 xor d
D162  AA                 xor d
D163  AA                 xor d
D164  AA                 xor d
D165  AA                 xor d
D166  AA                 xor d
D167  AA                 xor d
D168  AA                 xor d
D169  AA                 xor d
D16A  AA                 xor d
D16B  AA                 xor d
D16C  AA                 xor d
D16D  AA                 xor d
D16E  AA                 xor d
D16F  AA                 xor d
D170  AA                 xor d
D171  AA                 xor d
D172  AA                 xor d
D173  AA                 xor d
D174  AA                 xor d
D175  AA                 xor d
D176  AA                 xor d
D177  AA                 xor d
D178  AA                 xor d
D179  AA                 xor d
D17A  AA                 xor d
D17B  AA                 xor d
D17C  AA                 xor d
D17D  AA                 xor d
D17E  AA                 xor d
D17F  AA                 xor d
D180  AA                 xor d
D181  AA                 xor d
D182  AA                 xor d
D183  AA                 xor d
D184  AA                 xor d
D185  AA                 xor d
D186  AA                 xor d
D187  AA                 xor d
D188  AA                 xor d
D189  AA                 xor d
D18A  AA                 xor d
D18B  AA                 xor d
D18C  AA                 xor d
D18D  AA                 xor d
D18E  AA                 xor d
D18F  AA                 xor d
D190  AA                 xor d
D191  AA                 xor d
D192  AA                 xor d
D193  AA                 xor d
D194  AA                 xor d
D195  AA                 xor d
D196  AA                 xor d
D197  AA                 xor d
D198  AA                 xor d
D199  AA                 xor d
D19A  AA                 xor d
D19B  AA                 xor d
D19C  AA                 xor d
D19D  AA                 xor d
D19E  AA                 xor d
D19F  AA                 xor d
D1A0  AA                 xor d
D1A1  AA                 xor d
D1A2  AA                 xor d
D1A3  AA                 xor d
D1A4  AA                 xor d
D1A5  AA                 xor d
D1A6  AA                 xor d
D1A7  AA                 xor d
D1A8  AA                 xor d
D1A9  AA                 xor d
D1AA  AA                 xor d
D1AB  AA                 xor d
D1AC  AA                 xor d
D1AD  AA                 xor d
D1AE  AA                 xor d
D1AF  AA                 xor d
D1B0  AA                 xor d
D1B1  AA                 xor d
D1B2  AA                 xor d
D1B3  AA                 xor d
D1B4  AA                 xor d
D1B5  AA                 xor d
D1B6  AA                 xor d
D1B7  AA                 xor d
D1B8  AA                 xor d
D1B9  AA                 xor d
D1BA  AA                 xor d
D1BB  AA                 xor d
D1BC  AA                 xor d
D1BD  AA                 xor d
D1BE  AA                 xor d
D1BF  AA                 xor d
D1C0  AA                 xor d
D1C1  AA                 xor d
D1C2  AA                 xor d
D1C3  AA                 xor d
D1C4  AA                 xor d
D1C5  AA                 xor d
D1C6  AA                 xor d
D1C7  AA                 xor d
D1C8  AA                 xor d
D1C9  AA                 xor d
D1CA  AA                 xor d
D1CB  AA                 xor d
D1CC  AA                 xor d
D1CD  AA                 xor d
D1CE  AA                 xor d
D1CF  AA                 xor d
D1D0  AA                 xor d
D1D1  AA                 xor d
D1D2  AA                 xor d
D1D3  AA                 xor d
D1D4  AA                 xor d
D1D5  AA                 xor d
D1D6  AA                 xor d
D1D7  AA                 xor d
D1D8  AA                 xor d
D1D9  AA                 xor d
D1DA  AA                 xor d
D1DB  AA                 xor d
D1DC  AA                 xor d
D1DD  AA                 xor d
D1DE  AA                 xor d
D1DF  AA                 xor d
D1E0  AA                 xor d
D1E1  AA                 xor d
D1E2  AA                 xor d
D1E3  AA                 xor d
D1E4  AA                 xor d
D1E5  AA                 xor d
D1E6  AA                 xor d
D1E7  AA                 xor d
D1E8  AA                 xor d
D1E9  AA                 xor d
D1EA  AA                 xor d
D1EB  AA                 xor d
D1EC  AA                 xor d
D1ED  AA                 xor d
D1EE  AA                 xor d
D1EF  AA                 xor d
D1F0  AA                 xor d
D1F1  AA                 xor d
D1F2  AA                 xor d
D1F3  AA                 xor d
D1F4  AA                 xor d
D1F5  AA                 xor d
D1F6  AA                 xor d
D1F7  AA                 xor d
D1F8  AA                 xor d
D1F9  AA                 xor d
D1FA  AA                 xor d
D1FB  AA                 xor d
D1FC  AA                 xor d
D1FD  AA                 xor d
D1FE  AA                 xor d
D1FF  AA                 xor d
D200  AA                 xor d
D201  AA                 xor d
D202  AA                 xor d
D203  AA                 xor d
D204  AA                 xor d
D205  AA                 xor d
D206  AA                 xor d
D207  AA                 xor d
D208  AA                 xor d
D209  AA                 xor d
D20A  AA                 xor d
D20B  AA                 xor d
D20C  AA                 xor d
D20D  AA                 xor d
D20E  AA                 xor d
D20F  AA                 xor d
D210  FF                 rst 38h
D211  FF                 rst 38h
D212  FF                 rst 38h
D213  FF                 rst 38h
D214  FF                 rst 38h
D215  FF                 rst 38h
D216  FF                 rst 38h
D217  FF                 rst 38h
D218  FF                 rst 38h
D219  FF                 rst 38h
D21A  FF                 rst 38h
D21B  FF                 rst 38h
D21C  FF                 rst 38h
D21D  FF                 rst 38h
D21E  FF                 rst 38h
D21F  FF                 rst 38h
D220  FF                 rst 38h
D221  FF                 rst 38h
D222  FF                 rst 38h
D223  FF                 rst 38h
D224  FF                 rst 38h
D225  FF                 rst 38h
D226  FF                 rst 38h
D227  FF                 rst 38h
D228  FF                 rst 38h
D229  FF                 rst 38h
D22A  FF                 rst 38h
D22B  FF                 rst 38h
D22C  FF                 rst 38h
D22D  FF                 rst 38h
D22E  FF                 rst 38h
D22F  FF                 rst 38h
D230  FF                 rst 38h
D231  FF                 rst 38h
D232  FF                 rst 38h
D233  FF                 rst 38h
D234  FF                 rst 38h
D235  FF                 rst 38h
D236  FF                 rst 38h
D237  FF                 rst 38h
D238  FF                 rst 38h
D239  FF                 rst 38h
D23A  FF                 rst 38h
D23B  FF                 rst 38h
D23C  FF                 rst 38h
D23D  FF                 rst 38h
D23E  FF                 rst 38h
D23F  FF                 rst 38h
D240  FF                 rst 38h
D241  FF                 rst 38h
D242  FF                 rst 38h
D243  FF                 rst 38h
D244  FF                 rst 38h
D245  FF                 rst 38h
D246  FF                 rst 38h
D247  FF                 rst 38h
D248  FF                 rst 38h
D249  FF                 rst 38h
D24A  FF                 rst 38h
D24B  FF                 rst 38h
D24C  FF                 rst 38h
D24D  FF                 rst 38h
D24E  FF                 rst 38h
D24F  FF                 rst 38h
D250  FF                 rst 38h
D251  FF                 rst 38h
D252  FF                 rst 38h
D253  FF                 rst 38h
D254  FF                 rst 38h
D255  FF                 rst 38h
D256  FF                 rst 38h
D257  FF                 rst 38h
D258  FF                 rst 38h
D259  FF                 rst 38h
D25A  FF                 rst 38h
D25B  FF                 rst 38h
D25C  FF                 rst 38h
D25D  FF                 rst 38h
D25E  FF                 rst 38h
D25F  FF                 rst 38h
D260  FF                 rst 38h
D261  FF                 rst 38h
D262  FF                 rst 38h
D263  FF                 rst 38h
D264  FF                 rst 38h
D265  FF                 rst 38h
D266  FF                 rst 38h
D267  FF                 rst 38h
D268  FF                 rst 38h
D269  FF                 rst 38h
D26A  FF                 rst 38h
D26B  FF                 rst 38h
D26C  FF                 rst 38h
D26D  FF                 rst 38h
D26E  FF                 rst 38h
D26F  FF                 rst 38h
D270  FF                 rst 38h
D271  FF                 rst 38h
D272  FF                 rst 38h
D273  FF                 rst 38h
D274  FF                 rst 38h
D275  FF                 rst 38h
D276  FF                 rst 38h
D277  FF                 rst 38h
D278  FF                 rst 38h
D279  FF                 rst 38h
D27A  FF                 rst 38h
D27B  FF                 rst 38h
D27C  FF                 rst 38h
D27D  FF                 rst 38h
D27E  FF                 rst 38h
D27F  FF                 rst 38h
D280  FF                 rst 38h
D281  FF                 rst 38h
D282  FF                 rst 38h
D283  FF                 rst 38h
D284  FF                 rst 38h
D285  FF                 rst 38h
D286  FF                 rst 38h
D287  FF                 rst 38h
D288  FF                 rst 38h
D289  FF                 rst 38h
D28A  FF                 rst 38h
D28B  FF                 rst 38h
D28C  FF                 rst 38h
D28D  FF                 rst 38h
D28E  FF                 rst 38h
D28F  FF                 rst 38h
D290  FF                 rst 38h
D291  FF                 rst 38h
D292  FF                 rst 38h
D293  FF                 rst 38h
D294  FF                 rst 38h
D295  FF                 rst 38h
D296  FF                 rst 38h
D297  FF                 rst 38h
D298  FF                 rst 38h
D299  FF                 rst 38h
D29A  FF                 rst 38h
D29B  FF                 rst 38h
D29C  FF                 rst 38h
D29D  FF                 rst 38h
D29E  FF                 rst 38h
D29F  FF                 rst 38h
D2A0  FF                 rst 38h
D2A1  FF                 rst 38h
D2A2  FF                 rst 38h
D2A3  FF                 rst 38h
D2A4  FF                 rst 38h
D2A5  FF                 rst 38h
D2A6  FF                 rst 38h
D2A7  FF                 rst 38h
D2A8  FF                 rst 38h
D2A9  FF                 rst 38h
D2AA  FF                 rst 38h
D2AB  FF                 rst 38h
D2AC  FF                 rst 38h
D2AD  FF                 rst 38h
D2AE  FF                 rst 38h
D2AF  FF                 rst 38h
D2B0  FF                 rst 38h
D2B1  FF                 rst 38h
D2B2  FF                 rst 38h
D2B3  FF                 rst 38h
D2B4  FF                 rst 38h
D2B5  FF                 rst 38h
D2B6  FF                 rst 38h
D2B7  FF                 rst 38h
D2B8  FF                 rst 38h
D2B9  FF                 rst 38h
D2BA  FF                 rst 38h
D2BB  FF                 rst 38h
D2BC  FF                 rst 38h
D2BD  FF                 rst 38h
D2BE  FF                 rst 38h
D2BF  FF                 rst 38h
D2C0  FF                 rst 38h
D2C1  FF                 rst 38h
D2C2  FF                 rst 38h
D2C3  FF                 rst 38h
D2C4  FF                 rst 38h
D2C5  FF                 rst 38h
D2C6  FF                 rst 38h
D2C7  FF                 rst 38h
D2C8  FF                 rst 38h
D2C9  FF                 rst 38h
D2CA  FF                 rst 38h
D2CB  FF                 rst 38h
D2CC  FF                 rst 38h
D2CD  FF                 rst 38h
D2CE  FF                 rst 38h
D2CF  FF                 rst 38h
D2D0  FF                 rst 38h
D2D1  FF                 rst 38h
D2D2  FF                 rst 38h
D2D3  FF                 rst 38h
D2D4  FF                 rst 38h
D2D5  FF                 rst 38h
D2D6  FF                 rst 38h
D2D7  FF                 rst 38h
D2D8  FF                 rst 38h
D2D9  FF                 rst 38h
D2DA  FF                 rst 38h
D2DB  FF                 rst 38h
D2DC  FF                 rst 38h
D2DD  FF                 rst 38h
D2DE  FF                 rst 38h
D2DF  FF                 rst 38h
D2E0  FF                 rst 38h
D2E1  FF                 rst 38h
D2E2  00                 nop
D2E3  00                 nop
D2E4  00                 nop
D2E5  00                 nop
D2E6  00                 nop
D2E7  00                 nop
D2E8  00                 nop
D2E9  00                 nop
D2EA  00                 nop
D2EB  00                 nop
D2EC  00                 nop
D2ED  00                 nop
D2EE  00                 nop
D2EF  00                 nop
D2F0  00                 nop
D2F1  00                 nop
D2F2  00                 nop
D2F3  00                 nop
D2F4  00                 nop
D2F5  00                 nop
D2F6  00                 nop
D2F7  00                 nop
D2F8  00                 nop
D2F9  00                 nop
D2FA  00                 nop
D2FB  00                 nop
D2FC  00                 nop
D2FD  00                 nop
D2FE  00                 nop
D2FF  00                 nop














E700  03                 inc bc
E701  43                 ld b,e
E702  4F                 ld c,a
E703  4F                 ld c,a
E704  4F                 ld c,a
E705  4F                 ld c,a
E706  4F                 ld c,a
E707  4F                 ld c,a
E708  4A                 ld c,d
E709  4D                 ld c,l
E70A  6E                 ld l,(hl)
E70B  56                 ld d,(hl)
E70C  4E                 ld c,(hl)
E70D  56                 ld d,(hl)
E70E  5E                 ld e,(hl)
E70F  56                 ld d,(hl)
E710  4E                 ld c,(hl)
E711  56                 ld d,(hl)
E712  5E                 ld e,(hl)
E713  56                 ld d,(hl)
E714  5E                 ld e,(hl)
E715  6E                 ld l,(hl)
E716  4D                 ld c,l
E717  4A                 ld c,d
E718  4F                 ld c,a
E719  4F                 ld c,a
E71A  4F                 ld c,a
E71B  4F                 ld c,a
E71C  4F                 ld c,a
E71D  4F                 ld c,a
E71E  03                 inc bc
E71F  43                 ld b,e
E720  4C                 ld c,h
E721  4C                 ld c,h
E722  4E                 ld c,(hl)
E723  4E                 ld c,(hl)
E724  4E                 ld c,(hl)
E725  4E                 ld c,(hl)
E726  4E                 ld c,(hl)
E727  4E                 ld c,(hl)
E728  4A                 ld c,d
E729  47                 ld b,a
E72A  05                 dec b
E72B  07                 rlca
E72C  6E                 ld l,(hl)
E72D  47                 ld b,a
E72E  07                 rlca
E72F  6E                 ld l,(hl)
E730  6E                 ld l,(hl)
E731  47                 ld b,a
E732  07                 rlca
E733  6E                 ld l,(hl)
E734  05                 dec b
E735  07                 rlca
E736  47                 ld b,a
E737  4A                 ld c,d
E738  4E                 ld c,(hl)
E739  4E                 ld c,(hl)
E73A  4E                 ld c,(hl)
E73B  4E                 ld c,(hl)
E73C  4E                 ld c,(hl)
E73D  4E                 ld c,(hl)
E73E  4C                 ld c,h
E73F  4C                 ld c,h
E740  0C                 inc c
E741  0C                 inc c
E742  0C                 inc c
E743  0E 0E              ld c,14
E745  0F                 rrca
E746  0D                 dec c
E747  0B                 dec bc
E748  0B                 dec bc
E749  46                 ld b,(hl)
E74A  05                 dec b
E74B  45                 ld b,l
E74C  69                 ld l,c
E74D  47                 ld b,a
E74E  07                 rlca
E74F  69                 ld l,c
E750  69                 ld l,c
E751  47                 ld b,a
E752  07                 rlca
E753  69                 ld l,c
E754  05                 dec b
E755  45                 ld b,l
E756  46                 ld b,(hl)
E757  0B                 dec bc
E758  0B                 dec bc
E759  0D                 dec c
E75A  0F                 rrca
E75B  0E 0E              ld c,14
E75D  0C                 inc c
E75E  0C                 inc c
E75F  0C                 inc c
E760  44                 ld b,h
E761  44                 ld b,h
E762  46                 ld b,(hl)
E763  46                 ld b,(hl)
E764  45                 ld b,l
E765  07                 rlca
E766  45                 ld b,l
E767  03                 inc bc
E768  42                 ld b,d
E769  56                 ld d,(hl)
E76A  56                 ld d,(hl)
E76B  56                 ld d,(hl)
E76C  41                 ld b,c
E76D  56                 ld d,(hl)
E76E  56                 ld d,(hl)
E76F  41                 ld b,c
E770  41                 ld b,c
E771  56                 ld d,(hl)
E772  56                 ld d,(hl)
E773  41                 ld b,c
E774  56                 ld d,(hl)
E775  56                 ld d,(hl)
E776  56                 ld d,(hl)
E777  02                 ld (bc),a
E778  03                 inc bc
E779  05                 dec b
E77A  07                 rlca
E77B  05                 dec b
E77C  46                 ld b,(hl)
E77D  46                 ld b,(hl)
E77E  44                 ld b,h
E77F  44                 ld b,h
E780  00                 nop
E781  00                 nop
E782  00                 nop
E783  00                 nop
E784  00                 nop
E785  00                 nop
E786  00                 nop
E787  00                 nop
E788  00                 nop
E789  00                 nop
E78A  00                 nop
E78B  00                 nop
E78C  00                 nop
E78D  00                 nop
E78E  00                 nop
E78F  00                 nop
E790  00                 nop
E791  00                 nop
E792  00                 nop
E793  00                 nop
E794  00                 nop
E795  00                 nop
E796  00                 nop
E797  00                 nop
E798  00                 nop
E799  00                 nop
E79A  00                 nop
E79B  00                 nop
E79C  00                 nop
E79D  00                 nop
E79E  00                 nop
E79F  00                 nop
E7A0  00                 nop
E7A1  00                 nop
E7A2  00                 nop
E7A3  00                 nop
E7A4  00                 nop
E7A5  00                 nop
E7A6  00                 nop
E7A7  00                 nop
E7A8  00                 nop
E7A9  00                 nop
E7AA  00                 nop
E7AB  00                 nop
E7AC  00                 nop
E7AD  00                 nop
E7AE  00                 nop
E7AF  00                 nop
E7B0  00                 nop
E7B1  00                 nop
E7B2  00                 nop
E7B3  00                 nop
E7B4  00                 nop
E7B5  00                 nop
E7B6  00                 nop
E7B7  00                 nop
E7B8  00                 nop
E7B9  00                 nop
E7BA  00                 nop
E7BB  00                 nop
E7BC  00                 nop
E7BD  00                 nop
E7BE  00                 nop
E7BF  00                 nop
E7C0  00                 nop
E7C1  00                 nop
E7C2  00                 nop
E7C3  00                 nop
E7C4  00                 nop
E7C5  00                 nop
E7C6  00                 nop
E7C7  00                 nop
E7C8  00                 nop
E7C9  00                 nop
E7CA  00                 nop
E7CB  00                 nop
E7CC  00                 nop
E7CD  00                 nop
E7CE  00                 nop
E7CF  00                 nop
E7D0  00                 nop
E7D1  00                 nop
E7D2  00                 nop
E7D3  00                 nop
E7D4  00                 nop
E7D5  00                 nop
E7D6  00                 nop
E7D7  00                 nop
E7D8  00                 nop
E7D9  00                 nop
E7DA  00                 nop
E7DB  00                 nop
E7DC  00                 nop
E7DD  00                 nop
E7DE  00                 nop
E7DF  00                 nop
E7E0  00                 nop
E7E1  00                 nop
E7E2  00                 nop
E7E3  00                 nop
E7E4  00                 nop
E7E5  00                 nop
E7E6  00                 nop
E7E7  00                 nop
E7E8  00                 nop
E7E9  00                 nop
E7EA  00                 nop
E7EB  00                 nop
E7EC  00                 nop
E7ED  00                 nop
E7EE  00                 nop
E7EF  00                 nop
E7F0  00                 nop
E7F1  00                 nop
E7F2  00                 nop
E7F3  00                 nop
E7F4  00                 nop
E7F5  00                 nop
E7F6  00                 nop
E7F7  00                 nop
E7F8  00                 nop
E7F9  00                 nop
E7FA  00                 nop
E7FB  00                 nop
E7FC  00                 nop
E7FD  00                 nop
E7FE  00                 nop
E7FF  00                 nop
E800  00                 nop
E801  00                 nop
E802  00                 nop
E803  00                 nop
E804  00                 nop
E805  00                 nop
E806  00                 nop
E807  00                 nop
E808  00                 nop
E809  00                 nop
E80A  00                 nop
E80B  00                 nop
E80C  00                 nop
E80D  00                 nop
E80E  00                 nop
E80F  00                 nop
E810  00                 nop
E811  00                 nop
E812  00                 nop
E813  00                 nop
E814  00                 nop
E815  00                 nop
E816  00                 nop
E817  00                 nop
E818  00                 nop
E819  00                 nop
E81A  00                 nop
E81B  00                 nop
E81C  00                 nop
E81D  00                 nop
E81E  00                 nop
E81F  00                 nop
E820  00                 nop
E821  00                 nop
E822  00                 nop
E823  00                 nop
E824  00                 nop
E825  00                 nop
E826  00                 nop
E827  00                 nop
E828  00                 nop
E829  00                 nop
E82A  00                 nop
E82B  00                 nop
E82C  00                 nop
E82D  00                 nop
E82E  00                 nop
E82F  00                 nop
E830  00                 nop
E831  00                 nop
E832  00                 nop
E833  00                 nop
E834  00                 nop
E835  00                 nop
E836  00                 nop
E837  00                 nop
E838  00                 nop
E839  00                 nop
E83A  00                 nop
E83B  00                 nop
E83C  00                 nop
E83D  00                 nop
E83E  00                 nop
E83F  00                 nop
E840  00                 nop
E841  00                 nop
E842  00                 nop
E843  00                 nop
E844  00                 nop
E845  00                 nop
E846  00                 nop
E847  00                 nop
E848  00                 nop
E849  00                 nop
E84A  00                 nop
E84B  00                 nop
E84C  00                 nop
E84D  00                 nop
E84E  00                 nop
E84F  00                 nop
E850  00                 nop
E851  00                 nop
E852  00                 nop
E853  00                 nop
E854  00                 nop
E855  00                 nop
E856  00                 nop
E857  00                 nop
E858  00                 nop
E859  00                 nop
E85A  00                 nop
E85B  00                 nop
E85C  00                 nop
E85D  00                 nop
E85E  00                 nop
E85F  00                 nop
E860  00                 nop
E861  00                 nop
E862  00                 nop
E863  00                 nop
E864  00                 nop
E865  00                 nop
E866  00                 nop
E867  00                 nop
E868  00                 nop
E869  00                 nop
E86A  00                 nop
E86B  00                 nop
E86C  00                 nop
E86D  00                 nop
E86E  00                 nop
E86F  00                 nop
E870  00                 nop
E871  00                 nop
E872  00                 nop
E873  00                 nop
E874  00                 nop
E875  00                 nop
E876  00                 nop
E877  00                 nop
E878  00                 nop
E879  00                 nop
E87A  00                 nop
E87B  00                 nop
E87C  00                 nop
E87D  00                 nop
E87E  00                 nop
E87F  00                 nop
E880  00                 nop
E881  00                 nop
E882  00                 nop
E883  00                 nop
E884  00                 nop
E885  00                 nop
E886  00                 nop
E887  00                 nop
E888  00                 nop
E889  00                 nop
E88A  00                 nop
E88B  00                 nop
E88C  00                 nop
E88D  00                 nop
E88E  00                 nop
E88F  00                 nop
E890  00                 nop
E891  00                 nop
E892  00                 nop
E893  00                 nop
E894  00                 nop
E895  00                 nop
E896  00                 nop
E897  00                 nop
E898  00                 nop
E899  00                 nop
E89A  00                 nop
E89B  00                 nop
E89C  00                 nop
E89D  00                 nop
E89E  00                 nop
E89F  00                 nop
E8A0  00                 nop
E8A1  00                 nop
E8A2  00                 nop
E8A3  00                 nop
E8A4  00                 nop
E8A5  00                 nop
E8A6  00                 nop
E8A7  00                 nop
E8A8  00                 nop
E8A9  00                 nop
E8AA  00                 nop
E8AB  00                 nop
E8AC  00                 nop
E8AD  00                 nop
E8AE  00                 nop
E8AF  00                 nop
E8B0  00                 nop
E8B1  00                 nop
E8B2  00                 nop
E8B3  00                 nop
E8B4  00                 nop
E8B5  00                 nop
E8B6  00                 nop
E8B7  00                 nop
E8B8  00                 nop
E8B9  00                 nop
E8BA  00                 nop
E8BB  00                 nop
E8BC  00                 nop
E8BD  00                 nop
E8BE  00                 nop
E8BF  00                 nop
E8C0  00                 nop
E8C1  00                 nop
E8C2  00                 nop
E8C3  00                 nop
E8C4  00                 nop
E8C5  00                 nop
E8C6  00                 nop
E8C7  00                 nop
E8C8  00                 nop
E8C9  00                 nop
E8CA  00                 nop
E8CB  00                 nop
E8CC  00                 nop
E8CD  00                 nop
E8CE  00                 nop
E8CF  00                 nop
E8D0  00                 nop
E8D1  00                 nop
E8D2  00                 nop
E8D3  00                 nop
E8D4  00                 nop
E8D5  00                 nop
E8D6  00                 nop
E8D7  00                 nop
E8D8  00                 nop
E8D9  00                 nop
E8DA  00                 nop
E8DB  00                 nop
E8DC  00                 nop
E8DD  00                 nop
E8DE  00                 nop
E8DF  00                 nop
E8E0  00                 nop
E8E1  00                 nop
E8E2  00                 nop
E8E3  00                 nop
E8E4  00                 nop
E8E5  00                 nop
E8E6  00                 nop
E8E7  00                 nop
E8E8  00                 nop
E8E9  00                 nop
E8EA  00                 nop
E8EB  00                 nop
E8EC  00                 nop
E8ED  00                 nop
E8EE  00                 nop
E8EF  00                 nop
E8F0  00                 nop
E8F1  00                 nop
E8F2  00                 nop
E8F3  00                 nop
E8F4  00                 nop
E8F5  00                 nop
E8F6  00                 nop
E8F7  00                 nop
E8F8  00                 nop
E8F9  00                 nop
E8FA  00                 nop
E8FB  00                 nop
E8FC  00                 nop
E8FD  00                 nop
E8FE  00                 nop
E8FF  00                 nop
E900  17                 rla
E901  17                 rla
E902  4F                 ld c,a
E903  97                 sub a
E904  97                 sub a
E905  97                 sub a
E906  47                 ld b,a
E907  07                 rlca
E908  57                 ld d,a
E909  D7                 rst 10h
E90A  D7                 rst 10h
E90B  CF                 rst 08h
E90C  67                 ld h,a
E90D  67                 ld h,a
E90E  67                 ld h,a
E90F  67                 ld h,a
E910  47                 ld b,a
E911  07                 rlca
E912  47                 ld b,a
E913  06 07              ld b,7
E915  07                 rlca
E916  4F                 ld c,a
E917  47                 ld b,a
E918  07                 rlca
E919  07                 rlca
E91A  47                 ld b,a
E91B  07                 rlca
E91C  47                 ld b,a
E91D  4F                 ld c,a
E91E  4F                 ld c,a
E91F  07                 rlca
E920  07                 rlca
E921  47                 ld b,a
E922  06 07              ld b,7
E924  07                 rlca
E925  07                 rlca
E926  07                 rlca
E927  07                 rlca
E928  07                 rlca
E929  07                 rlca
E92A  07                 rlca
E92B  07                 rlca
E92C  07                 rlca
E92D  07                 rlca
E92E  07                 rlca
E92F  07                 rlca
E930  07                 rlca
E931  07                 rlca
E932  07                 rlca
E933  07                 rlca
E934  07                 rlca
E935  07                 rlca
E936  07                 rlca
E937  07                 rlca
E938  07                 rlca
E939  17                 rla
E93A  17                 rla
E93B  17                 rla
E93C  07                 rlca
E93D  07                 rlca
E93E  07                 rlca
E93F  07                 rlca
E940  07                 rlca
E941  07                 rlca
E942  07                 rlca
E943  07                 rlca
E944  07                 rlca
E945  07                 rlca
E946  07                 rlca
E947  07                 rlca
E948  07                 rlca
E949  07                 rlca
E94A  07                 rlca
E94B  07                 rlca
E94C  07                 rlca
E94D  07                 rlca
E94E  07                 rlca
E94F  07                 rlca
E950  07                 rlca
E951  07                 rlca
E952  07                 rlca
E953  57                 ld d,a
E954  57                 ld d,a
E955  57                 ld d,a
E956  57                 ld d,a
E957  D7                 rst 10h
E958  07                 rlca
E959  47                 ld b,a
E95A  07                 rlca
E95B  47                 ld b,a
E95C  47                 ld b,a
E95D  07                 rlca
E95E  47                 ld b,a
E95F  47                 ld b,a
E960  07                 rlca
E961  07                 rlca
E962  47                 ld b,a
E963  07                 rlca
E964  07                 rlca
E965  07                 rlca
E966  07                 rlca
E967  07                 rlca
E968  47                 ld b,a
E969  07                 rlca
E96A  17                 rla
E96B  17                 rla
E96C  07                 rlca
E96D  57                 ld d,a
E96E  07                 rlca
E96F  07                 rlca
E970  97                 sub a
E971  47                 ld b,a
E972  47                 ld b,a
E973  07                 rlca
E974  07                 rlca
E975  07                 rlca
E976  07                 rlca
E977  07                 rlca
E978  07                 rlca
E979  47                 ld b,a
E97A  47                 ld b,a
E97B  07                 rlca
E97C  07                 rlca
E97D  07                 rlca
E97E  07                 rlca
E97F  07                 rlca
E980  02                 ld (bc),a
E981  57                 ld d,a
E982  57                 ld d,a
E983  57                 ld d,a
E984  47                 ld b,a
E985  57                 ld d,a
E986  57                 ld d,a
E987  57                 ld d,a
E988  47                 ld b,a
E989  07                 rlca
E98A  07                 rlca
E98B  07                 rlca
E98C  07                 rlca
E98D  07                 rlca
E98E  07                 rlca
E98F  07                 rlca
E990  07                 rlca
E991  07                 rlca
E992  07                 rlca
E993  0A                 ld a,(bc)
E994  1B                 dec de
E995  20 2D              jr nz,$E9C4
E997  2E 2E              ld l,46
E999  2E 17              ld l,23
E99B  2C                 inc l
E99C  1A                 ld a,(de)
E99D  18 2D              jr $E9CC
E99F  1A                 ld a,(de)
E9A0  0B                 dec bc
E9A1  0D                 dec c
E9A2  0E 11              ld c,17
E9A4  23                 inc hl
E9A5  14                 inc d
E9A6  21 2F 29           ld hl,$292F
E9A9  26 32              ld h,50
E9AB  25                 dec h
E9AC  13                 inc de
E9AD  19                 add hl,de
E9AE  07                 rlca
E9AF  07                 rlca
E9B0  1B                 dec de
E9B1  1D                 dec e
E9B2  1C                 inc e
E9B3  7F                 ld a,a
E9B4  7F                 ld a,a
E9B5  00                 nop
E9B6  04                 inc b
E9B7  08                 ex af,af'
E9B8  07                 rlca
E9B9  00                 nop
E9BA  00                 nop
E9BB  00                 nop
E9BC  00                 nop
E9BD  00                 nop
E9BE  00                 nop
E9BF  00                 nop
E9C0  00                 nop
E9C1  00                 nop
E9C2  00                 nop
E9C3  00                 nop
E9C4  00                 nop
E9C5  00                 nop
E9C6  02                 ld (bc),a
E9C7  00                 nop
E9C8  04                 inc b
E9C9  02                 ld (bc),a
E9CA  00                 nop
E9CB  04                 inc b
E9CC  02                 ld (bc),a
E9CD  02                 ld (bc),a
E9CE  04                 inc b
E9CF  04                 inc b
E9D0  02                 ld (bc),a
E9D1  00                 nop
E9D2  00                 nop
E9D3  00                 nop
E9D4  07                 rlca
E9D5  07                 rlca
E9D6  00                 nop
E9D7  00                 nop
E9D8  00                 nop
E9D9  00                 nop
E9DA  00                 nop
E9DB  04                 inc b
E9DC  00                 nop
E9DD  04                 inc b
E9DE  3A FC FF           ld a,($FFFC)
E9E1  DD 21 FC FF        ld ix,$FFFC
E9E5  3D                 dec a
E9E6  DD AE 01           xor (ix+1)
E9E9  17                 rla
E9EA  DD CB 02 1E        rr (ix+2)
E9EE  DD 77 00           ld (ix+0),a
E9F1  2F                 cpl
E9F2  17                 rla
E9F3  EE 30              xor $30
E9F5  DD AE 01           xor (ix+1)
E9F8  DD 77 01           ld (ix+1),a
E9FB  DD AE 02           xor (ix+2)
E9FE  C9                 ret
E9FF  FF                 rst 38h
EA00  02                 ld (bc),a
EA01  06 42              ld b,66
EA03  03                 inc bc
EA04  0D                 dec c
EA05  05                 dec b
EA06  44                 ld b,h
EA07  04                 inc b
EA08  47                 ld b,a
EA09  4F                 ld c,a
EA0A  4D                 ld c,l
EA0B  4D                 ld c,l
EA0C  4F                 ld c,a
EA0D  4F                 ld c,a
EA0E  47                 ld b,a
EA0F  47                 ld b,a
EA10  0C                 inc c
EA11  0C                 inc c
EA12  4C                 ld c,h
EA13  0E 0D              ld c,13
EA15  0D                 dec c
EA16  47                 ld b,a
EA17  06 0C              ld b,12
EA19  04                 inc b
EA1A  4C                 ld c,h
EA1B  0E 45              ld c,69
EA1D  47                 ld b,a
EA1E  47                 ld b,a
EA1F  0C                 inc c
EA20  0A                 ld a,(bc)
EA21  4C                 ld c,h
EA22  0E 0D              ld c,13
EA24  02                 ld (bc),a
EA25  02                 ld (bc),a
EA26  02                 ld (bc),a
EA27  06 06              ld b,6
EA29  06 06              ld b,6
EA2B  06 06              ld b,6
EA2D  06 04              ld b,4
EA2F  04                 inc b
EA30  04                 inc b
EA31  04                 inc b
EA32  04                 inc b
EA33  04                 inc b
EA34  0F                 rrca
EA35  0F                 rrca
EA36  0F                 rrca
EA37  0F                 rrca
EA38  0F                 rrca
EA39  0A                 ld a,(bc)
EA3A  0A                 ld a,(bc)
EA3B  0A                 ld a,(bc)
EA3C  0E 0E              ld c,14
EA3E  0E 0E              ld c,14
EA40  0E 0E              ld c,14
EA42  0E 0E              ld c,14
EA44  0E 0E              ld c,14
EA46  0E 0E              ld c,14
EA48  0E 0E              ld c,14
EA4A  0E 0E              ld c,14
EA4C  0E 0E              ld c,14
EA4E  0E 0E              ld c,14
EA50  0E 0E              ld c,14
EA52  0E 0A              ld c,10
EA54  42                 ld b,d
EA55  46                 ld b,(hl)
EA56  46                 ld b,(hl)
EA57  42                 ld b,d
EA58  0A                 ld a,(bc)
EA59  4A                 ld c,d
EA5A  02                 ld (bc),a
EA5B  42                 ld b,d
EA5C  46                 ld b,(hl)
EA5D  0C                 inc c
EA5E  4E                 ld c,(hl)
EA5F  4C                 ld c,h
EA60  0C                 inc c
EA61  0C                 inc c
EA62  4C                 ld c,h
EA63  02                 ld (bc),a
EA64  02                 ld (bc),a
EA65  02                 ld (bc),a
EA66  02                 ld (bc),a
EA67  02                 ld (bc),a
EA68  42                 ld b,d
EA69  02                 ld (bc),a
EA6A  02                 ld (bc),a
EA6B  02                 ld (bc),a
EA6C  07                 rlca
EA6D  44                 ld b,h
EA6E  05                 dec b
EA6F  0C                 inc c
EA70  05                 dec b
EA71  4E                 ld c,(hl)
EA72  4E                 ld c,(hl)
EA73  0E 02              ld c,2
EA75  02                 ld (bc),a
EA76  02                 ld (bc),a
EA77  02                 ld (bc),a
EA78  02                 ld (bc),a
EA79  0C                 inc c
EA7A  0C                 inc c
EA7B  0C                 inc c
EA7C  0C                 inc c
EA7D  0C                 inc c
EA7E  0C                 inc c
EA7F  0C                 inc c
EA80  0F                 rrca
EA81  42                 ld b,d
EA82  42                 ld b,d
EA83  42                 ld b,d
EA84  42                 ld b,d
EA85  42                 ld b,d
EA86  42                 ld b,d
EA87  42                 ld b,d
EA88  42                 ld b,d
EA89  0F                 rrca
EA8A  0F                 rrca
EA8B  0F                 rrca
EA8C  0F                 rrca
EA8D  47                 ld b,a
EA8E  0F                 rrca
EA8F  0F                 rrca
EA90  47                 ld b,a
EA91  0F                 rrca
EA92  62                 ld h,d
EA93  CB C5              set 0,l
EA95  C7                 rst 00h
EA96  9D                 sbc a,l
EA97  CC A3 C9           call z,$C9A3
EA9A  09                 add hl,bc
EA9B  CA 30 C8           jp z,$C830
EA9E  30 C8              jr nc,$EA68
EAA0  30 C8              jr nc,$EA6A
EAA2  2F                 cpl
EAA3  CA 2F CA           jp z,$CA2F
EAA6  2F                 cpl
EAA7  CA E2 C9           jp z,$C9E2
EAAA  0C                 inc c
EAAB  CA 4F C6           jp z,$C64F
EAAE  52                 ld d,d
EAAF  C7                 rst 00h
EAB0  AE                 xor (hl)
EAB1  C7                 rst 00h
EAB2  AE                 xor (hl)
EAB3  C7                 rst 00h
EAB4  AE                 xor (hl)
EAB5  C7                 rst 00h
EAB6  7C                 ld a,h
EAB7  C7                 rst 00h
EAB8  A2                 and d
EAB9  C7                 rst 00h
EABA  AE                 xor (hl)
EABB  C7                 rst 00h
EABC  A2                 and d
EABD  C7                 rst 00h
EABE  7C                 ld a,h
EABF  C7                 rst 00h
EAC0  A2                 and d
EAC1  C7                 rst 00h
EAC2  AE                 xor (hl)
EAC3  C7                 rst 00h
EAC4  B4                 or h
EAC5  C7                 rst 00h
EAC6  DC C6 90           call c,$90C6
EAC9  C6 3C              add a,60
EACB  CB AE              res 5,(hl)
EACD  C7                 rst 00h
EACE  9E                 sbc a,(hl)
EACF  CA CE CA           jp z,$CACE
EAD2  E2 C7 59           jp po,$59C7
EAD5  C4 FB C7           call nz,$C7FB
EAD8  DA C3 01           jp c,$01C3
EADB  7C                 ld a,h
EADC  01 7C FB           ld bc,-1156
EADF  FC FC FC           call m,$FCFC
EAE2  FC FC FD           call m,$FDFC
EAE5  FD FD FD FE FE     cp -2
EAEA  FE FF              cp -1
EAEC  FF                 rst 38h
EAED  00                 nop
EAEE  00                 nop
EAEF  01 01 02           ld bc,513
EAF2  02                 ld (bc),a
EAF3  02                 ld (bc),a
EAF4  03                 inc bc
EAF5  03                 inc bc
EAF6  03                 inc bc
EAF7  03                 inc bc
EAF8  04                 inc b
EAF9  04                 inc b
EAFA  04                 inc b
EAFB  04                 inc b
EAFC  04                 inc b
EAFD  05                 dec b
EAFE  00                 nop
EAFF  07                 rlca
EB00  02                 ld (bc),a
EB01  02                 ld (bc),a
EB02  02                 ld (bc),a
EB03  02                 ld (bc),a
EB04  02                 ld (bc),a
EB05  02                 ld (bc),a
EB06  02                 ld (bc),a
EB07  02                 ld (bc),a
EB08  02                 ld (bc),a
EB09  02                 ld (bc),a
EB0A  02                 ld (bc),a
EB0B  02                 ld (bc),a
EB0C  02                 ld (bc),a
EB0D  02                 ld (bc),a
EB0E  02                 ld (bc),a
EB0F  02                 ld (bc),a
EB10  02                 ld (bc),a
EB11  02                 ld (bc),a
EB12  02                 ld (bc),a
EB13  02                 ld (bc),a
EB14  02                 ld (bc),a
EB15  02                 ld (bc),a
EB16  02                 ld (bc),a
EB17  02                 ld (bc),a
EB18  02                 ld (bc),a
EB19  02                 ld (bc),a
EB1A  02                 ld (bc),a
EB1B  02                 ld (bc),a
EB1C  02                 ld (bc),a
EB1D  02                 ld (bc),a
EB1E  02                 ld (bc),a
EB1F  02                 ld (bc),a
EB20  02                 ld (bc),a
EB21  02                 ld (bc),a
EB22  02                 ld (bc),a
EB23  02                 ld (bc),a
EB24  02                 ld (bc),a
EB25  02                 ld (bc),a
EB26  02                 ld (bc),a
EB27  02                 ld (bc),a
EB28  02                 ld (bc),a
EB29  02                 ld (bc),a
EB2A  02                 ld (bc),a
EB2B  02                 ld (bc),a
EB2C  02                 ld (bc),a
EB2D  02                 ld (bc),a
EB2E  02                 ld (bc),a
EB2F  02                 ld (bc),a
EB30  02                 ld (bc),a
EB31  02                 ld (bc),a
EB32  02                 ld (bc),a
EB33  02                 ld (bc),a
EB34  02                 ld (bc),a
EB35  02                 ld (bc),a
EB36  02                 ld (bc),a
EB37  02                 ld (bc),a
EB38  02                 ld (bc),a
EB39  02                 ld (bc),a
EB3A  02                 ld (bc),a
EB3B  02                 ld (bc),a
EB3C  02                 ld (bc),a
EB3D  80                 add a,b
EB3E  80                 add a,b
EB3F  80                 add a,b
EB40  80                 add a,b
EB41  80                 add a,b
EB42  80                 add a,b
EB43  80                 add a,b
EB44  80                 add a,b
EB45  80                 add a,b
EB46  80                 add a,b
EB47  80                 add a,b
EB48  80                 add a,b
EB49  80                 add a,b
EB4A  80                 add a,b
EB4B  80                 add a,b
EB4C  80                 add a,b
EB4D  80                 add a,b
EB4E  80                 add a,b
EB4F  80                 add a,b
EB50  80                 add a,b
EB51  80                 add a,b
EB52  80                 add a,b
EB53  80                 add a,b
EB54  02                 ld (bc),a
EB55  06 06              ld b,6
EB57  06 06              ld b,6
EB59  06 06              ld b,6
EB5B  06 06              ld b,6
EB5D  06 06              ld b,6
EB5F  06 06              ld b,6
EB61  06 06              ld b,6
EB63  06 06              ld b,6
EB65  02                 ld (bc),a
EB66  07                 rlca
EB67  06 06              ld b,6
EB69  02                 ld (bc),a
EB6A  80                 add a,b
EB6B  80                 add a,b
EB6C  80                 add a,b
EB6D  80                 add a,b
EB6E  80                 add a,b
EB6F  80                 add a,b
EB70  02                 ld (bc),a
EB71  07                 rlca
EB72  07                 rlca
EB73  07                 rlca
EB74  07                 rlca
EB75  07                 rlca
EB76  07                 rlca
EB77  07                 rlca
EB78  07                 rlca
EB79  07                 rlca
EB7A  07                 rlca
EB7B  07                 rlca
EB7C  07                 rlca
EB7D  02                 ld (bc),a
EB7E  02                 ld (bc),a
EB7F  02                 ld (bc),a
EB80  02                 ld (bc),a
EB81  80                 add a,b
EB82  80                 add a,b
EB83  80                 add a,b
EB84  80                 add a,b
EB85  80                 add a,b
EB86  80                 add a,b
EB87  80                 add a,b
EB88  80                 add a,b
EB89  80                 add a,b
EB8A  80                 add a,b
EB8B  80                 add a,b
EB8C  80                 add a,b
EB8D  80                 add a,b
EB8E  80                 add a,b
EB8F  80                 add a,b
EB90  80                 add a,b
EB91  80                 add a,b
EB92  80                 add a,b
EB93  80                 add a,b
EB94  80                 add a,b
EB95  80                 add a,b
EB96  80                 add a,b
EB97  80                 add a,b
EB98  07                 rlca
EB99  07                 rlca
EB9A  07                 rlca
EB9B  07                 rlca
EB9C  07                 rlca
EB9D  07                 rlca
EB9E  07                 rlca
EB9F  07                 rlca
EBA0  02                 ld (bc),a
EBA1  07                 rlca
EBA2  07                 rlca
EBA3  07                 rlca
EBA4  07                 rlca
EBA5  07                 rlca
EBA6  07                 rlca
EBA7  07                 rlca
EBA8  07                 rlca
EBA9  07                 rlca
EBAA  07                 rlca
EBAB  07                 rlca
EBAC  07                 rlca
EBAD  07                 rlca
EBAE  07                 rlca
EBAF  02                 ld (bc),a
EBB0  80                 add a,b
EBB1  80                 add a,b
EBB2  80                 add a,b
EBB3  80                 add a,b
EBB4  80                 add a,b
EBB5  80                 add a,b
EBB6  80                 add a,b
EBB7  80                 add a,b
EBB8  80                 add a,b
EBB9  80                 add a,b
EBBA  80                 add a,b
EBBB  80                 add a,b
EBBC  80                 add a,b
EBBD  80                 add a,b
EBBE  80                 add a,b
EBBF  80                 add a,b
EBC0  44                 ld b,h
EBC1  80                 add a,b
EBC2  80                 add a,b
EBC3  80                 add a,b
EBC4  80                 add a,b
EBC5  80                 add a,b
EBC6  80                 add a,b
EBC7  80                 add a,b
EBC8  80                 add a,b
EBC9  80                 add a,b
EBCA  80                 add a,b
EBCB  80                 add a,b
EBCC  80                 add a,b
EBCD  34                 inc (hl)
EBCE  35                 dec (hl)
EBCF  37                 scf
EBD0  38 80              jr c,$EB52
EBD2  80                 add a,b
EBD3  80                 add a,b
EBD4  02                 ld (bc),a
EBD5  07                 rlca
EBD6  06 06              ld b,6
EBD8  06 06              ld b,6
EBDA  06 06              ld b,6
EBDC  06 06              ld b,6
EBDE  06 06              ld b,6
EBE0  06 06              ld b,6
EBE2  06 06              ld b,6
EBE4  06 02              ld b,2
EBE6  07                 rlca
EBE7  06 06              ld b,6
EBE9  06 36              ld b,54
EBEB  37                 scf
EBEC  38 04              jr c,$EBF2
EBEE  80                 add a,b
EBEF  80                 add a,b
EBF0  06 07              ld b,7
EBF2  06 06              ld b,6
EBF4  06 06              ld b,6
EBF6  06 30              ld b,48
EBF8  31 32 33           ld sp,$3332
EBFB  06 06              ld b,6
EBFD  02                 ld (bc),a
EBFE  02                 ld (bc),a
EBFF  02                 ld (bc),a
EC00  02                 ld (bc),a
EC01  80                 add a,b
EC02  80                 add a,b
EC03  80                 add a,b
EC04  80                 add a,b
EC05  34                 inc (hl)
EC06  35                 dec (hl)
EC07  36 37              ld (hl),55
EC09  38 80              jr c,$EB8B
EC0B  80                 add a,b
EC0C  80                 add a,b
EC0D  80                 add a,b
EC0E  80                 add a,b
EC0F  80                 add a,b
EC10  80                 add a,b
EC11  80                 add a,b
EC12  80                 add a,b
EC13  80                 add a,b
EC14  80                 add a,b
EC15  80                 add a,b
EC16  34                 inc (hl)
EC17  35                 dec (hl)
EC18  06 06              ld b,6
EC1A  06 06              ld b,6
EC1C  06 06              ld b,6
EC1E  06 06              ld b,6
EC20  02                 ld (bc),a
EC21  07                 rlca
EC22  06 06              ld b,6
EC24  06 06              ld b,6
EC26  06 06              ld b,6
EC28  06 06              ld b,6
EC2A  06 06              ld b,6
EC2C  06 06              ld b,6
EC2E  06 02              ld b,2
EC30  80                 add a,b
EC31  80                 add a,b
EC32  80                 add a,b
EC33  80                 add a,b
EC34  80                 add a,b
EC35  80                 add a,b
EC36  80                 add a,b
EC37  80                 add a,b
EC38  80                 add a,b
EC39  80                 add a,b
EC3A  80                 add a,b
EC3B  80                 add a,b
EC3C  80                 add a,b
EC3D  80                 add a,b
EC3E  80                 add a,b
EC3F  80                 add a,b
EC40  44                 ld b,h
EC41  80                 add a,b
EC42  80                 add a,b
EC43  80                 add a,b
EC44  34                 inc (hl)
EC45  35                 dec (hl)
EC46  36 36              ld (hl),54
EC48  37                 scf
EC49  38 80              jr c,$EBCB
EC4B  80                 add a,b
EC4C  80                 add a,b
EC4D  80                 add a,b
EC4E  80                 add a,b
EC4F  80                 add a,b
EC50  80                 add a,b
EC51  80                 add a,b
EC52  80                 add a,b
EC53  80                 add a,b
EC54  02                 ld (bc),a
EC55  07                 rlca
EC56  06 06              ld b,6
EC58  06 06              ld b,6
EC5A  06 06              ld b,6
EC5C  06 06              ld b,6
EC5E  06 06              ld b,6
EC60  06 06              ld b,6
EC62  06 06              ld b,6
EC64  06 02              ld b,2
EC66  07                 rlca
EC67  06 06              ld b,6
EC69  06 80              ld b,-128
EC6B  80                 add a,b
EC6C  80                 add a,b
EC6D  80                 add a,b
EC6E  80                 add a,b
EC6F  80                 add a,b
EC70  06 06              ld b,6
EC72  06 06              ld b,6
EC74  06 06              ld b,6
EC76  06 2F              ld b,47
EC78  6C                 ld l,h
EC79  6C                 ld l,h
EC7A  2E 06              ld l,6
EC7C  06 02              ld b,2
EC7E  02                 ld (bc),a
EC7F  02                 ld (bc),a
EC80  02                 ld (bc),a
EC81  37                 scf
EC82  38 80              jr c,$EC04
EC84  80                 add a,b
EC85  80                 add a,b
EC86  80                 add a,b
EC87  80                 add a,b
EC88  80                 add a,b
EC89  80                 add a,b
EC8A  80                 add a,b
EC8B  80                 add a,b
EC8C  80                 add a,b
EC8D  80                 add a,b
EC8E  80                 add a,b
EC8F  80                 add a,b
EC90  80                 add a,b
EC91  80                 add a,b
EC92  80                 add a,b
EC93  80                 add a,b
EC94  80                 add a,b
EC95  80                 add a,b
EC96  80                 add a,b
EC97  80                 add a,b
EC98  06 06              ld b,6
EC9A  06 06              ld b,6
EC9C  06 00              ld b,0
EC9E  00                 nop
EC9F  00                 nop
ECA0  02                 ld (bc),a
ECA1  07                 rlca
ECA2  06 06              ld b,6
ECA4  06 06              ld b,6
ECA6  06 01              ld b,1
ECA8  00                 nop
ECA9  03                 inc bc
ECAA  00                 nop
ECAB  00                 nop
ECAC  00                 nop
ECAD  06 06              ld b,6
ECAF  02                 ld (bc),a
ECB0  80                 add a,b
ECB1  80                 add a,b
ECB2  80                 add a,b
ECB3  80                 add a,b
ECB4  80                 add a,b
ECB5  80                 add a,b
ECB6  80                 add a,b
ECB7  80                 add a,b
ECB8  80                 add a,b
ECB9  80                 add a,b
ECBA  80                 add a,b
ECBB  80                 add a,b
ECBC  80                 add a,b
ECBD  80                 add a,b
ECBE  80                 add a,b
ECBF  80                 add a,b
ECC0  44                 ld b,h
ECC1  80                 add a,b
ECC2  80                 add a,b
ECC3  80                 add a,b
ECC4  80                 add a,b
ECC5  80                 add a,b
ECC6  80                 add a,b
ECC7  80                 add a,b
ECC8  80                 add a,b
ECC9  80                 add a,b
ECCA  80                 add a,b
ECCB  80                 add a,b
ECCC  80                 add a,b
ECCD  80                 add a,b
ECCE  80                 add a,b
ECCF  80                 add a,b
ECD0  80                 add a,b
ECD1  80                 add a,b
ECD2  80                 add a,b
ECD3  80                 add a,b
ECD4  02                 ld (bc),a
ECD5  07                 rlca
ECD6  06 06              ld b,6
ECD8  06 06              ld b,6
ECDA  06 00              ld b,0
ECDC  00                 nop
ECDD  00                 nop
ECDE  06 06              ld b,6
ECE0  00                 nop
ECE1  03                 inc bc
ECE2  01 06 06           ld bc,1542
ECE5  02                 ld (bc),a
ECE6  07                 rlca
ECE7  06 06              ld b,6
ECE9  06 80              ld b,-128
ECEB  80                 add a,b
ECEC  80                 add a,b
ECED  04                 inc b
ECEE  80                 add a,b
ECEF  80                 add a,b
ECF0  06 06              ld b,6
ECF2  06 06              ld b,6
ECF4  06 06              ld b,6
ECF6  06 2F              ld b,47
ECF8  6C                 ld l,h
ECF9  6C                 ld l,h
ECFA  2E 06              ld l,6
ECFC  06 02              ld b,2
ECFE  02                 ld (bc),a
ECFF  02                 ld (bc),a
ED00  02                 ld (bc),a
ED01  80                 add a,b
ED02  80                 add a,b
ED03  80                 add a,b
ED04  80                 add a,b
ED05  80                 add a,b
ED06  80                 add a,b
ED07  80                 add a,b
ED08  80                 add a,b
ED09  80                 add a,b
ED0A  80                 add a,b
ED0B  80                 add a,b
ED0C  80                 add a,b
ED0D  80                 add a,b
ED0E  80                 add a,b
ED0F  80                 add a,b
ED10  00                 nop
ED11  00                 nop
ED12  00                 nop
ED13  80                 add a,b
ED14  80                 add a,b
ED15  00                 nop
ED16  00                 nop
ED17  00                 nop
ED18  03                 inc bc
ED19  00                 nop
ED1A  00                 nop
ED1B  06 06              ld b,6
ED1D  06 07              ld b,7
ED1F  07                 rlca
ED20  02                 ld (bc),a
ED21  00                 nop
ED22  00                 nop
ED23  00                 nop
ED24  00                 nop
ED25  06 06              ld b,6
ED27  02                 ld (bc),a
ED28  07                 rlca
ED29  07                 rlca
ED2A  07                 rlca
ED2B  07                 rlca
ED2C  07                 rlca
ED2D  07                 rlca
ED2E  06 02              ld b,2
ED30  80                 add a,b
ED31  80                 add a,b
ED32  80                 add a,b
ED33  80                 add a,b
ED34  80                 add a,b
ED35  80                 add a,b
ED36  80                 add a,b
ED37  80                 add a,b
ED38  80                 add a,b
ED39  80                 add a,b
ED3A  80                 add a,b
ED3B  80                 add a,b
ED3C  80                 add a,b
ED3D  80                 add a,b
ED3E  80                 add a,b
ED3F  80                 add a,b
ED40  44                 ld b,h
ED41  80                 add a,b
ED42  80                 add a,b
ED43  80                 add a,b
ED44  80                 add a,b
ED45  80                 add a,b
ED46  80                 add a,b
ED47  80                 add a,b
ED48  80                 add a,b
ED49  80                 add a,b
ED4A  80                 add a,b
ED4B  80                 add a,b
ED4C  80                 add a,b
ED4D  80                 add a,b
ED4E  34                 inc (hl)
ED4F  37                 scf
ED50  38 80              jr c,$ECD2
ED52  80                 add a,b
ED53  80                 add a,b
ED54  02                 ld (bc),a
ED55  07                 rlca
ED56  00                 nop
ED57  00                 nop
ED58  06 06              ld b,6
ED5A  00                 nop
ED5B  06 07              ld b,7
ED5D  07                 rlca
ED5E  07                 rlca
ED5F  06 06              ld b,6
ED61  07                 rlca
ED62  02                 ld (bc),a
ED63  07                 rlca
ED64  06 02              ld b,2
ED66  07                 rlca
ED67  06 06              ld b,6
ED69  01 80 80           ld bc,$8080
ED6C  80                 add a,b
ED6D  04                 inc b
ED6E  80                 add a,b
ED6F  80                 add a,b
ED70  01 6C 6C           ld bc,$6C6C
ED73  6C                 ld l,h
ED74  6C                 ld l,h
ED75  00                 nop
ED76  00                 nop
ED77  00                 nop
ED78  00                 nop
ED79  00                 nop
ED7A  00                 nop
ED7B  00                 nop
ED7C  00                 nop
ED7D  02                 ld (bc),a
ED7E  02                 ld (bc),a
ED7F  02                 ld (bc),a
ED80  02                 ld (bc),a
ED81  80                 add a,b
ED82  80                 add a,b
ED83  80                 add a,b
ED84  80                 add a,b
ED85  80                 add a,b
ED86  80                 add a,b
ED87  80                 add a,b
ED88  80                 add a,b
ED89  00                 nop
ED8A  80                 add a,b
ED8B  00                 nop
ED8C  80                 add a,b
ED8D  00                 nop
ED8E  80                 add a,b
ED8F  80                 add a,b
ED90  80                 add a,b
ED91  80                 add a,b
ED92  80                 add a,b
ED93  80                 add a,b
ED94  80                 add a,b
ED95  80                 add a,b
ED96  80                 add a,b
ED97  80                 add a,b
ED98  01 07 07           ld bc,1799
ED9B  07                 rlca
ED9C  06 06              ld b,6
ED9E  06 06              ld b,6
EDA0  02                 ld (bc),a
EDA1  07                 rlca
EDA2  07                 rlca
EDA3  07                 rlca
EDA4  07                 rlca
EDA5  07                 rlca
EDA6  00                 nop
EDA7  02                 ld (bc),a
EDA8  07                 rlca
EDA9  06 06              ld b,6
EDAB  06 06              ld b,6
EDAD  06 06              ld b,6
EDAF  02                 ld (bc),a
EDB0  80                 add a,b
EDB1  80                 add a,b
EDB2  80                 add a,b
EDB3  80                 add a,b
EDB4  80                 add a,b
EDB5  80                 add a,b
EDB6  80                 add a,b
EDB7  80                 add a,b
EDB8  80                 add a,b
EDB9  80                 add a,b
EDBA  80                 add a,b
EDBB  80                 add a,b
EDBC  80                 add a,b
EDBD  80                 add a,b
EDBE  80                 add a,b
EDBF  51                 ld d,c
EDC0  44                 ld b,h
EDC1  52                 ld d,d
EDC2  80                 add a,b
EDC3  80                 add a,b
EDC4  80                 add a,b
EDC5  80                 add a,b
EDC6  80                 add a,b
EDC7  80                 add a,b
EDC8  80                 add a,b
EDC9  80                 add a,b
EDCA  80                 add a,b
EDCB  80                 add a,b
EDCC  80                 add a,b
EDCD  80                 add a,b
EDCE  80                 add a,b
EDCF  80                 add a,b
EDD0  80                 add a,b
EDD1  80                 add a,b
EDD2  80                 add a,b
EDD3  80                 add a,b
EDD4  02                 ld (bc),a
EDD5  07                 rlca
EDD6  06 07              ld b,7
EDD8  07                 rlca
EDD9  06 06              ld b,6
EDDB  07                 rlca
EDDC  06 06              ld b,6
EDDE  06 06              ld b,6
EDE0  06 06              ld b,6
EDE2  02                 ld (bc),a
EDE3  07                 rlca
EDE4  06 02              ld b,2
EDE6  07                 rlca
EDE7  06 00              ld b,0
EDE9  02                 ld (bc),a
EDEA  00                 nop
EDEB  80                 add a,b
EDEC  80                 add a,b
EDED  04                 inc b
EDEE  80                 add a,b
EDEF  00                 nop
EDF0  02                 ld (bc),a
EDF1  6C                 ld l,h
EDF2  6C                 ld l,h
EDF3  6C                 ld l,h
EDF4  6C                 ld l,h
EDF5  6C                 ld l,h
EDF6  6C                 ld l,h
EDF7  24                 inc h
EDF8  25                 dec h
EDF9  26 6C              ld h,108
EDFB  6C                 ld l,h
EDFC  6C                 ld l,h
EDFD  02                 ld (bc),a
EDFE  02                 ld (bc),a
EDFF  02                 ld (bc),a
EE00  02                 ld (bc),a
EE01  80                 add a,b
EE02  80                 add a,b
EE03  80                 add a,b
EE04  80                 add a,b
EE05  80                 add a,b
EE06  80                 add a,b
EE07  80                 add a,b
EE08  80                 add a,b
EE09  80                 add a,b
EE0A  80                 add a,b
EE0B  80                 add a,b
EE0C  80                 add a,b
EE0D  80                 add a,b
EE0E  80                 add a,b
EE0F  80                 add a,b
EE10  80                 add a,b
EE11  80                 add a,b
EE12  80                 add a,b
EE13  80                 add a,b
EE14  80                 add a,b
EE15  80                 add a,b
EE16  80                 add a,b
EE17  80                 add a,b
EE18  02                 ld (bc),a
EE19  07                 rlca
EE1A  06 06              ld b,6
EE1C  06 06              ld b,6
EE1E  06 06              ld b,6
EE20  02                 ld (bc),a
EE21  07                 rlca
EE22  06 06              ld b,6
EE24  06 06              ld b,6
EE26  06 02              ld b,2
EE28  07                 rlca
EE29  06 06              ld b,6
EE2B  06 06              ld b,6
EE2D  06 06              ld b,6
EE2F  02                 ld (bc),a
EE30  80                 add a,b
EE31  80                 add a,b
EE32  80                 add a,b
EE33  80                 add a,b
EE34  80                 add a,b
EE35  80                 add a,b
EE36  80                 add a,b
EE37  80                 add a,b
EE38  80                 add a,b
EE39  80                 add a,b
EE3A  80                 add a,b
EE3B  80                 add a,b
EE3C  80                 add a,b
EE3D  80                 add a,b
EE3E  80                 add a,b
EE3F  51                 ld d,c
EE40  44                 ld b,h
EE41  52                 ld d,d
EE42  80                 add a,b
EE43  80                 add a,b
EE44  80                 add a,b
EE45  04                 inc b
EE46  04                 inc b
EE47  04                 inc b
EE48  04                 inc b
EE49  80                 add a,b
EE4A  80                 add a,b
EE4B  80                 add a,b
EE4C  80                 add a,b
EE4D  80                 add a,b
EE4E  80                 add a,b
EE4F  80                 add a,b
EE50  80                 add a,b
EE51  80                 add a,b
EE52  80                 add a,b
EE53  80                 add a,b
EE54  02                 ld (bc),a
EE55  07                 rlca
EE56  06 06              ld b,6
EE58  06 06              ld b,6
EE5A  06 00              ld b,0
EE5C  00                 nop
EE5D  00                 nop
EE5E  00                 nop
EE5F  00                 nop
EE60  00                 nop
EE61  00                 nop
EE62  02                 ld (bc),a
EE63  07                 rlca
EE64  06 02              ld b,2
EE66  07                 rlca
EE67  06 06              ld b,6
EE69  02                 ld (bc),a
EE6A  80                 add a,b
EE6B  80                 add a,b
EE6C  80                 add a,b
EE6D  04                 inc b
EE6E  34                 inc (hl)
EE6F  35                 dec (hl)
EE70  02                 ld (bc),a
EE71  6C                 ld l,h
EE72  6C                 ld l,h
EE73  6C                 ld l,h
EE74  6C                 ld l,h
EE75  6C                 ld l,h
EE76  6C                 ld l,h
EE77  27                 daa
EE78  28 29              jr z,$EEA3
EE7A  6C                 ld l,h
EE7B  6C                 ld l,h
EE7C  00                 nop
EE7D  02                 ld (bc),a
EE7E  02                 ld (bc),a
EE7F  02                 ld (bc),a
EE80  02                 ld (bc),a
EE81  80                 add a,b
EE82  80                 add a,b
EE83  80                 add a,b
EE84  80                 add a,b
EE85  80                 add a,b
EE86  00                 nop
EE87  00                 nop
EE88  80                 add a,b
EE89  80                 add a,b
EE8A  80                 add a,b
EE8B  80                 add a,b
EE8C  80                 add a,b
EE8D  80                 add a,b
EE8E  80                 add a,b
EE8F  80                 add a,b
EE90  80                 add a,b
EE91  80                 add a,b
EE92  80                 add a,b
EE93  80                 add a,b
EE94  80                 add a,b
EE95  80                 add a,b
EE96  80                 add a,b
EE97  80                 add a,b
EE98  02                 ld (bc),a
EE99  00                 nop
EE9A  00                 nop
EE9B  00                 nop
EE9C  00                 nop
EE9D  03                 inc bc
EE9E  06 06              ld b,6
EEA0  02                 ld (bc),a
EEA1  07                 rlca
EEA2  06 06              ld b,6
EEA4  06 00              ld b,0
EEA6  00                 nop
EEA7  02                 ld (bc),a
EEA8  07                 rlca
EEA9  06 06              ld b,6
EEAB  06 06              ld b,6
EEAD  06 06              ld b,6
EEAF  02                 ld (bc),a
EEB0  80                 add a,b
EEB1  80                 add a,b
EEB2  80                 add a,b
EEB3  80                 add a,b
EEB4  80                 add a,b
EEB5  80                 add a,b
EEB6  80                 add a,b
EEB7  80                 add a,b
EEB8  80                 add a,b
EEB9  80                 add a,b
EEBA  80                 add a,b
EEBB  80                 add a,b
EEBC  80                 add a,b
EEBD  80                 add a,b
EEBE  80                 add a,b
EEBF  51                 ld d,c
EEC0  44                 ld b,h
EEC1  52                 ld d,d
EEC2  80                 add a,b
EEC3  80                 add a,b
EEC4  80                 add a,b
EEC5  80                 add a,b
EEC6  80                 add a,b
EEC7  80                 add a,b
EEC8  80                 add a,b
EEC9  80                 add a,b
EECA  80                 add a,b
EECB  80                 add a,b
EECC  80                 add a,b
EECD  80                 add a,b
EECE  80                 add a,b
EECF  01 06 06           ld bc,1542
EED2  06 06              ld b,6
EED4  02                 ld (bc),a
EED5  07                 rlca
EED6  06 06              ld b,6
EED8  06 00              ld b,0
EEDA  06 06              ld b,6
EEDC  07                 rlca
EEDD  07                 rlca
EEDE  07                 rlca
EEDF  07                 rlca
EEE0  07                 rlca
EEE1  07                 rlca
EEE2  02                 ld (bc),a
EEE3  07                 rlca
EEE4  06 02              ld b,2
EEE6  07                 rlca
EEE7  06 00              ld b,0
EEE9  02                 ld (bc),a
EEEA  00                 nop
EEEB  80                 add a,b
EEEC  80                 add a,b
EEED  04                 inc b
EEEE  80                 add a,b
EEEF  00                 nop
EEF0  02                 ld (bc),a
EEF1  6C                 ld l,h
EEF2  6C                 ld l,h
EEF3  6C                 ld l,h
EEF4  6C                 ld l,h
EEF5  6C                 ld l,h
EEF6  6C                 ld l,h
EEF7  2A 2B 2C           ld hl,($2C2B)
EEFA  6C                 ld l,h
EEFB  6C                 ld l,h
EEFC  6C                 ld l,h
EEFD  02                 ld (bc),a
EEFE  02                 ld (bc),a
EEFF  02                 ld (bc),a
EF00  02                 ld (bc),a
EF01  80                 add a,b
EF02  80                 add a,b
EF03  00                 nop
EF04  80                 add a,b
EF05  80                 add a,b
EF06  80                 add a,b
EF07  80                 add a,b
EF08  80                 add a,b
EF09  80                 add a,b
EF0A  80                 add a,b
EF0B  80                 add a,b
EF0C  80                 add a,b
EF0D  80                 add a,b
EF0E  80                 add a,b
EF0F  80                 add a,b
EF10  80                 add a,b
EF11  80                 add a,b
EF12  80                 add a,b
EF13  80                 add a,b
EF14  34                 inc (hl)
EF15  35                 dec (hl)
EF16  37                 scf
EF17  38 02              jr c,$EF1B
EF19  07                 rlca
EF1A  07                 rlca
EF1B  07                 rlca
EF1C  07                 rlca
EF1D  07                 rlca
EF1E  07                 rlca
EF1F  06 02              ld b,2
EF21  07                 rlca
EF22  06 06              ld b,6
EF24  06 06              ld b,6
EF26  07                 rlca
EF27  02                 ld (bc),a
EF28  07                 rlca
EF29  06 06              ld b,6
EF2B  06 06              ld b,6
EF2D  06 06              ld b,6
EF2F  02                 ld (bc),a
EF30  80                 add a,b
EF31  80                 add a,b
EF32  80                 add a,b
EF33  80                 add a,b
EF34  80                 add a,b
EF35  80                 add a,b
EF36  80                 add a,b
EF37  80                 add a,b
EF38  80                 add a,b
EF39  80                 add a,b
EF3A  80                 add a,b
EF3B  80                 add a,b
EF3C  80                 add a,b
EF3D  80                 add a,b
EF3E  80                 add a,b
EF3F  44                 ld b,h
EF40  44                 ld b,h
EF41  44                 ld b,h
EF42  80                 add a,b
EF43  80                 add a,b
EF44  80                 add a,b
EF45  80                 add a,b
EF46  80                 add a,b
EF47  80                 add a,b
EF48  80                 add a,b
EF49  80                 add a,b
EF4A  80                 add a,b
EF4B  80                 add a,b
EF4C  80                 add a,b
EF4D  00                 nop
EF4E  00                 nop
EF4F  02                 ld (bc),a
EF50  07                 rlca
EF51  06 06              ld b,6
EF53  06 02              ld b,2
EF55  07                 rlca
EF56  06 06              ld b,6
EF58  06 06              ld b,6
EF5A  07                 rlca
EF5B  06 06              ld b,6
EF5D  06 06              ld b,6
EF5F  06 06              ld b,6
EF61  06 02              ld b,2
EF63  07                 rlca
EF64  06 02              ld b,2
EF66  07                 rlca
EF67  06 06              ld b,6
EF69  02                 ld (bc),a
EF6A  80                 add a,b
EF6B  80                 add a,b
EF6C  80                 add a,b
EF6D  04                 inc b
EF6E  80                 add a,b
EF6F  80                 add a,b
EF70  02                 ld (bc),a
EF71  6C                 ld l,h
EF72  6C                 ld l,h
EF73  6C                 ld l,h
EF74  6C                 ld l,h
EF75  6C                 ld l,h
EF76  6C                 ld l,h
EF77  6C                 ld l,h
EF78  6C                 ld l,h
EF79  2D                 dec l
EF7A  6C                 ld l,h
EF7B  6C                 ld l,h
EF7C  00                 nop
EF7D  02                 ld (bc),a
EF7E  02                 ld (bc),a
EF7F  02                 ld (bc),a
EF80  02                 ld (bc),a
EF81  80                 add a,b
EF82  80                 add a,b
EF83  80                 add a,b
EF84  80                 add a,b
EF85  80                 add a,b
EF86  80                 add a,b
EF87  80                 add a,b
EF88  80                 add a,b
EF89  80                 add a,b
EF8A  80                 add a,b
EF8B  80                 add a,b
EF8C  80                 add a,b
EF8D  34                 inc (hl)
EF8E  37                 scf
EF8F  38 80              jr c,$EF11
EF91  80                 add a,b
EF92  80                 add a,b
EF93  80                 add a,b
EF94  80                 add a,b
EF95  80                 add a,b
EF96  80                 add a,b
EF97  80                 add a,b
EF98  02                 ld (bc),a
EF99  07                 rlca
EF9A  06 06              ld b,6
EF9C  06 06              ld b,6
EF9E  06 06              ld b,6
EFA0  02                 ld (bc),a
EFA1  07                 rlca
EFA2  06 00              ld b,0
EFA4  00                 nop
EFA5  06 06              ld b,6
EFA7  02                 ld (bc),a
EFA8  07                 rlca
EFA9  06 06              ld b,6
EFAB  06 06              ld b,6
EFAD  06 06              ld b,6
EFAF  02                 ld (bc),a
EFB0  80                 add a,b
EFB1  80                 add a,b
EFB2  80                 add a,b
EFB3  80                 add a,b
EFB4  80                 add a,b
EFB5  80                 add a,b
EFB6  80                 add a,b
EFB7  80                 add a,b
EFB8  80                 add a,b
EFB9  80                 add a,b
EFBA  80                 add a,b
EFBB  80                 add a,b
EFBC  80                 add a,b
EFBD  80                 add a,b
EFBE  39                 add hl,sp
EFBF  3A 3A 3A           ld a,($3A3A)
EFC2  3B                 dec sp
EFC3  80                 add a,b
EFC4  80                 add a,b
EFC5  80                 add a,b
EFC6  80                 add a,b
EFC7  80                 add a,b
EFC8  80                 add a,b
EFC9  80                 add a,b
EFCA  80                 add a,b
EFCB  80                 add a,b
EFCC  80                 add a,b
EFCD  80                 add a,b
EFCE  80                 add a,b
EFCF  02                 ld (bc),a
EFD0  07                 rlca
EFD1  06 06              ld b,6
EFD3  06 02              ld b,2
EFD5  07                 rlca
EFD6  06 00              ld b,0
EFD8  00                 nop
EFD9  00                 nop
EFDA  00                 nop
EFDB  00                 nop
EFDC  00                 nop
EFDD  00                 nop
EFDE  00                 nop
EFDF  00                 nop
EFE0  00                 nop
EFE1  01 02 07           ld bc,1794
EFE4  06 02              ld b,2
EFE6  07                 rlca
EFE7  06 00              ld b,0
EFE9  02                 ld (bc),a
EFEA  00                 nop
EFEB  80                 add a,b
EFEC  80                 add a,b
EFED  04                 inc b
EFEE  80                 add a,b
EFEF  00                 nop
EFF0  02                 ld (bc),a
EFF1  6C                 ld l,h
EFF2  6C                 ld l,h
EFF3  6C                 ld l,h
EFF4  6C                 ld l,h
EFF5  6C                 ld l,h
EFF6  6C                 ld l,h
EFF7  6C                 ld l,h
EFF8  6C                 ld l,h
EFF9  2D                 dec l
EFFA  6C                 ld l,h
EFFB  6C                 ld l,h
EFFC  6C                 ld l,h
EFFD  02                 ld (bc),a
EFFE  02                 ld (bc),a
EFFF  02                 ld (bc),a
F000  02                 ld (bc),a
F001  80                 add a,b
F002  00                 nop
F003  80                 add a,b
F004  80                 add a,b
F005  80                 add a,b
F006  80                 add a,b
F007  80                 add a,b
F008  80                 add a,b
F009  80                 add a,b
F00A  80                 add a,b
F00B  80                 add a,b
F00C  80                 add a,b
F00D  80                 add a,b
F00E  80                 add a,b
F00F  80                 add a,b
F010  80                 add a,b
F011  80                 add a,b
F012  80                 add a,b
F013  80                 add a,b
F014  80                 add a,b
F015  80                 add a,b
F016  80                 add a,b
F017  80                 add a,b
F018  02                 ld (bc),a
F019  07                 rlca
F01A  06 00              ld b,0
F01C  00                 nop
F01D  00                 nop
F01E  00                 nop
F01F  03                 inc bc
F020  02                 ld (bc),a
F021  07                 rlca
F022  06 06              ld b,6
F024  07                 rlca
F025  07                 rlca
F026  06 02              ld b,2
F028  07                 rlca
F029  06 06              ld b,6
F02B  06 06              ld b,6
F02D  06 06              ld b,6
F02F  02                 ld (bc),a
F030  80                 add a,b
F031  80                 add a,b
F032  80                 add a,b
F033  80                 add a,b
F034  80                 add a,b
F035  80                 add a,b
F036  80                 add a,b
F037  80                 add a,b
F038  80                 add a,b
F039  80                 add a,b
F03A  80                 add a,b
F03B  80                 add a,b
F03C  80                 add a,b
F03D  80                 add a,b
F03E  3C                 inc a
F03F  44                 ld b,h
F040  44                 ld b,h
F041  44                 ld b,h
F042  43                 ld b,e
F043  80                 add a,b
F044  80                 add a,b
F045  80                 add a,b
F046  80                 add a,b
F047  80                 add a,b
F048  80                 add a,b
F049  80                 add a,b
F04A  00                 nop
F04B  00                 nop
F04C  80                 add a,b
F04D  80                 add a,b
F04E  80                 add a,b
F04F  02                 ld (bc),a
F050  07                 rlca
F051  06 00              ld b,0
F053  00                 nop
F054  02                 ld (bc),a
F055  07                 rlca
F056  06 06              ld b,6
F058  07                 rlca
F059  07                 rlca
F05A  07                 rlca
F05B  07                 rlca
F05C  07                 rlca
F05D  07                 rlca
F05E  07                 rlca
F05F  07                 rlca
F060  07                 rlca
F061  02                 ld (bc),a
F062  07                 rlca
F063  07                 rlca
F064  06 02              ld b,2
F066  07                 rlca
F067  06 06              ld b,6
F069  02                 ld (bc),a
F06A  80                 add a,b
F06B  80                 add a,b
F06C  80                 add a,b
F06D  04                 inc b
F06E  80                 add a,b
F06F  80                 add a,b
F070  02                 ld (bc),a
F071  00                 nop
F072  6C                 ld l,h
F073  6C                 ld l,h
F074  6C                 ld l,h
F075  6C                 ld l,h
F076  6C                 ld l,h
F077  6C                 ld l,h
F078  6C                 ld l,h
F079  2D                 dec l
F07A  6C                 ld l,h
F07B  6C                 ld l,h
F07C  00                 nop
F07D  02                 ld (bc),a
F07E  02                 ld (bc),a
F07F  02                 ld (bc),a
F080  02                 ld (bc),a
F081  80                 add a,b
F082  80                 add a,b
F083  80                 add a,b
F084  80                 add a,b
F085  34                 inc (hl)
F086  35                 dec (hl)
F087  36 37              ld (hl),55
F089  38 80              jr c,$F00B
F08B  80                 add a,b
F08C  80                 add a,b
F08D  80                 add a,b
F08E  80                 add a,b
F08F  80                 add a,b
F090  80                 add a,b
F091  80                 add a,b
F092  80                 add a,b
F093  80                 add a,b
F094  80                 add a,b
F095  80                 add a,b
F096  80                 add a,b
F097  80                 add a,b
F098  02                 ld (bc),a
F099  07                 rlca
F09A  06 06              ld b,6
F09C  07                 rlca
F09D  07                 rlca
F09E  07                 rlca
F09F  07                 rlca
F0A0  02                 ld (bc),a
F0A1  00                 nop
F0A2  06 06              ld b,6
F0A4  06 05              ld b,5
F0A6  05                 dec b
F0A7  02                 ld (bc),a
F0A8  05                 dec b
F0A9  05                 dec b
F0AA  05                 dec b
F0AB  06 06              ld b,6
F0AD  05                 dec b
F0AE  05                 dec b
F0AF  02                 ld (bc),a
F0B0  80                 add a,b
F0B1  80                 add a,b
F0B2  80                 add a,b
F0B3  80                 add a,b
F0B4  80                 add a,b
F0B5  80                 add a,b
F0B6  80                 add a,b
F0B7  80                 add a,b
F0B8  80                 add a,b
F0B9  80                 add a,b
F0BA  80                 add a,b
F0BB  80                 add a,b
F0BC  80                 add a,b
F0BD  3E 3F              ld a,63
F0BF  46                 ld b,(hl)
F0C0  47                 ld b,a
F0C1  49                 ld c,c
F0C2  40                 ld b,b
F0C3  41                 ld b,c
F0C4  80                 add a,b
F0C5  80                 add a,b
F0C6  80                 add a,b
F0C7  80                 add a,b
F0C8  80                 add a,b
F0C9  80                 add a,b
F0CA  80                 add a,b
F0CB  80                 add a,b
F0CC  80                 add a,b
F0CD  80                 add a,b
F0CE  00                 nop
F0CF  02                 ld (bc),a
F0D0  07                 rlca
F0D1  06 06              ld b,6
F0D3  07                 rlca
F0D4  02                 ld (bc),a
F0D5  07                 rlca
F0D6  06 06              ld b,6
F0D8  06 06              ld b,6
F0DA  06 06              ld b,6
F0DC  06 00              ld b,0
F0DE  01 07 06           ld bc,1543
F0E1  02                 ld (bc),a
F0E2  07                 rlca
F0E3  06 06              ld b,6
F0E5  02                 ld (bc),a
F0E6  07                 rlca
F0E7  06 00              ld b,0
F0E9  02                 ld (bc),a
F0EA  80                 add a,b
F0EB  00                 nop
F0EC  00                 nop
F0ED  80                 add a,b
F0EE  80                 add a,b
F0EF  00                 nop
F0F0  02                 ld (bc),a
F0F1  6C                 ld l,h
F0F2  6C                 ld l,h
F0F3  6C                 ld l,h
F0F4  6C                 ld l,h
F0F5  6C                 ld l,h
F0F6  6C                 ld l,h
F0F7  6C                 ld l,h
F0F8  6C                 ld l,h
F0F9  2D                 dec l
F0FA  6C                 ld l,h
F0FB  6C                 ld l,h
F0FC  6C                 ld l,h
F0FD  02                 ld (bc),a
F0FE  02                 ld (bc),a
F0FF  02                 ld (bc),a
F100  02                 ld (bc),a
F101  80                 add a,b
F102  00                 nop
F103  80                 add a,b
F104  80                 add a,b
F105  80                 add a,b
F106  80                 add a,b
F107  80                 add a,b
F108  80                 add a,b
F109  80                 add a,b
F10A  80                 add a,b
F10B  80                 add a,b
F10C  80                 add a,b
F10D  80                 add a,b
F10E  80                 add a,b
F10F  80                 add a,b
F110  80                 add a,b
F111  80                 add a,b
F112  80                 add a,b
F113  80                 add a,b
F114  80                 add a,b
F115  80                 add a,b
F116  80                 add a,b
F117  80                 add a,b
F118  02                 ld (bc),a
F119  07                 rlca
F11A  06 06              ld b,6
F11C  06 06              ld b,6
F11E  06 06              ld b,6
F120  02                 ld (bc),a
F121  07                 rlca
F122  07                 rlca
F123  06 06              ld b,6
F125  02                 ld (bc),a
F126  02                 ld (bc),a
F127  02                 ld (bc),a
F128  02                 ld (bc),a
F129  02                 ld (bc),a
F12A  01 07 06           ld bc,1543
F12D  01 01 02           ld bc,513
F130  80                 add a,b
F131  80                 add a,b
F132  80                 add a,b
F133  04                 inc b
F134  04                 inc b
F135  04                 inc b
F136  04                 inc b
F137  04                 inc b
F138  04                 inc b
F139  04                 inc b
F13A  80                 add a,b
F13B  80                 add a,b
F13C  80                 add a,b
F13D  3C                 inc a
F13E  3D                 dec a
F13F  80                 add a,b
F140  80                 add a,b
F141  80                 add a,b
F142  42                 ld b,d
F143  43                 ld b,e
F144  80                 add a,b
F145  80                 add a,b
F146  80                 add a,b
F147  80                 add a,b
F148  80                 add a,b
F149  80                 add a,b
F14A  80                 add a,b
F14B  80                 add a,b
F14C  80                 add a,b
F14D  80                 add a,b
F14E  80                 add a,b
F14F  02                 ld (bc),a
F150  07                 rlca
F151  06 06              ld b,6
F153  06 02              ld b,2
F155  01 03 01           ld bc,259
F158  01 03 01           ld bc,259
F15B  06 06              ld b,6
F15D  06 02              ld b,2
F15F  07                 rlca
F160  06 02              ld b,2
F162  07                 rlca
F163  06 06              ld b,6
F165  02                 ld (bc),a
F166  07                 rlca
F167  06 06              ld b,6
F169  02                 ld (bc),a
F16A  80                 add a,b
F16B  80                 add a,b
F16C  80                 add a,b
F16D  80                 add a,b
F16E  80                 add a,b
F16F  80                 add a,b
F170  02                 ld (bc),a
F171  00                 nop
F172  6C                 ld l,h
F173  6C                 ld l,h
F174  6C                 ld l,h
F175  6C                 ld l,h
F176  6C                 ld l,h
F177  6C                 ld l,h
F178  6C                 ld l,h
F179  2D                 dec l
F17A  6C                 ld l,h
F17B  6C                 ld l,h
F17C  00                 nop
F17D  02                 ld (bc),a
F17E  02                 ld (bc),a
F17F  02                 ld (bc),a
F180  02                 ld (bc),a
F181  80                 add a,b
F182  80                 add a,b
F183  80                 add a,b
F184  00                 nop
F185  80                 add a,b
F186  80                 add a,b
F187  00                 nop
F188  80                 add a,b
F189  80                 add a,b
F18A  00                 nop
F18B  00                 nop
F18C  80                 add a,b
F18D  80                 add a,b
F18E  80                 add a,b
F18F  80                 add a,b
F190  80                 add a,b
F191  80                 add a,b
F192  80                 add a,b
F193  80                 add a,b
F194  80                 add a,b
F195  80                 add a,b
F196  80                 add a,b
F197  79                 ld a,c
F198  02                 ld (bc),a
F199  00                 nop
F19A  00                 nop
F19B  00                 nop
F19C  00                 nop
F19D  00                 nop
F19E  06 06              ld b,6
F1A0  02                 ld (bc),a
F1A1  00                 nop
F1A2  06 06              ld b,6
F1A4  06 06              ld b,6
F1A6  07                 rlca
F1A7  07                 rlca
F1A8  07                 rlca
F1A9  02                 ld (bc),a
F1AA  07                 rlca
F1AB  07                 rlca
F1AC  06 06              ld b,6
F1AE  07                 rlca
F1AF  07                 rlca
F1B0  80                 add a,b
F1B1  80                 add a,b
F1B2  80                 add a,b
F1B3  80                 add a,b
F1B4  80                 add a,b
F1B5  80                 add a,b
F1B6  80                 add a,b
F1B7  80                 add a,b
F1B8  80                 add a,b
F1B9  80                 add a,b
F1BA  80                 add a,b
F1BB  80                 add a,b
F1BC  39                 add hl,sp
F1BD  3A 3A 3A           ld a,($3A3A)
F1C0  3A 3A 3A           ld a,($3A3A)
F1C3  3A 3B 80           ld a,($803B)
F1C6  80                 add a,b
F1C7  80                 add a,b
F1C8  34                 inc (hl)
F1C9  35                 dec (hl)
F1CA  36 37              ld (hl),55
F1CC  38 80              jr c,$F14E
F1CE  00                 nop
F1CF  02                 ld (bc),a
F1D0  07                 rlca
F1D1  06 00              ld b,0
F1D3  00                 nop
F1D4  02                 ld (bc),a
F1D5  02                 ld (bc),a
F1D6  02                 ld (bc),a
F1D7  02                 ld (bc),a
F1D8  02                 ld (bc),a
F1D9  02                 ld (bc),a
F1DA  02                 ld (bc),a
F1DB  07                 rlca
F1DC  06 00              ld b,0
F1DE  02                 ld (bc),a
F1DF  07                 rlca
F1E0  06 02              ld b,2
F1E2  07                 rlca
F1E3  06 00              ld b,0
F1E5  02                 ld (bc),a
F1E6  07                 rlca
F1E7  06 00              ld b,0
F1E9  02                 ld (bc),a
F1EA  00                 nop
F1EB  00                 nop
F1EC  80                 add a,b
F1ED  80                 add a,b
F1EE  80                 add a,b
F1EF  00                 nop
F1F0  02                 ld (bc),a
F1F1  6C                 ld l,h
F1F2  6C                 ld l,h
F1F3  6C                 ld l,h
F1F4  6C                 ld l,h
F1F5  6C                 ld l,h
F1F6  6C                 ld l,h
F1F7  6C                 ld l,h
F1F8  6C                 ld l,h
F1F9  2D                 dec l
F1FA  6C                 ld l,h
F1FB  6C                 ld l,h
F1FC  6C                 ld l,h
F1FD  02                 ld (bc),a
F1FE  02                 ld (bc),a
F1FF  02                 ld (bc),a
F200  02                 ld (bc),a
F201  80                 add a,b
F202  80                 add a,b
F203  80                 add a,b
F204  80                 add a,b
F205  80                 add a,b
F206  80                 add a,b
F207  80                 add a,b
F208  80                 add a,b
F209  80                 add a,b
F20A  80                 add a,b
F20B  80                 add a,b
F20C  80                 add a,b
F20D  80                 add a,b
F20E  00                 nop
F20F  00                 nop
F210  00                 nop
F211  80                 add a,b
F212  80                 add a,b
F213  00                 nop
F214  00                 nop
F215  79                 ld a,c
F216  7A                 ld a,d
F217  00                 nop
F218  02                 ld (bc),a
F219  07                 rlca
F21A  07                 rlca
F21B  07                 rlca
F21C  07                 rlca
F21D  07                 rlca
F21E  07                 rlca
F21F  06 02              ld b,2
F221  07                 rlca
F222  07                 rlca
F223  06 06              ld b,6
F225  06 06              ld b,6
F227  06 06              ld b,6
F229  02                 ld (bc),a
F22A  07                 rlca
F22B  06 06              ld b,6
F22D  06 06              ld b,6
F22F  06 80              ld b,-128
F231  80                 add a,b
F232  80                 add a,b
F233  80                 add a,b
F234  80                 add a,b
F235  80                 add a,b
F236  80                 add a,b
F237  80                 add a,b
F238  80                 add a,b
F239  80                 add a,b
F23A  04                 inc b
F23B  04                 inc b
F23C  3C                 inc a
F23D  44                 ld b,h
F23E  44                 ld b,h
F23F  44                 ld b,h
F240  44                 ld b,h
F241  44                 ld b,h
F242  44                 ld b,h
F243  44                 ld b,h
F244  43                 ld b,e
F245  80                 add a,b
F246  80                 add a,b
F247  80                 add a,b
F248  80                 add a,b
F249  80                 add a,b
F24A  80                 add a,b
F24B  80                 add a,b
F24C  80                 add a,b
F24D  80                 add a,b
F24E  80                 add a,b
F24F  02                 ld (bc),a
F250  07                 rlca
F251  06 06              ld b,6
F253  07                 rlca
F254  07                 rlca
F255  07                 rlca
F256  07                 rlca
F257  07                 rlca
F258  07                 rlca
F259  07                 rlca
F25A  07                 rlca
F25B  07                 rlca
F25C  06 06              ld b,6
F25E  02                 ld (bc),a
F25F  07                 rlca
F260  06 02              ld b,2
F262  07                 rlca
F263  06 06              ld b,6
F265  07                 rlca
F266  07                 rlca
F267  06 06              ld b,6
F269  02                 ld (bc),a
F26A  80                 add a,b
F26B  80                 add a,b
F26C  80                 add a,b
F26D  80                 add a,b
F26E  80                 add a,b
F26F  80                 add a,b
F270  02                 ld (bc),a
F271  00                 nop
F272  6C                 ld l,h
F273  6C                 ld l,h
F274  6C                 ld l,h
F275  6C                 ld l,h
F276  6C                 ld l,h
F277  6C                 ld l,h
F278  6C                 ld l,h
F279  2D                 dec l
F27A  6C                 ld l,h
F27B  6C                 ld l,h
F27C  00                 nop
F27D  02                 ld (bc),a
F27E  02                 ld (bc),a
F27F  02                 ld (bc),a
F280  02                 ld (bc),a
F281  80                 add a,b
F282  80                 add a,b
F283  80                 add a,b
F284  80                 add a,b
F285  80                 add a,b
F286  80                 add a,b
F287  80                 add a,b
F288  79                 ld a,c
F289  7A                 ld a,d
F28A  7B                 ld a,e
F28B  7C                 ld a,h
F28C  80                 add a,b
F28D  80                 add a,b
F28E  80                 add a,b
F28F  80                 add a,b
F290  80                 add a,b
F291  80                 add a,b
F292  80                 add a,b
F293  79                 ld a,c
F294  7A                 ld a,d
F295  7F                 ld a,a
F296  7F                 ld a,a
F297  7F                 ld a,a
F298  02                 ld (bc),a
F299  07                 rlca
F29A  06 06              ld b,6
F29C  06 06              ld b,6
F29E  06 06              ld b,6
F2A0  02                 ld (bc),a
F2A1  00                 nop
F2A2  00                 nop
F2A3  00                 nop
F2A4  00                 nop
F2A5  00                 nop
F2A6  00                 nop
F2A7  06 06              ld b,6
F2A9  02                 ld (bc),a
F2AA  07                 rlca
F2AB  06 06              ld b,6
F2AD  06 06              ld b,6
F2AF  06 80              ld b,-128
F2B1  80                 add a,b
F2B2  80                 add a,b
F2B3  80                 add a,b
F2B4  80                 add a,b
F2B5  80                 add a,b
F2B6  80                 add a,b
F2B7  80                 add a,b
F2B8  80                 add a,b
F2B9  80                 add a,b
F2BA  80                 add a,b
F2BB  4B                 ld c,e
F2BC  4C                 ld c,h
F2BD  45                 ld b,l
F2BE  46                 ld b,(hl)
F2BF  47                 ld b,a
F2C0  80                 add a,b
F2C1  48                 ld c,b
F2C2  49                 ld c,c
F2C3  4A                 ld c,d
F2C4  4F                 ld c,a
F2C5  50                 ld d,b
F2C6  80                 add a,b
F2C7  80                 add a,b
F2C8  80                 add a,b
F2C9  80                 add a,b
F2CA  80                 add a,b
F2CB  80                 add a,b
F2CC  80                 add a,b
F2CD  80                 add a,b
F2CE  00                 nop
F2CF  02                 ld (bc),a
F2D0  07                 rlca
F2D1  06 06              ld b,6
F2D3  06 06              ld b,6
F2D5  06 06              ld b,6
F2D7  06 06              ld b,6
F2D9  06 06              ld b,6
F2DB  06 06              ld b,6
F2DD  00                 nop
F2DE  02                 ld (bc),a
F2DF  07                 rlca
F2E0  06 02              ld b,2
F2E2  07                 rlca
F2E3  06 06              ld b,6
F2E5  06 06              ld b,6
F2E7  06 00              ld b,0
F2E9  02                 ld (bc),a
F2EA  80                 add a,b
F2EB  00                 nop
F2EC  00                 nop
F2ED  80                 add a,b
F2EE  80                 add a,b
F2EF  00                 nop
F2F0  02                 ld (bc),a
F2F1  6C                 ld l,h
F2F2  6C                 ld l,h
F2F3  6C                 ld l,h
F2F4  6C                 ld l,h
F2F5  6C                 ld l,h
F2F6  6C                 ld l,h
F2F7  6C                 ld l,h
F2F8  6C                 ld l,h
F2F9  2D                 dec l
F2FA  6C                 ld l,h
F2FB  6C                 ld l,h
F2FC  6C                 ld l,h
F2FD  02                 ld (bc),a
F2FE  02                 ld (bc),a
F2FF  02                 ld (bc),a
F300  02                 ld (bc),a
F301  79                 ld a,c
F302  7A                 ld a,d
F303  7B                 ld a,e
F304  7C                 ld a,h
F305  80                 add a,b
F306  79                 ld a,c
F307  7A                 ld a,d
F308  7F                 ld a,a
F309  7F                 ld a,a
F30A  7F                 ld a,a
F30B  7F                 ld a,a
F30C  7D                 ld a,l
F30D  7B                 ld a,e
F30E  7C                 ld a,h
F30F  80                 add a,b
F310  79                 ld a,c
F311  7A                 ld a,d
F312  7D                 ld a,l
F313  7E                 ld a,(hl)
F314  7F                 ld a,a
F315  7F                 ld a,a
F316  7F                 ld a,a
F317  00                 nop
F318  02                 ld (bc),a
F319  07                 rlca
F31A  06 03              ld b,3
F31C  00                 nop
F31D  03                 inc bc
F31E  00                 nop
F31F  00                 nop
F320  02                 ld (bc),a
F321  07                 rlca
F322  07                 rlca
F323  07                 rlca
F324  07                 rlca
F325  07                 rlca
F326  07                 rlca
F327  07                 rlca
F328  00                 nop
F329  02                 ld (bc),a
F32A  07                 rlca
F32B  06 06              ld b,6
F32D  06 06              ld b,6
F32F  06 7A              ld b,122
F331  7B                 ld a,e
F332  7C                 ld a,h
F333  80                 add a,b
F334  80                 add a,b
F335  80                 add a,b
F336  80                 add a,b
F337  80                 add a,b
F338  80                 add a,b
F339  80                 add a,b
F33A  4B                 ld c,e
F33B  4C                 ld c,h
F33C  4D                 ld c,l
F33D  80                 add a,b
F33E  80                 add a,b
F33F  80                 add a,b
F340  01 80 80           ld bc,$8080
F343  80                 add a,b
F344  4E                 ld c,(hl)
F345  4F                 ld c,a
F346  50                 ld d,b
F347  80                 add a,b
F348  80                 add a,b
F349  00                 nop
F34A  00                 nop
F34B  00                 nop
F34C  80                 add a,b
F34D  80                 add a,b
F34E  80                 add a,b
F34F  02                 ld (bc),a
F350  07                 rlca
F351  06 06              ld b,6
F353  06 06              ld b,6
F355  06 06              ld b,6
F357  06 06              ld b,6
F359  06 06              ld b,6
F35B  06 06              ld b,6
F35D  06 02              ld b,2
F35F  07                 rlca
F360  06 02              ld b,2
F362  01 03 01           ld bc,259
F365  01 03 01           ld bc,259
F368  01 02 80           ld bc,$8002
F36B  80                 add a,b
F36C  80                 add a,b
F36D  80                 add a,b
F36E  80                 add a,b
F36F  80                 add a,b
F370  02                 ld (bc),a
F371  00                 nop
F372  6C                 ld l,h
F373  6C                 ld l,h
F374  6C                 ld l,h
F375  6C                 ld l,h
F376  6C                 ld l,h
F377  6C                 ld l,h
F378  6C                 ld l,h
F379  6C                 ld l,h
F37A  6C                 ld l,h
F37B  6C                 ld l,h
F37C  00                 nop
F37D  02                 ld (bc),a
F37E  02                 ld (bc),a
F37F  02                 ld (bc),a
F380  02                 ld (bc),a
F381  7F                 ld a,a
F382  7F                 ld a,a
F383  7F                 ld a,a
F384  7F                 ld a,a
F385  7D                 ld a,l
F386  7E                 ld a,(hl)
F387  7F                 ld a,a
F388  7F                 ld a,a
F389  7F                 ld a,a
F38A  7F                 ld a,a
F38B  7F                 ld a,a
F38C  7F                 ld a,a
F38D  7F                 ld a,a
F38E  7F                 ld a,a
F38F  7D                 ld a,l
F390  7E                 ld a,(hl)
F391  7F                 ld a,a
F392  7F                 ld a,a
F393  7F                 ld a,a
F394  7F                 ld a,a
F395  7F                 ld a,a
F396  7F                 ld a,a
F397  7F                 ld a,a
F398  02                 ld (bc),a
F399  07                 rlca
F39A  06 06              ld b,6
F39C  07                 rlca
F39D  07                 rlca
F39E  07                 rlca
F39F  07                 rlca
F3A0  07                 rlca
F3A1  07                 rlca
F3A2  06 06              ld b,6
F3A4  06 06              ld b,6
F3A6  06 06              ld b,6
F3A8  06 02              ld b,2
F3AA  07                 rlca
F3AB  06 06              ld b,6
F3AD  06 06              ld b,6
F3AF  06 7F              ld b,127
F3B1  7F                 ld a,a
F3B2  7F                 ld a,a
F3B3  7D                 ld a,l
F3B4  7B                 ld a,e
F3B5  7C                 ld a,h
F3B6  80                 add a,b
F3B7  80                 add a,b
F3B8  80                 add a,b
F3B9  4B                 ld c,e
F3BA  4C                 ld c,h
F3BB  4D                 ld c,l
F3BC  80                 add a,b
F3BD  80                 add a,b
F3BE  80                 add a,b
F3BF  80                 add a,b
F3C0  02                 ld (bc),a
F3C1  80                 add a,b
F3C2  80                 add a,b
F3C3  80                 add a,b
F3C4  80                 add a,b
F3C5  4E                 ld c,(hl)
F3C6  4F                 ld c,a
F3C7  50                 ld d,b
F3C8  80                 add a,b
F3C9  80                 add a,b
F3CA  80                 add a,b
F3CB  80                 add a,b
F3CC  80                 add a,b
F3CD  80                 add a,b
F3CE  80                 add a,b
F3CF  02                 ld (bc),a
F3D0  07                 rlca
F3D1  06 06              ld b,6
F3D3  06 06              ld b,6
F3D5  06 06              ld b,6
F3D7  06 06              ld b,6
F3D9  06 06              ld b,6
F3DB  06 00              ld b,0
F3DD  00                 nop
F3DE  02                 ld (bc),a
F3DF  07                 rlca
F3E0  06 06              ld b,6
F3E2  07                 rlca
F3E3  07                 rlca
F3E4  07                 rlca
F3E5  07                 rlca
F3E6  07                 rlca
F3E7  07                 rlca
F3E8  07                 rlca
F3E9  07                 rlca
F3EA  80                 add a,b
F3EB  80                 add a,b
F3EC  80                 add a,b
F3ED  80                 add a,b
F3EE  80                 add a,b
F3EF  00                 nop
F3F0  02                 ld (bc),a
F3F1  6C                 ld l,h
F3F2  6C                 ld l,h
F3F3  6C                 ld l,h
F3F4  6C                 ld l,h
F3F5  6C                 ld l,h
F3F6  6C                 ld l,h
F3F7  6C                 ld l,h
F3F8  6C                 ld l,h
F3F9  6C                 ld l,h
F3FA  6C                 ld l,h
F3FB  6C                 ld l,h
F3FC  6C                 ld l,h
F3FD  02                 ld (bc),a
F3FE  02                 ld (bc),a
F3FF  02                 ld (bc),a
F400  02                 ld (bc),a
F401  00                 nop
F402  00                 nop
F403  00                 nop
F404  00                 nop
F405  00                 nop
F406  00                 nop
F407  00                 nop
F408  00                 nop
F409  00                 nop
F40A  00                 nop
F40B  00                 nop
F40C  00                 nop
F40D  00                 nop
F40E  00                 nop
F40F  00                 nop
F410  00                 nop
F411  00                 nop
F412  00                 nop
F413  00                 nop
F414  00                 nop
F415  00                 nop
F416  00                 nop
F417  00                 nop
F418  02                 ld (bc),a
F419  07                 rlca
F41A  06 06              ld b,6
F41C  06 06              ld b,6
F41E  06 06              ld b,6
F420  06 06              ld b,6
F422  06 06              ld b,6
F424  06 06              ld b,6
F426  06 00              ld b,0
F428  00                 nop
F429  02                 ld (bc),a
F42A  07                 rlca
F42B  06 06              ld b,6
F42D  06 06              ld b,6
F42F  06 7F              ld b,127
F431  7F                 ld a,a
F432  7F                 ld a,a
F433  7F                 ld a,a
F434  7F                 ld a,a
F435  7F                 ld a,a
F436  7B                 ld a,e
F437  7C                 ld a,h
F438  4B                 ld c,e
F439  4C                 ld c,h
F43A  4D                 ld c,l
F43B  80                 add a,b
F43C  80                 add a,b
F43D  80                 add a,b
F43E  80                 add a,b
F43F  80                 add a,b
F440  02                 ld (bc),a
F441  80                 add a,b
F442  80                 add a,b
F443  80                 add a,b
F444  80                 add a,b
F445  80                 add a,b
F446  4E                 ld c,(hl)
F447  4F                 ld c,a
F448  50                 ld d,b
F449  80                 add a,b
F44A  80                 add a,b
F44B  80                 add a,b
F44C  80                 add a,b
F44D  80                 add a,b
F44E  00                 nop
F44F  02                 ld (bc),a
F450  07                 rlca
F451  06 06              ld b,6
F453  06 06              ld b,6
F455  06 06              ld b,6
F457  06 06              ld b,6
F459  06 06              ld b,6
F45B  06 06              ld b,6
F45D  07                 rlca
F45E  02                 ld (bc),a
F45F  07                 rlca
F460  06 06              ld b,6
F462  06 06              ld b,6
F464  06 06              ld b,6
F466  06 06              ld b,6
F468  06 06              ld b,6
F46A  80                 add a,b
F46B  80                 add a,b
F46C  80                 add a,b
F46D  80                 add a,b
F46E  80                 add a,b
F46F  80                 add a,b
F470  02                 ld (bc),a
F471  00                 nop
F472  00                 nop
F473  00                 nop
F474  00                 nop
F475  00                 nop
F476  00                 nop
F477  00                 nop
F478  00                 nop
F479  00                 nop
F47A  00                 nop
F47B  00                 nop
F47C  00                 nop
F47D  02                 ld (bc),a
F47E  02                 ld (bc),a
F47F  02                 ld (bc),a
F480  02                 ld (bc),a
F481  01 01 01           ld bc,257
F484  01 01 01           ld bc,257
F487  01 01 01           ld bc,257
F48A  01 01 01           ld bc,257
F48D  01 01 01           ld bc,257
F490  01 01 01           ld bc,257
F493  01 01 01           ld bc,257
F496  01 01 02           ld bc,513
F499  03                 inc bc
F49A  00                 nop
F49B  00                 nop
F49C  00                 nop
F49D  03                 inc bc
F49E  00                 nop
F49F  00                 nop
F4A0  03                 inc bc
F4A1  00                 nop
F4A2  00                 nop
F4A3  03                 inc bc
F4A4  00                 nop
F4A5  00                 nop
F4A6  00                 nop
F4A7  00                 nop
F4A8  01 02 01           ld bc,258
F4AB  03                 inc bc
F4AC  01 01 01           ld bc,257
F4AF  01 00 00           ld bc,0
F4B2  00                 nop
F4B3  03                 inc bc
F4B4  03                 inc bc
F4B5  00                 nop
F4B6  00                 nop
F4B7  00                 nop
F4B8  00                 nop
F4B9  00                 nop
F4BA  00                 nop
F4BB  03                 inc bc
F4BC  03                 inc bc
F4BD  03                 inc bc
F4BE  03                 inc bc
F4BF  03                 inc bc
F4C0  02                 ld (bc),a
F4C1  03                 inc bc
F4C2  03                 inc bc
F4C3  03                 inc bc
F4C4  03                 inc bc
F4C5  03                 inc bc
F4C6  00                 nop
F4C7  00                 nop
F4C8  00                 nop
F4C9  00                 nop
F4CA  00                 nop
F4CB  00                 nop
F4CC  00                 nop
F4CD  00                 nop
F4CE  00                 nop
F4CF  02                 ld (bc),a
F4D0  00                 nop
F4D1  03                 inc bc
F4D2  00                 nop
F4D3  00                 nop
F4D4  03                 inc bc
F4D5  00                 nop
F4D6  00                 nop
F4D7  00                 nop
F4D8  00                 nop
F4D9  00                 nop
F4DA  00                 nop
F4DB  00                 nop
F4DC  00                 nop
F4DD  00                 nop
F4DE  02                 ld (bc),a
F4DF  00                 nop
F4E0  03                 inc bc
F4E1  00                 nop
F4E2  00                 nop
F4E3  03                 inc bc
F4E4  00                 nop
F4E5  00                 nop
F4E6  03                 inc bc
F4E7  00                 nop
F4E8  00                 nop
F4E9  00                 nop
F4EA  00                 nop
F4EB  00                 nop
F4EC  00                 nop
F4ED  00                 nop
F4EE  00                 nop
F4EF  00                 nop
F4F0  02                 ld (bc),a
F4F1  02                 ld (bc),a
F4F2  02                 ld (bc),a
F4F3  02                 ld (bc),a
F4F4  02                 ld (bc),a
F4F5  02                 ld (bc),a
F4F6  02                 ld (bc),a
F4F7  02                 ld (bc),a
F4F8  02                 ld (bc),a
F4F9  02                 ld (bc),a
F4FA  02                 ld (bc),a
F4FB  02                 ld (bc),a
F4FC  02                 ld (bc),a
F4FD  02                 ld (bc),a
F4FE  02                 ld (bc),a
F4FF  02                 ld (bc),a
F500  F3                 di
F501  3E 01              ld a,1
F503  32 FF FF           ld ($FFFF),a
F506  3E 00              ld a,0
F508  F3                 di
F509  21 00 6F           ld hl,$6F00
F50C  11 00 5B           ld de,$5B00
F50F  01 00 78           ld bc,$7800
F512  ED B0              ldir
F514  32 96 B0           ld ($B096),a
F517  31 60 B0           ld sp,$B060
F51A  CD 30 F5           call $F530
F51D  FB                 ei
F51E  AF                 xor a
F51F  D3 FE              out ($FE),a
F521  CD 93 F5           call $F593
F524  CD 7F F5           call $F57F
F527  CD 5E F5           call $F55E
F52A  CD 44 F5           call $F544
F52D  C3 16 B3           jp $B316
F530  21 00 AF           ld hl,$AF00
F533  11 01 AF           ld de,$AF01
F536  01 00 01           ld bc,256
F539  36 B0              ld (hl),-80
F53B  ED B0              ldir
F53D  3E AF              ld a,-81
F53F  ED 47              ld i,a
F541  ED 5E              im 2
F543  C9                 ret
F544  F3                 di
F545  01 00 50           ld bc,$5000
F548  DB 1F              in a,($1F)
F54A  B1                 or c
F54B  4F                 ld c,a
F54C  FB                 ei
F54D  76                 halt
F54E  F3                 di
F54F  10 F7              djnz $F548
F551  79                 ld a,c
F552  E6 E0              and $E0
F554  3E 00              ld a,0
F556  28 01              jr z,$F559
F558  3D                 dec a
F559  32 86 B0           ld ($B086),a
F55C  FB                 ei
F55D  C9                 ret
F55E  21 00 E8           ld hl,$E800
F561  7D                 ld a,l
F562  1F                 rra
F563  CB 11              rl c
F565  1F                 rra
F566  CB 11              rl c
F568  1F                 rra
F569  CB 11              rl c
F56B  1F                 rra
F56C  CB 11              rl c
F56E  1F                 rra
F56F  CB 11              rl c
F571  1F                 rra
F572  CB 11              rl c
F574  1F                 rra
F575  CB 11              rl c
F577  1F                 rra
F578  CB 11              rl c
F57A  71                 ld (hl),c
F57B  2C                 inc l
F57C  20 E3              jr nz,$F561
F57E  C9                 ret
F57F  11 00 E7           ld de,$E700
F582  2E 00              ld l,0
F584  26 80              ld h,-128
F586  AF                 xor a
F587  06 2F              ld b,47
F589  AE                 xor (hl)
F58A  24                 inc h
F58B  10 FC              djnz $F589
F58D  12                 ld (de),a
F58E  1C                 inc e
F58F  2C                 inc l
F590  20 F2              jr nz,$F584
F592  C9                 ret
F593  21 4B F7           ld hl,$F74B
F596  11 80 50           ld de,$5080
F599  06 20              ld b,32
F59B  C5                 push bc
F59C  D5                 push de
F59D  01 20 00           ld bc,32
F5A0  ED B0              ldir
F5A2  D1                 pop de
F5A3  C1                 pop bc
F5A4  CD B5 F5           call $F5B5
F5A7  10 F2              djnz $F59B
F5A9  21 00 E7           ld hl,$E700
F5AC  11 80 5A           ld de,$5A80
F5AF  01 80 00           ld bc,128
F5B2  ED B0              ldir
F5B4  C9                 ret
F5B5  14                 inc d
F5B6  7A                 ld a,d
F5B7  E6 07              and $07
F5B9  C0                 ret nz
F5BA  7B                 ld a,e
F5BB  C6 20              add a,32
F5BD  5F                 ld e,a
F5BE  D8                 ret c
F5BF  7A                 ld a,d
F5C0  D6 08              sub 8
F5C2  57                 ld d,a
F5C3  C9                 ret
F5C4  02                 ld (bc),a
F5C5  00                 nop
F5C6  00                 nop
F5C7  48                 ld c,b
F5C8  45                 ld b,l
F5C9  52                 ld d,d
F5CA  45                 ld b,l
F5CB  27                 daa
F5CC  53                 ld d,e
F5CD  20 41              jr nz,$F610
F5CF  20 4E              jr nz,$F61F
F5D1  4F                 ld c,a
F5D2  54                 ld d,h
F5D3  45                 ld b,l
F5D4  20 46              jr nz,$F61C
F5D6  4F                 ld c,a
F5D7  52                 ld d,d
F5D8  20 41              jr nz,$F61B
F5DA  4C                 ld c,h
F5DB  4C                 ld c,h
F5DC  20 54              jr nz,$F632
F5DE  48                 ld c,b
F5DF  4F                 ld c,a
F5E0  53                 ld d,e
F5E1  45                 ld b,l
F5E2  20 50              jr nz,$F634
F5E4  4F                 ld c,a
F5E5  4F                 ld c,a
F5E6  52                 ld d,d
F5E7  00                 nop
F5E8  53                 ld d,e
F5E9  4F                 ld c,a
F5EA  55                 ld d,l
F5EB  4C                 ld c,h
F5EC  53                 ld d,e
F5ED  20 57              jr nz,$F646
F5EF  49                 ld c,c
F5F0  54                 ld d,h
F5F1  48                 ld c,b
F5F2  20 52              jr nz,$F646
F5F4  55                 ld d,l
F5F5  42                 ld b,d
F5F6  42                 ld b,d
F5F7  45                 ld b,l
F5F8  52                 ld d,d
F5F9  20 4B              jr nz,$F646
F5FB  45                 ld b,l
F5FC  59                 ld e,c
F5FD  45                 ld b,l
F5FE  44                 ld b,h
F5FF  20 53              jr nz,$F654
F601  50                 ld d,b
F602  45                 ld b,l
F603  43                 ld b,e
F604  43                 ld b,e
F605  59                 ld e,c
F606  53                 ld d,e
F607  00                 nop
F608  4F                 ld c,a
F609  52                 ld d,d
F60A  20 54              jr nz,$F660
F60C  57                 ld d,a
F60D  49                 ld c,c
F60E  4E                 ld c,(hl)
F60F  20 53              jr nz,$F664
F611  49                 ld c,c
F612  4E                 ld c,(hl)
F613  43                 ld b,e
F614  4C                 ld c,h
F615  41                 ld b,c
F616  49                 ld c,c
F617  52                 ld d,d
F618  20 53              jr nz,$F66D
F61A  54                 ld d,h
F61B  49                 ld c,c
F61C  43                 ld b,e
F61D  4B                 ld c,e
F61E  53                 ld d,e
F61F  2C                 inc l
F620  46                 ld b,(hl)
F621  4F                 ld c,a
F622  52                 ld d,d
F623  20 57              jr nz,$F67C
F625  48                 ld c,b
F626  4F                 ld c,a
F627  4D                 ld c,l
F628  00                 nop
F629  54                 ld d,h
F62A  49                 ld c,c
F62B  4D                 ld c,l
F62C  53                 ld d,e
F62D  20 53              jr nz,$F682
F62F  55                 ld d,l
F630  50                 ld d,b
F631  45                 ld b,l
F632  52                 ld d,d
F633  20 44              jr nz,$F679
F635  55                 ld d,l
F636  50                 ld d,b
F637  45                 ld b,l
F638  52                 ld d,d
F639  20 4E              jr nz,$F689
F63B  4F                 ld c,a
F63C  2D                 dec l
F63D  43                 ld b,e
F63E  4C                 ld c,h
F63F  41                 ld b,c
F640  53                 ld d,e
F641  48                 ld c,b
F642  20 4B              jr nz,$F68F
F644  45                 ld b,l
F645  59                 ld e,c
F646  53                 ld d,e
F647  00                 nop
F648  57                 ld d,a
F649  4F                 ld c,a
F64A  4E                 ld c,(hl)
F64B  27                 daa
F64C  54                 ld d,h
F64D  20 57              jr nz,$F6A6
F64F  4F                 ld c,a
F650  52                 ld d,d
F651  4B                 ld c,e
F652  2E 00              ld l,0
F654  20 20              jr nz,$F676
F656  20 20              jr nz,$F678
F658  20 20              jr nz,$F67A
F65A  20 20              jr nz,$F67C
F65C  20 20              jr nz,$F67E
F65E  20 20              jr nz,$F680
F660  4D                 ld c,l
F661  4F                 ld c,a
F662  52                 ld d,d
F663  45                 ld b,l
F664  20 3E              jr nz,$F6A4
F666  3E 3E              ld a,62
F668  BE                 cp (hl)
F669  49                 ld c,c
F66A  46                 ld b,(hl)
F66B  20 59              jr nz,$F6C6
F66D  4F                 ld c,a
F66E  55                 ld d,l
F66F  20 54              jr nz,$F6C5
F671  59                 ld e,c
F672  50                 ld d,b
F673  45                 ld b,l
F674  20 27              jr nz,$F69D
F676  54                 ld d,h
F677  52                 ld d,d
F678  55                 ld d,l
F679  4B                 ld c,e
F67A  45                 ld b,l
F67B  59                 ld e,c
F67C  53                 ld d,e
F67D  27                 daa
F67E  20 53              jr nz,$F6D3
F680  54                 ld d,h
F681  52                 ld d,d
F682  41                 ld b,c
F683  49                 ld c,c
F684  47                 ld b,a
F685  48                 ld c,b
F686  54                 ld d,h
F687  00                 nop
F688  41                 ld b,c
F689  46                 ld b,(hl)
F68A  54                 ld d,h
F68B  45                 ld b,l
F68C  52                 ld d,d
F68D  20 4C              jr nz,$F6DB
F68F  4F                 ld c,a
F690  41                 ld b,c
F691  44                 ld b,h
F692  49                 ld c,c
F693  4E                 ld c,(hl)
F694  47                 ld b,a
F695  20 54              jr nz,$F6EB
F697  48                 ld c,b
F698  45                 ld b,l
F699  20 47              jr nz,$F6E2
F69B  41                 ld b,c
F69C  4D                 ld c,l
F69D  45                 ld b,l
F69E  2C                 inc l
F69F  20 54              jr nz,$F6F5
F6A1  48                 ld c,b
F6A2  45                 ld b,l
F6A3  4E                 ld c,(hl)
F6A4  00                 nop
F6A5  4B                 ld c,e
F6A6  45                 ld b,l
F6A7  59                 ld e,c
F6A8  20 4F              jr nz,$F6F9
F6AA  50                 ld d,b
F6AB  54                 ld d,h
F6AC  49                 ld c,c
F6AD  4F                 ld c,a
F6AE  4E                 ld c,(hl)
F6AF  53                 ld d,e
F6B0  20 36              jr nz,$F6E8
F6B2  20 41              jr nz,$F6F5
F6B4  4E                 ld c,(hl)
F6B5  44                 ld b,h
F6B6  20 37              jr nz,$F6EF
F6B8  20 41              jr nz,$F6FB
F6BA  52                 ld d,d
F6BB  45                 ld b,l
F6BC  20 54              jr nz,$F712
F6BE  48                 ld c,b
F6BF  45                 ld b,l
F6C0  01 4F 52           ld bc,$524F
F6C3  49                 ld c,c
F6C4  47                 ld b,a
F6C5  49                 ld c,c
F6C6  4E                 ld c,(hl)
F6C7  41                 ld b,c
F6C8  4C                 ld c,h
F6C9  20 43              jr nz,$F70E
F6CB  4C                 ld c,h
F6CC  41                 ld b,c
F6CD  53                 ld d,e
F6CE  48                 ld c,b
F6CF  49                 ld c,c
F6D0  4E                 ld c,(hl)
F6D1  47                 ld b,a
F6D2  20 44              jr nz,$F718
F6D4  45                 ld b,l
F6D5  56                 ld d,(hl)
F6D6  45                 ld b,l
F6D7  4C                 ld c,h
F6D8  4F                 ld c,a
F6D9  50                 ld d,b
F6DA  4D                 ld c,l
F6DB  45                 ld b,l
F6DC  4E                 ld c,(hl)
F6DD  54                 ld d,h
F6DE  00                 nop
F6DF  53                 ld d,e
F6E0  45                 ld b,l
F6E1  54                 ld d,h
F6E2  2C                 inc l
F6E3  20 41              jr nz,$F726
F6E5  4E                 ld c,(hl)
F6E6  44                 ld b,h
F6E7  20 44              jr nz,$F72D
F6E9  4F                 ld c,a
F6EA  55                 ld d,l
F6EB  42                 ld b,d
F6EC  4C                 ld c,h
F6ED  45                 ld b,l
F6EE  20 53              jr nz,$F743
F6F0  54                 ld d,h
F6F1  49                 ld c,c
F6F2  58                 ld e,b
F6F3  2E 20              ld l,32
F6F5  54                 ld d,h
F6F6  48                 ld c,b
F6F7  49                 ld c,c
F6F8  53                 ld d,e
F6F9  20 4E              jr nz,$F749
F6FB  45                 ld b,l
F6FC  45                 ld b,l
F6FD  44                 ld b,h
F6FE  53                 ld d,e
F6FF  00                 nop
F700  54                 ld d,h
F701  48                 ld c,b
F702  45                 ld b,l
F703  20 50              jr nz,$F755
F705  41                 ld b,c
F706  55                 ld d,l
F707  53                 ld d,e
F708  45                 ld b,l
F709  20 41              jr nz,$F74C
F70B  4E                 ld c,(hl)
F70C  44                 ld b,h
F70D  20 51              jr nz,$F760
F70F  55                 ld d,l
F710  49                 ld c,c
F711  54                 ld d,h
F712  20 4B              jr nz,$F75F
F714  45                 ld b,l
F715  59                 ld e,c
F716  53                 ld d,e
F717  20 54              jr nz,$F76D
F719  4F                 ld c,a
F71A  20 42              jr nz,$F75E
F71C  45                 ld b,l
F71D  01 27 53           ld bc,$5327
F720  59                 ld e,c
F721  4D                 ld c,l
F722  42                 ld b,d
F723  4F                 ld c,a
F724  4C                 ld c,h
F725  2D                 dec l
F726  53                 ld d,e
F727  48                 ld c,b
F728  49                 ld c,c
F729  46                 ld b,(hl)
F72A  54                 ld d,h
F72B  45                 ld b,l
F72C  44                 ld b,h
F72D  27                 daa
F72E  01 20 20           ld bc,$2020
F731  20 20              jr nz,$F753
F733  20 20              jr nz,$F755
F735  20 20              jr nz,$F757
F737  20 52              jr nz,$F78B
F739  20 46              jr nz,$F781
F73B  52                 ld d,d
F73C  45                 ld b,l
F73D  44                 ld b,h
F73E  20 57              jr nz,$F797
F740  20 20              jr nz,$F762
F742  32 39 2D           ld ($2D39),a
F745  31 2D 39           ld sp,$392D
F748  31 20 A0           ld sp,$A020
F74B  FF                 rst 38h
F74C  FF                 rst 38h
F74D  60                 ld h,b
F74E  00                 nop
F74F  00                 nop
F750  00                 nop
F751  C0                 ret nz
F752  00                 nop
F753  03                 inc bc
F754  26 26              ld h,38
F756  FF                 rst 38h
F757  FF                 rst 38h
F758  FF                 rst 38h
F759  FF                 rst 38h
F75A  FF                 rst 38h
F75B  FF                 rst 38h
F75C  FF                 rst 38h
F75D  FF                 rst 38h
F75E  FF                 rst 38h
F75F  FF                 rst 38h
F760  64                 ld h,h
F761  64                 ld h,h
F762  E0                 ret po
F763  00                 nop
F764  01 80 00           ld bc,128
F767  00                 nop
F768  03                 inc bc
F769  FF                 rst 38h
F76A  FF                 rst 38h
F76B  F6 51              or $51
F76D  F8                 ret m
F76E  00                 nop
F76F  00                 nop
F770  03                 inc bc
F771  E0                 ret po
F772  00                 nop
F773  02                 ld (bc),a
F774  0B                 dec bc
F775  0B                 dec bc
F776  88                 adc a,b
F777  B1                 or c
F778  F1                 pop af
F779  71                 ld (hl),c
F77A  15                 dec d
F77B  15                 dec d
F77C  18 A8              jr $F726
F77E  88                 adc a,b
F77F  8F                 adc a,a
F780  D0                 ret nc
F781  D0                 ret nc
F782  A0                 and b
F783  00                 nop
F784  03                 inc bc
F785  E0                 ret po
F786  00                 nop
F787  00                 nop
F788  0F                 rrca
F789  C6 51              add a,81
F78B  E6 52              and $52
F78D  7F                 ld a,a
F78E  80                 add a,b
F78F  00                 nop
F790  0B                 dec bc
F791  EC 00 03           call pe,$0300
F794  26 26              ld h,38
F796  AD                 xor l
F797  D7                 rst 10h
F798  F7                 rst 30h
F799  77                 ld (hl),a
F79A  55                 ld d,l
F79B  55                 ld d,l
F79C  BA                 cp d
F79D  AD                 xor l
F79E  DA BF 64           jp c,$64BF
F7A1  64                 ld h,h
F7A2  E0                 ret po
F7A3  00                 nop
F7A4  1B                 dec de
F7A5  E8                 ret pe
F7A6  00                 nop
F7A7  00                 nop
F7A8  FF                 rst 38h
F7A9  92                 sub d
F7AA  52                 ld d,d
F7AB  C6 52              add a,82
F7AD  F9                 ld sp,hl
F7AE  F0                 ret p
F7AF  00                 nop
F7B0  1D                 dec e
F7B1  9E                 sbc a,(hl)
F7B2  E0                 ret po
F7B3  02                 ld (bc),a
F7B4  0B                 dec bc
F7B5  0B                 dec bc
F7B6  BD                 cp l
F7B7  F1                 pop af
F7B8  F1                 pop af
F7B9  71                 ld (hl),c
F7BA  11 11 B8           ld de,$B811
F7BD  8D                 adc a,l
F7BE  DB 8F              in a,($8F)
F7C0  D0                 ret nc
F7C1  D0                 ret nc
F7C2  A0                 and b
F7C3  03                 inc bc
F7C4  BC                 cp h
F7C5  DC 00 07           call c,$0700
F7C8  CF                 rst 08h
F7C9  F2 52 E6           jp p,$E652
F7CC  51                 ld d,c
F7CD  00                 nop
F7CE  00                 nop
F7CF  00                 nop
F7D0  6F                 ld l,a
F7D1  FE F8              cp -8
F7D3  05                 dec b
F7D4  26 26              ld h,38
F7D6  BD                 cp l
F7D7  FD F7              rst 30h
F7D9  77                 ld (hl),a
F7DA  75                 ld (hl),l
F7DB  51                 ld d,c
F7DC  BA                 cp d
F7DD  8D                 adc a,l
F7DE  DB EF              in a,($EF)
F7E0  64                 ld h,h
F7E1  64                 ld h,h
F7E2  D0                 ret nc
F7E3  0F                 rrca
F7E4  BF                 cp a
F7E5  FB                 ei
F7E6  00                 nop
F7E7  00                 nop
F7E8  00                 nop
F7E9  C6 51              add a,81
F7EB  E6 53              and $53
F7ED  00                 nop
F7EE  00                 nop
F7EF  9B                 sbc a,e
F7F0  FC E1 E0           call m,$E0E1
F7F3  06 0B              ld b,11
F7F5  0B                 dec bc
F7F6  AD                 xor l
F7F7  FD F7              rst 30h
F7F9  77                 ld (hl),a
F7FA  75                 ld (hl),l
F7FB  55                 ld d,l
F7FC  BA                 cp d
F7FD  AD                 xor l
F7FE  DA EF D0           jp c,$D0EF
F801  D0                 ret nc
F802  B0                 or b
F803  03                 inc bc
F804  C3 9F EC           jp $EC9F
F807  80                 add a,b
F808  00                 nop
F809  9E                 sbc a,(hl)
F80A  53                 ld d,e
F80B  C3 33 00           jp $0033
F80E  00                 nop
F80F  00                 nop
F810  00                 nop
F811  00                 nop
F812  00                 nop
F813  05                 dec b
F814  26 26              ld h,38
F816  89                 adc a,c
F817  F1                 pop af
F818  F1                 pop af
F819  11 75 55           ld de,$5575
F81C  BA                 cp d
F81D  AD                 xor l
F81E  88                 adc a,b
F81F  8F                 adc a,a
F820  64                 ld h,h
F821  64                 ld h,h
F822  D0                 ret nc
F823  00                 nop
F824  00                 nop
F825  00                 nop
F826  00                 nop
F827  00                 nop
F828  00                 nop
F829  83                 add a,e
F82A  33                 inc sp
F82B  FF                 rst 38h
F82C  FF                 rst 38h
F82D  FF                 rst 38h
F82E  FF                 rst 38h
F82F  FF                 rst 38h
F830  FF                 rst 38h
F831  FF                 rst 38h
F832  FF                 rst 38h
F833  86                 add a,(hl)
F834  0B                 dec bc
F835  0B                 dec bc
F836  FF                 rst 38h
F837  FF                 rst 38h
F838  FF                 rst 38h
F839  FF                 rst 38h
F83A  FF                 rst 38h
F83B  FF                 rst 38h
F83C  FF                 rst 38h
F83D  FF                 rst 38h
F83E  FF                 rst 38h
F83F  FF                 rst 38h
F840  D0                 ret nc
F841  D0                 ret nc
F842  B0                 or b
F843  FF                 rst 38h
F844  FF                 rst 38h
F845  FF                 rst 38h
F846  FF                 rst 38h
F847  FF                 rst 38h
F848  FF                 rst 38h
F849  FF                 rst 38h
F84A  FF                 rst 38h
F84B  00                 nop
F84C  00                 nop
F84D  7F                 ld a,a
F84E  FF                 rst 38h
F84F  FF                 rst 38h
F850  FF                 rst 38h
F851  FF                 rst 38h
F852  FF                 rst 38h
F853  84                 add a,h
F854  00                 nop
F855  00                 nop
F856  00                 nop
F857  FF                 rst 38h
F858  00                 nop
F859  00                 nop
F85A  7F                 ld a,a
F85B  FF                 rst 38h
F85C  00                 nop
F85D  00                 nop
F85E  7F                 ld a,a
F85F  80                 add a,b
F860  00                 nop
F861  00                 nop
F862  11 FF FF           ld de,-1
F865  FF                 rst 38h
F866  FF                 rst 38h
F867  FF                 rst 38h
F868  FF                 rst 38h
F869  80                 add a,b
F86A  00                 nop
F86B  E0                 ret po
F86C  00                 nop
F86D  7F                 ld a,a
F86E  FF                 rst 38h
F86F  FF                 rst 38h
F870  FF                 rst 38h
F871  FF                 rst 38h
F872  FF                 rst 38h
F873  8E                 adc a,(hl)
F874  00                 nop
F875  00                 nop
F876  00                 nop
F877  55                 ld d,l
F878  3D                 dec a
F879  BC                 cp h
F87A  2A AA 3D           ld hl,($3DAA)
F87D  BC                 cp h
F87E  2A 00 00           ld hl,($0000)
F881  00                 nop
F882  39                 add hl,sp
F883  FF                 rst 38h
F884  FF                 rst 38h
F885  FF                 rst 38h
F886  FF                 rst 38h
F887  FF                 rst 38h
F888  FF                 rst 38h
F889  E0                 ret po
F88A  00                 nop
F88B  99                 sbc a,c
F88C  80                 add a,b
F88D  7F                 ld a,a
F88E  FF                 rst 38h
F88F  FF                 rst 38h
F890  FF                 rst 38h
F891  FF                 rst 38h
F892  FF                 rst 38h
F893  8C                 adc a,h
F894  00                 nop
F895  01 F0 7F           ld bc,$7FF0
F898  77                 ld (hl),a
F899  EE 7F              xor $7F
F89B  FF                 rst 38h
F89C  77                 ld (hl),a
F89D  EE 7F              xor $7F
F89F  07                 rlca
F8A0  C0                 ret nz
F8A1  00                 nop
F8A2  19                 add hl,de
F8A3  FF                 rst 38h
F8A4  FF                 rst 38h
F8A5  FF                 rst 38h
F8A6  FF                 rst 38h
F8A7  FF                 rst 38h
F8A8  FF                 rst 38h
F8A9  99                 sbc a,c
F8AA  80                 add a,b
F8AB  8D                 adc a,l
F8AC  00                 nop
F8AD  7F                 ld a,a
F8AE  FF                 rst 38h
F8AF  FF                 rst 38h
F8B0  FF                 rst 38h
F8B1  FF                 rst 38h
F8B2  FF                 rst 38h
F8B3  8A                 adc a,d
F8B4  00                 nop
F8B5  B2                 or d
F8B6  0C                 inc c
F8B7  2A 6D B6           ld hl,($B66D)
F8BA  2A AA 6D           ld hl,($6DAA)
F8BD  B6                 or (hl)
F8BE  55                 ld d,l
F8BF  18 26              jr $F8E7
F8C1  80                 add a,b
F8C2  29                 add hl,hl
F8C3  FF                 rst 38h
F8C4  FF                 rst 38h
F8C5  FF                 rst 38h
F8C6  FF                 rst 38h
F8C7  FF                 rst 38h
F8C8  FF                 rst 38h
F8C9  8D                 adc a,l
F8CA  00                 nop
F8CB  24                 inc h
F8CC  00                 nop
F8CD  7F                 ld a,a
F8CE  FF                 rst 38h
F8CF  FF                 rst 38h
F8D0  FF                 rst 38h
F8D1  FF                 rst 38h
F8D2  FF                 rst 38h
F8D3  98                 sbc a,b
F8D4  05                 dec b
F8D5  AD                 xor l
F8D6  22 55 6B           ld ($6B55),hl
F8D9  D6 55              sub 85
F8DB  55                 ld d,l
F8DC  6B                 ld l,e
F8DD  D6 2A              sub 42
F8DF  22 5A D0           ld ($D05A),hl
F8E2  0D                 dec c
F8E3  FF                 rst 38h
F8E4  FF                 rst 38h
F8E5  FF                 rst 38h
F8E6  FF                 rst 38h
F8E7  FF                 rst 38h
F8E8  FF                 rst 38h
F8E9  A4                 and h
F8EA  00                 nop
F8EB  01 DC 7F           ld bc,$7FDC
F8EE  FF                 rst 38h
F8EF  FF                 rst 38h
F8F0  FF                 rst 38h
F8F1  FF                 rst 38h
F8F2  FF                 rst 38h
F8F3  9A                 sbc a,d
F8F4  08                 ex af,af'
F8F5  7D                 ld a,l
F8F6  52                 ld d,d
F8F7  00                 nop
F8F8  69                 ld l,c
F8F9  96                 sub (hl)
F8FA  00                 nop
F8FB  00                 nop
F8FC  69                 ld l,c
F8FD  96                 sub (hl)
F8FE  00                 nop
F8FF  25                 dec h
F900  5F                 ld e,a
F901  08                 ex af,af'
F902  2D                 dec l
F903  FF                 rst 38h
F904  FF                 rst 38h
F905  FF                 rst 38h
F906  FF                 rst 38h
F907  FF                 rst 38h
F908  FF                 rst 38h
F909  81                 add a,c
F90A  DC 07 7F           call c,$7F07
F90D  7F                 ld a,a
F90E  FF                 rst 38h
F90F  FF                 rst 38h
F910  FF                 rst 38h
F911  FF                 rst 38h
F912  FF                 rst 38h
F913  98                 sbc a,b
F914  20 9B              jr nz,$F8B1
F916  29                 add hl,hl
F917  2A 75 AE           ld hl,($AE75)
F91A  55                 ld d,l
F91B  55                 ld d,l
F91C  75                 ld (hl),l
F91D  AE                 xor (hl)
F91E  55                 ld d,l
F91F  4A                 ld c,d
F920  6C                 ld l,h
F921  82                 add a,d
F922  0D                 dec c
F923  FF                 rst 38h
F924  FF                 rst 38h
F925  FF                 rst 38h
F926  FF                 rst 38h
F927  FF                 rst 38h
F928  FF                 rst 38h
F929  87                 add a,a
F92A  7F                 ld a,a
F92B  46                 ld b,(hl)
F92C  81                 add a,c
F92D  7F                 ld a,a
F92E  FF                 rst 38h
F92F  FF                 rst 38h
F930  FF                 rst 38h
F931  FF                 rst 38h
F932  FF                 rst 38h
F933  80                 add a,b
F934  08                 ex af,af'
F935  06 55              ld b,85
F937  00                 nop
F938  39                 add hl,sp
F939  9C                 sbc a,h
F93A  00                 nop
F93B  00                 nop
F93C  39                 add hl,sp
F93D  9C                 sbc a,h
F93E  00                 nop
F93F  55                 ld d,l
F940  30 08              jr nc,$F94A
F942  01 FF FF           ld bc,-1
F945  FF                 rst 38h
F946  FF                 rst 38h
F947  FF                 rst 38h
F948  FF                 rst 38h
F949  C6 81              add a,-127
F94B  24                 inc h
F94C  3A 00 00           ld a,($0000)
F94F  00                 nop
F950  00                 nop
F951  00                 nop
F952  00                 nop
F953  00                 nop
F954  02                 ld (bc),a
F955  1D                 dec e
F956  F9                 ld sp,hl
F957  00                 nop
F958  15                 dec d
F959  A8                 xor b
F95A  00                 nop
F95B  00                 nop
F95C  15                 dec d
F95D  A8                 xor b
F95E  00                 nop
F95F  4F                 ld c,a
F960  DC 20 3E           call c,$3E20
F963  00                 nop
F964  00                 nop
F965  00                 nop
F966  00                 nop
F967  00                 nop
F968  00                 nop
F969  24                 inc h
F96A  3A 93 DD           ld a,($DD93)
F96D  60                 ld h,b
F96E  01 00 FB           ld bc,-1280
F971  FF                 rst 38h
F972  F8                 ret m
F973  FE 20              cp 32
F975  17                 rla
F976  F5                 push af
F977  00                 nop
F978  2D                 dec l
F979  B4                 or h
F97A  00                 nop
F97B  00                 nop
F97C  2D                 dec l
F97D  B4                 or h
F97E  00                 nop
F97F  57                 ld d,a
F980  F4 02 3F           call p,$3F02
F983  8F                 adc a,a
F984  FF                 rst 38h
F985  FF                 rst 38h
F986  00                 nop
F987  40                 ld b,b
F988  03                 inc bc
F989  93                 sub e
F98A  DD 94              sub ixh
F98C  EF                 rst 28h
F98D  B0                 or b
F98E  03                 inc bc
F98F  80                 add a,b
F990  82                 add a,d
F991  5B                 ld e,e
F992  80                 add a,b
F993  00                 nop
F994  05                 dec b
F995  1F                 rra
F996  F9                 ld sp,hl
F997  AA                 xor d
F998  4E                 ld c,(hl)
F999  72                 ld (hl),d
F99A  55                 ld d,l
F99B  55                 ld d,l
F99C  4E                 ld c,(hl)
F99D  72                 ld (hl),d
F99E  55                 ld d,l
F99F  4F                 ld c,a
F9A0  FC 50 00           call m,$0050
F9A3  00                 nop
F9A4  ED 21              ???
F9A6  00                 nop
F9A7  E0                 ret po
F9A8  06 94              ld b,-108
F9AA  EF                 rst 28h
F9AB  88                 adc a,b
F9AC  FE C8              cp -56
F9AE  07                 rlca
F9AF  00                 nop
F9B0  00                 nop
F9B1  1B                 dec de
F9B2  00                 nop
F9B3  CA 0E 0F           jp z,$0F0E
F9B6  FA 00 27           jp m,$2700
F9B9  E4 00 00           call po,$0000
F9BC  27                 daa
F9BD  E4 00 2F           call po,$2F00
F9C0  F8                 ret m
F9C1  38 29              jr c,$F9EC
F9C3  80                 add a,b
F9C4  6C                 ld l,h
F9C5  00                 nop
F9C6  00                 nop
F9C7  70                 ld (hl),b
F9C8  09                 add hl,bc
F9C9  88                 adc a,b
F9CA  FE 0A              cp 10
F9CC  5C                 ld e,h
F9CD  2E 0F              ld l,15
F9CF  A0                 and b
F9D0  80                 add a,b
F9D1  1A                 ld a,(de)
F9D2  00                 nop
F9D3  C0                 ret nz
F9D4  07                 rlca
F9D5  0F                 rrca
F9D6  F2 55 09           jp p,$0955
F9D9  90                 sub b
F9DA  55                 ld d,l
F9DB  55                 ld d,l
F9DC  09                 add hl,bc
F9DD  90                 sub b
F9DE  AA                 xor d
F9DF  27                 daa
F9E0  F8                 ret m
F9E1  70                 ld (hl),b
F9E2  01 80 2C           ld bc,$2C80
F9E5  01 02 F8           ld bc,-2046
F9E8  3A 0A 5C           ld a,($5C0A)
F9EB  98                 sbc a,b
F9EC  EE 78              xor $78
F9EE  1F                 rra
F9EF  00                 nop
F9F0  80                 add a,b
F9F1  12                 ld (de),a
F9F2  00                 nop
F9F3  00                 nop
F9F4  20 07              jr nz,$F9FD
F9F6  CC AA 0E           call z,$0EAA
F9F9  70                 ld (hl),b
F9FA  2A AA 0E           ld hl,($0EAA)
F9FD  70                 ld (hl),b
F9FE  55                 ld d,l
F9FF  19                 add hl,de
FA00  F0                 ret p
FA01  02                 ld (bc),a
FA02  00                 nop
FA03  00                 nop
FA04  24                 inc h
FA05  01 00 7C           ld bc,$7C00
FA08  0F                 rrca
FA09  98                 sbc a,b
FA0A  EE 73              xor $73
FA0C  BF                 cp a
FA0D  50                 ld d,b
FA0E  3F                 ccf
FA0F  A8                 xor b
FA10  03                 inc bc
FA11  92                 sub d
FA12  0F                 rrca
FA13  FE 04              cp 4
FA15  41                 ld b,c
FA16  F0                 ret p
FA17  FF                 rst 38h
FA18  0F                 rrca
FA19  70                 ld (hl),b
FA1A  7F                 ld a,a
FA1B  FF                 rst 38h
FA1C  0F                 rrca
FA1D  70                 ld (hl),b
FA1E  FF                 rst 38h
FA1F  07                 rlca
FA20  C1                 pop bc
FA21  10 3F              djnz $FA62
FA23  F8                 ret m
FA24  24                 inc h
FA25  E0                 ret po
FA26  0A                 ld a,(bc)
FA27  FE 05              cp 5
FA29  73                 ld (hl),e
FA2A  BF                 cp a
FA2B  23                 inc hl
FA2C  40                 ld b,b
FA2D  0E 7F              ld c,127
FA2F  40                 ld b,b
FA30  05                 dec b
FA31  42                 ld b,d
FA32  00                 nop
FA33  00                 nop
FA34  00                 nop
FA35  00                 nop
FA36  00                 nop
FA37  55                 ld d,l
FA38  00                 nop
FA39  00                 nop
FA3A  2A AA 00           ld hl,($00AA)
FA3D  00                 nop
FA3E  AA                 xor d
FA3F  00                 nop
FA40  00                 nop
FA41  00                 nop
FA42  00                 nop
FA43  00                 nop
FA44  21 50 01           ld hl,336
FA47  7F                 ld a,a
FA48  38 23              jr c,$FA6D
FA4A  40                 ld b,b
FA4B  2A 1D 63           ld hl,($631D)
FA4E  AA                 xor d
FA4F  A2                 and d
FA50  07                 rlca
FA51  C0                 ret nz
FA52  07                 rlca
FA53  FE 00              cp 0
FA55  00                 nop
FA56  00                 nop
FA57  FF                 rst 38h
FA58  00                 nop
FA59  00                 nop
FA5A  7F                 ld a,a
FA5B  FF                 rst 38h
FA5C  00                 nop
FA5D  00                 nop
FA5E  FF                 rst 38h
FA5F  00                 nop
FA60  00                 nop
FA61  00                 nop
FA62  3F                 ccf
FA63  F0                 ret p
FA64  01 F0 22           ld bc,$22F0
FA67  AA                 xor d
FA68  E3                 ex (sp),hl
FA69  2A 1D 69           ld hl,($691D)
FA6C  EE 10              xor $10
FA6E  11 00 03           ld de,768
FA71  82                 add a,d
FA72  07                 rlca
FA73  54                 ld d,h
FA74  00                 nop
FA75  00                 nop
FA76  00                 nop
FA77  FF                 rst 38h
FA78  00                 nop
FA79  00                 nop
FA7A  7F                 ld a,a
FA7B  FF                 rst 38h
FA7C  00                 nop
FA7D  00                 nop
FA7E  FF                 rst 38h
FA7F  00                 nop
FA80  00                 nop
FA81  00                 nop
FA82  15                 dec d
FA83  70                 ld (hl),b
FA84  20 E0              jr nz,$FA66
FA86  00                 nop
FA87  44                 ld b,h
FA88  04                 inc b
FA89  69                 ld l,c
FA8A  EE E2              xor $E2
FA8C  77                 ld (hl),a
FA8D  13                 inc de
FA8E  84                 add a,h
FA8F  00                 nop
FA90  06 C0              ld b,-64
FA92  0D                 dec c
FA93  00                 nop
FA94  00                 nop
FA95  00                 nop
FA96  00                 nop
FA97  AA                 xor d
FA98  00                 nop
FA99  00                 nop
FA9A  55                 ld d,l
FA9B  55                 ld d,l
FA9C  00                 nop
FA9D  00                 nop
FA9E  55                 ld d,l
FA9F  00                 nop
FAA0  00                 nop
FAA1  00                 nop
FAA2  00                 nop
FAA3  58                 ld e,b
FAA4  01 B0 00           ld bc,176
FAA7  10 E4              djnz $FA8D
FAA9  E2 77 ED           jp po,$ED77
FAAC  7F                 ld a,a
FAAD  1A                 ld a,(de)
FAAE  F0                 ret p
FAAF  78                 ld a,b
FAB0  0B                 dec bc
FAB1  D0                 ret nc
FAB2  0A                 ld a,(bc)
FAB3  6A                 ld l,d
FAB4  00                 nop
FAB5  00                 nop
FAB6  00                 nop
FAB7  FF                 rst 38h
FAB8  00                 nop
FAB9  00                 nop
FABA  7F                 ld a,a
FABB  FF                 rst 38h
FABC  00                 nop
FABD  00                 nop
FABE  FF                 rst 38h
FABF  00                 nop
FAC0  00                 nop
FAC1  00                 nop
FAC2  2B                 dec hl
FAC3  28 05              jr z,$FACA
FAC5  E8                 ret pe
FAC6  0F                 rrca
FAC7  07                 rlca
FAC8  AC                 xor h
FAC9  ED 7F              ???
FACB  AD                 xor l
FACC  AE                 xor (hl)
FACD  09                 add hl,bc
FACE  07                 rlca
FACF  BF                 cp a
FAD0  0A                 ld a,(bc)
FAD1  D8                 ret c
FAD2  1A                 ld a,(de)
FAD3  00                 nop
FAD4  00                 nop
FAD5  00                 nop
FAD6  00                 nop
FAD7  55                 ld d,l
FAD8  00                 nop
FAD9  00                 nop
FADA  2A AA 00           ld hl,($00AA)
FADD  00                 nop
FADE  AA                 xor d
FADF  00                 nop
FAE0  00                 nop
FAE1  00                 nop
FAE2  00                 nop
FAE3  2C                 inc l
FAE4  0D                 dec c
FAE5  A8                 xor b
FAE6  7E                 ld a,(hl)
FAE7  F0                 ret p
FAE8  48                 ld c,b
FAE9  AD                 xor l
FAEA  AE                 xor (hl)
FAEB  B6                 or (hl)
FAEC  80                 add a,b
FAED  18 7F              jr $FB6E
FAEF  2F                 cpl
FAF0  C7                 rst 00h
FAF1  E8                 ret pe
FAF2  34                 inc (hl)
FAF3  00                 nop
FAF4  00                 nop
FAF5  00                 nop
FAF6  00                 nop
FAF7  AA                 xor d
FAF8  00                 nop
FAF9  00                 nop
FAFA  55                 ld d,l
FAFB  55                 ld d,l
FAFC  00                 nop
FAFD  00                 nop
FAFE  55                 ld d,l
FAFF  00                 nop
FB00  00                 nop
FB01  00                 nop
FB02  00                 nop
FB03  16 0B              ld d,11
FB05  F1                 pop af
FB06  FA 7F 0C           jp m,$0C7F
FB09  B6                 or (hl)
FB0A  80                 add a,b
FB0B  2D                 dec l
FB0C  00                 nop
FB0D  03                 inc bc
FB0E  FD 25              dec iyh
FB10  B8                 cp b
FB11  EE 00              xor $00
FB13  00                 nop
FB14  00                 nop
FB15  00                 nop
FB16  00                 nop
FB17  00                 nop
FB18  00                 nop
FB19  00                 nop
FB1A  00                 nop
FB1B  00                 nop
FB1C  00                 nop
FB1D  00                 nop
FB1E  00                 nop
FB1F  00                 nop
FB20  00                 nop
FB21  00                 nop
FB22  00                 nop
FB23  00                 nop
FB24  3B                 dec sp
FB25  8E                 adc a,(hl)
FB26  D2 5F E0           jp nc,$E05F
FB29  2D                 dec l
FB2A  00                 nop
FB2B  4E                 ld c,(hl)
FB2C  80                 add a,b
FB2D  7E                 ld a,(hl)
FB2E  AA                 xor d
FB2F  00                 nop
FB30  37                 scf
FB31  67                 ld h,a
FB32  F8                 ret m
FB33  00                 nop
FB34  00                 nop
FB35  00                 nop
FB36  00                 nop
FB37  55                 ld d,l
FB38  00                 nop
FB39  00                 nop
FB3A  2A AA 00           ld hl,($00AA)
FB3D  00                 nop
FB3E  AA                 xor d
FB3F  00                 nop
FB40  00                 nop
FB41  00                 nop
FB42  00                 nop
FB43  0F                 rrca
FB44  F3                 di
FB45  76                 halt
FB46  00                 nop
FB47  2A BF 4E           ld hl,($4EBF)
FB4A  80                 add a,b
FB4B  CD 22 74           call $7422
FB4E  ED 4B 92 FB        ld bc,($FB92)
FB52  B9                 cp c
FB53  C8                 ret z
FB54  32 92 FB           ld ($FB92),a
FB57  FE 2A              cp 42
FB59  C8                 ret z
FB5A  2A 93 FB           ld hl,($FB93)
FB5D  7B                 ld a,e
FB5E  BE                 cp (hl)
FB5F  28 09              jr z,$FB6A
FB61  21 95 FB           ld hl,-1131
FB64  22 93 FB           ld ($FB93),hl
FB67  C3 B7 FB           jp $FBB7
FB6A  23                 inc hl
FB6B  22 93 FB           ld ($FB93),hl
FB6E  7E                 ld a,(hl)
FB6F  FE FF              cp -1
FB71  C2 B7 FB           jp nz,$FBB7
FB74  AF                 xor a
FB75  32 76 B3           ld ($B376),a
FB78  3E FF              ld a,-1
FB7A  32 14 B3           ld ($B314),a
FB7D  01 02 04           ld bc,1026
FB80  CD 29 B5           call $B529
FB83  21 9F FB           ld hl,-1121
FB86  CD CE B2           call $B2CE
FB89  CD 6F 74           call $746F
FB8C  CD 17 B5           call $B517
FB8F  C3 6C 78           jp $786C
FB92  2A 95 FB           ld hl,($FB95)
FB95  20 05              jr nz,$FB9C
FB97  0E 05              ld c,5
FB99  26 07              ld h,7
FB9B  20 05              jr nz,$FBA2
FB9D  0E FF              ld c,-1
FB9F  02                 ld (bc),a
FBA0  08                 ex af,af'
FBA1  0B                 dec bc
FBA2  43                 ld b,e
FBA3  48                 ld c,b
FBA4  45                 ld b,l
FBA5  41                 ld b,c
FBA6  54                 ld d,h
FBA7  20 4D              jr nz,$FBF6
FBA9  4F                 ld c,a
FBAA  44                 ld b,h
FBAB  45                 ld b,l
FBAC  01 20 45           ld bc,$4520
FBAF  4E                 ld c,(hl)
FBB0  41                 ld b,c
FBB1  42                 ld b,d
FBB2  4C                 ld c,h
FBB3  45                 ld b,l
FBB4  44                 ld b,h
FBB5  20 A1              jr nz,$FB58
FBB7  2A E8 FB           ld hl,($FBE8)
FBBA  7B                 ld a,e
FBBB  BE                 cp (hl)
FBBC  28 07              jr z,$FBC5
FBBE  21 EA FB           ld hl,-1046
FBC1  22 E8 FB           ld ($FBE8),hl
FBC4  C9                 ret
FBC5  23                 inc hl
FBC6  22 E8 FB           ld ($FBE8),hl
FBC9  7E                 ld a,(hl)
FBCA  FE FF              cp -1
FBCC  C0                 ret nz
FBCD  3E 07              ld a,7
FBCF  32 2E 78           ld ($782E),a
FBD2  01 02 04           ld bc,1026
FBD5  CD 29 B5           call $B529
FBD8  21 F2 FB           ld hl,-1038
FBDB  CD CE B2           call $B2CE
FBDE  CD 6F 74           call $746F
FBE1  CD 17 B5           call $B517
FBE4  C3 6C 78           jp $786C
FBE7  2A EA FB           ld hl,($FBEA)
FBEA  0E 0D              ld c,13
FBEC  1C                 inc e
FBED  20 0C              jr nz,$FBFB
FBEF  1D                 dec e
FBF0  06 FF              ld b,-1
FBF2  02                 ld (bc),a
FBF3  08                 ex af,af'
FBF4  0B                 dec bc
FBF5  54                 ld d,h
FBF6  52                 ld d,d
FBF7  55                 ld d,l
FBF8  45                 ld b,l
FBF9  20 4B              jr nz,$FC46
FBFB  45                 ld b,l
FBFC  59                 ld e,c
FBFD  53                 ld d,e
FBFE  01 45 4E           ld bc,$4E45
FC01  41                 ld b,c
FC02  42                 ld b,d
FC03  4C                 ld c,h
FC04  45                 ld b,l
FC05  44                 ld b,h
FC06  20 A1              jr nz,$FBA9
FC08  00                 nop
FC09  00                 nop
FC0A  00                 nop
FC0B  00                 nop
FC0C  00                 nop
FC0D  00                 nop
FC0E  00                 nop
FC0F  00                 nop
FC10  00                 nop
FC11  00                 nop
FC12  00                 nop
FC13  00                 nop
FC14  00                 nop
FC15  00                 nop
FC16  00                 nop
FC17  00                 nop
FC18  00                 nop
FC19  00                 nop
FC1A  00                 nop
FC1B  00                 nop
FC1C  00                 nop
FC1D  00                 nop
FC1E  00                 nop
FC1F  00                 nop
FC20  00                 nop
FC21  00                 nop
FC22  00                 nop
FC23  00                 nop
FC24  00                 nop
FC25  00                 nop
FC26  00                 nop
FC27  00                 nop
FC28  00                 nop
FC29  00                 nop
FC2A  00                 nop
FC2B  00                 nop
FC2C  00                 nop
FC2D  00                 nop
FC2E  00                 nop
FC2F  00                 nop
FC30  00                 nop
FC31  00                 nop
FC32  00                 nop
FC33  00                 nop
FC34  00                 nop
FC35  00                 nop
FC36  00                 nop
FC37  00                 nop
FC38  00                 nop
FC39  00                 nop
FC3A  00                 nop
FC3B  00                 nop
FC3C  00                 nop
FC3D  00                 nop
FC3E  00                 nop
FC3F  00                 nop
FC40  00                 nop
FC41  00                 nop
FC42  00                 nop
FC43  00                 nop
FC44  00                 nop
FC45  00                 nop
FC46  00                 nop
FC47  00                 nop
FC48  00                 nop
FC49  00                 nop
FC4A  00                 nop
FC4B  00                 nop
FC4C  00                 nop
FC4D  00                 nop
FC4E  00                 nop
FC4F  00                 nop
FC50  00                 nop
FC51  00                 nop
FC52  00                 nop
FC53  00                 nop
FC54  00                 nop
FC55  00                 nop
FC56  00                 nop
FC57  00                 nop
FC58  00                 nop
FC59  00                 nop
FC5A  00                 nop
FC5B  00                 nop
FC5C  00                 nop
FC5D  00                 nop
FC5E  00                 nop
FC5F  00                 nop
FC60  00                 nop
FC61  00                 nop
FC62  00                 nop
FC63  00                 nop
FC64  00                 nop
FC65  00                 nop
FC66  00                 nop
FC67  00                 nop
FC68  00                 nop
FC69  00                 nop
FC6A  00                 nop
FC6B  00                 nop
FC6C  00                 nop
FC6D  00                 nop
FC6E  00                 nop
FC6F  00                 nop
FC70  00                 nop
FC71  00                 nop
FC72  00                 nop
FC73  00                 nop
FC74  00                 nop
FC75  00                 nop
FC76  00                 nop
FC77  00                 nop
FC78  00                 nop
FC79  00                 nop
FC7A  00                 nop
FC7B  00                 nop
FC7C  00                 nop
FC7D  00                 nop
FC7E  00                 nop
FC7F  00                 nop
FC80  00                 nop
FC81  00                 nop
FC82  00                 nop
FC83  00                 nop
FC84  00                 nop
FC85  00                 nop
FC86  00                 nop
FC87  00                 nop
FC88  00                 nop
FC89  00                 nop
FC8A  00                 nop
FC8B  00                 nop
FC8C  00                 nop
FC8D  00                 nop
FC8E  00                 nop
FC8F  00                 nop
FC90  00                 nop
FC91  00                 nop
FC92  00                 nop
FC93  00                 nop
FC94  00                 nop
FC95  00                 nop
FC96  00                 nop
FC97  00                 nop
FC98  00                 nop
FC99  00                 nop
FC9A  00                 nop
FC9B  00                 nop
FC9C  00                 nop
FC9D  00                 nop
FC9E  00                 nop
FC9F  00                 nop
FCA0  00                 nop
FCA1  00                 nop
FCA2  00                 nop
FCA3  00                 nop
FCA4  00                 nop
FCA5  00                 nop
FCA6  00                 nop
FCA7  00                 nop
FCA8  00                 nop
FCA9  00                 nop
FCAA  00                 nop
FCAB  00                 nop
FCAC  00                 nop
FCAD  00                 nop
FCAE  00                 nop
FCAF  00                 nop
FCB0  00                 nop
FCB1  00                 nop
FCB2  00                 nop
FCB3  00                 nop
FCB4  00                 nop
FCB5  00                 nop
FCB6  00                 nop
FCB7  00                 nop
FCB8  00                 nop
FCB9  00                 nop
FCBA  00                 nop
FCBB  00                 nop
FCBC  00                 nop
FCBD  00                 nop
FCBE  00                 nop
FCBF  00                 nop
FCC0  00                 nop
FCC1  00                 nop
FCC2  00                 nop
FCC3  00                 nop
FCC4  00                 nop
FCC5  00                 nop
FCC6  00                 nop
FCC7  00                 nop
FCC8  00                 nop
FCC9  00                 nop
FCCA  00                 nop
FCCB  00                 nop
FCCC  00                 nop
FCCD  00                 nop
FCCE  00                 nop
FCCF  00                 nop
FCD0  00                 nop
FCD1  00                 nop
FCD2  00                 nop
FCD3  00                 nop
FCD4  00                 nop
FCD5  00                 nop
FCD6  00                 nop
FCD7  00                 nop
FCD8  00                 nop
FCD9  00                 nop
FCDA  00                 nop
FCDB  00                 nop
FCDC  00                 nop
FCDD  00                 nop
FCDE  00                 nop
FCDF  00                 nop
FCE0  00                 nop
FCE1  00                 nop
FCE2  00                 nop
FCE3  00                 nop
FCE4  00                 nop
FCE5  00                 nop
FCE6  00                 nop
FCE7  00                 nop
FCE8  00                 nop
FCE9  00                 nop
FCEA  00                 nop
FCEB  00                 nop
FCEC  00                 nop
FCED  00                 nop
FCEE  00                 nop
FCEF  00                 nop
FCF0  00                 nop
FCF1  00                 nop
FCF2  00                 nop
FCF3  00                 nop
FCF4  00                 nop
FCF5  00                 nop
FCF6  00                 nop
FCF7  00                 nop
FCF8  00                 nop
FCF9  00                 nop
FCFA  00                 nop
FCFB  00                 nop
FCFC  00                 nop
FCFD  00                 nop
FCFE  00                 nop
FCFF  00                 nop
FD00  00                 nop
FD01  00                 nop
FD02  00                 nop
FD03  00                 nop
FD04  00                 nop
FD05  00                 nop
FD06  00                 nop
FD07  00                 nop
FD08  00                 nop
FD09  00                 nop
FD0A  00                 nop
FD0B  00                 nop
FD0C  00                 nop
FD0D  00                 nop
FD0E  00                 nop
FD0F  00                 nop
FD10  00                 nop
FD11  00                 nop
FD12  00                 nop
FD13  00                 nop
FD14  00                 nop
FD15  00                 nop
FD16  00                 nop
FD17  00                 nop
FD18  00                 nop
FD19  00                 nop
FD1A  00                 nop
FD1B  00                 nop
FD1C  00                 nop
FD1D  00                 nop
FD1E  00                 nop
FD1F  00                 nop
FD20  00                 nop
FD21  00                 nop
FD22  00                 nop
FD23  00                 nop
FD24  00                 nop
FD25  00                 nop
FD26  00                 nop
FD27  00                 nop
FD28  00                 nop
FD29  00                 nop
FD2A  00                 nop
FD2B  00                 nop
FD2C  00                 nop
FD2D  00                 nop
FD2E  00                 nop
FD2F  00                 nop
FD30  00                 nop
FD31  00                 nop
FD32  00                 nop
FD33  00                 nop
FD34  00                 nop
FD35  00                 nop
FD36  00                 nop
FD37  00                 nop
FD38  00                 nop
FD39  00                 nop
FD3A  00                 nop
FD3B  00                 nop
FD3C  00                 nop
FD3D  00                 nop
FD3E  00                 nop
FD3F  00                 nop
FD40  00                 nop
FD41  00                 nop
FD42  00                 nop
FD43  00                 nop
FD44  00                 nop
FD45  00                 nop
FD46  00                 nop
FD47  00                 nop
FD48  00                 nop
FD49  00                 nop
FD4A  00                 nop
FD4B  00                 nop
FD4C  00                 nop
FD4D  00                 nop
FD4E  00                 nop
FD4F  00                 nop
FD50  00                 nop
FD51  00                 nop
FD52  00                 nop
FD53  00                 nop
FD54  00                 nop
FD55  00                 nop
FD56  00                 nop
FD57  00                 nop
FD58  00                 nop
FD59  00                 nop
FD5A  00                 nop
FD5B  00                 nop
FD5C  00                 nop
FD5D  00                 nop
FD5E  00                 nop
FD5F  00                 nop
FD60  00                 nop
FD61  00                 nop
FD62  00                 nop
FD63  00                 nop
FD64  00                 nop
FD65  00                 nop
FD66  00                 nop
FD67  00                 nop
FD68  00                 nop
FD69  00                 nop
FD6A  00                 nop
FD6B  00                 nop
FD6C  00                 nop
FD6D  00                 nop
FD6E  00                 nop
FD6F  00                 nop
FD70  00                 nop
FD71  00                 nop
FD72  00                 nop
FD73  00                 nop
FD74  00                 nop
FD75  00                 nop
FD76  00                 nop
FD77  00                 nop
FD78  00                 nop
FD79  00                 nop
FD7A  00                 nop
FD7B  00                 nop
FD7C  00                 nop
FD7D  00                 nop
FD7E  00                 nop
FD7F  00                 nop
FD80  00                 nop
FD81  00                 nop
FD82  00                 nop
FD83  00                 nop
FD84  00                 nop
FD85  00                 nop
FD86  00                 nop
FD87  00                 nop
FD88  00                 nop
FD89  00                 nop
FD8A  00                 nop
FD8B  00                 nop
FD8C  00                 nop
FD8D  00                 nop
FD8E  00                 nop
FD8F  00                 nop
FD90  00                 nop
FD91  00                 nop
FD92  00                 nop
FD93  00                 nop
FD94  00                 nop
FD95  00                 nop
FD96  00                 nop
FD97  00                 nop
FD98  00                 nop
FD99  00                 nop
FD9A  00                 nop
FD9B  00                 nop
FD9C  00                 nop
FD9D  00                 nop
FD9E  00                 nop
FD9F  00                 nop
FDA0  00                 nop
FDA1  00                 nop
FDA2  00                 nop
FDA3  00                 nop
FDA4  00                 nop
FDA5  00                 nop
FDA6  00                 nop
FDA7  00                 nop
FDA8  00                 nop
FDA9  00                 nop
FDAA  00                 nop
FDAB  00                 nop
FDAC  00                 nop
FDAD  00                 nop
FDAE  00                 nop
FDAF  00                 nop
FDB0  00                 nop
FDB1  00                 nop
FDB2  00                 nop
FDB3  00                 nop
FDB4  00                 nop
FDB5  00                 nop
FDB6  00                 nop
FDB7  00                 nop
FDB8  00                 nop
FDB9  00                 nop
FDBA  00                 nop
FDBB  00                 nop
FDBC  00                 nop
FDBD  00                 nop
FDBE  00                 nop
FDBF  00                 nop
FDC0  00                 nop
FDC1  00                 nop
FDC2  00                 nop
FDC3  00                 nop
FDC4  00                 nop
FDC5  00                 nop
FDC6  00                 nop
FDC7  00                 nop
FDC8  00                 nop
FDC9  00                 nop
FDCA  00                 nop
FDCB  00                 nop
FDCC  00                 nop
FDCD  00                 nop
FDCE  00                 nop
FDCF  00                 nop
FDD0  00                 nop
FDD1  00                 nop
FDD2  00                 nop
FDD3  00                 nop
FDD4  00                 nop
FDD5  00                 nop
FDD6  00                 nop
FDD7  00                 nop
FDD8  00                 nop
FDD9  00                 nop
FDDA  00                 nop
FDDB  00                 nop
FDDC  00                 nop
FDDD  00                 nop
FDDE  00                 nop
FDDF  00                 nop
FDE0  00                 nop
FDE1  00                 nop
FDE2  00                 nop
FDE3  00                 nop
FDE4  00                 nop
FDE5  00                 nop
FDE6  00                 nop
FDE7  00                 nop
FDE8  00                 nop
FDE9  00                 nop
FDEA  00                 nop
FDEB  00                 nop
FDEC  00                 nop
FDED  00                 nop
FDEE  00                 nop
FDEF  00                 nop
FDF0  00                 nop
FDF1  00                 nop
FDF2  00                 nop
FDF3  00                 nop
FDF4  00                 nop
FDF5  00                 nop
FDF6  00                 nop
FDF7  00                 nop
FDF8  00                 nop
FDF9  00                 nop
FDFA  00                 nop
FDFB  00                 nop
FDFC  00                 nop
FDFD  00                 nop
FDFE  00                 nop
FDFF  00                 nop
FE00  00                 nop
FE01  00                 nop
FE02  00                 nop
FE03  00                 nop
FE04  00                 nop
FE05  00                 nop
FE06  00                 nop
FE07  00                 nop
FE08  00                 nop
FE09  00                 nop
FE0A  00                 nop
FE0B  00                 nop
FE0C  00                 nop
FE0D  00                 nop
FE0E  00                 nop
FE0F  00                 nop
FE10  00                 nop
FE11  00                 nop
FE12  00                 nop
FE13  00                 nop
FE14  00                 nop
FE15  00                 nop
FE16  00                 nop
FE17  00                 nop
FE18  00                 nop
FE19  00                 nop
FE1A  00                 nop
FE1B  00                 nop
FE1C  00                 nop
FE1D  00                 nop
FE1E  00                 nop
FE1F  00                 nop
FE20  00                 nop
FE21  00                 nop
FE22  00                 nop
FE23  00                 nop
FE24  00                 nop
FE25  00                 nop
FE26  00                 nop
FE27  00                 nop
FE28  00                 nop
FE29  00                 nop
FE2A  00                 nop
FE2B  00                 nop
FE2C  00                 nop
FE2D  00                 nop
FE2E  00                 nop
FE2F  00                 nop
FE30  00                 nop
FE31  00                 nop
FE32  00                 nop
FE33  00                 nop
FE34  00                 nop
FE35  00                 nop
FE36  00                 nop
FE37  00                 nop
FE38  00                 nop
FE39  00                 nop
FE3A  00                 nop
FE3B  00                 nop
FE3C  00                 nop
FE3D  00                 nop
FE3E  00                 nop
FE3F  00                 nop
FE40  00                 nop
FE41  00                 nop
FE42  00                 nop
FE43  00                 nop
FE44  00                 nop
FE45  00                 nop
FE46  00                 nop
FE47  00                 nop
FE48  00                 nop
FE49  00                 nop
FE4A  00                 nop
FE4B  00                 nop
FE4C  00                 nop
FE4D  00                 nop
FE4E  00                 nop
FE4F  00                 nop
FE50  00                 nop
FE51  00                 nop
FE52  00                 nop
FE53  00                 nop
FE54  00                 nop
FE55  00                 nop
FE56  00                 nop
FE57  00                 nop
FE58  00                 nop
FE59  00                 nop
FE5A  00                 nop
FE5B  00                 nop
FE5C  00                 nop
FE5D  00                 nop
FE5E  00                 nop
FE5F  00                 nop
FE60  00                 nop
FE61  00                 nop
FE62  00                 nop
FE63  00                 nop
FE64  00                 nop
FE65  00                 nop
FE66  00                 nop
FE67  00                 nop
FE68  00                 nop
FE69  00                 nop
FE6A  00                 nop
FE6B  00                 nop
FE6C  00                 nop
FE6D  00                 nop
FE6E  00                 nop
FE6F  00                 nop
FE70  00                 nop
FE71  00                 nop
FE72  00                 nop
FE73  00                 nop
FE74  00                 nop
FE75  00                 nop
FE76  00                 nop
FE77  00                 nop
FE78  00                 nop
FE79  00                 nop
FE7A  00                 nop
FE7B  00                 nop
FE7C  00                 nop
FE7D  00                 nop
FE7E  00                 nop
FE7F  00                 nop
FE80  00                 nop
FE81  00                 nop
FE82  00                 nop
FE83  00                 nop
FE84  00                 nop
FE85  00                 nop
FE86  00                 nop
FE87  00                 nop
FE88  00                 nop
FE89  00                 nop
FE8A  00                 nop
FE8B  00                 nop
FE8C  00                 nop
FE8D  00                 nop
FE8E  00                 nop
FE8F  00                 nop
FE90  00                 nop
FE91  00                 nop
FE92  00                 nop
FE93  00                 nop
FE94  00                 nop
FE95  00                 nop
FE96  00                 nop
FE97  00                 nop
FE98  00                 nop
FE99  00                 nop
FE9A  00                 nop
FE9B  00                 nop
FE9C  00                 nop
FE9D  00                 nop
FE9E  00                 nop
FE9F  00                 nop
FEA0  00                 nop
FEA1  00                 nop
FEA2  00                 nop
FEA3  00                 nop
FEA4  00                 nop
FEA5  00                 nop
FEA6  00                 nop
FEA7  00                 nop
FEA8  00                 nop
FEA9  00                 nop
FEAA  00                 nop
FEAB  00                 nop
FEAC  00                 nop
FEAD  00                 nop
FEAE  00                 nop
FEAF  00                 nop
FEB0  00                 nop
FEB1  00                 nop
FEB2  00                 nop
FEB3  00                 nop
FEB4  00                 nop
FEB5  00                 nop
FEB6  00                 nop
FEB7  00                 nop
FEB8  00                 nop
FEB9  00                 nop
FEBA  00                 nop
FEBB  00                 nop
FEBC  00                 nop
FEBD  00                 nop
FEBE  00                 nop
FEBF  00                 nop
FEC0  00                 nop
FEC1  00                 nop
FEC2  00                 nop
FEC3  00                 nop
FEC4  00                 nop
FEC5  00                 nop
FEC6  00                 nop
FEC7  00                 nop
FEC8  00                 nop
FEC9  00                 nop
FECA  00                 nop
FECB  00                 nop
FECC  00                 nop
FECD  00                 nop
FECE  00                 nop
FECF  00                 nop
FED0  00                 nop
FED1  00                 nop
FED2  00                 nop
FED3  00                 nop
FED4  00                 nop
FED5  00                 nop
FED6  00                 nop
FED7  00                 nop
FED8  00                 nop
FED9  00                 nop
FEDA  00                 nop
FEDB  00                 nop
FEDC  00                 nop
FEDD  00                 nop
FEDE  00                 nop
FEDF  00                 nop
FEE0  00                 nop
FEE1  00                 nop
FEE2  00                 nop
FEE3  00                 nop
FEE4  00                 nop
FEE5  00                 nop
FEE6  00                 nop
FEE7  00                 nop
FEE8  00                 nop
FEE9  00                 nop
FEEA  00                 nop
FEEB  00                 nop
FEEC  00                 nop
FEED  00                 nop
FEEE  00                 nop
FEEF  00                 nop
FEF0  00                 nop
FEF1  00                 nop
FEF2  00                 nop
FEF3  00                 nop
FEF4  00                 nop
FEF5  00                 nop
FEF6  00                 nop
FEF7  00                 nop
FEF8  00                 nop
FEF9  00                 nop
FEFA  00                 nop
FEFB  00                 nop
FEFC  00                 nop
FEFD  00                 nop
FEFE  00                 nop
FEFF  00                 nop
FF00  00                 nop
FF01  00                 nop
FF02  0A                 ld a,(bc)
FF03  1B                 dec de
FF04  20 2D              jr nz,$FF33
FF06  2E 2E              ld l,46
FF08  2E 17              ld l,23
FF0A  2C                 inc l
FF0B  1A                 ld a,(de)
FF0C  18 2D              jr $FF3B
FF0E  1A                 ld a,(de)
FF0F  0B                 dec bc
FF10  0D                 dec c
FF11  0E 11              ld c,17
FF13  23                 inc hl
FF14  14                 inc d
FF15  21 2F 29           ld hl,$292F
FF18  26 31              ld h,49
FF1A  25                 dec h
FF1B  13                 inc de
FF1C  19                 add hl,de
FF1D  00                 nop
FF1E  00                 nop
FF1F  1B                 dec de
FF20  1D                 dec e
FF21  1C                 inc e
FF22  7F                 ld a,a
FF23  7F                 ld a,a
FF24  00                 nop
FF25  04                 inc b
FF26  07                 rlca
FF27  00                 nop
FF28  00                 nop
FF29  00                 nop
FF2A  06 02              ld b,2
FF2C  00                 nop
FF2D  00                 nop
FF2E  00                 nop
FF2F  06 06              ld b,6
FF31  06 06              ld b,6
FF33  02                 ld (bc),a
FF34  00                 nop
FF35  00                 nop
FF36  03                 inc bc
FF37  03                 inc bc
FF38  03                 inc bc
FF39  02                 ld (bc),a
FF3A  03                 inc bc
FF3B  03                 inc bc
FF3C  03                 inc bc
FF3D  02                 ld (bc),a
FF3E  03                 inc bc
FF3F  03                 inc bc
FF40  03                 inc bc
FF41  00                 nop
FF42  00                 nop
FF43  00                 nop
FF44  03                 inc bc
FF45  00                 nop
FF46  00                 nop
FF47  00                 nop
FF48  00                 nop
FF49  00                 nop
FF4A  00                 nop
FF4B  00                 nop
FF4C  00                 nop
FF4D  08                 ex af,af'
FF4E  0A                 ld a,(bc)
FF4F  1B                 dec de
FF50  20 2D              jr nz,$FF7F
FF52  2E 2E              ld l,46
FF54  2E 17              ld l,23
FF56  2C                 inc l
FF57  1A                 ld a,(de)
FF58  18 2D              jr $FF87
FF5A  1A                 ld a,(de)
FF5B  0C                 inc c
FF5C  0D                 dec c
FF5D  0E 11              ld c,17
FF5F  24                 inc h
FF60  16 21              ld d,33
FF62  2F                 cpl
FF63  29                 add hl,hl
FF64  26 32              ld h,50
FF66  25                 dec h
FF67  13                 inc de
FF68  19                 add hl,de
FF69  08                 ex af,af'
FF6A  08                 ex af,af'
FF6B  1B                 dec de
FF6C  1D                 dec e
FF6D  1C                 inc e
FF6E  7F                 ld a,a
FF6F  7F                 ld a,a
FF70  00                 nop
FF71  06 08              ld b,8
FF73  04                 inc b
FF74  0A                 ld a,(bc)
FF75  1B                 dec de
FF76  20 2D              jr nz,$FFA5
FF78  2E 2E              ld l,46
FF7A  2E 17              ld l,23
FF7C  2C                 inc l
FF7D  1A                 ld a,(de)
FF7E  18 2D              jr $FFAD
FF80  1A                 ld a,(de)
FF81  0B                 dec bc
FF82  0D                 dec c
FF83  0E 11              ld c,17
FF85  23                 inc hl
FF86  14                 inc d
FF87  21 31 29           ld hl,$2931
FF8A  28 32              jr z,$FFBE
FF8C  25                 dec h
FF8D  13                 inc de
FF8E  19                 add hl,de
FF8F  04                 inc b
FF90  04                 inc b
FF91  1B                 dec de
FF92  1D                 dec e
FF93  1C                 inc e
FF94  7F                 ld a,a
FF95  7F                 ld a,a
FF96  00                 nop
FF97  05                 dec b
FF98  08                 ex af,af'
FF99  04                 inc b
FF9A  FF                 rst 38h
FF9B  FF                 rst 38h
FF9C  FF                 rst 38h
FF9D  FF                 rst 38h
FF9E  FF                 rst 38h
FF9F  FF                 rst 38h
FFA0  FF                 rst 38h
FFA1  FF                 rst 38h
FFA2  FF                 rst 38h
FFA3  FF                 rst 38h
FFA4  FF                 rst 38h
FFA5  FF                 rst 38h
FFA6  FF                 rst 38h
FFA7  FF                 rst 38h
FFA8  FF                 rst 38h
FFA9  FF                 rst 38h
FFAA  FF                 rst 38h
FFAB  FF                 rst 38h
FFAC  08                 ex af,af'
FFAD  FF                 rst 38h
FFAE  0A                 ld a,(bc)
FFAF  FF                 rst 38h
FFB0  09                 add hl,bc
FFB1  FF                 rst 38h
FFB2  0A                 ld a,(bc)
FFB3  0B                 dec bc
FFB4  FF                 rst 38h
FFB5  04                 inc b
FFB6  04                 inc b
FFB7  FF                 rst 38h
FFB8  FF                 rst 38h
FFB9  FF                 rst 38h
FFBA  0D                 dec c
FFBB  FF                 rst 38h
FFBC  0C                 inc c
FFBD  FF                 rst 38h
FFBE  0A                 ld a,(bc)
FFBF  08                 ex af,af'
FFC0  00                 nop
FFC1  00                 nop
FFC2  00                 nop
FFC3  00                 nop
FFC4  00                 nop
FFC5  00                 nop
FFC6  00                 nop
FFC7  00                 nop
FFC8  00                 nop
FFC9  00                 nop
FFCA  00                 nop
FFCB  01 01 01           ld bc,257
FFCE  02                 ld (bc),a
FFCF  01 00 01           ld bc,256
FFD2  01 03 01           ld bc,259
FFD5  03                 inc bc
FFD6  01 01 02           ld bc,513
FFD9  01 01 08           ld bc,$0801
FFDC  08                 ex af,af'
FFDD  00                 nop
FFDE  00                 nop
FFDF  00                 nop
FFE0  00                 nop
FFE1  00                 nop
FFE2  18 18              jr $FFFC
FFE4  18 32              jr $10018
FFE6  96                 sub (hl)
FFE7  FF                 rst 38h
FFE8  28 8C              jr z,$FF76
FFEA  00                 nop
FFEB  14                 inc d
FFEC  50                 ld d,b
FFED  01 00 00           ld bc,0
FFF0  00                 nop
FFF1  00                 nop
FFF2  00                 nop
FFF3  00                 nop
FFF4  25                 dec h
FFF5  7E                 ld a,(hl)
FFF6  87                 add a,a
FFF7  7E                 ld a,(hl)
FFF8  EB                 ex de,hl
FFF9  7E                 ld a,(hl)
FFFA  62                 ld h,d
FFFB  7F                 ld a,a
FFFC  59                 ld e,c
FFFD  A3                 and e
FFFE  13                 inc de
```
