import React from "react";
import { RiCloseLine } from "react-icons/ri";
import styles from "./Modal.module.css";


const Modal = ({ setIsOpen }) => {
  return (
    <>
      <div className={styles.darkBG} onClick={() => setIsOpen(false)} />
      <div className={styles.centered}>
        <div className={styles.modal}>
          <div className={styles.modalHeader}>
            <h5 className={styles.heading}>Dialog</h5>
          </div>
          <button className={styles.closeBtn} onClick={() => setIsOpen(false)}>
            <RiCloseLine style={{ marginBottom: "-3px" }} />
          </button>
          <div className={styles.modalContent}>
          One of the classes we're going to have is called .darkBG that's because once the modal is open I'll add a background color to slightly hide all the other components that are on the page. This is to focus the user's attention only on the modal.

          Then our component will be divided into three areas, the first one will be the header, where you can put the modal title. The second part will be the content, here you can put the message you want.

          The third and last part will be the actions that can be performed in the modal, that is, cancel the modal in order to close it and another action (save, update, delete, etc).
           </div>
           <div className={styles.modalImage}>IMG</div>
           <div className={styles.modalBottom}> </div>
          <div className={styles.modalActions}>
            <div className={styles.actionsContainer}>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
export default Modal;
