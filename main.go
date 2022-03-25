// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	chart "github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const probableMageryGainFactor = 2.0 / 3.0

func main() {
	for i := 1.0; i < 9.0; i += 1.0 {
		fmt.Printf("Minimum/maximum skill for circle %.0f spells: %.1f / %.1f\n", i, minSkillToCast(i-1.0), maxSkillToCast(i-1.0))
	}
	fmt.Println()

	makeSkillGainChart()
	makeMageryCirclesChart()
	printSimulations()
}

func makeSkillGainChart() {
	xVals := make([]float64, 0, 500)
	yVals := make([]float64, 0, 500)
	for i := 0.0; i <= 1.0; i += 0.01 {
		y := chanceToGain(i, 100.0, probableMageryGainFactor)
		if y < 0.0 {
			break
		}
		xVals = append(xVals, i)
		yVals = append(yVals, y)
	}

	whiteOnBlackBackgroundStyle := chart.Style{
		FillColor: drawing.Color{0, 0, 0, 255},
		FontColor: drawing.Color{128, 128, 128, 255},
	}
	graph := chart.Chart{
		Title:      "% chance to gain skill, based on % chance to succeed at skill",
		TitleStyle: whiteOnBlackBackgroundStyle,
		Background: whiteOnBlackBackgroundStyle,
		Canvas:     whiteOnBlackBackgroundStyle,
		XAxis: chart.XAxis{
			Name:           "% chance to succeed",
			NameStyle:      whiteOnBlackBackgroundStyle,
			Style:          whiteOnBlackBackgroundStyle,
			TickStyle:      whiteOnBlackBackgroundStyle,
			GridMajorStyle: whiteOnBlackBackgroundStyle,
			GridMinorStyle: whiteOnBlackBackgroundStyle,
		},
		YAxis: chart.YAxis{
			Name:           "% chance to gain",
			NameStyle:      whiteOnBlackBackgroundStyle,
			Style:          whiteOnBlackBackgroundStyle,
			TickStyle:      whiteOnBlackBackgroundStyle,
			GridMajorStyle: whiteOnBlackBackgroundStyle,
			GridMinorStyle: whiteOnBlackBackgroundStyle,
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xVals,
				YValues: yVals,
			},
		},
	}
	fd, err := os.Create("gains-vs-success.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	if err := graph.Render(chart.PNG, fd); err != nil {
		panic(err)
	}
}

func printSimulations() {
	// Let's model gaining magery from 80.0 -> 100.0
	// Assumptions:
	// - each 7th circle cast costs 6gp (2x flamestrike regs @ 3gp/ea)
	// - each 8th circle cast costs 10gp (2x resurrect regs @ 3gp/ea, 1x @ 4gp/ea)
	// - minimum skill to cast a spell is ((100.0 / 7.0) * circle) - 20.0, where circles are numbered starting at 0
	// - maximum skill to always cast a spell is ((100.0 / 7.0) * circle) + 20.0
	// - chance to cast a spell is (SKILL - minSkill) / (maxSkill - minSkill)
	// - probability of gaining from a cast-attempt is -0.3chance^2 + 0.1chance + 0.2

	fmt.Println("SIMULATION: from 65.8 to 70.0 skill, casting EB(7gp) vs FS(6gp)")
	sixthCircleCasts := 0.0
	for i := 65.8; i < 69.9; i += 0.1 {
		sixthCircleCasts += expectedCastsForGain(i, 5.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 6th circle casts: %f, total cost: %.0f\n", sixthCircleCasts, sixthCircleCasts*7.0)
	seventhCircleCasts := 0.0
	for i := 65.8; i < 69.9; i += 0.1 {
		seventhCircleCasts += expectedCastsForGain(i, 6.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 7th circle casts: %f, total cost: %.0f\n", seventhCircleCasts, seventhCircleCasts*6.0)
	fmt.Println()

	fmt.Println("SIMULATION: from 80.0 to 100.0 skill, casting FS vs Resurrect(10gp)")
	seventhCircleCasts = 0.0
	for i := 80.0; i < 99.9; i += 0.1 {
		seventhCircleCasts += expectedCastsForGain(i, 6.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 7th circle casts: %f, total cost: %.0f\n", seventhCircleCasts, seventhCircleCasts*6.0)
	eightCircleCasts := 0.0
	for i := 80.0; i < 99.9; i += 0.1 {
		eightCircleCasts += expectedCastsForGain(i, 7.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 8th circle casts: %f, total cost: %.0f\n", eightCircleCasts, eightCircleCasts*10.0)
	fmt.Println()

	fmt.Println("SIMULATION: from 80.0 to 81.0 skill, casting FS vs Resurrect(10gp)")
	seventhCircleCasts = 0.0
	for i := 80.0; i < 80.9; i += 0.1 {
		seventhCircleCasts += expectedCastsForGain(i, 6.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 7th circle casts: %f, total cost: %.0f\n", seventhCircleCasts, seventhCircleCasts*6.0)
	eightCircleCasts = 0.0
	for i := 80.0; i < 80.9; i += 0.1 {
		eightCircleCasts += expectedCastsForGain(i, 7.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 8th circle casts: %f, total cost: %.0f\n", eightCircleCasts, eightCircleCasts*10.0)
	fmt.Println()

	fmt.Println("SIMULATION: from 90.0 to 100.0 skill, casting FS vs Resurrect(10gp)")
	seventhCircleCasts = 0.0
	for i := 90.0; i < 99.9; i += 0.1 {
		seventhCircleCasts += expectedCastsForGain(i, 6.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 7th circle casts: %f, total cost: %.0f\n", seventhCircleCasts, seventhCircleCasts*6.0)
	eightCircleCasts = 0.0
	for i := 90.0; i < 99.9; i += 0.1 {
		eightCircleCasts += expectedCastsForGain(i, 7.0, probableMageryGainFactor)
	}
	fmt.Printf("Total 8th circle casts: %f, total cost: %.0f\n", eightCircleCasts, eightCircleCasts*10.0)
	fmt.Println()
}

func chanceToGain(chanceToSucceed, currentSkill, gainFactor float64) float64 {
	// From https://github.com/runuo/runuo/blob/4298b27264c9fd94bdb4b74853d0eb1499d3fe08/Scripts/Misc/SkillCheck.cs#L130
	//    double gc = (double)(from.Skills.Cap - from.Skills.Total) / from.Skills.Cap;
	//    gc += ( skill.Cap - skill.Base ) / skill.Cap;
	//    gc /= 2;
	//
	//    gc += ( 1.0 - chance ) * ( success ? 0.5 : (Core.AOS ? 0.0 : 0.2) );
	//    gc /= 2;
	//
	//    gc *= skill.Info.GainFactor;
	//
	// This bit has been explicitly refuted by Chris/Telamon as being present on UOR:
	//    double gc = (double)(from.Skills.Cap - from.Skills.Total) / from.Skills.Cap;
	// The effect there would've been to make skill gain slower the closer your base skill total is to 700.0.
	gainChance := 0.0

	// This bit is straightforward:
	//    gc += ( skill.Cap - skill.Base ) / skill.Cap;
	// The effect here, if it's active on UOR (and I believe it is), is to make skill gain much faster
	// the lower your skill, regardless of your chance to succeed in the attempt.
	gainChance = (100.0 - currentSkill) / 100.0
	gainChance /= 2.0

	// This bit is more interesting:
	//    gc += ( 1.0 - chance ) * ( success ? 0.5 : (Core.AOS ? 0.0 : 0.2) )
	// We can omit the Core.AOS ternary, leaving us with just
	//    (1.0 - chance) * (success ? 0.5 : 0.2)
	// Since we're looking for probability, not absolutes, we can replace the success ternary with this expression:
	//    ((chance * 0.5) + ((1-chance) * 0.2))
	// That leaves us with this overall:
	//    (1.0 - chance) * (0.5 * chance + (0.2 * (1 - chance)))
	//
	gainChance += (1.0 - chanceToSucceed) * ((0.5 * chanceToSucceed) + (0.2 * (1 - chanceToSucceed)))
	gainChance /= 2.0
	// EXTRA EXTRA EXTRA: The graph for the immediate above equation is an inverted parabola... couldn't we find the
	// peak, and thus know the optimal thing to attempt to gain skill? WHY YES WE CAN!
	// First, let's make it look a bit more like a friendly polynomial by cross-multiplying everything out, giving:
	//    -0.3chance^2 + 0.1chance + 0.2
	// That let's us easily do a tiny bit of calculus here, applying the Power Rule to find the first derivative:
	//    y = -0.6chance + 0.1
	// If we find where that line crosses the X axis, we'll have the value of X where the parabola peaks.
	// So we set y to 0, giving chance = 0.1 / 0.6, or 1/6, or about 16.66% chance to succeed.

	// Finally, and CRITICALLY for evaluating if I'm off my rocker, note line 133:
	//    gc *= skill.Info.GainFactor;
	// This value is configurable per-skill, applying a flat multiplier to all skill-gain checks.
	// So if it's 0.5, it takes 2x as long to gain the skill, regardless of all other factors.
	// By default it's 1.0... I suspect that on UOR, for Magery at least, it's somewhat lower, hence
	// the much higher *actual* number of casts to get a gain we see in-game versus what these formulas
	// predict.
	gainChance *= gainFactor

	return gainChance
}

func minSkillToCast(circle float64) float64 {
	// from https://github.com/runuo/runuo/blob/4298b27264c9fd94bdb4b74853d0eb1499d3fe08/Scripts/Spells/Base/MagerySpell.cs#L29
	return 100.0/7.0*circle - 20.0
}
func maxSkillToCast(circle float64) float64 {
	// from https://github.com/runuo/runuo/blob/4298b27264c9fd94bdb4b74853d0eb1499d3fe08/Scripts/Spells/Base/MagerySpell.cs#L29
	return 100.0/7.0*circle + 20.0
}

func expectedCastsForGain(currentSkill, circle, gainFactor float64) float64 {
	// from https://github.com/runuo/runuo/blob/4298b27264c9fd94bdb4b74853d0eb1499d3fe08/Scripts/Misc/SkillCheck.cs#L98
	//     double chance = (value - minSkill) / (maxSkill - minSkill)
	chanceToSucceed := (currentSkill - minSkillToCast(circle)) / (maxSkillToCast(circle) - minSkillToCast(circle))

	return 1.0 / chanceToGain(chanceToSucceed, currentSkill, gainFactor)
}

func makeMageryCirclesChart() {
	whiteOnBlackBackgroundStyle := chart.Style{
		FillColor: drawing.Color{0, 0, 0, 255},
		FontColor: drawing.Color{128, 128, 128, 255},
	}
	var xTicks, yTicks []chart.Tick
	var xLines, yLines []chart.GridLine
	for i := 0.0; i < 110.0; i += 10.0 {
		xTicks = append(xTicks, chart.Tick{
			Value: i,
			Label: fmt.Sprintf("%.1f", i),
		})
		xLines = append(xLines, chart.GridLine{
			//IsMinor: false,
			//Style: whiteOnBlackBackgroundStyle,
			Value: i,
		})
	}
	for i := 4.0; i < 40.0; i += 1.0 {
		yTicks = append(yTicks, chart.Tick{
			Value: i,
			Label: fmt.Sprintf("%.1f", i),
		})
		yLines = append(yLines, chart.GridLine{
			//IsMinor: false,
			//Style: whiteOnBlackBackgroundStyle,
			Value: i,
		})
	}

	graph := chart.Chart{
		Title: "Expected casts to gain 0.1 magery skill, by spell circle",
		XAxis: chart.XAxis{
			Name:           "Magery skill",
			NameStyle:      whiteOnBlackBackgroundStyle,
			Style:          whiteOnBlackBackgroundStyle,
			TickStyle:      whiteOnBlackBackgroundStyle,
			GridMajorStyle: whiteOnBlackBackgroundStyle,
			GridMinorStyle: whiteOnBlackBackgroundStyle,
			Ticks:          xTicks,
			GridLines:      xLines,
		},
		YAxis: chart.YAxis{
			Name:           "Expected casts",
			NameStyle:      whiteOnBlackBackgroundStyle,
			Style:          whiteOnBlackBackgroundStyle,
			TickStyle:      whiteOnBlackBackgroundStyle,
			GridMajorStyle: whiteOnBlackBackgroundStyle,
			GridMinorStyle: whiteOnBlackBackgroundStyle,
			Ticks:          yTicks,
			GridLines:      yLines,
		},
		TitleStyle: whiteOnBlackBackgroundStyle,
		Background: whiteOnBlackBackgroundStyle,
		Canvas:     whiteOnBlackBackgroundStyle,
	}
	// 1st circle, 0-20
	graph.Series = append(graph.Series, makeMageryCircleSeries("1st circle", 0.0, 0.0, 20.0, chart.GetDefaultColor(0)))
	// 2nd circle, 0-34.3
	graph.Series = append(graph.Series, makeMageryCircleSeries("2nd circle", 1.0, 0.0, 34.3, drawing.ColorGreen))
	// 3rd circle, 8.6-48.6
	graph.Series = append(graph.Series, makeMageryCircleSeries("3rd circle", 2.0, 8.6, 48.6, chart.GetDefaultColor(2)))
	// 4th circle, 22.9-62.9
	graph.Series = append(graph.Series, makeMageryCircleSeries("4th circle", 3.0, 22.9, 62.9, chart.GetDefaultColor(3)))
	// 5th circle, 37.1-77.1
	graph.Series = append(graph.Series, makeMageryCircleSeries("5th circle", 4.0, 37.1, 77.1, chart.GetDefaultColor(4)))
	// 6th circle, 51.4-91.4
	graph.Series = append(graph.Series, makeMageryCircleSeries("6th circle", 5.0, 51.4, 91.4, drawing.ColorWhite))
	// 7th circle, 65.7-100.0
	graph.Series = append(graph.Series, makeMageryCircleSeries("7th circle", 6.0, 65.7, 99.9, chart.GetDefaultColor(6)))
	// 8th circle, 80.0-100.0
	graph.Series = append(graph.Series, makeMageryCircleSeries("8th circle", 7.0, 80.0, 99.9, drawing.ColorRed))

	graph.Elements = []chart.Renderable{chart.Legend(&graph, whiteOnBlackBackgroundStyle)}
	graph.Width = 2048
	graph.Height = 800

	fd, err := os.Create("magery-circles-gains.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	if err := graph.Render(chart.PNG, fd); err != nil {
		panic(err)
	}
}

func makeMageryCircleSeries(name string, circle, domainStart, domainEnd float64, color drawing.Color) chart.ContinuousSeries {
	numPoints := int((domainEnd - domainStart) / 0.1)
	xVals := make([]float64, 0, numPoints)
	yVals := make([]float64, 0, numPoints)
	for x := domainStart; x < domainEnd; x += 0.1 {
		y := expectedCastsForGain(x, circle, probableMageryGainFactor)
		//if y > 30.0 {
		//	break
		//}
		xVals = append(xVals, x)
		yVals = append(yVals, y)
	}
	return chart.ContinuousSeries{
		Name: name,
		Style: chart.Style{
			StrokeColor: color,
		},
		XValues: xVals,
		YValues: yVals,
	}
}
