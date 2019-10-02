import React, {useCallback} from 'react';
import {useState} from 'react';
import {useDropzone} from 'react-dropzone';
import CloudUploadIcon from '@material-ui/icons/CloudUpload';
import DeleteIcon from '@material-ui/icons/Delete';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(theme => ({
  button: {
    margin: theme.spacing(1),
    width: "150px",
  },
  leftIcon: {
    marginRight: theme.spacing(1),
  },
  rightIcon: {
    marginLeft: theme.spacing(1),
  },
  iconSmall: {
    fontSize: 20,
  },
}));

export default function Drop(props) {
  const [isVisible, setVisibility] = useState(false);
  const onDrop = useCallback(acceptedFiles => {
    setVisibility(true)
  }, [])
  const {acceptedFiles, getRootProps, getInputProps, isDragActive} = useDropzone({onDrop});
  const files = acceptedFiles.map(file => (
    <li key={file.name}>
      <div className="App-filesdrop-outer"><div className="App-filesdrop-filepath">{file.path}</div> <div className="App-filesdrop-filezise">{file.size} bytes</div></div>
    </li>
  ));
  const classes = useStyles();

  function hideButtons() {
    setVisibility(false)
  }

  function uploadFiles() {
    for (var i = 0; i < acceptedFiles.length; i++) { //for multiple files          
      (function(file) {

        var name = file.name;
        //var path = file.path;
        //var size = file.size;
        var reader = new FileReader();  
        reader.onload = function(e) {  
            // get file content  
            var binaryStr = e.target.result; 
            // pass it to backend
            window.backend.FilesHandling.UploadFile(name, binaryStr)
        }
        reader.readAsBinaryString(file);
      })(acceptedFiles[i]);
    }

    hideButtons()
    //setList([])
  }

  return (
    <section className="container">
      <div {...getRootProps({className: 'dropzone'})}>
        <input {...getInputProps()} />
        {
          isDragActive ?
            <p>Drop the files here ...</p> :
            <p>Drag 'n' drop a file here, or click for window dialog</p>
        }
      </div>
      <div className="App-button-div">
        {isVisible && <Button onClick={uploadFiles} className={classes.button} variant="contained" color="primary">
              Upload
              <CloudUploadIcon />
            </Button>}
        {isVisible && <Button onClick={hideButtons} className={classes.button} variant="contained" color="secondary">
              Clear 
              <DeleteIcon />
            </Button>}
      </div>
      <aside>
        <ul>{isVisible && files}</ul>
      </aside>
    </section>
  );
}
