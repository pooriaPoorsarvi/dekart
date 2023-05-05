import { useSelector, useDispatch } from 'react-redux'
import message from 'antd/es/message'
import { useEffect } from 'react'
import {slowDownLimit, finishAlertLimit} from './actions'

function ExcessDataMessage () {
  return (<span>More than {slowDownLimit} rows detected, can slow down the app</span>)
}

export default function Excess () {
  const excessRows = useSelector(state => state.excessRows)
  const [api, contextHolder] = message.useMessage();
  const show = excessRows.length > 0;
  const dispatch = useDispatch();
  useEffect(() => {
    
    if(show){
      api.info({
        content: <ExcessDataMessage />,
        duration: 10,
        onClose: () => {
          for(let excessRow of excessRows){
            dispatch(finishAlertLimit(excessRow.rows, excessRow.datasetId));   
          }
        }
      })
    }
  }, [api, show, dispatch, excessRows])
  return contextHolder
}
