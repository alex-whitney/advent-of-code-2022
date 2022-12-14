package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type packetData struct {
	data  []packetData
	value int
}

func (p packetData) toList() packetData {
	if p.data != nil {
		panic("invalid op")
	}

	return packetData{
		data: []packetData{p},
	}
}

func (p *packetData) toString() string {
	if p.data == nil {
		return strconv.Itoa(p.value)
	}

	out := "["
	for i, d := range p.data {
		out = out + d.toString()
		if i != len(p.data)-1 {
			out += ","
		}
	}
	return out + "]"
}

type sortablePacketData []packetData

func (p sortablePacketData) Len() int {
	return len(p)
}

func (p sortablePacketData) Less(i, j int) bool {
	return compare(p[i], p[j]) == -1
}

func (p sortablePacketData) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Regret for not choosing python or JS, where a simple eval would be a lot simpler...
func newPacketData(s string) packetData {
	if !strings.ContainsAny(s, "[],") {
		val, err := strconv.Atoi(s)
		if err != nil {
			panic("yeah, this isn't right: " + s)
		}
		return packetData{
			value: val,
		}
	}

	packet := packetData{
		data: make([]packetData, 0),
	}

	depth := 0
	pos := 0
	tokenStart := -1
	for pos < len(s) {
		if s[pos] == '[' {
			depth += 1
			if depth == 2 {
				tokenStart = pos
			}
		} else if s[pos] == ']' {
			if depth == 1 && tokenStart > -1 {
				packet.data = append(packet.data, newPacketData(s[tokenStart:pos]))
			} else if depth == 2 {
				packet.data = append(packet.data, newPacketData(s[tokenStart:pos+1]))
				tokenStart = -1
			}
			depth -= 1
		} else if s[pos] == ',' && depth == 1 {
			if tokenStart > -1 {
				packet.data = append(packet.data, newPacketData(s[tokenStart:pos]))
				tokenStart = -1
			}
		} else if tokenStart == -1 {
			tokenStart = pos
		}
		pos += 1
	}

	return packet
}

func compare(left packetData, right packetData) int {
	if left.data == nil && right.data == nil {
		if left.value < right.value {
			return -1
		} else if left.value > right.value {
			return 1
		} else {
			return 0
		}
	}

	if left.data == nil {
		return compare(left.toList(), right)
	}
	if right.data == nil {
		return compare(left, right.toList())
	}

	for i := 0; i < len(left.data) && i < len(right.data); i++ {
		result := compare(left.data[i], right.data[i])
		if result != 0 {
			return result
		}
	}

	if len(left.data) < len(right.data) {
		return -1
	} else if len(left.data) > len(right.data) {
		return 1
	}

	return 0
}

type Today struct {
	packets [][]packetData
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadFile(input)
	if err != nil {
		return err
	}

	groups := strings.Split(raw, "\n\n")
	d.packets = make([][]packetData, len(groups))
	for i, group := range groups {
		packets := strings.Split(group, "\n")
		d.packets[i] = []packetData{
			newPacketData(packets[0]),
			newPacketData(packets[1]),
		}
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	sum := 0
	for i, packet := range d.packets {
		result := compare(packet[0], packet[1])

		if result == -1 {
			sum += i + 1
		}
	}

	return strconv.Itoa(sum), nil
}

func (d *Today) Part2() (string, error) {
	allPackets := make(sortablePacketData, 0)
	allPackets = append(allPackets, newPacketData("[[2]]"), newPacketData("[[6]]"))
	for _, group := range d.packets {
		allPackets = append(allPackets, group[0], group[1])
	}

	sort.Sort(allPackets)

	val := 1
	for i, packet := range allPackets {
		if packet.toString() == "[[2]]" || packet.toString() == "[[6]]" {
			val = val * (i + 1)
		}
	}

	return strconv.Itoa(val), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
