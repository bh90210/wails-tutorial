import React, { useState } from 'react';
import { useEffect } from 'react';
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

export default function InteractiveList() {
  const classes = useStyles();
  const dense = true;
  const secondary = true;
  const [list, setList] = useState([["shit.go", "poop/shit.go", "666"],["lmao.mp3", "lol/lmao.mp3", "1654"]]);

  useEffect(() => {
    window.backend.FH.ListFiles()
 
    window.wails.Events.On("filesList", (list) => {
      setList(list)
      console.log(list)
      //setList([["shit.go", "poop/shit.go", "666"],["aaa", "aaa", "aaa"]])
    });
  }, []);

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
            {list.map((value, index) => {
            const path = `${value[0]}`;
            const name = `${value[1]}`;
            const size = `${value[2]}`;
        

        return (
                <ListItem key={index}>
                  <ListItemAvatar>
                    <Avatar>
                      <FileCopy />
                    </Avatar>
                  </ListItemAvatar>
                  <ListItemText
                    primary={name}
                    secondary={secondary ? `path: ${path} - size: ${size} bytes` : null}
                  />
                  <ListItemSecondaryAction>
                    <IconButton edge="end" aria-label="download">
                      <DownloadIcon />
                    </IconButton>
                    <IconButton edge="end" aria-label="delete">
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
                );
              })}
            </List>
          </div>
        </Grid>
      </Grid>
    </div>
  );
}