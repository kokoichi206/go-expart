package jp.mydns.kokoichi206.catgame;

import android.os.Bundle;

import androidx.appcompat.app.AppCompatActivity;

import go.Seq;
import jp.mydns.kokoichi206.game.cat.mobile.EbitenView;

public class MainActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Seq.setContext(getApplicationContext());
    }

    private EbitenView getEbitenView() {
        return (EbitenView) findViewById(R.id.ebitenview);
    }

    @Override
    protected void onPause() {
        super.onPause();
        this.getEbitenView().suspendGame();
    }

    @Override
    protected void onResume() {
        super.onResume();
        this.getEbitenView().resumeGame();
    }
}