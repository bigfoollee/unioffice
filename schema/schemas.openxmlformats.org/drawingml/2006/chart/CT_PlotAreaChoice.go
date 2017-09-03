// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_PlotAreaChoice struct {
	AreaChart      *CT_AreaChart
	Area3DChart    *CT_Area3DChart
	LineChart      *CT_LineChart
	Line3DChart    *CT_Line3DChart
	StockChart     *CT_StockChart
	RadarChart     *CT_RadarChart
	ScatterChart   *CT_ScatterChart
	PieChart       *CT_PieChart
	Pie3DChart     *CT_Pie3DChart
	DoughnutChart  *CT_DoughnutChart
	BarChart       *CT_BarChart
	Bar3DChart     *CT_Bar3DChart
	OfPieChart     *CT_OfPieChart
	SurfaceChart   *CT_SurfaceChart
	Surface3DChart *CT_Surface3DChart
	BubbleChart    *CT_BubbleChart
}

func NewCT_PlotAreaChoice() *CT_PlotAreaChoice {
	ret := &CT_PlotAreaChoice{}
	return ret
}

func (m *CT_PlotAreaChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AreaChart != nil {
		seareaChart := xml.StartElement{Name: xml.Name{Local: "c:areaChart"}}
		e.EncodeElement(m.AreaChart, seareaChart)
	}
	if m.Area3DChart != nil {
		searea3DChart := xml.StartElement{Name: xml.Name{Local: "c:area3DChart"}}
		e.EncodeElement(m.Area3DChart, searea3DChart)
	}
	if m.LineChart != nil {
		selineChart := xml.StartElement{Name: xml.Name{Local: "c:lineChart"}}
		e.EncodeElement(m.LineChart, selineChart)
	}
	if m.Line3DChart != nil {
		seline3DChart := xml.StartElement{Name: xml.Name{Local: "c:line3DChart"}}
		e.EncodeElement(m.Line3DChart, seline3DChart)
	}
	if m.StockChart != nil {
		sestockChart := xml.StartElement{Name: xml.Name{Local: "c:stockChart"}}
		e.EncodeElement(m.StockChart, sestockChart)
	}
	if m.RadarChart != nil {
		seradarChart := xml.StartElement{Name: xml.Name{Local: "c:radarChart"}}
		e.EncodeElement(m.RadarChart, seradarChart)
	}
	if m.ScatterChart != nil {
		sescatterChart := xml.StartElement{Name: xml.Name{Local: "c:scatterChart"}}
		e.EncodeElement(m.ScatterChart, sescatterChart)
	}
	if m.PieChart != nil {
		sepieChart := xml.StartElement{Name: xml.Name{Local: "c:pieChart"}}
		e.EncodeElement(m.PieChart, sepieChart)
	}
	if m.Pie3DChart != nil {
		sepie3DChart := xml.StartElement{Name: xml.Name{Local: "c:pie3DChart"}}
		e.EncodeElement(m.Pie3DChart, sepie3DChart)
	}
	if m.DoughnutChart != nil {
		sedoughnutChart := xml.StartElement{Name: xml.Name{Local: "c:doughnutChart"}}
		e.EncodeElement(m.DoughnutChart, sedoughnutChart)
	}
	if m.BarChart != nil {
		sebarChart := xml.StartElement{Name: xml.Name{Local: "c:barChart"}}
		e.EncodeElement(m.BarChart, sebarChart)
	}
	if m.Bar3DChart != nil {
		sebar3DChart := xml.StartElement{Name: xml.Name{Local: "c:bar3DChart"}}
		e.EncodeElement(m.Bar3DChart, sebar3DChart)
	}
	if m.OfPieChart != nil {
		seofPieChart := xml.StartElement{Name: xml.Name{Local: "c:ofPieChart"}}
		e.EncodeElement(m.OfPieChart, seofPieChart)
	}
	if m.SurfaceChart != nil {
		sesurfaceChart := xml.StartElement{Name: xml.Name{Local: "c:surfaceChart"}}
		e.EncodeElement(m.SurfaceChart, sesurfaceChart)
	}
	if m.Surface3DChart != nil {
		sesurface3DChart := xml.StartElement{Name: xml.Name{Local: "c:surface3DChart"}}
		e.EncodeElement(m.Surface3DChart, sesurface3DChart)
	}
	if m.BubbleChart != nil {
		sebubbleChart := xml.StartElement{Name: xml.Name{Local: "c:bubbleChart"}}
		e.EncodeElement(m.BubbleChart, sebubbleChart)
	}
	return nil
}

func (m *CT_PlotAreaChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PlotAreaChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "areaChart":
				m.AreaChart = NewCT_AreaChart()
				if err := d.DecodeElement(m.AreaChart, &el); err != nil {
					return err
				}
			case "area3DChart":
				m.Area3DChart = NewCT_Area3DChart()
				if err := d.DecodeElement(m.Area3DChart, &el); err != nil {
					return err
				}
			case "lineChart":
				m.LineChart = NewCT_LineChart()
				if err := d.DecodeElement(m.LineChart, &el); err != nil {
					return err
				}
			case "line3DChart":
				m.Line3DChart = NewCT_Line3DChart()
				if err := d.DecodeElement(m.Line3DChart, &el); err != nil {
					return err
				}
			case "stockChart":
				m.StockChart = NewCT_StockChart()
				if err := d.DecodeElement(m.StockChart, &el); err != nil {
					return err
				}
			case "radarChart":
				m.RadarChart = NewCT_RadarChart()
				if err := d.DecodeElement(m.RadarChart, &el); err != nil {
					return err
				}
			case "scatterChart":
				m.ScatterChart = NewCT_ScatterChart()
				if err := d.DecodeElement(m.ScatterChart, &el); err != nil {
					return err
				}
			case "pieChart":
				m.PieChart = NewCT_PieChart()
				if err := d.DecodeElement(m.PieChart, &el); err != nil {
					return err
				}
			case "pie3DChart":
				m.Pie3DChart = NewCT_Pie3DChart()
				if err := d.DecodeElement(m.Pie3DChart, &el); err != nil {
					return err
				}
			case "doughnutChart":
				m.DoughnutChart = NewCT_DoughnutChart()
				if err := d.DecodeElement(m.DoughnutChart, &el); err != nil {
					return err
				}
			case "barChart":
				m.BarChart = NewCT_BarChart()
				if err := d.DecodeElement(m.BarChart, &el); err != nil {
					return err
				}
			case "bar3DChart":
				m.Bar3DChart = NewCT_Bar3DChart()
				if err := d.DecodeElement(m.Bar3DChart, &el); err != nil {
					return err
				}
			case "ofPieChart":
				m.OfPieChart = NewCT_OfPieChart()
				if err := d.DecodeElement(m.OfPieChart, &el); err != nil {
					return err
				}
			case "surfaceChart":
				m.SurfaceChart = NewCT_SurfaceChart()
				if err := d.DecodeElement(m.SurfaceChart, &el); err != nil {
					return err
				}
			case "surface3DChart":
				m.Surface3DChart = NewCT_Surface3DChart()
				if err := d.DecodeElement(m.Surface3DChart, &el); err != nil {
					return err
				}
			case "bubbleChart":
				m.BubbleChart = NewCT_BubbleChart()
				if err := d.DecodeElement(m.BubbleChart, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PlotAreaChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PlotAreaChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PlotAreaChoice and its children
func (m *CT_PlotAreaChoice) Validate() error {
	return m.ValidateWithPath("CT_PlotAreaChoice")
}

// ValidateWithPath validates the CT_PlotAreaChoice and its children, prefixing error messages with path
func (m *CT_PlotAreaChoice) ValidateWithPath(path string) error {
	if m.AreaChart != nil {
		if err := m.AreaChart.ValidateWithPath(path + "/AreaChart"); err != nil {
			return err
		}
	}
	if m.Area3DChart != nil {
		if err := m.Area3DChart.ValidateWithPath(path + "/Area3DChart"); err != nil {
			return err
		}
	}
	if m.LineChart != nil {
		if err := m.LineChart.ValidateWithPath(path + "/LineChart"); err != nil {
			return err
		}
	}
	if m.Line3DChart != nil {
		if err := m.Line3DChart.ValidateWithPath(path + "/Line3DChart"); err != nil {
			return err
		}
	}
	if m.StockChart != nil {
		if err := m.StockChart.ValidateWithPath(path + "/StockChart"); err != nil {
			return err
		}
	}
	if m.RadarChart != nil {
		if err := m.RadarChart.ValidateWithPath(path + "/RadarChart"); err != nil {
			return err
		}
	}
	if m.ScatterChart != nil {
		if err := m.ScatterChart.ValidateWithPath(path + "/ScatterChart"); err != nil {
			return err
		}
	}
	if m.PieChart != nil {
		if err := m.PieChart.ValidateWithPath(path + "/PieChart"); err != nil {
			return err
		}
	}
	if m.Pie3DChart != nil {
		if err := m.Pie3DChart.ValidateWithPath(path + "/Pie3DChart"); err != nil {
			return err
		}
	}
	if m.DoughnutChart != nil {
		if err := m.DoughnutChart.ValidateWithPath(path + "/DoughnutChart"); err != nil {
			return err
		}
	}
	if m.BarChart != nil {
		if err := m.BarChart.ValidateWithPath(path + "/BarChart"); err != nil {
			return err
		}
	}
	if m.Bar3DChart != nil {
		if err := m.Bar3DChart.ValidateWithPath(path + "/Bar3DChart"); err != nil {
			return err
		}
	}
	if m.OfPieChart != nil {
		if err := m.OfPieChart.ValidateWithPath(path + "/OfPieChart"); err != nil {
			return err
		}
	}
	if m.SurfaceChart != nil {
		if err := m.SurfaceChart.ValidateWithPath(path + "/SurfaceChart"); err != nil {
			return err
		}
	}
	if m.Surface3DChart != nil {
		if err := m.Surface3DChart.ValidateWithPath(path + "/Surface3DChart"); err != nil {
			return err
		}
	}
	if m.BubbleChart != nil {
		if err := m.BubbleChart.ValidateWithPath(path + "/BubbleChart"); err != nil {
			return err
		}
	}
	return nil
}
