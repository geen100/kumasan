import { useState } from "react";
import "./menu.styles.css";

function MenuComponent({ isOpen }: { isOpen: boolean }) {
  const [activeTab, setActiveTab] = useState("news");

  return (
    <div className={`menu ${isOpen ? "open" : ""}`}>
      <div className="tab-buttons">
        <button
          className={`tab-button ${activeTab === "news" ? "active" : ""}`}
          onClick={() => setActiveTab("news")}
        >
          News
        </button>
        <button
          className={`tab-button ${activeTab === "comments" ? "active" : ""}`}
          onClick={() => setActiveTab("comments")}
        >
          Comments
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
  );
}

export default MenuComponent;
