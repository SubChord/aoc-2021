package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	l := readLines("inp")[0]

	binaryString := ""
	for _, v := range l {
		hex, _ := strconv.ParseInt(string(v), 16, 64)
		binaryString += fmt.Sprintf("%04s", strconv.FormatInt(hex, 2))
	}

	packet, _, _ := parsePacket(binaryString)
	fmt.Printf("part 1: %d\n", pt1(packet))
	fmt.Printf("part 2: %d\n", pt2(packet))
}

func pt2(packet Packet) int {
	v := 0
	switch packet.Type {
	case 0:
		for _, p := range packet.Packets {
			v += pt2(p)
		}
	case 1:
		if len(packet.Packets) > 1 {
			prod := pt2(packet.Packets[0])
			for _, p := range packet.Packets[1:] {
				prod *= pt2(p)
			}
			v = prod
		} else {
			v = pt2(packet.Packets[0])
		}
	case 2:
		min := math.MaxInt64
		for _, p := range packet.Packets {
			res := pt2(p)
			if res < min {
				min = res
			}
		}
		v = min
	case 3:
		max := 0
		for _, p := range packet.Packets {
			res := pt2(p)
			if res > max {
				max = res
			}
		}
		v = max
	case 4:
		v = int(packet.Value)
	case 5:
		if pt2(packet.Packets[0]) > pt2(packet.Packets[1]) {
			v = 1
		} else {
			v = 0
		}
	case 6:
		if pt2(packet.Packets[0]) < pt2(packet.Packets[1]) {
			v = 1
		} else {
			v = 0
		}
	case 7:
		if pt2(packet.Packets[0]) == pt2(packet.Packets[1]) {
			v = 1
		} else {
			v = 0
		}
	}

	return v
}

func pt1(packet Packet) int {
	s := 0
	s += packet.Version
	for _, p := range packet.Packets {
		s += pt1(p)
	}

	return s
}

func parsePacket(packet string) (Packet, string, int) {
	versionStr := packet[0:3]
	version, _ := strconv.ParseInt(versionStr, 2, 64)

	typeStr := packet[3:6]
	typeInt, _ := strconv.ParseInt(typeStr, 2, 64)

	p := Packet{
		Version: int(version),
		Type:    int(typeInt),
	}

	bitsParsed := 6

	// parse literal packet
	if p.Type == 4 {
		valueStr := ""
		lastGroup := false
		for !lastGroup {
			take := packet[bitsParsed : bitsParsed+5]
			if take[0] == '0' {
				lastGroup = true
			}
			bitsParsed += 5
			valueStr += take[1:]
		}
		parsed, err := strconv.ParseInt(valueStr, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		p.Value = parsed
	}

	if p.Type != 4 {
		lengthId := packet[bitsParsed : bitsParsed+1]
		bitsParsed += 1

		switch lengthId {
		case "0":
			next15Bits := packet[bitsParsed : bitsParsed+15]
			bitsParsed += 15

			length, err := strconv.ParseInt(next15Bits, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			nextBitsByLength := packet[bitsParsed : bitsParsed+int(length)]
			bitsParsed += int(length)

			sub, remainder, n := parsePacket(nextBitsByLength)
			p.Packets = append(p.Packets, sub)
			length -= int64(n)

			for length > 0 {
				sub, remainder, n = parsePacket(remainder)
				p.Packets = append(p.Packets, sub)
				length -= int64(n)
			}
		case "1":
			next11Bits := packet[bitsParsed : bitsParsed+11]
			bitsParsed += 11

			numOfPackets, err := strconv.ParseInt(next11Bits, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			remainder := packet[bitsParsed:]
			n := 0
			sub := Packet{}
			for numOfPackets > 0 {
				sub, remainder, n = parsePacket(remainder)
				p.Packets = append(p.Packets, sub)
				numOfPackets -= 1
				bitsParsed += n
			}
		}
	}

	return p, packet[bitsParsed:], bitsParsed
}

type Packet struct {
	Version int
	Type    int
	Value   int64
	Packets []Packet
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
