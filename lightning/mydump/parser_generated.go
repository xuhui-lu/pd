// Code generated by ragel DO NOT EDIT.

//.... lightning/mydump/parser.rl:1
// Please edit `parser.rl` if you want to modify this file. To generate
// `parser_generated.go`, please execute
//
// ```sh
// make data_parsers
// ```

package mydump

import (
	"io"

	"github.com/pingcap/tidb-lightning/lightning/common"
	"github.com/pkg/errors"
)


//.... lightning/mydump/parser.rl:76



//.... tmp_parser.go:25
const chunk_parser_start int = 21
const chunk_parser_first_final int = 21
const chunk_parser_error int = 0

const chunk_parser_en_main int = 21


//.... lightning/mydump/parser.rl:79

func (parser *ChunkParser) lex() (token, []byte, error) {
	var cs, ts, te, act, p int
	
//.... tmp_parser.go:38
	{
	cs = chunk_parser_start
	ts = 0
	te = 0
	act = 0
	}

//.... lightning/mydump/parser.rl:83

	for {
		data := parser.buf
		consumedToken := tokNil
		pe := len(data)
		eof := -1
		if parser.isLastChunk {
			eof = pe
		}

		
//.... tmp_parser.go:58
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 0:
		goto st_case_0
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 25:
		goto st_case_25
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 29:
		goto st_case_29
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	}
	goto st_out
tr0:
//.... NONE:1
	switch act {
	case 0:
	{{goto st0 }}
	case 2:
	{p = (te) - 1

		consumedToken = tokValues
		{p++; cs = 21; goto _out }
	}
	case 4:
	{p = (te) - 1

		consumedToken = tokName
		{p++; cs = 21; goto _out }
	}
	default:
	{p = (te) - 1
}
	}
	
	goto st21
tr8:
//.... lightning/mydump/parser.rl:64
te = p+1
{
		consumedToken = tokRow
		{p++; cs = 21; goto _out }
	}
	goto st21
tr12:
//.... lightning/mydump/parser.rl:69
p = (te) - 1
{
		consumedToken = tokName
		{p++; cs = 21; goto _out }
	}
	goto st21
tr14:
//.... lightning/mydump/parser.rl:57
te = p+1

	goto st21
tr34:
//.... lightning/mydump/parser.rl:69
te = p
p--
{
		consumedToken = tokName
		{p++; cs = 21; goto _out }
	}
	goto st21
tr35:
//.... lightning/mydump/parser.rl:57
te = p
p--

	goto st21
	st21:
//.... NONE:1
ts = 0

//.... NONE:1
act = 0

		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//.... NONE:1
ts = p

//.... tmp_parser.go:221
		switch data[p] {
		case 32:
			goto tr14
		case 34:
			goto st1
		case 40:
			goto st4
		case 44:
			goto tr14
		case 45:
			goto tr30
		case 47:
			goto tr31
		case 59:
			goto tr14
		case 73:
			goto tr32
		case 86:
			goto tr33
		case 96:
			goto st3
		case 105:
			goto tr32
		case 118:
			goto tr33
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto st0
			}
		case data[p] >= 9:
			goto tr14
		}
		goto tr2
tr2:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st22
tr37:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:57
act = 1;
	goto st22
tr47:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:59
act = 2;
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//.... tmp_parser.go:283
		switch data[p] {
		case 32:
			goto tr0
		case 34:
			goto st1
		case 44:
			goto tr0
		case 59:
			goto tr0
		case 96:
			goto st3
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr0
			}
		case data[p] >= 9:
			goto tr0
		}
		goto tr2
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 34:
			goto tr2
		case 92:
			goto st2
		}
		goto st1
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		goto st1
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 96 {
			goto tr2
		}
		goto st3
st_case_0:
	st0:
		cs = 0
		goto _out
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 34:
			goto st5
		case 39:
			goto st7
		case 41:
			goto tr8
		case 96:
			goto st9
		}
		goto st4
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 34:
			goto st4
		case 92:
			goto st6
		}
		goto st5
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st5
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 39:
			goto st4
		case 92:
			goto st8
		}
		goto st7
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		goto st7
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 96 {
			goto st4
		}
		goto st9
tr30:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//.... tmp_parser.go:409
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 45:
			goto tr17
		case 59:
			goto tr34
		case 96:
			goto st3
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr17:
//.... NONE:1
te = p+1

	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//.... tmp_parser.go:443
		switch data[p] {
		case 10:
			goto tr14
		case 32:
			goto st10
		case 34:
			goto st11
		case 44:
			goto st10
		case 59:
			goto st10
		case 96:
			goto st13
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto st10
			}
		case data[p] >= 9:
			goto st10
		}
		goto tr17
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 10 {
			goto tr14
		}
		goto st10
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 10:
			goto tr16
		case 34:
			goto tr17
		case 92:
			goto st12
		}
		goto st11
tr16:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:57
act = 1;
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//.... tmp_parser.go:502
		switch data[p] {
		case 34:
			goto tr2
		case 92:
			goto st2
		}
		goto st1
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 10 {
			goto tr16
		}
		goto st11
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 10:
			goto tr20
		case 96:
			goto tr17
		}
		goto st13
tr20:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:57
act = 1;
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//.... tmp_parser.go:543
		if data[p] == 96 {
			goto tr2
		}
		goto st3
tr31:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//.... tmp_parser.go:560
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 42:
			goto tr24
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 96:
			goto st3
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr24:
//.... NONE:1
te = p+1

	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//.... tmp_parser.go:594
		switch data[p] {
		case 32:
			goto st14
		case 34:
			goto st16
		case 42:
			goto tr36
		case 44:
			goto st14
		case 59:
			goto st14
		case 96:
			goto st19
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto st14
			}
		case data[p] >= 9:
			goto st14
		}
		goto tr24
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 42 {
			goto st15
		}
		goto st14
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 42:
			goto st15
		case 47:
			goto tr14
		}
		goto st14
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 34:
			goto tr24
		case 42:
			goto st17
		case 92:
			goto st18
		}
		goto st16
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 34:
			goto tr24
		case 42:
			goto st17
		case 47:
			goto tr16
		case 92:
			goto st18
		}
		goto st16
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 42 {
			goto st17
		}
		goto st16
tr36:
//.... NONE:1
te = p+1

	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//.... tmp_parser.go:688
		switch data[p] {
		case 32:
			goto st14
		case 34:
			goto st16
		case 42:
			goto tr36
		case 44:
			goto st14
		case 47:
			goto tr37
		case 59:
			goto st14
		case 96:
			goto st19
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto st14
			}
		case data[p] >= 9:
			goto st14
		}
		goto tr24
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 42:
			goto st20
		case 96:
			goto tr24
		}
		goto st19
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 42:
			goto st20
		case 47:
			goto tr20
		case 96:
			goto tr24
		}
		goto st19
tr32:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//.... tmp_parser.go:752
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 78:
			goto tr38
		case 96:
			goto st3
		case 110:
			goto tr38
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr38:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//.... tmp_parser.go:790
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 83:
			goto tr39
		case 84:
			goto tr40
		case 96:
			goto st3
		case 115:
			goto tr39
		case 116:
			goto tr40
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr39:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//.... tmp_parser.go:832
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 69:
			goto tr41
		case 96:
			goto st3
		case 101:
			goto tr41
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr41:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//.... tmp_parser.go:870
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 82:
			goto tr42
		case 96:
			goto st3
		case 114:
			goto tr42
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr42:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//.... tmp_parser.go:908
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 84:
			goto tr37
		case 96:
			goto st3
		case 116:
			goto tr37
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr40:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//.... tmp_parser.go:946
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 79:
			goto tr37
		case 96:
			goto st3
		case 111:
			goto tr37
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr33:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//.... tmp_parser.go:984
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 65:
			goto tr43
		case 96:
			goto st3
		case 97:
			goto tr43
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr43:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//.... tmp_parser.go:1022
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 76:
			goto tr44
		case 96:
			goto st3
		case 108:
			goto tr44
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr44:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//.... tmp_parser.go:1060
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 85:
			goto tr45
		case 96:
			goto st3
		case 117:
			goto tr45
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr45:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//.... tmp_parser.go:1098
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 69:
			goto tr46
		case 96:
			goto st3
		case 101:
			goto tr46
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
tr46:
//.... NONE:1
te = p+1

//.... lightning/mydump/parser.rl:69
act = 4;
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//.... tmp_parser.go:1136
		switch data[p] {
		case 32:
			goto tr34
		case 34:
			goto st1
		case 44:
			goto tr34
		case 59:
			goto tr34
		case 83:
			goto tr47
		case 96:
			goto st3
		case 115:
			goto tr47
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto tr34
			}
		case data[p] >= 9:
			goto tr34
		}
		goto tr2
	st_out:
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 22:
			goto tr0
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		case 23:
			goto tr34
		case 24:
			goto tr34
		case 10:
			goto tr12
		case 11:
			goto tr12
		case 25:
			goto tr35
		case 12:
			goto tr12
		case 13:
			goto tr12
		case 26:
			goto tr35
		case 27:
			goto tr34
		case 28:
			goto tr34
		case 14:
			goto tr12
		case 15:
			goto tr12
		case 16:
			goto tr12
		case 17:
			goto tr12
		case 18:
			goto tr12
		case 29:
			goto tr34
		case 19:
			goto tr12
		case 20:
			goto tr12
		case 30:
			goto tr34
		case 31:
			goto tr34
		case 32:
			goto tr34
		case 33:
			goto tr34
		case 34:
			goto tr34
		case 35:
			goto tr34
		case 36:
			goto tr34
		case 37:
			goto tr34
		case 38:
			goto tr34
		case 39:
			goto tr34
		case 40:
			goto tr34
		}
	}

	_out: {}
	}

//.... lightning/mydump/parser.rl:94

		if cs == 0 {
			common.AppLogger.Errorf("Syntax error near byte %d, content is «%s»", parser.pos, string(data))
			return tokNil, nil, errors.New("Syntax error")
		}

		if consumedToken != tokNil {
			result := data[ts:te]
			parser.buf = data[te:]
			parser.pos += int64(te)
			return consumedToken, result, nil
		}

		if parser.isLastChunk {
			return tokNil, nil, io.EOF
		}

		parser.buf = parser.buf[ts:]
		parser.pos += int64(ts)
		p -= ts
		te -= ts
		ts = 0
		if err := parser.readBlock(); err != nil {
			return tokNil, nil, errors.Trace(err)
		}
	}

	return tokNil, nil, nil
}