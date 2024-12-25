import "./styles.css";
import KUMA_LOGO from "../../assets/icons/kumasan-logo.png";

//これはヘッダーのコンポーネントです。
function HeaderComponent() {
  return (
    <>
      <div className="header">
        <div className="header-container">
          <div className="header-left">
            <button className="menu-icon">
              <img
                src="https://img.icons8.com/ios-filled/500/menu--v6.png"
                alt="menu"
                className="menu-icon-img"
              />
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
