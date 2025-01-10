import { useState } from "react";
import "./header.styles.css";
import KUMA_LOGO from "../../assets/icons/kumasan-logo.png";
import MenuComponent from "../menu-component/menu-component";

//これはヘッダーのコンポーネントです。
function HeaderComponent() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  return (
    <>
      <div className="header">
        <div className="header-container">
          <div className="header-left">
            <button className="menu-icon" onClick={toggleMenu}>
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
        <MenuComponent isOpen={isMenuOpen} />
      </div>
    </>
  );
}

export default HeaderComponent;
