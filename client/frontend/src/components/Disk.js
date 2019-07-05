import React, { Component } from "react";
import Button from "@material-ui/core/Button";
import { styled } from "@material-ui/styles";

const DiskButton = styled(Button)({
  background: "linear-gradient(45deg, #2196F3 30%, #21CBF3 90%)",
  border: 0,
  borderRadius: 3,

  color: "white",
  height: 48,
  padding: "0 30px",
  margin: "10px 10px"
});

class Disk extends Component {
  render() {
    return (
      <DiskButton variant="contained" color="primary">
        Disk
      </DiskButton>
    );
  }
}

export default Disk;
