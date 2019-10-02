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
  const [list, setList] = useState([[]]);
  const [showList, setShowList] = useState(false);

  // functional equivalent of componentDidMout()
  useEffect(() => {
    // fetch files list when frontend inits
    window.backend.FilesHandling.ListFiles()
    
    // start listening for events coming from backend
    window.wails.Events.On("filesList", (list) => {
      console.log(list)
      if (list != null) {
        setShowList(true)
        setList(list)
      }
      if (list == null) {
        setShowList(false)
      } 
    });
  }, []);

  function deleteFile(e) {
    console.log(e);
    window.backend.FilesHandling.DeleteFile(e);
  }

  function downloadFile(e) {
    console.log(e)
    window.backend.FilesHandling.DownloadFile(e)
  }

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
          {showList ?   
          <div className={classes.demo}>
            <List dense={dense}>
            {list.map((value, index) => {
            //const path = `${value[0]}`;
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
                    secondary={secondary ? `size: ${size} bytes` : null}
                  />
                  <ListItemSecondaryAction>
                    <IconButton onClick={(e) => downloadFile(name)} edge="end" aria-label="download">
                      <DownloadIcon />
                    </IconButton>
                    <IconButton onClick={(e) => deleteFile(name)} edge="end" aria-label="delete">
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
                );
              })}
            </List>
          </div>
          // watch out here
          // as we set content to 'nothing'
          // when showList is empty (false)
          : '' }
        </Grid>
      </Grid>
    </div>
  );
}