import React from 'react';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core/styles';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import ComputerIcon from '@material-ui/icons/Computer';
import StorageIcon from '@material-ui/icons/Storage';
import MemoryIcon from '@material-ui/icons/Memory';
import LoadIcon from '@material-ui/icons/Reorder'
import Cpu from "./Cpu";
import Load from "./Load";
import Mem from "./Mem";
import PropTypes from 'prop-types';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  paper: {
    flexGrow: 1,
    elevation: 0,
    square: true,
    padding: '0 20px',
    boxShadow: '0 0px 0px 0px rgba(0, 0, 0, 0)',
  },
  cpu: {
  },
  load: {
  },
  mem: {
  },
});

function TabContainer(props) {
  return (
    <Typography component="div" style={{ padding: 8 * 3 }}>
      {props.children}
    </Typography>
  );
}

function LinkTab(props) {
  return (
    <Tab
      component="a"
      onClick={event => {
        event.preventDefault();
      }}
      {...props}
    />
  );
}

TabContainer.propTypes = {
  children: PropTypes.node.isRequired,
};

export default function IconLabelTabs() {
  const classes = useStyles();
  const [value, setValue] = React.useState(0);

  function handleChange(event, newValue) {
    setValue(newValue);
  }

  function handleClickCPU(e) {
    e.preventDefault();
    window.backend.ServiceChooser.Choose("1").then(console.log)
  }
  function handleClickLOAD(e) {
    e.preventDefault();
    window.backend.ServiceChooser.Choose("3").then(console.log)
  }
  function handleClickMEM(e) {
    e.preventDefault();
    window.backend.ServiceChooser.Choose("4").then(console.log)
  }

  return (
    <Paper square className={classes.paper}>
      <Tabs
        value={value}
        onChange={handleChange}
        variant="fullWidth"
        indicatorColor="secondary"
        textColor="secondary"
        centered
      >
        <Tab className={classes.cpu} icon={<ComputerIcon />} label="CPU" onClick={handleClickCPU} />
        <Tab className={classes.load} icon={<LoadIcon />} label="LOAD" onClick={handleClickLOAD} />
        <Tab className={classes.mem} icon={<MemoryIcon />} label="MEMORY" onClick={handleClickMEM} />
      </Tabs>
      {value === 0 && <Cpu />}
      {value === 1 && <Load />}
      {value === 2 && <Mem />}
    </Paper>
  );
}
