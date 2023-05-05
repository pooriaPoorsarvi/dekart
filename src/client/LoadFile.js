
import Button from 'antd/es/button'
import Modal from 'antd/es/modal'
import { useState } from 'react'
import styles from './LoadFile.module.css'
import { useSelector, useDispatch } from 'react-redux'
import {downloadDataset} from './actions'


export default function LoadFileButton () {
    const [modalOpen, setModalOpen] = useState(true)
    const fileLoadNeeded = useSelector(state => state.fileLoadNeeded);
    const dispatch = useDispatch();
    if (fileLoadNeeded.length>0){
        return (
        // <>
            <Modal
            title='Load report'
            visible={modalOpen}
            onOk={() => setModalOpen(false)}
            onCancel={() => setModalOpen(false)}
            bodyStyle={{ padding: '0px' }}
            footer={
                <div className={styles.modalFooter}>
                <Button type='danger' onClick={() => setModalOpen(false)} >
                    Skip
                </Button>
                <Button type='primary' onClick={() => {
                    // dispatch();
                    // console.log(fileLoadNeeded);
                    for(let file of fileLoadNeeded){
                        // console.log(file);
                        dispatch(downloadDataset(
                            file.dataset, 
                            file.sourceId, 
                            file.extension, 
                            file.label
                            )
                        )
                    }
                    setModalOpen(false);
                }} >
                    Load
                </Button>
                </div>
            }
            style={{ textAlign: "center" }}
            >
                File too large as a variable, load in alternate format? it takes more space, you should close your other applications as we can not garauntee your memory will allow for this dataset.
            </Modal>
        // </>
        );
    }else{
        return <></>;
    }
}
  