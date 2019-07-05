import React, { Component } from "react";
import Button from "@material-ui/core/Button";
import { styled } from "@material-ui/styles";

const MemButton = styled(Button)({
  background: "linear-gradient(45deg, #4caf50 30%, #00c853 90%)",
  border: 0,
  borderRadius: 3,

  color: "white",
  height: 48,
  padding: "0 30px",
  margin: "10px 10px"
});

class Mem extends Component {
  render() {
    return (
      <MemButton variant="contained" color="primary">
        Memory
      </MemButton>
    );
  }
}

export default Mem;
