import { useState, useEffect, useRef } from "react";
import "./header.styles.css";
import KUMA_LOGO from "../../assets/icons/kumasan-logo.png";
import MenuComponent from "../menu-component/menu-component";

//これはヘッダーのコンポーネントです。
function HeaderComponent() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const headerRef = useRef<HTMLDivElement | null>(null);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const handleClickOutside = (event: MouseEvent) => {
    if (
      headerRef.current &&
      !headerRef.current.contains(event.target as Node)
    ) {
      setIsMenuOpen(false);
    }
  };

  useEffect(() => {
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
    <>
      <div className="header" ref={headerRef}>
        <div className="header-container">
          <div className="header-left">
            <button
              className={`menu-icon ${isMenuOpen ? "active" : ""}`}
              onClick={toggleMenu}
            >
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
