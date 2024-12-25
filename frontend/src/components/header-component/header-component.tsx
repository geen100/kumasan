import "./styles.css";
import KUMA_LOGO from "../../assets/icons/kumasan-logo.png";
import MENU_ICON from "../../assets/icons/menubar-icon.png";

//これはヘッダーのコンポーネントです。
function HeaderComponent() {
  return (
    <>
      <div className="header">
        <div className="header-container">
          <div className="header-left">
            <button className="menu-icon">
              <img src={MENU_ICON} alt="menu" />
            </button>
            <img src={KUMA_LOGO} alt="logo" className="kumasan-logo" />
          </div>
          <div className="header-right"></div>
        </div>
      </div>
    </>
  );
}

export default HeaderComponent;
