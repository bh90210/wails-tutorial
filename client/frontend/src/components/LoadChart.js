import React, { Component } from "react";
import { Line, Bubble } from "react-chartjs-2";
import Grid from '@material-ui/core/Grid';

class LoadChart extends Component {
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
        <Grid container spacing={3}>
          <Grid item item xs={12}>
            <Line
          data={this.state.chartData}
          options={{
            title:{
              display:this.props.displayTitle,
              text:'Largest Cities In '+this.props.location,
              fontSize:25
            },
            legend:{
              display:this.props.displayLegend,
              position:this.props.legendPosition
            }
          }}
        />
          </Grid>
          <Grid item item xs={12}>
            <Bubble
          data={this.state.chartData}
          options={{
            title:{
              display:this.props.displayTitle,
              text:'Largest Cities In '+this.props.location,
              fontSize:25
            },
            legend:{
              display:this.props.displayLegend,
              position:this.props.legendPosition
            }
          }}
        />
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default LoadChart;