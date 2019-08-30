import React, {Component} from 'react';
import Dropzone from 'react-dropzone';

class FilesDrop extends Component {
  constructor() {
    super();
    this.onDrop = (files) => {
      this.setState({files})
      const reader = new FileReader()

      reader.onabort = () => console.log('file reading was aborted')
      reader.onerror = () => console.log('file reading has failed')
      reader.onload = () => {
        // Do whatever you want with the file contents
        const binaryStr = reader.result
        console.log(binaryStr)
      }

      files.forEach(file => reader.readAsBinaryString(file))
      //window.wails.Events.On("filesDropped", {file.path});
    };
    this.state = {
      files: []
    };
  }

  render() {
    const files = this.state.files.map(file => (
      <li key={file.name}>
        {file.path} - {file.name} - {file.size} bytes
      </li>
    ));

    return (
      <Dropzone onDrop={this.onDrop}>
        {({getRootProps, getInputProps}) => (
          <section className="container">
            <div {...getRootProps({className: 'dropzone'})}>
              <input {...getInputProps()} />
              <p>Drag 'n' drop some files here, or click to select files</p>
            </div>
            <aside>
              <h4>Files</h4>
              <ul>{files}</ul>
            </aside>
          </section>
        )}
      </Dropzone>
    );
  }
}

export default FilesDrop;