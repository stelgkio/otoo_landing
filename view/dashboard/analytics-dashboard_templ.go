// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AnalyticsDashboard() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"container-fluid px-3 py-5 p-lg-6 p-xxl-8\"><div class=\"mb-6 mb-xl-10\"><div class=\"row g-3 align-items-center\"><div class=\"col\"><h1 class=\"ls-tight\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := `Orders`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"col\"><div class=\"hstack gap-2 justify-content-end\"><button hx-get=\"/product\" hx-target=\"#dashboard-content\" type=\"button\" class=\"btn btn-sm btn-neutral d-sm-inline-flex\"><span class=\"pe-2\"></span><span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := `Products`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></button> <button hx-get=\"/customer\" hx-target=\"#dashboard-content\" type=\"button\" class=\"btn btn-sm btn-neutral  d-sm-inline-flex\"><span class=\"pe-2\"></span><span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `Customers`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></button></div></div></div></div><div class=\"vstack gap-3 gap-xl-6 mt-8\"><div class=\"row row-cols-sm-2 row-cols-md-4 g-3\"><div class=\"col\"><div class=\"card border-primary-hover\"><div class=\"card-body d-flex gap-3\"><div class=\"\"><span class=\"d-block text-muted mb-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var5 := `Total Orders`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"d-block text-lg fw-bold text-heading\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var6 := `123.456`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var6)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div></div></div><div class=\"col\"><div class=\"card\"><div class=\"card-body\"><div class=\"d-flex gap-3 mb-1\"><span class=\"text-muted\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var7 := `Orders`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"badge text-bg-light rounded\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := `24H`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"text-muted text-opacity-60 text-opacity-100-hover\" tabindex=\"0\" data-bs-toggle=\"popover\" data-bs-trigger=\"hover focus\" data-bs-placement=\"right\" data-bs-html=\"true\" data-bs-content=\"The total market value of a cryptocurrency&#39;s circulating supply. It is analogous to the free-float capitalization in the stock market.&lt;br /&gt;&lt;br /&gt;Market Cap = Current Price x Circulating Supply.\" role=\"button\"><i class=\"bi bi-info-circle\"></i></span></div><div class=\"d-flex align-items-center\"><span class=\"text-lg text-heading fw-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var9 := `1.987 `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var9)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><span class=\"ms-auto text-success text-sm fw-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var10 := `+23%`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var10)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div></div></div><div class=\"col\"><div class=\"card\"><div class=\"card-body\"><div class=\"d-flex gap-3 mb-1\"><span class=\"text-muted\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var11 := `Orders`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var11)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"badge text-bg-light rounded\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var12 := `7D`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var12)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"text-muted text-opacity-60 text-opacity-100-hover\" tabindex=\"0\" data-bs-toggle=\"popover\" data-bs-trigger=\"hover focus\" data-bs-placement=\"right\" data-bs-html=\"true\" data-bs-content=\"The total market value of a cryptocurrency&#39;s circulating supply. It is analogous to the free-float capitalization in the stock market.&lt;br /&gt;&lt;br /&gt;Market Cap = Current Price x Circulating Supply.\" role=\"button\"><i class=\"bi bi-info-circle\"></i></span></div><div class=\"d-flex align-items-center\"><span class=\"text-lg text-heading fw-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var13 := `20.987 `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var13)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><span class=\"ms-auto text-success text-sm fw-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var14 := `+23%`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var14)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div></div></div><div class=\"col\"><div class=\"card\"><div class=\"card-body\"><div class=\"d-flex gap-3 mb-1\"><span class=\"text-muted\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var15 := `Orders`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var15)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"badge text-bg-light rounded\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var16 := `30D`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var16)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"text-muted text-opacity-60 text-opacity-100-hover\" tabindex=\"0\" data-bs-toggle=\"popover\" data-bs-trigger=\"hover focus\" data-bs-placement=\"right\" data-bs-html=\"true\" data-bs-content=\"The total market value of a cryptocurrency&#39;s circulating supply. It is analogous to the free-float capitalization in the stock market.&lt;br /&gt;&lt;br /&gt;Market Cap = Current Price x Circulating Supply.\" role=\"button\"><i class=\"bi bi-info-circle\"></i></span></div><div class=\"d-flex align-items-center\"><span class=\"text-lg text-heading fw-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var17 := `34.000 `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var17)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><span class=\"ms-auto text-success text-sm fw-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var18 := `+23%`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var18)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div></div></div></div><div class=\"row g-3 g-xl-6\"><div class=\"col-xxl-12\"><div class=\"card\"><div class=\"card-body pb-0\"><div class=\"row g-3 justify-content-between align-items-center\"><div class=\"col-12 col-sm\"><h5>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var19 := `Orders chart`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var19)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h5></div><div class=\"col-12 col-sm-auto\"><div class=\"d-flex scrollable-x justify-content-between gap-1 p-1 align-items-center bg-body-secondary rounded text-xs fw-semibold\"><a href=\"#\" class=\"px-3 py-1 text-muted  bg-white-hover bg-opacity-50-hover rounded\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var20 := `1W`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var20)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a> <a href=\"#\" class=\"px-3 py-1 text-muted bg-white-hover bg-opacity-50-hover rounded\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var21 := `1M`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var21)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a> <a href=\"#\" class=\"d-none d-sm-inline-block px-3 py-1 text-muted bg-white-hover bg-opacity-50-hover rounded\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var22 := `1Y`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var22)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div></div></div><div class=\"mx-n4\"><div id=\"chart\" data-height=\"420\"></div></div></div></div></div></div></div></main></div><script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var23 := `
        // Assuming you have an HTML element with id "chart" to render the chart
        var chartElement = document.getElementById("chart");

        var options = {
            chart: {
                zoom: {
                    enabled: false
                },
                toolbar: {
                    show: false
                },
                shadow: {
                    enabled: false
                },
                animations: {
                    enabled: false
                },
                height: 390,
            },
            colors: ['#007bff', '#8957ff', '#ffc107', '#dc3545'], // Using the same colors as the original example
            stroke: {
                width: 4,
                curve: "smooth"
            },
            series: [
                { name: "Card", type: "line", data: [30, 50, 70, 90, 80, 70, 90, 120, 100, 120, 140] },
                { name: "Cash", type: "line", data: [50, 70, 30, 20, 10, 10, 40, 100, 90, 100, 120] }
            ],
            markers: {
                size: 0
            },
            xaxis: {
                axisBorder: {
                    show: false
                },
                axisTicks: {
                    show: false
                },
                categories: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
                labels: {
                    style: {
                        colors: "#e3dede",
                        fontSize: "13px",
                        fontFamily: "#333",
                        cssClass: "apexcharts-xaxis-label"
                    }
                }
            },
            yaxis: {
                axisBorder: {
                    show: false
                },
                axisTicks: {
                    show: false
                },
                labels: {
                    style: {
                        colors: "#e3dede",
                        fontSize: "13px",
                        fontFamily: "inherit",
                        cssClass: "apexcharts-xaxis-label"
                    }
                }
            },
            legend: {
                show: false
            },
            grid: {
                borderColor: "#807c7c",
                strokeDashArray: 3
            },
            dataLabels: {
                enabled: false
            },
            tooltip: {
                shared: true,
                intersect: false,
                y: {
                    formatter: function (e) {
                        return e !== undefined ? e.toFixed(0) + " orders" : e;
                    }
                }
            }
        };

        var lineChart = new ApexCharts(chartElement, options);
        lineChart.render();


    `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var23)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
