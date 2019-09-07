import React, {Component} from 'react';
import Dropzone from 'react-dropzone';

class FilesDrop extends Component {
  constructor() {
    super();
    this.onDrop = (files) => {
      this.setState({files})

      for (var i = 0; i < files.length; i++) { //for multiple files          
        (function(file) {
            var name = file.name;
            var path = file.path;
            var size = file.size;
            var reader = new FileReader();  
            reader.onload = function(e) {  
                // get file content  
                var binaryStr = e.target.result; 
                window.backend.basic(name, path, size, binaryStr)
            }
            reader.readAsBinaryString(file);
        })(files[i]);
      }
    };
    this.state = {
      files: []
    };
  }

  render() {
    const files = this.state.files.map(file => (
      <li key={file.name}>
        {file.path} - {file.size} bytes
      </li>
    ));

    return (
      <Dropzone onDrop={this.onDrop}>
        {({getRootProps, getInputProps}) => (
          <section className="container">
            <div {...getRootProps({className: 'dropzone'})}>
              <p className="App-header-text">Drag 'n' drop some files here, or click to select files</p>
              <input {...getInputProps()} />
            </div>
            <aside>
              <ul>{files}</ul>
            </aside>
          </section>
        )}
      </Dropzone>
    );
  }
}

export default FilesDrop;