import { useState } from "react";
import logo from "../../assets/icons/kumasan-logo.png";
import "./menu.styles.css";

function MenuComponent({ isOpen }: { isOpen: boolean }) {
  const [activeTab, setActiveTab] = useState("news");

  return (
    <div className={`menu ${isOpen ? "open" : ""}`}>
      <div className="menu-content">
        <img src={logo} alt="Kumasan" />
        <p>
          「Kumasan」は、クマ🐻を発見した際にその情報を報告できるサービスです。クマ🐻が近くにいる場合には通知機能を通じて警告を受け取ることができます。このサービスは、クマ🐻の出没情報を共有し、地域の安全を守るために役立ちます。
        </p>
      </div>
      <div className="menu-tabs">
        <div className="tab-buttons">
          <button
            className={`tab-button ${activeTab === "news" ? "active" : ""}`}
            onClick={() => setActiveTab("news")}
          >
            最近のニュース
          </button>
          <button
            className={`tab-button ${activeTab === "comments" ? "active" : ""}`}
            onClick={() => setActiveTab("comments")}
          >
            最近のコメント
          </button>
        </div>
        <div className="tab-content">
          {activeTab === "news" && (
            <div className="menu-list">
              <ul>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
              </ul>
            </div>
          )}
          {activeTab === "comments" && (
            <div className="menu-list">
              <ul>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
                <li>
                  <p>2025年01月01日</p>
                  <p>ホゲホゲ</p>
                </li>
              </ul>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default MenuComponent;
