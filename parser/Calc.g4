grammar Calc;

@parser::header {
	import "os"
	import "time"

	var Result Interval

	func Atoi(location, v string) int {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Could not parse %s [%s]: %v\n", location, v, err)
			os.Exit(1)
		}
		return i
	}
	func ParseDuration(v string) time.Duration {
		d, err := time.ParseDuration(v)
		if err != nil {
			fmt.Printf("Could not parse duration [%s]: %v\n", v, err)
			os.Exit(1)
		}
		return d
	}
}

///////////////////////////////////////////////////////////////////////////////
// Lexer
MUL: '*';
ZONE: 'Z'[1-5][abc]?;
DURSUFFIX: [smh];
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

///////////////////////////////////////////////////////////////////////////////
// Parser
start:
	e=expression EOF { Result = localctx.GetE().GetRet() };

expression returns [Interval ret]:
	r=repeat {
		$ret = localctx.GetR().GetRet()
	} |
	a=atom (',' e=expression)* {
		all := localctx.AllExpression()
		if len(all) == 0 {
			$ret = localctx.GetA().GetRet()
		} else {
			u := &UnionInterval{}
			u.interval = []Interval{localctx.GetA().GetRet()}
			for _, e := range localctx.AllExpression() {
				u.interval = append(u.interval, e.GetRet())
			}
			$ret = u
		}
	};

repeat returns [*RepeatInterval ret]:
	n=NUMBER '*' '(' e=expression ')' {
		$ret = &RepeatInterval{
			count: Atoi("count", $n.GetText()),
			interval: localctx.GetE().GetRet(),
		}
	};

atom returns [Interval ret]:
	{var pow Power}
	(d=dur | i=dist | n=NUMBER) (z=zone | r=range) {
		if localctx.GetZ() != nil {
			pow = localctx.GetZ().GetRet()
		} else if localctx.GetR() != nil {
			pow = localctx.GetR().GetRet()
		}
		if localctx.GetD() != nil {
			$ret = &DurationInterval{
				dur: localctx.GetD().GetRet(),
				pow: pow,
			}
		} else if localctx.GetI() != nil {
			$ret = &DistanceInterval{
				dist: localctx.GetI().GetRet().Distance,
				unit: localctx.GetI().GetRet().Unit,
				pow: pow,
			}
		} else if localctx.GetN() != nil {
			$ret = &DistanceInterval{
				dist: Atoi("constant", $n.GetText()),
				unit: "yards",
				pow: pow,
			}
		}
	};

zone returns [*ZonePower ret]:
	z=ZONE {
		switch $z.GetText() {
		case "Z1":
		case "Z2":
		case "Z3":
		case "Z4":
		case "Z5a":
		case "Z5b":
		case "Z5c":
		default:
			fmt.Printf("Unknown zone: %s\n", $z)
			os.Exit(1)
		}
		$ret = &ZonePower{$z.GetText()}
	};

range returns [*RangePower ret]:
	l=NUMBER '%' ('-' h=NUMBER '%')? {
		$ret = &RangePower{
			low: Atoi("range low", $l.GetText()),
		}
		if localctx.GetH() != nil {
			$ret.high = Atoi("range high", $h.GetText())
		} else {
			$ret.high = $ret.low
		}
	};

dur returns [time.Duration ret]:
	(NUMBER DURSUFFIX)+ {
		$ret = ParseDuration(localctx.GetText())
	};

dist returns [*Distance ret]:
	n=NUMBER u=('mile'|'miles'|'yard'|'yards'|'meter'|'meters'|'km') {
		$ret = &Distance{
			Distance: Atoi("distance", $n.GetText()),
			Unit: $u.GetText(),
		}
	};

