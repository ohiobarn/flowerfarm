<?php

namespace KaliForms\Inc\Backend\Views\Metaboxes;

if (!defined('WPINC')) {
    die;
}

/**
 * Class Form_Builder
 *
 * @package Inc\Backend\Views\Metaboxes;
 */
class Form_Builder extends Metabox
{

    /**
     * Render the metabox
     */
    public function render_box()
    {
        wp_nonce_field(KALIFORMS_BASE, 'kaliforms_fields');
        echo '<div id="kaliforms-container">';
        // echo '<div class="kaliforms-preloader"><span class="focus-in-contract">' . esc_html__('Loading', 'kaliforms') . '</span></div>';
        echo '</div>';
    }
}
