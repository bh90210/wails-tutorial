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
import Disk from "./Disk";
import Load from "./Load";
import Mem from "./Mem";
import PropTypes from 'prop-types';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  root: {
    flexGrow: 1,
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

  return (
    <Paper square className={classes.root}>
      <Tabs
        value={value}
        onChange={handleChange}
        variant="fullWidth"
        indicatorColor="secondary"
        textColor="secondary"
        centered
      >
        <Tab icon={<ComputerIcon />} label="CPU" />
        <Tab icon={<StorageIcon />} label="DISK" />
        <Tab icon={<LoadIcon />} label="LOAD" />
        <Tab icon={<MemoryIcon />} label="MEMORY" />
      </Tabs>
      {value === 0 && <Cpu />}
      {value === 1 && <Disk />}
      {value === 2 && <Load />}
      {value === 3 && <Mem />}
    </Paper>
  );
}
