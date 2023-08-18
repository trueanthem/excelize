// Copyright 2016 - 2023 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.16 or later.

package excelize

import (
	"fmt"
	"github.com/xuri/excelize/v2/xencoding/xml"
	"strconv"
	"strings"
)

// ChartType is the type of supported chart types.
type ChartType byte

// This section defines the currently supported chart types enumeration.
const (
	Area ChartType = iota
	AreaStacked
	AreaPercentStacked
	Area3D
	Area3DStacked
	Area3DPercentStacked
	Bar
	BarStacked
	BarPercentStacked
	Bar3DClustered
	Bar3DStacked
	Bar3DPercentStacked
	Bar3DConeClustered
	Bar3DConeStacked
	Bar3DConePercentStacked
	Bar3DPyramidClustered
	Bar3DPyramidStacked
	Bar3DPyramidPercentStacked
	Bar3DCylinderClustered
	Bar3DCylinderStacked
	Bar3DCylinderPercentStacked
	Col
	ColStacked
	ColPercentStacked
	Col3D
	Col3DClustered
	Col3DStacked
	Col3DPercentStacked
	Col3DCone
	Col3DConeClustered
	Col3DConeStacked
	Col3DConePercentStacked
	Col3DPyramid
	Col3DPyramidClustered
	Col3DPyramidStacked
	Col3DPyramidPercentStacked
	Col3DCylinder
	Col3DCylinderClustered
	Col3DCylinderStacked
	Col3DCylinderPercentStacked
	Doughnut
	Line
	Line3D
	Pie
	Pie3D
	PieOfPie
	BarOfPie
	Radar
	Scatter
	Surface3D
	WireframeSurface3D
	Contour
	WireframeContour
	Bubble
	Bubble3D
)

// This section defines the default value of chart properties.
var (
	chartView3DRotX = map[ChartType]int{
		Area:                        0,
		AreaStacked:                 0,
		AreaPercentStacked:          0,
		Area3D:                      15,
		Area3DStacked:               15,
		Area3DPercentStacked:        15,
		Bar:                         0,
		BarStacked:                  0,
		BarPercentStacked:           0,
		Bar3DClustered:              15,
		Bar3DStacked:                15,
		Bar3DPercentStacked:         15,
		Bar3DConeClustered:          15,
		Bar3DConeStacked:            15,
		Bar3DConePercentStacked:     15,
		Bar3DPyramidClustered:       15,
		Bar3DPyramidStacked:         15,
		Bar3DPyramidPercentStacked:  15,
		Bar3DCylinderClustered:      15,
		Bar3DCylinderStacked:        15,
		Bar3DCylinderPercentStacked: 15,
		Col:                         0,
		ColStacked:                  0,
		ColPercentStacked:           0,
		Col3D:                       15,
		Col3DClustered:              15,
		Col3DStacked:                15,
		Col3DPercentStacked:         15,
		Col3DCone:                   15,
		Col3DConeClustered:          15,
		Col3DConeStacked:            15,
		Col3DConePercentStacked:     15,
		Col3DPyramid:                15,
		Col3DPyramidClustered:       15,
		Col3DPyramidStacked:         15,
		Col3DPyramidPercentStacked:  15,
		Col3DCylinder:               15,
		Col3DCylinderClustered:      15,
		Col3DCylinderStacked:        15,
		Col3DCylinderPercentStacked: 15,
		Doughnut:                    0,
		Line:                        0,
		Line3D:                      20,
		Pie:                         0,
		Pie3D:                       30,
		PieOfPie:                    0,
		BarOfPie:                    0,
		Radar:                       0,
		Scatter:                     0,
		Surface3D:                   15,
		WireframeSurface3D:          15,
		Contour:                     90,
		WireframeContour:            90,
	}
	chartView3DRotY = map[ChartType]int{
		Area:                        0,
		AreaStacked:                 0,
		AreaPercentStacked:          0,
		Area3D:                      20,
		Area3DStacked:               20,
		Area3DPercentStacked:        20,
		Bar:                         0,
		BarStacked:                  0,
		BarPercentStacked:           0,
		Bar3DClustered:              20,
		Bar3DStacked:                20,
		Bar3DPercentStacked:         20,
		Bar3DConeClustered:          20,
		Bar3DConeStacked:            20,
		Bar3DConePercentStacked:     20,
		Bar3DPyramidClustered:       20,
		Bar3DPyramidStacked:         20,
		Bar3DPyramidPercentStacked:  20,
		Bar3DCylinderClustered:      20,
		Bar3DCylinderStacked:        20,
		Bar3DCylinderPercentStacked: 20,
		Col:                         0,
		ColStacked:                  0,
		ColPercentStacked:           0,
		Col3D:                       20,
		Col3DClustered:              20,
		Col3DStacked:                20,
		Col3DPercentStacked:         20,
		Col3DCone:                   20,
		Col3DConeClustered:          20,
		Col3DConeStacked:            20,
		Col3DConePercentStacked:     20,
		Col3DPyramid:                20,
		Col3DPyramidClustered:       20,
		Col3DPyramidStacked:         20,
		Col3DPyramidPercentStacked:  20,
		Col3DCylinder:               20,
		Col3DCylinderClustered:      20,
		Col3DCylinderStacked:        20,
		Col3DCylinderPercentStacked: 20,
		Doughnut:                    0,
		Line:                        0,
		Line3D:                      15,
		Pie:                         0,
		Pie3D:                       0,
		PieOfPie:                    0,
		BarOfPie:                    0,
		Radar:                       0,
		Scatter:                     0,
		Surface3D:                   20,
		WireframeSurface3D:          20,
		Contour:                     0,
		WireframeContour:            0,
	}
	plotAreaChartOverlap = map[ChartType]int{
		BarStacked:        100,
		BarPercentStacked: 100,
		ColStacked:        100,
		ColPercentStacked: 100,
	}
	chartView3DPerspective = map[ChartType]int{
		Line3D:           30,
		Contour:          0,
		WireframeContour: 0,
	}
	chartView3DRAngAx = map[ChartType]int{
		Area:                        0,
		AreaStacked:                 0,
		AreaPercentStacked:          0,
		Area3D:                      1,
		Area3DStacked:               1,
		Area3DPercentStacked:        1,
		Bar:                         0,
		BarStacked:                  0,
		BarPercentStacked:           0,
		Bar3DClustered:              1,
		Bar3DStacked:                1,
		Bar3DPercentStacked:         1,
		Bar3DConeClustered:          1,
		Bar3DConeStacked:            1,
		Bar3DConePercentStacked:     1,
		Bar3DPyramidClustered:       1,
		Bar3DPyramidStacked:         1,
		Bar3DPyramidPercentStacked:  1,
		Bar3DCylinderClustered:      1,
		Bar3DCylinderStacked:        1,
		Bar3DCylinderPercentStacked: 1,
		Col:                         0,
		ColStacked:                  0,
		ColPercentStacked:           0,
		Col3D:                       1,
		Col3DClustered:              1,
		Col3DStacked:                1,
		Col3DPercentStacked:         1,
		Col3DCone:                   1,
		Col3DConeClustered:          1,
		Col3DConeStacked:            1,
		Col3DConePercentStacked:     1,
		Col3DPyramid:                1,
		Col3DPyramidClustered:       1,
		Col3DPyramidStacked:         1,
		Col3DPyramidPercentStacked:  1,
		Col3DCylinder:               1,
		Col3DCylinderClustered:      1,
		Col3DCylinderStacked:        1,
		Col3DCylinderPercentStacked: 1,
		Doughnut:                    0,
		Line:                        0,
		Line3D:                      0,
		Pie:                         0,
		Pie3D:                       0,
		PieOfPie:                    0,
		BarOfPie:                    0,
		Radar:                       0,
		Scatter:                     0,
		Surface3D:                   0,
		WireframeSurface3D:          0,
		Contour:                     0,
		Bubble:                      0,
		Bubble3D:                    0,
	}
	chartLegendPosition = map[string]string{
		"bottom":    "b",
		"left":      "l",
		"right":     "r",
		"top":       "t",
		"top_right": "tr",
	}
	chartValAxNumFmtFormatCode = map[ChartType]string{
		Area:                        "General",
		AreaStacked:                 "General",
		AreaPercentStacked:          "0%",
		Area3D:                      "General",
		Area3DStacked:               "General",
		Area3DPercentStacked:        "0%",
		Bar:                         "General",
		BarStacked:                  "General",
		BarPercentStacked:           "0%",
		Bar3DClustered:              "General",
		Bar3DStacked:                "General",
		Bar3DPercentStacked:         "0%",
		Bar3DConeClustered:          "General",
		Bar3DConeStacked:            "General",
		Bar3DConePercentStacked:     "0%",
		Bar3DPyramidClustered:       "General",
		Bar3DPyramidStacked:         "General",
		Bar3DPyramidPercentStacked:  "0%",
		Bar3DCylinderClustered:      "General",
		Bar3DCylinderStacked:        "General",
		Bar3DCylinderPercentStacked: "0%",
		Col:                         "General",
		ColStacked:                  "General",
		ColPercentStacked:           "0%",
		Col3D:                       "General",
		Col3DClustered:              "General",
		Col3DStacked:                "General",
		Col3DPercentStacked:         "0%",
		Col3DCone:                   "General",
		Col3DConeClustered:          "General",
		Col3DConeStacked:            "General",
		Col3DConePercentStacked:     "0%",
		Col3DPyramid:                "General",
		Col3DPyramidClustered:       "General",
		Col3DPyramidStacked:         "General",
		Col3DPyramidPercentStacked:  "0%",
		Col3DCylinder:               "General",
		Col3DCylinderClustered:      "General",
		Col3DCylinderStacked:        "General",
		Col3DCylinderPercentStacked: "0%",
		Doughnut:                    "General",
		Line:                        "General",
		Line3D:                      "General",
		Pie:                         "General",
		Pie3D:                       "General",
		PieOfPie:                    "General",
		BarOfPie:                    "General",
		Radar:                       "General",
		Scatter:                     "General",
		Surface3D:                   "General",
		WireframeSurface3D:          "General",
		Contour:                     "General",
		WireframeContour:            "General",
		Bubble:                      "General",
		Bubble3D:                    "General",
	}
	chartValAxCrossBetween = map[ChartType]string{
		Area:                        "midCat",
		AreaStacked:                 "midCat",
		AreaPercentStacked:          "midCat",
		Area3D:                      "midCat",
		Area3DStacked:               "midCat",
		Area3DPercentStacked:        "midCat",
		Bar:                         "between",
		BarStacked:                  "between",
		BarPercentStacked:           "between",
		Bar3DClustered:              "between",
		Bar3DStacked:                "between",
		Bar3DPercentStacked:         "between",
		Bar3DConeClustered:          "between",
		Bar3DConeStacked:            "between",
		Bar3DConePercentStacked:     "between",
		Bar3DPyramidClustered:       "between",
		Bar3DPyramidStacked:         "between",
		Bar3DPyramidPercentStacked:  "between",
		Bar3DCylinderClustered:      "between",
		Bar3DCylinderStacked:        "between",
		Bar3DCylinderPercentStacked: "between",
		Col:                         "between",
		ColStacked:                  "between",
		ColPercentStacked:           "between",
		Col3D:                       "between",
		Col3DClustered:              "between",
		Col3DStacked:                "between",
		Col3DPercentStacked:         "between",
		Col3DCone:                   "between",
		Col3DConeClustered:          "between",
		Col3DConeStacked:            "between",
		Col3DConePercentStacked:     "between",
		Col3DPyramid:                "between",
		Col3DPyramidClustered:       "between",
		Col3DPyramidStacked:         "between",
		Col3DPyramidPercentStacked:  "between",
		Col3DCylinder:               "between",
		Col3DCylinderClustered:      "between",
		Col3DCylinderStacked:        "between",
		Col3DCylinderPercentStacked: "between",
		Doughnut:                    "between",
		Line:                        "between",
		Line3D:                      "between",
		Pie:                         "between",
		Pie3D:                       "between",
		PieOfPie:                    "between",
		BarOfPie:                    "between",
		Radar:                       "between",
		Scatter:                     "between",
		Surface3D:                   "midCat",
		WireframeSurface3D:          "midCat",
		Contour:                     "midCat",
		WireframeContour:            "midCat",
		Bubble:                      "midCat",
		Bubble3D:                    "midCat",
	}
	plotAreaChartGrouping = map[ChartType]string{
		Area:                        "standard",
		AreaStacked:                 "stacked",
		AreaPercentStacked:          "percentStacked",
		Area3D:                      "standard",
		Area3DStacked:               "stacked",
		Area3DPercentStacked:        "percentStacked",
		Bar:                         "clustered",
		BarStacked:                  "stacked",
		BarPercentStacked:           "percentStacked",
		Bar3DClustered:              "clustered",
		Bar3DStacked:                "stacked",
		Bar3DPercentStacked:         "percentStacked",
		Bar3DConeClustered:          "clustered",
		Bar3DConeStacked:            "stacked",
		Bar3DConePercentStacked:     "percentStacked",
		Bar3DPyramidClustered:       "clustered",
		Bar3DPyramidStacked:         "stacked",
		Bar3DPyramidPercentStacked:  "percentStacked",
		Bar3DCylinderClustered:      "clustered",
		Bar3DCylinderStacked:        "stacked",
		Bar3DCylinderPercentStacked: "percentStacked",
		Col:                         "clustered",
		ColStacked:                  "stacked",
		ColPercentStacked:           "percentStacked",
		Col3D:                       "standard",
		Col3DClustered:              "clustered",
		Col3DStacked:                "stacked",
		Col3DPercentStacked:         "percentStacked",
		Col3DCone:                   "standard",
		Col3DConeClustered:          "clustered",
		Col3DConeStacked:            "stacked",
		Col3DConePercentStacked:     "percentStacked",
		Col3DPyramid:                "standard",
		Col3DPyramidClustered:       "clustered",
		Col3DPyramidStacked:         "stacked",
		Col3DPyramidPercentStacked:  "percentStacked",
		Col3DCylinder:               "standard",
		Col3DCylinderClustered:      "clustered",
		Col3DCylinderStacked:        "stacked",
		Col3DCylinderPercentStacked: "percentStacked",
		Line:                        "standard",
		Line3D:                      "standard",
	}
	plotAreaChartBarDir = map[ChartType]string{
		Bar:                         "bar",
		BarStacked:                  "bar",
		BarPercentStacked:           "bar",
		Bar3DClustered:              "bar",
		Bar3DStacked:                "bar",
		Bar3DPercentStacked:         "bar",
		Bar3DConeClustered:          "bar",
		Bar3DConeStacked:            "bar",
		Bar3DConePercentStacked:     "bar",
		Bar3DPyramidClustered:       "bar",
		Bar3DPyramidStacked:         "bar",
		Bar3DPyramidPercentStacked:  "bar",
		Bar3DCylinderClustered:      "bar",
		Bar3DCylinderStacked:        "bar",
		Bar3DCylinderPercentStacked: "bar",
		Col:                         "col",
		ColStacked:                  "col",
		ColPercentStacked:           "col",
		Col3D:                       "col",
		Col3DClustered:              "col",
		Col3DStacked:                "col",
		Col3DPercentStacked:         "col",
		Col3DCone:                   "col",
		Col3DConeStacked:            "col",
		Col3DConeClustered:          "col",
		Col3DConePercentStacked:     "col",
		Col3DPyramid:                "col",
		Col3DPyramidClustered:       "col",
		Col3DPyramidStacked:         "col",
		Col3DPyramidPercentStacked:  "col",
		Col3DCylinder:               "col",
		Col3DCylinderClustered:      "col",
		Col3DCylinderStacked:        "col",
		Col3DCylinderPercentStacked: "col",
		Line:                        "standard",
		Line3D:                      "standard",
	}
	orientation = map[bool]string{
		true:  "maxMin",
		false: "minMax",
	}
	catAxPos = map[bool]string{
		true:  "t",
		false: "b",
	}
	valAxPos = map[bool]string{
		true:  "r",
		false: "l",
	}
	valTickLblPos = map[ChartType]string{
		Contour:          "none",
		WireframeContour: "none",
	}
)

// parseChartOptions provides a function to parse the format settings of the
// chart with default value.
func parseChartOptions(opts *Chart) (*Chart, error) {
	if opts == nil {
		return nil, ErrParameterInvalid
	}
	if opts.Dimension.Width == 0 {
		opts.Dimension.Width = defaultChartDimensionWidth
	}
	if opts.Dimension.Height == 0 {
		opts.Dimension.Height = defaultChartDimensionHeight
	}
	if opts.Format.PrintObject == nil {
		opts.Format.PrintObject = boolPtr(true)
	}
	if opts.Format.Locked == nil {
		opts.Format.Locked = boolPtr(false)
	}
	if opts.Format.ScaleX == 0 {
		opts.Format.ScaleX = defaultPictureScale
	}
	if opts.Format.ScaleY == 0 {
		opts.Format.ScaleY = defaultPictureScale
	}
	if opts.Legend.Position == "" {
		opts.Legend.Position = defaultChartLegendPosition
	}
	for i := range opts.Title {
		if opts.Title[i].Font == nil {
			opts.Title[i].Font = &Font{}
		}
		if opts.Title[i].Font.Color == "" {
			opts.Title[i].Font.Color = "595959"
		}
		if opts.Title[i].Font.Size == 0 {
			opts.Title[i].Font.Size = 14
		}
	}
	if opts.VaryColors == nil {
		opts.VaryColors = boolPtr(true)
	}
	if opts.ShowBlanksAs == "" {
		opts.ShowBlanksAs = defaultChartShowBlanksAs
	}
	return opts, nil
}

// AddChart provides the method to add chart in a sheet by given chart format
// set (such as offset, scale, aspect ratio setting and print settings) and
// properties set. For example, create 3D clustered column chart with data
// Sheet1!$E$1:$L$15:
//
//	package main
//
//	import (
//	    "fmt"
//
//	    "github.com/xuri/excelize/v2"
//	)
//
//	func main() {
//	    f := excelize.NewFile()
//	    defer func() {
//	        if err := f.Close(); err != nil {
//	            fmt.Println(err)
//	        }
//	    }()
//	    for idx, row := range [][]interface{}{
//	        {nil, "Apple", "Orange", "Pear"}, {"Small", 2, 3, 3},
//	        {"Normal", 5, 2, 4}, {"Large", 6, 7, 8},
//	    } {
//	        cell, err := excelize.CoordinatesToCellName(1, idx+1)
//	        if err != nil {
//	            fmt.Println(err)
//	            return
//	        }
//	        f.SetSheetRow("Sheet1", cell, &row)
//	    }
//	    if err := f.AddChart("Sheet1", "E1", &excelize.Chart{
//	        Type: excelize.Col3DClustered,
//	        Series: []excelize.ChartSeries{
//	            {
//	                Name:       "Sheet1!$A$2",
//	                Categories: "Sheet1!$B$1:$D$1",
//	                Values:     "Sheet1!$B$2:$D$2",
//	            },
//	            {
//	                Name:       "Sheet1!$A$3",
//	                Categories: "Sheet1!$B$1:$D$1",
//	                Values:     "Sheet1!$B$3:$D$3",
//	            },
//	            {
//	                Name:       "Sheet1!$A$4",
//	                Categories: "Sheet1!$B$1:$D$1",
//	                Values:     "Sheet1!$B$4:$D$4",
//	            },
//	        },
//	        Title: []excelize.RichTextRun{
//	            {
//	                Text: "Fruit 3D Clustered Column Chart",
//	            },
//	        },
//	        Legend: excelize.ChartLegend{
//	            ShowLegendKey: false,
//	        },
//	        PlotArea: excelize.ChartPlotArea{
//	            ShowBubbleSize:  true,
//	            ShowCatName:     false,
//	            ShowLeaderLines: false,
//	            ShowPercent:     true,
//	            ShowSerName:     true,
//	            ShowVal:         true,
//	        },
//	    }); err != nil {
//	        fmt.Println(err)
//	        return
//	    }
//	    // Save spreadsheet by the given path.
//	    if err := f.SaveAs("Book1.xlsx"); err != nil {
//	        fmt.Println(err)
//	    }
//	}
//
// The following shows the type of chart supported by excelize:
//
//	 ID | Enumeration                 | Chart
//	----+-----------------------------+------------------------------
//	 0  | Area                        | 2D area chart
//	 1  | AreaStacked                 | 2D stacked area chart
//	 2  | AreaPercentStacked          | 2D 100% stacked area chart
//	 3  | Area3D                      | 3D area chart
//	 4  | Area3DStacked               | 3D stacked area chart
//	 5  | Area3DPercentStacked        | 3D 100% stacked area chart
//	 6  | Bar                         | 2D clustered bar chart
//	 7  | BarStacked                  | 2D stacked bar chart
//	 8  | BarPercentStacked           | 2D 100% stacked bar chart
//	 9  | Bar3DClustered              | 3D clustered bar chart
//	 10 | Bar3DStacked                | 3D stacked bar chart
//	 11 | Bar3DPercentStacked         | 3D 100% stacked bar chart
//	 12 | Bar3DConeClustered          | 3D cone clustered bar chart
//	 13 | Bar3DConeStacked            | 3D cone stacked bar chart
//	 14 | Bar3DConePercentStacked     | 3D cone percent bar chart
//	 15 | Bar3DPyramidClustered       | 3D pyramid clustered bar chart
//	 16 | Bar3DPyramidStacked         | 3D pyramid stacked bar chart
//	 17 | Bar3DPyramidPercentStacked  | 3D pyramid percent stacked bar chart
//	 18 | Bar3DCylinderClustered      | 3D cylinder clustered bar chart
//	 19 | Bar3DCylinderStacked        | 3D cylinder stacked bar chart
//	 20 | Bar3DCylinderPercentStacked | 3D cylinder percent stacked bar chart
//	 21 | Col                         | 2D clustered column chart
//	 22 | ColStacked                  | 2D stacked column chart
//	 23 | ColPercentStacked           | 2D 100% stacked column chart
//	 24 | Col3DClustered              | 3D clustered column chart
//	 25 | Col3D                       | 3D column chart
//	 26 | Col3DStacked                | 3D stacked column chart
//	 27 | Col3DPercentStacked         | 3D 100% stacked column chart
//	 28 | Col3DCone                   | 3D cone column chart
//	 29 | Col3DConeClustered          | 3D cone clustered column chart
//	 30 | Col3DConeStacked            | 3D cone stacked column chart
//	 31 | Col3DConePercentStacked     | 3D cone percent stacked column chart
//	 32 | Col3DPyramid                | 3D pyramid column chart
//	 33 | Col3DPyramidClustered       | 3D pyramid clustered column chart
//	 34 | Col3DPyramidStacked         | 3D pyramid stacked column chart
//	 35 | Col3DPyramidPercentStacked  | 3D pyramid percent stacked column chart
//	 36 | Col3DCylinder               | 3D cylinder column chart
//	 37 | Col3DCylinderClustered      | 3D cylinder clustered column chart
//	 38 | Col3DCylinderStacked        | 3D cylinder stacked column chart
//	 39 | Col3DCylinderPercentStacked | 3D cylinder percent stacked column chart
//	 40 | Doughnut                    | doughnut chart
//	 41 | Line                        | line chart
//	 42 | Line3D                      | 3D line chart
//	 43 | Pie                         | pie chart
//	 44 | Pie3D                       | 3D pie chart
//	 45 | PieOfPie                    | pie of pie chart
//	 46 | BarOfPie                    | bar of pie chart
//	 47 | Radar                       | radar chart
//	 48 | Scatter                     | scatter chart
//	 49 | Surface3D                   | 3D surface chart
//	 50 | WireframeSurface3D          | 3D wireframe surface chart
//	 51 | Contour                     | contour chart
//	 52 | WireframeContour            | wireframe contour chart
//	 53 | Bubble                      | bubble chart
//	 54 | Bubble3D                    | 3D bubble chart
//
// In Excel a chart series is a collection of information that defines which
// data is plotted such as values, axis labels and formatting.
//
// The series options that can be set are:
//
//	Name
//	Categories
//	Sizes
//	Values
//	Fill
//	Line
//	Marker
//
// Name: Set the name for the series. The name is displayed in the chart legend
// and in the formula bar. The 'Name' property is optional and if it isn't
// supplied it will default to Series 1..n. The name can also be a formula such
// as Sheet1!$A$1
//
// Categories: This sets the chart category labels. The category is more or less
// the same as the X axis. In most chart types the 'Categories' property is
// optional and the chart will just assume a sequential series from 1..n.
//
// Sizes: This sets the bubble size in a data series.
//
// Values: This is the most important property of a series and is the only
// mandatory option for every chart object. This option links the chart with
// the worksheet data that it displays.
//
// Fill: This set the format for the data series fill.
//
// Line: This sets the line format of the line chart. The 'Line' property is
// optional and if it isn't supplied it will default style. The options that
// can be set are width and color. The range of width is 0.25pt - 999pt. If the
// value of width is outside the range, the default width of the line is 2pt.
//
// Marker: This sets the marker of the line chart and scatter chart. The range
// of optional field 'Size' is 2-72 (default value is 5). The enumeration value
// of optional field 'Symbol' are (default value is 'auto'):
//
//	circle
//	dash
//	diamond
//	dot
//	none
//	picture
//	plus
//	square
//	star
//	triangle
//	x
//	auto
//
// Set properties of the chart legend. The options that can be set are:
//
//	Position
//	ShowLegendKey
//
// Position: Set the position of the chart legend. The default legend position
// is bottom. The available positions are:
//
//	none
//	top
//	bottom
//	left
//	right
//	top_right
//
// ShowLegendKey: Set the legend keys shall be shown in data labels. The default
// value is false.
//
// Set properties of the chart title. The properties that can be set are:
//
//	Title
//
// Title: Set the name (title) for the chart. The name is displayed above the
// chart. The name can also be a formula such as Sheet1!$A$1 or a list with a
// sheet name. The name property is optional. The default is to have no chart
// title.
//
// Specifies how blank cells are plotted on the chart by 'ShowBlanksAs'. The
// default value is gap. The options that can be set are:
//
//	gap
//	span
//	zero
//
// gap: Specifies that blank values shall be left as a gap.
//
// span: Specifies that blank values shall be spanned with a line.
//
// zero: Specifies that blank values shall be treated as zero.
//
// Specifies that each data marker in the series has a different color by
// 'VaryColors'. The default value is true.
//
// Set chart offset, scale, aspect ratio setting and print settings by format,
// same as function 'AddPicture'.
//
// Set the position of the chart plot area by PlotArea. The properties that can
// be set are:
//
//	SecondPlotValues
//	ShowBubbleSize
//	ShowCatName
//	ShowLeaderLines
//	ShowPercent
//	ShowSerName
//	ShowVal
//
// SecondPlotValues: Specifies the values in second plot for the 'pieOfPie' and
// 'barOfPie' chart.
//
// ShowBubbleSize: Specifies the bubble size shall be shown in a data label. The
// 'ShowBubbleSize' property is optional. The default value is false.
//
// ShowCatName: Specifies that the category name shall be shown in the data
// label. The 'ShowCatName' property is optional. The default value is true.
//
// ShowLeaderLines: Specifies leader lines shall be shown for data labels. The
// 'ShowLeaderLines' property is optional. The default value is false.
//
// ShowPercent: Specifies that the percentage shall be shown in a data label.
// The 'ShowPercent' property is optional. The default value is false.
//
// ShowSerName: Specifies that the series name shall be shown in a data label.
// The 'ShowSerName' property is optional. The default value is false.
//
// ShowVal: Specifies that the value shall be shown in a data label.
// The 'ShowVal' property is optional. The default value is false.
//
// Set the primary horizontal and vertical axis options by 'XAxis' and 'YAxis'.
// The properties of 'XAxis' that can be set are:
//
//	None
//	MajorGridLines
//	MinorGridLines
//	TickLabelSkip
//	ReverseOrder
//	Maximum
//	Minimum
//	Font
//	NumFmt
//	Title
//
// The properties of 'YAxis' that can be set are:
//
//	None
//	MajorGridLines
//	MinorGridLines
//	MajorUnit
//	Secondary
//	ReverseOrder
//	Maximum
//	Minimum
//	Font
//	LogBase
//	NumFmt
//	Title
//
// None: Disable axes.
//
// MajorGridLines: Specifies major grid lines.
//
// MinorGridLines: Specifies minor grid lines.
//
// MajorUnit: Specifies the distance between major ticks. Shall contain a
// positive floating-point number. The 'MajorUnit' property is optional. The
// default value is auto.
//
// Secondary: Specifies the current series vertical axis as the secondary axis,
// this only works for the second and later chart in the combo chart. The
// default value is false.
//
// TickLabelSkip: Specifies how many tick labels to skip between label that is
// drawn. The 'TickLabelSkip' property is optional. The default value is auto.
//
// ReverseOrder: Specifies that the categories or values on reverse order
// (orientation of the chart). The 'ReverseOrder' property is optional. The
// default value is false.
//
// Maximum: Specifies that the fixed maximum, 0 is auto. The 'Maximum' property
// is optional. The default value is auto.
//
// Minimum: Specifies that the fixed minimum, 0 is auto. The 'Minimum' property
// is optional. The default value is auto.
//
// Font: Specifies that the font of the horizontal and vertical axis. The
// properties of font that can be set are:
//
//	Bold
//	Italic
//	Underline
//	Family
//	Size
//	Strike
//	Color
//	VertAlign
//
// LogBase: Specifies logarithmic scale base number of the vertical axis.
//
// NumFmt: Specifies that if linked to source and set custom number format code
// for axis. The 'NumFmt' property is optional. The default format code is
// 'General'.
//
// Title: Specifies that the primary horizontal or vertical axis title and
// resize chart. The 'Title' property is optional.
//
// Set chart size by 'Dimension' property. The 'Dimension' property is optional.
// The default width is 480, and height is 260.
//
// combo: Specifies the create a chart that combines two or more chart types in
// a single chart. For example, create a clustered column - line chart with
// data Sheet1!$E$1:$L$15:
//
//	package main
//
//	import (
//	    "fmt"
//
//	    "github.com/xuri/excelize/v2"
//	)
//
//	func main() {
//	    f := excelize.NewFile()
//	    defer func() {
//	        if err := f.Close(); err != nil {
//	            fmt.Println(err)
//	        }
//	    }()
//	    for idx, row := range [][]interface{}{
//	        {nil, "Apple", "Orange", "Pear"}, {"Small", 2, 3, 3},
//	        {"Normal", 5, 2, 4}, {"Large", 6, 7, 8},
//	    } {
//	        cell, err := excelize.CoordinatesToCellName(1, idx+1)
//	        if err != nil {
//	            fmt.Println(err)
//	            return
//	        }
//	        f.SetSheetRow("Sheet1", cell, &row)
//	    }
//	    enable, disable := true, false
//	    if err := f.AddChart("Sheet1", "E1", &excelize.Chart{
//	        Type: "col",
//	        Series: []excelize.ChartSeries{
//	            {
//	                Name:       "Sheet1!$A$2",
//	                Categories: "Sheet1!$B$1:$D$1",
//	                Values:     "Sheet1!$B$2:$D$2",
//	            },
//	        },
//	        Format: excelize.GraphicOptions{
//	            ScaleX:          1,
//	            ScaleY:          1,
//	            OffsetX:         15,
//	            OffsetY:         10,
//	            PrintObject:     &enable,
//	            LockAspectRatio: false,
//	            Locked:          &disable,
//	        },
//	        Title: []excelize.RichTextRun{
//	            {
//	                Text: "Clustered Column - Line Chart",
//	            },
//	        },
//	        Legend: excelize.ChartLegend{
//	            Position:      "left",
//	            ShowLegendKey: false,
//	        },
//	        PlotArea: excelize.ChartPlotArea{
//	            ShowCatName:     false,
//	            ShowLeaderLines: false,
//	            ShowPercent:     true,
//	            ShowSerName:     true,
//	            ShowVal:         true,
//	        },
//	    }, &excelize.Chart{
//	        Type: "line",
//	        Series: []excelize.ChartSeries{
//	            {
//	                Name:       "Sheet1!$A$4",
//	                Categories: "Sheet1!$B$1:$D$1",
//	                Values:     "Sheet1!$B$4:$D$4",
//	                Marker: excelize.ChartMarker{
//	                    Symbol: "none", Size: 10,
//	                },
//	            },
//	        },
//	        Format: excelize.GraphicOptions{
//	            ScaleX:          1,
//	            ScaleY:          1,
//	            OffsetX:         15,
//	            OffsetY:         10,
//	            PrintObject:     &enable,
//	            LockAspectRatio: false,
//	            Locked:          &disable,
//	        },
//	        Legend: excelize.ChartLegend{
//	            Position:      "right",
//	            ShowLegendKey: false,
//	        },
//	        PlotArea: excelize.ChartPlotArea{
//	            ShowCatName:     false,
//	            ShowLeaderLines: false,
//	            ShowPercent:     true,
//	            ShowSerName:     true,
//	            ShowVal:         true,
//	        },
//	    }); err != nil {
//	        fmt.Println(err)
//	        return
//	    }
//	    // Save spreadsheet by the given path.
//	    if err := f.SaveAs("Book1.xlsx"); err != nil {
//	        fmt.Println(err)
//	    }
//	}
func (f *File) AddChart(sheet, cell string, chart *Chart, combo ...*Chart) error {
	// Read worksheet data
	ws, err := f.workSheetReader(sheet)
	if err != nil {
		return err
	}
	opts, comboCharts, err := f.getChartOptions(chart, combo)
	if err != nil {
		return err
	}
	// Add first picture for given sheet, create xl/drawings/ and xl/drawings/_rels/ folder.
	drawingID := f.countDrawings() + 1
	chartID := f.countCharts() + 1
	drawingXML := "xl/drawings/drawing" + strconv.Itoa(drawingID) + ".xml"
	drawingID, drawingXML = f.prepareDrawing(ws, drawingID, sheet, drawingXML)
	drawingRels := "xl/drawings/_rels/drawing" + strconv.Itoa(drawingID) + ".xml.rels"
	drawingRID := f.addRels(drawingRels, SourceRelationshipChart, "../charts/chart"+strconv.Itoa(chartID)+".xml", "")
	err = f.addDrawingChart(sheet, drawingXML, cell, int(opts.Dimension.Width), int(opts.Dimension.Height), drawingRID, &opts.Format)
	if err != nil {
		return err
	}
	f.addChart(opts, comboCharts)
	if err = f.addContentTypePart(chartID, "chart"); err != nil {
		return err
	}
	_ = f.addContentTypePart(drawingID, "drawings")
	f.addSheetNameSpace(sheet, SourceRelationship)
	return err
}

// AddChartSheet provides the method to create a chartsheet by given chart
// format set (such as offset, scale, aspect ratio setting and print settings)
// and properties set. In Excel a chartsheet is a worksheet that only contains
// a chart.
func (f *File) AddChartSheet(sheet string, chart *Chart, combo ...*Chart) error {
	// Check if the worksheet already exists
	idx, err := f.GetSheetIndex(sheet)
	if err != nil {
		return err
	}
	if idx != -1 {
		return ErrExistsSheet
	}
	opts, comboCharts, err := f.getChartOptions(chart, combo)
	if err != nil {
		return err
	}
	cs := xlsxChartsheet{
		SheetViews: &xlsxChartsheetViews{
			SheetView: []*xlsxChartsheetView{{ZoomScaleAttr: 100, ZoomToFitAttr: true}},
		},
	}
	f.SheetCount++
	wb, _ := f.workbookReader()
	sheetID := 0
	for _, v := range wb.Sheets.Sheet {
		if v.SheetID > sheetID {
			sheetID = v.SheetID
		}
	}
	sheetID++
	path := "xl/chartsheets/sheet" + strconv.Itoa(sheetID) + ".xml"
	f.sheetMap[sheet] = path
	f.Sheet.Store(path, nil)
	drawingID := f.countDrawings() + 1
	chartID := f.countCharts() + 1
	drawingXML := "xl/drawings/drawing" + strconv.Itoa(drawingID) + ".xml"
	f.prepareChartSheetDrawing(&cs, drawingID, sheet)
	drawingRels := "xl/drawings/_rels/drawing" + strconv.Itoa(drawingID) + ".xml.rels"
	drawingRID := f.addRels(drawingRels, SourceRelationshipChart, "../charts/chart"+strconv.Itoa(chartID)+".xml", "")
	if err = f.addSheetDrawingChart(drawingXML, drawingRID, &opts.Format); err != nil {
		return err
	}
	f.addChart(opts, comboCharts)
	if err = f.addContentTypePart(chartID, "chart"); err != nil {
		return err
	}
	_ = f.addContentTypePart(sheetID, "chartsheet")
	_ = f.addContentTypePart(drawingID, "drawings")
	// Update workbook.xml.rels
	rID := f.addRels(f.getWorkbookRelsPath(), SourceRelationshipChartsheet, fmt.Sprintf("/xl/chartsheets/sheet%d.xml", sheetID), "")
	// Update workbook.xml
	f.setWorkbook(sheet, sheetID, rID)
	chartsheet, _ := xml.Marshal(cs)
	f.addSheetNameSpace(sheet, NameSpaceSpreadSheet)
	f.saveFileList(path, replaceRelationshipsBytes(f.replaceNameSpaceBytes(path, chartsheet)))
	return err
}

// getChartOptions provides a function to check format set of the chart and
// create chart format.
func (f *File) getChartOptions(opts *Chart, combo []*Chart) (*Chart, []*Chart, error) {
	var comboCharts []*Chart
	options, err := parseChartOptions(opts)
	if err != nil {
		return options, comboCharts, err
	}
	for _, comboFormat := range combo {
		comboChart, err := parseChartOptions(comboFormat)
		if err != nil {
			return options, comboCharts, err
		}
		if _, ok := chartValAxNumFmtFormatCode[comboChart.Type]; !ok {
			return options, comboCharts, newUnsupportedChartType(comboChart.Type)
		}
		comboCharts = append(comboCharts, comboChart)
	}
	if _, ok := chartValAxNumFmtFormatCode[options.Type]; !ok {
		return options, comboCharts, newUnsupportedChartType(options.Type)
	}
	return options, comboCharts, err
}

// DeleteChart provides a function to delete chart in spreadsheet by given
// worksheet name and cell reference.
func (f *File) DeleteChart(sheet, cell string) error {
	col, row, err := CellNameToCoordinates(cell)
	if err != nil {
		return err
	}
	col--
	row--
	ws, err := f.workSheetReader(sheet)
	if err != nil {
		return err
	}
	if ws.Drawing == nil {
		return err
	}
	drawingXML := strings.ReplaceAll(f.getSheetRelationshipsTargetByID(sheet, ws.Drawing.RID), "..", "xl")
	return f.deleteDrawing(col, row, drawingXML, "Chart")
}

// countCharts provides a function to get chart files count storage in the
// folder xl/charts.
func (f *File) countCharts() int {
	count := 0
	f.Pkg.Range(func(k, v interface{}) bool {
		if strings.Contains(k.(string), "xl/charts/chart") {
			count++
		}
		return true
	})
	return count
}

// ptToEMUs provides a function to convert pt to EMUs, 1 pt = 12700 EMUs. The
// range of pt is 0.25pt - 999pt. If the value of pt is outside the range, the
// default EMUs will be returned.
func (f *File) ptToEMUs(pt float64) int {
	if 0.25 > pt || pt > 999 {
		return 25400
	}
	return int(12700 * pt)
}
