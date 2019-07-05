import React, { Component } from "react";
import { Polar } from "react-chartjs-2";
import Grid from "@material-ui/core/Grid";

class MemChart extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chartData: props.chartData
    };
  }

  static defaultProps = {
    displayTitle: false,
    displayLegend: true,
    legendPosition: "right",
    location: "City"
  };

  render() {
    return (
      <div className="Top-Div">
        <Grid container justify="center" alignItems="center" spacing={3}>
          <Grid item xs={12}>
            <Polar
              data={this.state.chartData}
              options={{
                title: {
                  display: this.props.displayTitle,
                  text: "Largest Cities In " + this.props.location,
                  fontSize: 25
                },
                legend: {
                  display: this.props.displayLegend,
                  position: this.props.legendPosition
                }
              }}
            />
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default MemChart;
