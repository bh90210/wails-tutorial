import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import FileCopy from '@material-ui/icons/FileCopy';
import DeleteIcon from '@material-ui/icons/Delete';
import DownloadIcon from '@material-ui/icons/VerticalAlignBottom';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    maxWidth: "99%",    
  },
  demo: {
    backgroundColor: theme.palette.background.paper,
  },
  title: {
    margin: theme.spacing(4, 0, 2),
  },
}));

var arrr = [0,1,2]
//var arrr

function generate(element) {
  return arrr.map(value =>
    React.cloneElement(element, {
      key: value,
    }),
  );
}

export default function InteractiveList() {
  const classes = useStyles();
  const [dense, setDense] = React.useState(false);
  const [secondary, setSecondary] = React.useState(true);

  
  window.wails.Events.On("cpu_usage", (cpu_usage) => {
    arrr = cpu_usage
  });

  return (
    <div className={classes.root}>
      <Grid container 
        direction="row"
        justify="center"
        alignItems="center" spacing={3}>
        <Grid item xs={9} md={6}>
          <Typography variant="h6" className={classes.title}>
            Remote Files
          </Typography>
          <div className={classes.demo}>
            <List dense={dense}>
              {
               generate(
                <ListItem>
                  <ListItemAvatar>
                    <Avatar>
                      <FileCopy />
                    </Avatar>
                  </ListItemAvatar>
                  <ListItemText
                    primary="Single-line item"
                    secondary={secondary ? 'Secondary text' : null}
                  />
                  <ListItemSecondaryAction>
                    <IconButton edge="end" aria-label="download">
                      <DownloadIcon />
                    </IconButton>
                    <IconButton edge="end" aria-label="delete">
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>,
              )}
            </List>
          </div>
        </Grid>
      </Grid>
    </div>
  );
}