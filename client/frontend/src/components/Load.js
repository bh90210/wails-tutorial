import React, { Component } from "react";
import Button from "@material-ui/core/Button";
import { styled } from "@material-ui/styles";

const LoadButton = styled(Button)({
  background: "linear-gradient(45deg, #fdd835 30%, #ffeb3b 90%)",
  border: 0,
  borderRadius: 3,

  color: "white",
  height: 48,
  padding: "0 30px",
  margin: "10px 10px"
});

class Load extends Component {
  render() {
    return (
      <LoadButton variant="contained" color="primary">
        Load
      </LoadButton>
    );
  }
}

export default Load;
